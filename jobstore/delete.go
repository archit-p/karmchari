package jobstore

// delete a job in the store
func (jh *JobStore) Delete(id string) (error) {
	// create a connection with redis instance
	conn := jh.ConnPool.Get()

	// defer closing the connection for clean exit
	defer conn.Close()

	// send command to redis instance
	_, err := conn.Do("DEL", "job:" + id)
	if err != nil {
		return err
	}

	return nil
}
