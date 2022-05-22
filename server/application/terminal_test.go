package application

import (
	"testing"

	"github.com/argoproj/gitops-engine/pkg/utils/kube"

	appv1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func TestPodExists(t *testing.T) {
	for _, tcase := range []struct {
		name           string
		podName        string
		namespace      string
		treeNodes      []appv1.ResourceNode
		expectedResult bool
	}{
		{
			name:           "empty tree nodes",
			podName:        "test-pod",
			namespace:      "test",
			treeNodes:      []appv1.ResourceNode{},
			expectedResult: false,
		},
		{
			name:           "matched Pod but empty UID",
			podName:        "test-pod",
			namespace:      "test",
			treeNodes:      []appv1.ResourceNode{{ResourceRef: appv1.ResourceRef{Name: "test-pod", Namespace: "test", UID: "", Kind: kube.PodKind}}},
			expectedResult: false,
		},
		{
			name:           "matched Pod",
			podName:        "test-pod",
			namespace:      "test",
			treeNodes:      []appv1.ResourceNode{{ResourceRef: appv1.ResourceRef{Name: "test-pod", Namespace: "test", UID: "testUID", Kind: kube.PodKind}}},
			expectedResult: true,
		},
		{
			name:           "unmatched Pod Namespace",
			podName:        "test-pod",
			namespace:      "test",
			treeNodes:      []appv1.ResourceNode{{ResourceRef: appv1.ResourceRef{Name: "test-pod", Namespace: "test-A", UID: "testUID", Kind: kube.PodKind}}},
			expectedResult: false,
		},
		{
			name:           "unmatched Kind",
			podName:        "test-pod",
			namespace:      "test",
			treeNodes:      []appv1.ResourceNode{{ResourceRef: appv1.ResourceRef{Name: "test-pod", Namespace: "test-A", UID: "testUID", Kind: kube.DeploymentKind}}},
			expectedResult: false,
		},
		{
			name:           "unmatched Group",
			podName:        "test-pod",
			namespace:      "test",
			treeNodes:      []appv1.ResourceNode{{ResourceRef: appv1.ResourceRef{Name: "test-pod", Namespace: "test", UID: "testUID", Group: "A", Kind: kube.PodKind}}},
			expectedResult: false,
		},
		{
			name:           "unmatched Pod Name",
			podName:        "test-pod",
			namespace:      "test",
			treeNodes:      []appv1.ResourceNode{{ResourceRef: appv1.ResourceRef{Name: "test", Namespace: "test", UID: "testUID", Kind: kube.PodKind}}},
			expectedResult: false,
		},
	} {
		t.Run(tcase.name, func(t *testing.T) {
			result := podExists(tcase.treeNodes, tcase.podName, tcase.namespace)
			if result != tcase.expectedResult {
				t.Errorf("Expected result %v, but got %v", tcase.expectedResult, result)
			}
		})
	}
}

func TestIsValidPodName(t *testing.T) {
	for _, tcase := range []struct {
		name           string
		resourceName   string
		expectedResult bool
	}{
		{
			name:           "valid pod name",
			resourceName:   "argocd-server-794644486d-r8v9d",
			expectedResult: true,
		},
		{
			name:           "not valid contains spaces",
			resourceName:   "kubectl delete pods",
			expectedResult: false,
		},
		{
			name:           "not valid",
			resourceName:   "kubectl -n kube-system delete pods --all",
			expectedResult: false,
		},
		{
			name:           "not valid contains special characters",
			resourceName:   "delete+*+from+etcd%3b",
			expectedResult: false,
		},
	} {
		t.Run(tcase.name, func(t *testing.T) {
			result := isValidPodName(tcase.resourceName)
			if result != tcase.expectedResult {
				t.Errorf("Expected result %v, but got %v", tcase.expectedResult, result)
			}
		})
	}
}

func TestIsValidNamespaceName(t *testing.T) {
	for _, tcase := range []struct {
		name           string
		resourceName   string
		expectedResult bool
	}{
		{
			name:           "valid pod namespace name",
			resourceName:   "argocd",
			expectedResult: true,
		},
		{
			name:           "not valid contains spaces",
			resourceName:   "kubectl delete ns argocd",
			expectedResult: false,
		},
		{
			name:           "not valid contains special characters",
			resourceName:   "delete+*+from+etcd%3b",
			expectedResult: false,
		},
	} {
		t.Run(tcase.name, func(t *testing.T) {
			result := isValidNamespaceName(tcase.resourceName)
			if result != tcase.expectedResult {
				t.Errorf("Expected result %v, but got %v", tcase.expectedResult, result)
			}
		})
	}
}

func TestIsValidContainerNameName(t *testing.T) {
	for _, tcase := range []struct {
		name           string
		resourceName   string
		expectedResult bool
	}{
		{
			name:           "valid container name",
			resourceName:   "argocd-server",
			expectedResult: true,
		},
		{
			name:           "not valid contains spaces",
			resourceName:   "kubectl delete pods",
			expectedResult: false,
		},
		{
			name:           "not valid contains special characters",
			resourceName:   "delete+*+from+etcd%3b",
			expectedResult: false,
		},
	} {
		t.Run(tcase.name, func(t *testing.T) {
			result := isValidContainerName(tcase.resourceName)
			if result != tcase.expectedResult {
				t.Errorf("Expected result %v, but got %v", tcase.expectedResult, result)
			}
		})
	}
}
