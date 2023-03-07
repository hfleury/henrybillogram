package main

import (
	"os"
	"testing"

	"github.com/hfleury/henrybillogram/database/postgres/support"
)

type TestAppMain struct {
}

func TestMain(m *testing.M) {
	os.Setenv("BILLO_ENVIRONMENT", "TEST")
	support.TestMain(m)
	m.Run()
}

// TODO: test server is up and running
