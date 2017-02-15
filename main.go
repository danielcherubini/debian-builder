package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/danmademe/debian-builder/models"
	"github.com/ghodss/yaml"
)

var configFile string

func loadConfigToModel(file string) (models.Config, string, error) {
	config := models.Config{}
	dat, err := ioutil.ReadFile(file)
	yaml.Unmarshal([]byte(dat), &config)

	location, _ := filepath.Abs(file)
	return config, location, err
}

func buildDebianPackage(control models.Control, fileName string) *exec.Cmd {
	args := []string{"-s", "dir", "-t", "deb"}

	//Setup Name
	name := control.Package + "-" + control.Distribution
	args = append(args, "-n", name)

	//Setup version
	version := control.Version
	args = append(args, "-v", version)

	//Setup After Install
	afterInstall := ".postinst"
	args = append(args, "--after-install", afterInstall)

	//Setup Config
	config := fileName + "=/etc/docker-service/services.d/" + control.Package + "_" + control.Version + ".yaml"
	args = append(args, config)

	//Exec fpm
	cmd := exec.Command("fpm", args...)
	return cmd
}

func main() {
	flag.StringVar(&configFile, "config", "test.yaml", "a string var")
	flag.Parse()

	config, filepath, err := loadConfigToModel(configFile)
	if err != nil {
		fmt.Println("Error loading config: " + err.Error())
		return
	}
	output, err := buildDebianPackage(config.Control, filepath).Output()
	if err != nil {
		fmt.Println("Error Building Package: " + err.Error())
		return
	}
	fmt.Println("Building Package")
	fmt.Printf("%s", output)
}
