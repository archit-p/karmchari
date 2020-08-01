package jobstore

// update state for a job in store
func (jh *JobStore) UpdateState(id string, newState string) (error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	// defer closing the connection for a clean exit
	defer conn.Close()

	// send update command to redis
	_, err := conn.Do("HSET", "job:" + id, "state", newState)
	if err != nil {
		return err
	}

	return nil
}

// update type for a job in store
func (jh *JobStore) UpdateType(id string, newType string) (error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	// defer closing the connection for a clean exit
	defer conn.Close()

	// send update command to redis
	_, err := conn.Do("HSET", "job:" + id, "type", newType)
	if err != nil {
		return err
	}

	return nil
}
