/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package archived

import (
	"encoding/json"
	"time"
)

type Task struct {
	Model
	Plugin        string          `gorm:"index"`
	Options       json.RawMessage `gorm:"type:json"`
	Status        string
	Message       string
	Progress      float32
	FailedSubTask string
	PipelineId    uint64 `gorm:"index"`
	PipelineRow   int
	PipelineCol   int
	BeganAt       *time.Time
	FinishedAt    *time.Time `gorm:"index"`
	SpentSeconds  int
}

func (Task) TableName() string {
	return "_devlake_tasks"
}
