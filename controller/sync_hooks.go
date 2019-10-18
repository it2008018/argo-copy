package controller

import (
	"fmt"

	"github.com/argoproj/argo-cd/util/health"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/apis/batch"

	"github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
)

// getOperationPhase returns a hook status from an _live_ unstructured object
func getOperationPhase(hook *unstructured.Unstructured, resourceOverrides map[string]appv1.ResourceOverride) (operation v1alpha1.OperationPhase, message string) {
	gvk := hook.GroupVersionKind()
	if isBatchJob(gvk) {
		return getStatusFromBatchJob(hook)
	} else if isArgoWorkflow(gvk) {
		return health.GetStatusFromArgoWorkflow(hook)
	} else if isPod(gvk) {
		return getStatusFromPod(hook)
	} else {
		// TODO users should opt-in to hook resource health check via annotation
		checkHealth = true
		if checkHealth {
			return getResourceHealth(hook, resourceOverrides)
		} else {
			return v1alpha1.OperationSucceeded, fmt.Sprintf("%s created", hook.GetName())
		}
	}
}

// isRunnable returns if the resource object is a runnable type which needs to be terminated
func isRunnable(gvk schema.GroupVersionKind) bool {
	return isBatchJob(gvk) || isArgoWorkflow(gvk) || isPod(gvk)
}

func isBatchJob(gvk schema.GroupVersionKind) bool {
	return gvk.Group == "batch" && gvk.Kind == "Job"
}

func getResourceHealth(hook *unstructured.Unstructured, resourceOverrides map[string]appv1.ResourceOverride) (operation v1alpha.OperationPhase, message string) {
	health, err := health.GetResourceHealth(hook, resourceOverrides)
	if err != nil {
		return v1alpha1.OperationError, err.Error()
	}

	switch health.Status {
	case v1alpha1.HealthStatusProgressing:
		return v1alpha.OperationRunning, health.Message
	case v1alpha1.HealthStatusHealthy:
		return v1alpha.OperationSucceeded, health.Message
	case v1alpha1.HealthStatusDegraded:
		fallthrough
	case v1alpha1.HealthStatusMissing:
		fallthrough
	case v1alpha1.HealthStatusUnknown:
		fallthrough
	case v1alpha1.HealthStatusSuspended:
		fallthrough
	case v1alpha1.HealthStatusMissing:
		return v1alpha.OperationFailed, health.Message
	}

}

// TODO this is a copy-and-paste of health.getJobHealth(), refactor out?
func getStatusFromBatchJob(hook *unstructured.Unstructured) (operation v1alpha1.OperationPhase, message string) {
	var job batch.Job
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(hook.Object, &job)
	if err != nil {
		return v1alpha1.OperationError, err.Error()
	}
	failed := false
	var failMsg string
	complete := false
	for _, condition := range job.Status.Conditions {
		switch condition.Type {
		case batch.JobFailed:
			failed = true
			complete = true
			failMsg = condition.Message
		case batch.JobComplete:
			complete = true
			message = condition.Message
		}
	}
	if !complete {
		return v1alpha1.OperationRunning, message
	} else if failed {
		return v1alpha1.OperationFailed, failMsg
	} else {
		return v1alpha1.OperationSucceeded, message
	}
}

func isArgoWorkflow(gvk schema.GroupVersionKind) bool {
	return gvk.Group == "argoproj.io" && gvk.Kind == "Workflow"
}

func isPod(gvk schema.GroupVersionKind) bool {
	return gvk.Group == "" && gvk.Kind == "Pod"
}

// TODO - this is very similar to health.getPodHealth() should we use that instead?
func getStatusFromPod(hook *unstructured.Unstructured) (v1alpha1.OperationPhase, string) {
	var pod apiv1.Pod
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(hook.Object, &pod)
	if err != nil {
		return v1alpha1.OperationError, err.Error()
	}
	getFailMessage := func(ctr *apiv1.ContainerStatus) string {
		if ctr.State.Terminated != nil {
			if ctr.State.Terminated.Message != "" {
				return ctr.State.Terminated.Message
			}
			if ctr.State.Terminated.Reason == "OOMKilled" {
				return ctr.State.Terminated.Reason
			}
			if ctr.State.Terminated.ExitCode != 0 {
				return fmt.Sprintf("container %q failed with exit code %d", ctr.Name, ctr.State.Terminated.ExitCode)
			}
		}
		return ""
	}

	switch pod.Status.Phase {
	case apiv1.PodPending, apiv1.PodRunning:
		return v1alpha1.OperationRunning, ""
	case apiv1.PodSucceeded:
		return v1alpha1.OperationSucceeded, ""
	case apiv1.PodFailed:
		if pod.Status.Message != "" {
			// Pod has a nice error message. Use that.
			return v1alpha1.OperationFailed, pod.Status.Message
		}
		for _, ctr := range append(pod.Status.InitContainerStatuses, pod.Status.ContainerStatuses...) {
			if msg := getFailMessage(&ctr); msg != "" {
				return v1alpha1.OperationFailed, msg
			}
		}
		return v1alpha1.OperationFailed, ""
	case apiv1.PodUnknown:
		return v1alpha1.OperationError, ""
	}
	return v1alpha1.OperationRunning, ""
}
