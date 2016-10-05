// Copyright 2014 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

package main

import (
	"db"

	"github.com/robfig/cron"
)

// 后台运行的任务
func ServeBackGround() {

	if db.MasterDB == nil {
		return
	}

	c := cron.New()

	c.Start()
}
