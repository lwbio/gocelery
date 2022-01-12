// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package main

import (
	"github.com/lwbio/gocelery"
)

// Run Celery Worker First!
// celery -A worker worker --loglevel=debug --without-heartbeat --without-mingle
func main() {
	var err error

	// redisPool := &redis.Pool{
	// 	MaxIdle:     3,                 // maximum number of idle connections in the pool
	// 	MaxActive:   0,                 // maximum number of connections allocated by the pool at a given time
	// 	IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
	// 	Dial: func() (redis.Conn, error) {
	// 		c, err := redis.DialURL("redis://:temp12138@127.0.0.1:6379/8")
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return c, err
	// 	},
	// 	TestOnBorrow: func(c redis.Conn, t time.Time) error {
	// 		_, err := c.Do("PING")
	// 		return err
	// 	},
	// }

	// initialize celery client

	celery, err := gocelery.NewCeleryClientUri("redis://:temp12138@127.0.0.1:6379/8", nil, 1)
	if err != nil {
		panic(err)
	}

	celery.Lock()
	defer celery.Unlock()

	// run task
	celery.SetQueueName("tasks.qiye")
	taskName := "qiye.tasks.notify_attraction_activity_create"
	storeId := 199
	activityId := 40
	_, err = celery.Delay(taskName, storeId, activityId)
	if err != nil {
		panic(err)
	}

}
