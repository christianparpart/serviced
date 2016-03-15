package marathon

type TaskStatus string

const (
	TaskStaging  = TaskStatus("TASK_STAGING")
	TaskStarting = TaskStatus("TASK_STARTING")
	TaskRunning  = TaskStatus("TASK_RUNNING")
	TaskFinished = TaskStatus("TASK_FINISHED")
	TaskFailed   = TaskStatus("TASK_FAILED")
	TaskKilled   = TaskStatus("TASK_KILLED")
	TaskLost     = TaskStatus("TASK_LOST")
)

type StatusUpdateEvent struct {
	SlaveId     string
	TaskId      string
	TaskStatus  TaskStatus
	Message     string
	AppId       string
	Host        string
	IpAddresses []string
	Ports       []int
	Version     string
}

type HealthStatusChangedEvent struct {
	AppId  string
	TaskId string
	Alive  bool
}
