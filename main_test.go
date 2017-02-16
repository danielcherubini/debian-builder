package main

import "testing"

func TestLoadConfigToModel(t *testing.T) {
	config, _, err := loadConfigToModel("test.yaml")
	if err != nil {
		t.Fail()
	}
	if config.Control.Package != "test-yaml" {
		t.Fail()
	}
}

func TestBuildDebianPackage(t *testing.T) {
	config, filename, err := loadConfigToModel("test.yaml")
	if err != nil {
		t.Fail()
	}
	cmd := buildDebianPackage(config.Control, filename)
	if cmd == nil {
		t.Fail()
	}
}

func TestRunEverything(t *testing.T) {
	config, filepath := setupEnvironment()
	if config.Package == "" {
		t.Fail()
	}
	if filepath == "" {
		t.Fail()
	}
}
