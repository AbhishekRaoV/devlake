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

package bridge

import (
	"context"

	"github.com/apache/incubator-devlake/core/config"
	ctx "github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/log"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/impls/logruslog"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var DefaultContext = NewRemoteContext(logruslog.Global, config.GetConfig())

type RemoteProgress struct {
	Current   int `json:"current"`
	Total     int `json:"total"`
	Increment int `json:"increment"`
}

type RemoteContext interface {
	plugin.ExecContext
	GetSettings() map[string]any
}

type remoteContextImpl struct {
	parent   plugin.ExecContext
	logger   log.Logger
	ctx      context.Context
	Settings map[string]any `json:"settings"`
}

func (r remoteContextImpl) GetSettings() map[string]any {
	return r.Settings
}

func (r remoteContextImpl) GetConfigReader() config.ConfigReader {
	return r.parent.GetConfigReader()
}

func (r remoteContextImpl) ReplaceLogger(logger log.Logger) ctx.BasicRes {
	return &remoteContextImpl{
		parent:   r.parent,
		logger:   logger,
		ctx:      r.ctx,
		Settings: r.Settings,
	}
}

func (r remoteContextImpl) NestedLogger(name string) ctx.BasicRes {
	return r.ReplaceLogger(r.logger.Nested(name))
}

func NewRemoteContext(logger log.Logger, cfg *viper.Viper) RemoteContext {
	return &remoteContextImpl{
		logger:   logger,
		Settings: cfg.AllSettings(),
		ctx:      context.Background(),
	}
}

func NewChildRemoteContext(ec plugin.ExecContext) RemoteContext {
	return &remoteContextImpl{
		parent:   ec,
		logger:   ec.GetLogger(),
		ctx:      ec.GetContext(),
		Settings: DefaultContext.GetSettings(),
	}
}

func (r remoteContextImpl) GetConfig(name string) string {
	val, ok := r.Settings[name]
	if !ok {
		return ""
	}
	return cast.ToString(val)
}

func (r remoteContextImpl) GetLogger() log.Logger {
	return r.logger
}

func (r remoteContextImpl) GetDal() dal.Dal {
	if r.parent != nil {
		return r.parent.GetDal()
	}
	return nil
}

func (r remoteContextImpl) GetName() string {
	if r.parent != nil {
		return r.parent.GetName()
	}
	return "default_remote"
}

func (r remoteContextImpl) GetContext() context.Context {
	return r.ctx
}

func (r remoteContextImpl) GetData() interface{} {
	if r.parent != nil {
		return r.parent.GetData()
	}
	return nil
}

func (r remoteContextImpl) SetProgress(current int, total int) {
	if r.parent != nil {
		r.parent.SetProgress(current, total)
	}
}

func (r remoteContextImpl) IncProgress(quantity int) {
	if r.parent != nil {
		r.parent.IncProgress(quantity)
	}
}

var _ RemoteContext = (*remoteContextImpl)(nil)
