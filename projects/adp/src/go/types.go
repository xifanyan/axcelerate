package adp

import "fmt"

type TaskResponse struct {
	ExecutionID         string  `json:"executionId"`
	TaskType            string  `json:"taskType"`
	LoggingEnabled      string  `json:"loggingEnabled"`
	ProgressMax         int     `json:"progressMax"`
	ExecutionStatus     string  `json:"executionStatus"`
	ExecutionRootDir    string  `json:"executionRootDir"`
	ContextID           string  `json:"contextId"`
	ExecutionPersistent string  `json:"executionPersistent"`
	ProgressCurrent     int     `json:"progressCurrent"`
	ProgressPercentage  float64 `json:"progressPercentage"`
	TaskDisplayName     string  `json:"taskDisplayName"`
	ExecutionMetaData   any     `json:"executionMetaData"`
	ErrorMessage        string  `json:"errorMessage,omitempty"`
}

type TaskExecutionError struct {
	ExecutionID       string
	TaskType          string
	ExecutionStatus   string
	ErrorMessage      string
	ExecutionMetaData any
}

func (e *TaskExecutionError) Error() string {
	return fmt.Sprintf("task %s failed: %s (executionId=%s)", e.TaskType, e.ErrorMessage, e.ExecutionID)
}
