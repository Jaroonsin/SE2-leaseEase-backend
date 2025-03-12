package main

import (
	"testing"

	"LeaseEase/test/steps"

	"github.com/cucumber/godog"
)

func TestMain(t *testing.T) {
	opts := godog.Options{
		Format:   "pretty",                  // Output style
		Paths:    []string{"./../features"}, // Path to feature files
		Strict:   true,                      // Fail if there are undefined steps
		TestingT: t,                         // Use Go's testing package
	}

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: steps.InitializeHandler,
		Options:             &opts,
	}.Run()

	if status != 0 {
		t.Fail()
	}
}
