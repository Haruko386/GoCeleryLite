package task

type Status string

const (
	StatusQueued   = "queued"
	StatusRunning  = "running"
	StatusSuccess  = "success"
	StatusFailed   = "failure"
	StatusError    = "error"
	StatusCanceled = "canceled"
	StatusTimeout  = "timeout"
	StatusRetrying = "Retrying"
)
