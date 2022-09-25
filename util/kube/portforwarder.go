package kube

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"k8s.io/kubectl/pkg/util/podutils"

	"github.com/argoproj/argo-cd/v2/util/io"
)

func PortForward(targetPort int, namespace string, overrides *clientcmd.ConfigOverrides, podSelectors ...string) (int, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	clientConfig := clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, overrides, os.Stdin)
	config, err := clientConfig.ClientConfig()
	if err != nil {
		return -1, err
	}

	if namespace == "" {
		namespace, _, err = clientConfig.Namespace()
		if err != nil {
			return -1, err
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return -1, err
	}

	var pod *corev1.Pod

	for _, podSelector := range podSelectors {
		pods, err := clientSet.CoreV1().Pods(namespace).List(context.Background(), v1.ListOptions{
			LabelSelector: podSelector,
		})
		if err != nil {
			return -1, err
		}

		for _, po := range pods.Items {
			if po.Status.Phase == corev1.PodRunning && podutils.IsPodReady(&po) {
				pod = &po
				break
			}
		}
	}

	if pod == nil {
		return -1, fmt.Errorf("cannot find ready pod with selector: %v", podSelectors)
	}

	url := clientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(pod.Namespace).
		Name(pod.Name).
		SubResource("portforward").URL()

	transport, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return -1, errors.Wrap(err, "Could not create round tripper")
	}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", url)

	readyChan := make(chan struct{}, 1)
	failedChan := make(chan error, 1)
	out := new(bytes.Buffer)
	errOut := new(bytes.Buffer)

	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return -1, err
	}
	port := ln.Addr().(*net.TCPAddr).Port
	io.Close(ln)

	forwarder, err := portforward.New(dialer, []string{fmt.Sprintf("%d:%d", port, targetPort)}, context.Background().Done(), readyChan, out, errOut)
	if err != nil {
		return -1, err
	}

	go func() {
		err = forwarder.ForwardPorts()
		if err != nil {
			failedChan <- err
		}
	}()
	select {
	case err = <-failedChan:
		return -1, err
	case <-readyChan:
	}
	if len(errOut.String()) != 0 {
		return -1, fmt.Errorf(errOut.String())
	}
	return port, nil
}
