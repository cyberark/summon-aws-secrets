package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RunCommand(name string, arg ...string) (bytes.Buffer, bytes.Buffer, error) {
	cmd := exec.Command(name, arg...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout, stderr, err
}

const PackageName = "summon-aws-secrets"

func TestPackage(t *testing.T) {

	Path := os.Getenv("PATH")

	t.Run("Compiled summon-aws-secrets package without params", func(t *testing.T) {
		e := ClearEnv()
		defer e.RestoreEnv()
		os.Setenv("PATH", Path)

		_, stderr, err := RunCommand(PackageName)

		assert.Error(t, err)
		assert.Contains(t, stderr.String(), "A variable ID or version flag must be given as the first and only argument!")
	})
}

func Test_getValueByKey(t *testing.T) {
	t.Run("Valid JSON format stored secret", func(t *testing.T) {
		secret := `{ "username": "USERNAME", "password": "PASSWORD", "port": 8000 }`

		stdout, err := getValueByKey("username", []byte(secret))
		assert.NoError(t, err)
		assert.Equal(t, "USERNAME", string(stdout))

		stdout, err = getValueByKey("port", []byte(secret))
		assert.NoError(t, err)
		assert.Equal(t, "8000", string(stdout))
	})
}
