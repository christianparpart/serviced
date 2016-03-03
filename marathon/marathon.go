package marathon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PortMapping struct {
	ContainerPort int
	HostPort      int
	ServicePort   int
	Protocol      string
}

type DockerContainer struct {
	Image          string
	Network        string
	PortMappings   []PortMapping
	Privileged     bool
	Parameters     [][]string
	ForcePullImage bool
}

type HealthCheck struct {
	Protocol               string
	Path                   string
	PortIndex              uint
	GracePeriodSeconds     uint
	IntervalSeconds        uint
	TimeoutSeconds         uint
	MaxConsecutiveFailures uint
	IgnoreHttp1xx          bool
}

type ContainerVolume struct {
	ContainerPath string
	HostPath      string
	Mode          string
}

type AppContainer struct {
	Type    string
	Volumes []ContainerVolume
	Docker  DockerContainer
}

type UpgradeStrategy struct {
	MinimumHealthCapacity float64
	MaximumOverCapacity   float64
}

type App struct {
	Id                    string
	Cmd                   *string
	Args                  *string
	User                  *string
	Env                   map[string]string
	Instances             int
	Cpus                  int
	Mem                   int
	Disk                  int
	Executor              string
	Constraints           [][]string
	Uris                  []string
	Fetch                 []string
	StoreUrls             []string
	Ports                 []int
	RequirePorts          bool
	BackoffSeconds        uint
	BackoffFactor         float64
	MaxLaunchDelaySeconds uint
	Container             AppContainer
	HealthChecks          []HealthCheck
	// TODO Dependencies []
	UpgradeStrategy UpgradeStrategy
	Labels          map[string]string
}

type jsonResponse struct {
	App *App
}

func GetApp(host string, port int, appPath string) (*App, error) {
	var url = fmt.Sprintf("http://%v:%v/v2/apps%s", host, port, appPath)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)

	var v jsonResponse
	err = json.Unmarshal(output, &v)
	if err != nil {
		return nil, err
	}

	return v.App, nil
}
