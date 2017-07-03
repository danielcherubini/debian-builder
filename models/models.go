package models

//Config struct is the object that comes from the yaml file
type Config struct {
	Containers []Containers `json:"containers"`
	Network    string       `json:"network"`
	Control    Control      `json:"control"`
}

//Containers ldfdf
type Containers struct {
	Name  string   `json:"name"`
	Image string   `json:"image"`
	Ports []Port   `json:"ports"`
	Links []string `json:"links"`
	Env   []Env    `json:"env"`
}

//Control is the control file
type Control struct {
	Package      string `json:"package"`
	Version      string `json:"version"`
	Section      string `json:"section"`
	Priority     string `json:"priority"`
	Architecture string `json:"architecture"`
	Maintainer   string `json:"maintainer"`
	Description  string `json:"description"`
	Distribution string `json:"distribution"`
}

//Port blah blah
type Port struct {
	HostPort      int    `json:"hostPort"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

//Env is the config object for environment
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
