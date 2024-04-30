package main

import (
	"fmt"
	"github.com/openim-sigs/oimws/pkg/common/cmd"
	"os"
	"path/filepath"
)

func main() {
	if err := cmd.NewSdkCmd().Exec(); err != nil {
		ExitWithError(err)
	}
}

// TODO
func ExitWithError(err error) {
	progName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "%s exit -1: %+v\n", progName, err)
	os.Exit(-1)
}
