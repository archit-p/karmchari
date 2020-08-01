package jobstore

import (
	"errors"

	"github.com/archit-p/karmchari/job"
	
	"github.com/gomodule/redigo/redis"
)

// read a job from store
func (jh *JobStore) Read(id string) (*job.Job, error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	// defer closing connection for clean exit
	defer conn.Close()

	// send command to redis
	values, err := redis.Values(conn.Do("HGETALL", "job:" + id))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, errors.New("No job found")
	}

	// scan the values into a job struct
	var j job.Job
	err = redis.ScanStruct(values, &j)
	if err != nil {
		return nil, err
	}

	return &j, nil
}
