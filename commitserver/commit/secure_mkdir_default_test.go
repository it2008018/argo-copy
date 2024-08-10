//go:build !linux

package commit

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecureMkdirAllDefault(t *testing.T) {
	root := t.TempDir()

	unsafePath := "test/dir"
	fullPath, err := SecureMkdirAll(root, unsafePath, os.ModePerm)
	require.NoError(t, err)

	expectedPath := path.Join(root, unsafePath)
	assert.Equal(t, expectedPath, fullPath)
}

func TestSecureMkdirAllWithExistingDir(t *testing.T) {
	root := t.TempDir()
	unsafePath := "existing/dir"

	fullPath, err := SecureMkdirAll(root, unsafePath, os.ModePerm)
	require.NoError(t, err)

	newPath, err := SecureMkdirAll(root, unsafePath, os.ModePerm)
	require.NoError(t, err)
	assert.Equal(t, fullPath, newPath)
}
