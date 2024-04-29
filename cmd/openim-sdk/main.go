package main

import (
	"github.com/openim-sigs/oimws/pkg/common/cmd"
	"github.com/openimsdk/tools/system/program"
)

func main() {
	if err := cmd.NewSdkCmd().Exec(); err != nil {
		program.ExitWithError(err)
	}
}
