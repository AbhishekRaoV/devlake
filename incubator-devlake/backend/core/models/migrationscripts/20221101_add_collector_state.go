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

package migrationscripts

import (
	"encoding/json"

	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	commonArchived "github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
)

type createCollectorState struct{}

type CollectorState20221101 struct {
	commonArchived.GenericModel[string]
	Type  string
	Value json.RawMessage `gorm:"type:json"`
}

func (CollectorState20221101) TableName() string {
	return "_devlake_collector_state"
}

func (*createCollectorState) Up(basicRes context.BasicRes) errors.Error {
	return basicRes.GetDal().AutoMigrate(CollectorState20221101{})
}

func (*createCollectorState) Version() uint64 {
	return 20221101000001
}

func (*createCollectorState) Name() string {
	return "Create collector state table"
}
