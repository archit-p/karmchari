package jobstore

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type JobStore struct {
	// pool of redis connections for use by threads
	ConnPool *redis.Pool
	// copy of ids of all jobs stored in redis instance
	JobIds []string
}

func New(host string) (*JobStore, error) {
	var jh JobStore

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
