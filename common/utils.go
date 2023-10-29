package common

import (
	"github.com/bwmarrin/snowflake"
	log "github.com/xuexihuang/new_log15"
	"runtime/debug"
)

// 用来捕获panic并且打印堆栈信息的方法
func TryRecoverAndDebugPrint() {
	errs := recover()
	if errs == nil {
		return
	}
	log.Crit("[Panic]", "err", errs, "stackInfo", string(debug.Stack()))

}

var G_flakeNode *snowflake.Node

func GetRandomSessionId() string {

	return G_flakeNode.Generate().String()
}
