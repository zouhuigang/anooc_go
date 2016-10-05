package jobqueue

type HandleFunc func(interface{})

type Job struct {
	Input   interface{}
	Handler HandleFunc
}

type JobQueue chan Job
