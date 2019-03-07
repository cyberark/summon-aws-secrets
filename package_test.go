package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

	Convey("Given a compiled summon-aws-secrets package", t, func() {
		Convey("Given no configuration information", func() {
			e := ClearEnv()
			defer e.RestoreEnv()
			os.Setenv("PATH", Path)

			Convey("Given summon-aws-secrets is run with no arguments", func() {
				_, stderr, err := RunCommand(PackageName)

				Convey("Returns with error", func() {
					So(err, ShouldNotBeNil)
					So(stderr.String(), ShouldStartWith, "A variable ID or version flag must be given as the first and only argument!")
				})
			})

			Convey("Given summon-aws-secrets is retrieving a secret with key value format", func() {
				secret := `{ "username": "USERNAME", "password": "PASSWORD", "port": 8000 }`

				Convey("Returns with the value of the key", func() {
					stdout, err := getValueByKey("username", []byte(secret))
					So(err, ShouldBeNil)
					So(string(stdout), ShouldEqual, "USERNAME")
				})

				Convey("Always returns as a string", func() {
					stdout, err := getValueByKey("port", []byte(secret))
					So(err, ShouldBeNil)
					So(string(stdout), ShouldEqual, "8000")
				})
			})
		})
	})
}
