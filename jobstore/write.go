package jobstore

import (
	"github.com/archit-p/karmchari/enc"
)

// write a new job to the store
func (jh *JobStore) Write(Type string, State string) (*string, error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	// defer closing the connection for a clean exit
	defer conn.Close()

	// generate a new id for the job
	id := enc.GetHashString()

	// send a hash set command to redis
	_, err := conn.Do("HMSET", "job:" + *id, "type", Type, "state", State)
	if err != nil {
		return nil, err
	}

	return id, nil
}
