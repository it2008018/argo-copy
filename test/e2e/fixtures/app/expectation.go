package app

import (
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"

	. "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
)

type state = string

const (
	failed    = "failed"
	pending   = "pending"
	succeeded = "succeeded"
)

type Expectation func(c *Consequences) (state state, message string)

func Not(e Expectation) Expectation {
	return func(c *Consequences) (state, string) {
		state, message := e(c)
		message = "not " + message
		switch state {
		case succeeded:
			return failed, message
		case failed:
			return succeeded, message
		default:
			return state, message
		}
	}
}

func OperationPhaseIs(expected OperationPhase) Expectation {
	return func(c *Consequences) (state, string) {
		actual := c.app().Status.OperationState.Phase
		return simple(actual == expected, fmt.Sprintf("expect app %s's operation phase to be %s, is %s", c.context.name, expected, actual))
	}
}

func simple(success bool, message string) (state, string) {
	if success {
		return succeeded, message
	} else {
		return pending, message
	}
}

func SyncStatusIs(expected SyncStatusCode) Expectation {
	return func(c *Consequences) (state, string) {
		actual := c.app().Status.Sync.Status
		return simple(actual == expected, fmt.Sprintf("expect app %s's sync status to be %s, is %s", c.context.name, expected, actual))
	}
}

func Condition(conditionType ApplicationConditionType) Expectation {
	return func(c *Consequences) (state, string) {
		message := fmt.Sprintf("condition of type %s", conditionType)
		for _, condition := range c.app().Status.Conditions {
			if conditionType == condition.Type {
				return succeeded, message
			}
		}
		return failed, message
	}
}

func HealthIs(expected HealthStatusCode) Expectation {
	return func(c *Consequences) (state, string) {
		actual := c.app().Status.Health.Status
		return simple(actual == expected, fmt.Sprintf("health to be %s, is %s", expected, actual))
	}
}

func ResourceSyncStatusIs(resource string, expected SyncStatusCode) Expectation {
	return func(c *Consequences) (state, string) {
		actual := c.resource(resource).Status
		return simple(actual == expected, fmt.Sprintf("resource '%s' sync status to be %s, is %s", resource, expected, actual))
	}
}

func ResourceHealthIs(resource string, expected HealthStatusCode) Expectation {
	return func(c *Consequences) (state, string) {
		actual := c.resource(resource).Health.Status
		return simple(actual == expected, fmt.Sprintf("resource '%s' health to be %s, is %s", resource, expected, actual))
	}
}

func DoesNotExist() Expectation {
	return func(c *Consequences) (state, string) {
		_, err := c.get()
		if err != nil {
			if apierrors.IsNotFound(err) {
				return succeeded, "app does not exist"
			}
			return failed, err.Error()
		}
		return pending, "app does not exist"
	}
}

func Pod(predicate func(p v1.Pod) bool) Expectation {
	return func(c *Consequences) (state, string) {
		c.context.fixture.KubeClientset.CoreV1()
		pods, err := c.context.fixture.KubeClientset.CoreV1().Pods(c.context.fixture.DeploymentNamespace).List(metav1.ListOptions{})
		if err != nil {
			return failed, err.Error()
		}
		for _, pod := range pods.Items {
			if predicate(pod) {
				return succeeded, fmt.Sprintf("pod predicate matched pod named '%s'", pod.GetName())
			}
		}
		return pending, fmt.Sprintf("pod predicate did not match pods: %v", pods.Items)
	}
}

func Event(reason string, message string) Expectation {
	return func(c *Consequences) (state, string) {
		list, err := c.context.fixture.KubeClientset.CoreV1().Events(c.context.fixture.ArgoCDNamespace).List(metav1.ListOptions{
			FieldSelector: fields.SelectorFromSet(map[string]string{
				"involvedObject.name":      c.context.name,
				"involvedObject.namespace": c.context.fixture.ArgoCDNamespace,
			}).String(),
		})
		if err != nil {
			return failed, err.Error()
		}

		for i := range list.Items {
			event := list.Items[i]
			if event.Reason == reason && strings.Contains(event.Message, message) {
				return succeeded, fmt.Sprintf("found event with reason=%s; message=%s", reason, message)
			}
		}
		return failed, fmt.Sprintf("unable to find event with reason=%s; message=%s", reason, message)
	}
}
