package main

import (
	"strings"
	"os"
)

func splitEq(s string) (string, string) {
	a := strings.SplitN(s, "=", 2)
	return a[0], a[1]
}

type envSnapshot struct {
	env []string
}

func ClearEnv() *envSnapshot {
	e := os.Environ()

	for _, s := range e {
		k, _ := splitEq(s)
		os.Setenv(k, "")
	}
	return &envSnapshot{env: e}
}

func (e *envSnapshot) RestoreEnv() {
	ClearEnv()
	for _, s := range e.env {
		k, v := splitEq(s)
		os.Setenv(k, v)
	}
}

