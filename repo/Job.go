package repo

type Job struct {
	Type string		`redis:"type"`
	State string	`redis:"state"`
}

func NewJob(JobType string) *Job {
	j := Job{ Type : JobType, State : "start" }

	return &j
}
