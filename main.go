package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/danmademe/debian-builder/models"
	"github.com/ghodss/yaml"
)

var (
	configFile string
	postInst   string
)

func check(err error, what string) {
	if err != nil {
		fmt.Println("debian-builder | " + what + ": " + err.Error())
		return
	}
}

func getDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err, "Error getting current directory")
	return dir
}

func loadConfigToModel(file string) (models.Config, string, error) {
	config := models.Config{}
	dat, err := ioutil.ReadFile(file)
	yaml.Unmarshal([]byte(dat), &config)

	location, _ := filepath.Abs(file)
	return config, location, err
}

func buildPostInst() error {
	//Setup default postinst
	postInst = "#!/bin/sh \nsudo docker-service provision"

	//Get current directory
	filepath := getDirectory()

	postinstByte := []byte(postInst)
	postinstLocation := filepath + "/.postinst"
	err := ioutil.WriteFile(postinstLocation, postinstByte, 0777)
	return err
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
	afterInstall := getDirectory() + "/.postinst"
	args = append(args, "--after-install", afterInstall)

	//Disable Warning for no config
	noConfigWarningDisable := "--deb-no-default-config-files"
	args = append(args, noConfigWarningDisable)

	//Setup Config
	config := fileName + "=/etc/docker-service/services.d/" + control.Package + "_" + control.Version + ".yaml"
	args = append(args, config)

	//Exec fpm
	fmt.Printf("debian-builder | %s", args)
	cmd := exec.Command("fpm", args...)
	return cmd
}
func setupEnvironment() (models.Control, string) {
	//Read flags
	flag.StringVar(&configFile, "config", "test.yaml", "a string var")
	flag.Parse()

	//Load config
	config, filepath, err := loadConfigToModel(configFile)
	check(err, "Error loading config")

	//Write file
	postinstErr := buildPostInst()
	check(postinstErr, "Error installing postinst File")

	return config.Control, filepath
}

func main() {
	//Setup Environment
	control, filepath := setupEnvironment()

	//Build Debian Package
	output, err := buildDebianPackage(control, filepath).Output()
	check(err, "Error Building Package")

	fmt.Println("debian-builder | Building Package")
	fmt.Printf("debian-builder | %s\n", output)
}
