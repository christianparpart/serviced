package marathon

import "time"

type TaskStatus string

const (
	// TaskStaging means, this task is in staging (such as 'docker pull'-state)
	TaskStaging  = TaskStatus("TASK_STAGING")
	TaskStarting = TaskStatus("TASK_STARTING")
	TaskRunning  = TaskStatus("TASK_RUNNING")
	TaskFinished = TaskStatus("TASK_FINISHED")
	TaskFailed   = TaskStatus("TASK_FAILED")
	TaskKilling  = TaskStatus("TASK_KILLING")
	TaskKilled   = TaskStatus("TASK_KILLED")
	TaskLost     = TaskStatus("TASK_LOST")
)

type IpAddr struct {
	IpAddress string
	Protocol  string
}

/*
{
	"slaveId":"9e1a18f2-011c-44fe-9715-be1cac1d5f41-S8",
	"taskId":"production_lovemag_app.03eb79d1-058d-11e6-a243-72491c981fcc",
	"taskStatus":"TASK_RUNNING",
	"message":"",
	"appId":"/production/lovemag/app",
	"host":"rack2-compute5.dawanda.services",
	"ipAddresses": [
		{
			"ipAddress":"172.17.0.7",
			"protocol":"IPv4"
		}
	],
	"ports": [47755],
	"version":"2016-04-14T12:52:12.465Z",
	"eventType":"status_update_event",
	"timestamp":"2016-04-18T17:43:10.580Z"
}
*/
type StatusUpdateEvent struct {
	EventType   string
	Timestamp   time.Time
	SlaveId     string
	TaskId      string
	TaskStatus  TaskStatus
	Message     string
	AppId       string
	Host        string
	IpAddresses []IpAddr
	Ports       []int
	Version     string
}

type HealthStatusChangedEvent struct {
	AppId  string
	TaskId string
	Alive  bool
}
