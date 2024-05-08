package common

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	log "github.com/xuexihuang/new_log15"
	"runtime/debug"
)

// Method used to capture panic and print stack information.
func TryRecoverAndDebugPrint() {
	errs := recover()
	if errs == nil {
		return
	}
	fmt.Printf("panic: %+v\n%s", errs, debug.Stack())
	log.Crit("[Panic]", "err", errs, "stackInfo", debug.Stack())

}

var G_flakeNode snowflake.Node

func GetRandomSessionId() string {

	return G_flakeNode.Generate().String()
}
