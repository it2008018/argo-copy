package commit

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
	"os/exec"
	"path"
	"sigs.k8s.io/yaml"
	"time"
)

/**
The commit package provides a way for the controller to push manifests to git.
*/

type Service interface {
	Commit(ManifestsRequest) (ManifestsResponse, error)
}

type ManifestsRequest struct {
	RepoURL       string
	TargetBranch  string
	DrySHA        string
	CommitAuthor  string
	CommitMessage string
	CommitTime    time.Time
	Paths         []PathDetails
}

type PathDetails struct {
	Path      string
	Manifests []ManifestDetails
	ReadmeDetails
}

type ManifestDetails struct {
	Manifest unstructured.Unstructured
}

type ReadmeDetails struct {
}

type ManifestsResponse struct {
	RequestId string
}

func NewService() Service {
	return &service{}
}

type service struct {
}

func (s *service) Commit(r ManifestsRequest) (ManifestsResponse, error) {
	logCtx := log.WithFields(log.Fields{"repo": r.RepoURL, "branch": r.TargetBranch, "drySHA": r.DrySHA})
	logCtx.Info("committing")

	// Create a temp dir with a UUID
	dirName, err := uuid.NewRandom()
	err = os.MkdirAll(path.Join("/tmp/_commit-service", dirName.String()), os.ModePerm)
	if err != nil {
		return ManifestsResponse{}, fmt.Errorf("failed to create temp dir: %w", err)
	}

	// Clone the repo into the temp dir using the git CLI
	err = exec.Command("git", "clone", r.RepoURL, dirName.String()).Run()
	if err != nil {
		return ManifestsResponse{}, fmt.Errorf("failed to clone repo: %w", err)
	}

	// Write the manifests to the temp dir
	for _, p := range r.Paths {
		err = os.MkdirAll(path.Join(dirName.String(), p.Path), os.ModePerm)
		if err != nil {
			return ManifestsResponse{}, fmt.Errorf("failed to create path: %w", err)
		}
		for _, m := range p.Manifests {
			// Marshal the manifests
			mYaml, err := yaml.Marshal(m.Manifest)
			if err != nil {
				return ManifestsResponse{}, fmt.Errorf("failed to marshal manifest: %w", err)
			}
			// Write the yaml to manifest.yaml
			err = os.WriteFile(path.Join(dirName.String(), p.Path, "manifest.yaml"), mYaml, os.ModePerm)
			if err != nil {
				return ManifestsResponse{}, fmt.Errorf("failed to write manifest: %w", err)
			}
		}
	}

	// Commit the changes
	err = exec.Command("git", "add", ".").Run()
	if err != nil {
		return ManifestsResponse{}, fmt.Errorf("failed to add files: %w", err)
	}

	commitCmd := exec.Command("git", "commit", "-m", r.CommitMessage, "--author", r.CommitAuthor)
	out, err := commitCmd.CombinedOutput()
	if err != nil {
		log.WithError(err).WithField("output", string(out)).Error("failed to commit files")
		return ManifestsResponse{}, fmt.Errorf("failed to commit: %w", err)
	}

	err = exec.Command("git", "push", "origin", r.TargetBranch).Run()
	if err != nil {
		return ManifestsResponse{}, fmt.Errorf("failed to push: %w", err)
	}

	return ManifestsResponse{}, nil
}
