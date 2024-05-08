// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"github.com/openim-sigs/oimws/internal/sdk"
	"github.com/openim-sigs/oimws/pkg/common/config"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type SdkCmd struct {
	*RootCmd
	ctx       context.Context
	configMap map[string]any
	sdkConfig sdk.Config
}

func NewSdkCmd() *SdkCmd {
	var ret SdkCmd
	ret.configMap = map[string]any{
		FileName: &ret.sdkConfig.SdkConfig,
	}
	// TODO program.GetProcessName()
	var processName string
	if args := os.Args; len(args) > 0 {
		segments := strings.Split(args[0], "/")
		processName = segments[len(segments)-1]
	}
	//TODO
	ret.RootCmd = NewRootCmd(processName, WithConfigMap(ret.configMap))
	//ret.RootCmd = NewRootCmd(program.GetProcessName(), WithConfigMap(ret.configMap))
	ret.ctx = context.WithValue(context.Background(), "version", config.Version)
	ret.Command.RunE = func(cmd *cobra.Command, args []string) error {
		return ret.runE()
	}
	return &ret
}

func (a *SdkCmd) Exec() error {
	return a.Execute()
}

func (a *SdkCmd) runE() error {
	return sdk.Start(a.ctx, a.Index(), &a.sdkConfig, &a.RootCmd.log)
}
