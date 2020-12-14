// Copyright 2020 beego-dev
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

func main() {
	// create a task
	tk1 := task.NewTask("tk1", "0/3 * * * * *", func(ctx context.Context) error {
		logs.Info("tk1")
		return nil
	})

	// check task
	err := tk1.Run(context.Background())
	if err != nil {
		logs.Error(err)
	}

	// add task to global todolist
	task.AddTask("tk1", tk1)

	// start tasks
	task.StartTask()

	// wait 12 second
	time.Sleep(12 * time.Second)
	defer task.StopTask()
}
