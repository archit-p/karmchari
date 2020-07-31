package repo

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/archit-p/karmchari/enc"
)

type JobHandler struct {
	// pool of redis connections for use by threads
	ConnPool *redis.Pool
	// copy of ids of all jobs stored in redis instance
	JobIds []string
}

func NewJobHandler(host string) (*JobHandler, error) {
	var jh JobHandler

	// initialize the connection pool
	jh.ConnPool = &redis.Pool {
		MaxIdle: 10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}

	// initialize the job ids
	// get a new connection from pool
	conn := jh.ConnPool.Get()

	// defer closing the connection for clean exit
	defer conn.Close()

	// iterator for keys from redis
	iter := 0
	for {
		// get the key fields starting at the iterator
		if resp, err := redis.MultiBulk(conn.Do("SCAN", iter, "MATCH", "job*",
			"COUNT", 10)); err != nil {
			return nil, err
		} else {
			// get the next location for iterator
			iter, err = redis.Int(resp[0], nil)
			if err != nil {
				return nil, err
			}

			// get the keys
			keys, err := redis.Strings(resp[1], nil)
			if err != nil {
				return nil, err
			}

			// append keys to local copy
			jh.JobIds = append(jh.JobIds, keys...)
		}

		if iter == 0 {
			break
		}
	}

	return &jh, nil
}

func (jh *JobHandler) WriteJob(Type string, State string) (*string, error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	defer conn.Close()

	id := enc.GetHashString()

	_, err := conn.Do("HMSET", "job:" + *id, "type", Type, "state", State)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (jh *JobHandler) ReadJob(id string) (*Job, error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", "job:" + id))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, errors.New("No job found")
	}

	var job Job
	err = redis.ScanStruct(values, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (jh *JobHandler) DeleteJob(id string) (error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	defer conn.Close()

	_, err := conn.Do("DEL", "job:" + id)
	if err != nil {
		return err
	}

	return nil
}

func (jh *JobHandler) UpdateJobState(id string, newState string) (error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	defer conn.Close()

	_, err := conn.Do("HSET", "job:" + id, "state", newState)
	if err != nil {
		return err
	}

	return nil
}

func (jh *JobHandler) UpdateJobType(id string, newType string) (error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	defer conn.Close()

	_, err := conn.Do("HSET", "job:" + id, "type", newType)
	if err != nil {
		return err
	}

	return nil
}
