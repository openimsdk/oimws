package common

import (
	"runtime/debug"

	"github.com/bwmarrin/snowflake"
	log "github.com/xuexihuang/new_log15"
)

// Method used to capture panic and print stack information.
func TryRecoverAndDebugPrint() {
	errs := recover()
	if errs == nil {
		return
	}
	log.Crit("[Panic]", "err", errs, "stackInfo", debug.Stack())

}

var G_flakeNode snowflake.Node

func GetRandomSessionId() string {

	return G_flakeNode.Generate().String()
}
