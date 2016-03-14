package marathon

import (
	"fmt"
	"strings"
	"time"
)

type PortMapping struct {
	ContainerPort int
	HostPort      int
	ServicePort   int
	Protocol      string
}

type KeyValuePair struct {
	Key   string
	Value *string
}

type DockerContainer struct {
	Image          string
	Network        string
	PortMappings   []PortMapping
	Privileged     bool
	Parameters     []KeyValuePair
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

type HealthCheckResult struct {
	Alive               bool
	ConsecutiveFailures uint
	FirstSuccess        *time.Time
	LastFailure         *time.Time
	LastSuccess         *time.Time
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

type Task struct {
	Id                 string
	Host               string
	Ports              []int
	StartedAt          time.Time
	StagedAt           time.Time
	Version            time.Time
	SlaveId            string
	AppId              string
	HealthCheckResults []HealthCheckResult
}

type FetchInfo struct {
	Uri        string
	Extract    bool
	Executable bool
	Cache      bool
}

type App struct {
	service               *Service
	Id                    string
	Cmd                   *string
	Args                  *string
	User                  *string
	Env                   map[string]string
	Instances             int
	Cpus                  float64
	Mem                   int
	Disk                  int
	Executor              string
	Constraints           [][]string
	Uris                  []string
	Fetch                 []FetchInfo
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
	Tasks           []Task
}

func (app *App) Scale(instance_count uint) error {
	// TODO
	_, err := app.service.HttpPost(
		"/v2/apps/TODO/scale",
		strings.NewReader(fmt.Sprintf("%v", instance_count)))

	return err
}
