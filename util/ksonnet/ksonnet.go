package ksonnet

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/ksonnet/ksonnet/metadata"
	"github.com/ksonnet/ksonnet/metadata/app"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	diffSeparator = regexp.MustCompile(`\n---`)
	lineSeparator = regexp.MustCompile(`\n`)
)

// KsonnetApp represents a ksonnet application directory and provides wrapper functionality around
// the `ks` command.
type KsonnetApp interface {
	// Root is the root path ksonnet application directory
	Root() string

	// App is the Ksonnet application
	App() app.App

	// Show returns a list of unstructured objects that would be applied to an environment
	Show(environment string) ([]*unstructured.Unstructured, error)
	ListEnvParams(environment string) (*unstructured.Unstructured, error)
}

type ksonnetApp struct {
	manager metadata.Manager
	app     app.App
}

// NewKsonnetApp tries to create a new wrapper to run commands on the `ks` command-line tool.
func NewKsonnetApp(path string) (KsonnetApp, error) {
	ksApp := ksonnetApp{}
	mgr, err := metadata.Find(path)
	if err != nil {
		return nil, err
	}
	ksApp.manager = mgr
	app, err := ksApp.manager.App()
	if err != nil {
		return nil, err
	}
	ksApp.app = app
	return &ksApp, nil
}

func (k *ksonnetApp) ksCmd(args ...string) (string, error) {
	cmd := exec.Command("ks", args...)
	cmd.Dir = k.Root()

	cmdStr := strings.Join(cmd.Args, " ")
	log.Debug(cmdStr)
	out, err := cmd.Output()
	if err != nil {
		exErr, ok := err.(*exec.ExitError)
		if !ok {
			return "", err
		}
		errOutput := string(exErr.Stderr)
		log.Errorf("`%s` failed: %s", cmdStr, errOutput)
		return "", fmt.Errorf(strings.TrimSpace(errOutput))
	}
	return string(out), nil
}

func (k *ksonnetApp) Root() string {
	return k.manager.Root()
}

// Spec is the Ksonnet application spec (app.yaml)
func (k *ksonnetApp) App() app.App {
	return k.app
}

// Show generates a concatenated list of Kubernetes manifests in the given environment.
func (k *ksonnetApp) Show(environment string) ([]*unstructured.Unstructured, error) {
	out, err := k.ksCmd("show", environment)
	if err != nil {
		return nil, err
	}
	parts := diffSeparator.Split(out, -1)
	objs := make([]*unstructured.Unstructured, 0)
	for _, part := range parts {
		if strings.TrimSpace(part) == "" {
			continue
		}
		var obj unstructured.Unstructured
		err = yaml.Unmarshal([]byte(part), &obj)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal manifest from `ks show`")
		}
		objs = append(objs, &obj)
	}
	// TODO(jessesuen): we need to sort objects based on their dependency order of creation
	return objs, nil
}

// Show generates a concatenated list of Kubernetes manifests in the given environment.
func (k *ksonnetApp) ListEnvParams(environment string) (*unstructured.Unstructured, error) {
	out, err := k.ksCmd("param", "list", "--env", environment)
	if err != nil {
		return nil, err
	}
	rows := lineSeparator.Split(out, -1)[2:]
	obj := new(unstructured.Unstructured)
	obj.Object = make(map[string](interface{}))
	for _, row := range rows[2:] {
		if strings.TrimSpace(row) == "" {
			continue
		}
		fields := strings.Fields(row)
		param, valueStr := fields[1], fields[2]
		var value interface{}
		// try 64-bit int, 64-bit float, bool, then default to string
		if i, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
			value = (interface{})(i)
		} else if f, err := strconv.ParseFloat(valueStr, 64); err == nil {
			value = interface{}(f)
		} else if b, err := strconv.ParseBool(valueStr); err == nil {
			value = interface{}(b)
		} else if v, err := strconv.Unquote(valueStr); err == nil {
			value = interface{}(v)
		}
		obj.Object[param] = value
	}
	return obj, nil
}
