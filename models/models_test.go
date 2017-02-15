package models_test

import (
	"encoding/json"
	"testing"

	"github.com/danmademe/docker-service/models"
)

func TestConfig(t *testing.T) {
	config := models.Config{}
	input := []byte(`{"containers":[{"env":[{"name":"ENV","value":"test"}],"image":"docker.tidal.com/tidal-test-yaml:1.0.0","name":"tidal-test-yaml","ports":[{"containerPort":9000,"hostPort":9000,"protocol":"tcp"}]}],"network":"host"}`)
	err := json.Unmarshal(input, &config)
	if err != nil {
		return
	}
	container := config.Containers[0]
	testEnv := container.Env
	testPorts := container.Ports
	testImage := container.Image

	if testImage != "docker.tidal.com/tidal-test-yaml:1.0.0" {
		t.Errorf("Config Model Image Error", testImage)
	}

	if testPorts[0].HostPort != 9000 || testPorts[0].ContainerPort != 9000 {
		t.Errorf("Config Model Ports Error", testPorts)
	}

	exampleEnvs := "ENV"
	if exampleEnvs != testEnv[0].Name {
		t.Errorf("Config Model Env Error", testEnv)
	}
}
