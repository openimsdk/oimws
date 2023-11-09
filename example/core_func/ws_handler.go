package core_func

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/ccontext"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
)

const (
	Success = "OnSuccess"
	Failed  = "OnError"
)

type EventData struct {
	Event       string `json:"event"`
	ErrCode     int32  `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
	Data        string `json:"data"`
	OperationID string `json:"operationID"`
}

type FuncRouter struct {
	userForSDK  *open_im_sdk.LoginMgr
	respMessage *RespMessage
	sessionId   string
}

// NewFuncRouter creates a new instance of FuncRouter with the provided session id and channel for event data.
func NewFuncRouter(respMessagesChan chan *EventData, sessionId string) *FuncRouter {
	return &FuncRouter{respMessage: NewRespMessage(respMessagesChan),
		userForSDK: new(open_im_sdk.LoginMgr), sessionId: sessionId}
}

// call is an asynchronous wrapper that invokes SDK functions and handles their responses.
func (f *FuncRouter) call(operationID string, fn any, args ...any) {
	go func() {
		funcPtr := reflect.ValueOf(fn).Pointer()
		funcName := runtime.FuncForPC(funcPtr).Name()
		parts := strings.Split(funcName, ".")
		var trimFuncName string
		if trimFuncNameList := strings.Split(parts[len(parts)-1], "-"); len(trimFuncNameList) == 0 {
			f.respMessage.sendOnErrorResp(operationID, "FuncError", errors.New("call function trimFuncNameList is empty"))
			return
		} else {
			trimFuncName = trimFuncNameList[0]
		}
		res, err := f.call_(operationID, fn, funcName, args...)
		if err != nil {
			f.respMessage.sendOnErrorResp(operationID, trimFuncName, err)
			return
		}
		data, err := json.Marshal(res)
		if err != nil {
			f.respMessage.sendOnErrorResp(operationID, trimFuncName, err)
			return
		} else {
			f.respMessage.sendOnSuccessResp(operationID, trimFuncName, string(data))
		}
	}()
}

// CheckResourceLoad checks the SDK is resource load status.
func CheckResourceLoad(uSDK *open_im_sdk.LoginMgr, funcName string) error {
	if uSDK == nil {
		return utils.Wrap(errors.New("CheckResourceLoad failed uSDK == nil "), "")
	}
	if funcName == "" {
		return nil
	}
	parts := strings.Split(funcName, ".")
	if parts[len(parts)-1] == "Login-fm" {
		return nil
	}
	if uSDK.Friend() == nil || uSDK.User() == nil || uSDK.Group() == nil || uSDK.Conversation() == nil ||
		uSDK.Full() == nil {
		return utils.Wrap(errors.New("CheckResourceLoad failed, resource nil "), "")
	}
	return nil
}

// call_ is the internal function that actually invokes the SDK functions.
func (f *FuncRouter) call_(operationID string, fn any, funcName string, args ...any) (res any, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %+v\n%s", r, debug.Stack())
			err = fmt.Errorf("call panic: %+v", r)
		}
	}()
	if operationID == "" {
		return nil, sdkerrs.ErrArgs.Wrap("call function operationID is empty")
	}
	if err := CheckResourceLoad(f.userForSDK, funcName); err != nil {
		return nil, sdkerrs.ErrResourceLoad.Wrap("not load resource")
	}

	ctx := ccontext.WithOperationID(f.userForSDK.BaseCtx(), operationID)

	fnv := reflect.ValueOf(fn)
	if fnv.Kind() != reflect.Func {
		return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("call function fn is not function, is %T", fn))
	}
	fnt := fnv.Type()
	nin := fnt.NumIn()
	if len(args)+1 != nin {
		return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("go code error: fn in args num is not match"))
	}
	t := time.Now()
	log.ZInfo(ctx, "input req", "function name", funcName, "args", args)
	ins := make([]reflect.Value, 0, nin)
	ins = append(ins, reflect.ValueOf(ctx))
	for i := 0; i < len(args); i++ {
		inFnField := fnt.In(i + 1)
		arg := reflect.TypeOf(args[i])
		if arg.String() == inFnField.String() || inFnField.Kind() == reflect.Interface {
			ins = append(ins, reflect.ValueOf(args[i]))
			continue
		}
		//convert float64 to int when javascript call with number,because javascript only have double
		//precision floating-point format
		if arg.String() == "float64" && isInteger(inFnField) {
			ins = append(ins, reflect.ValueOf(convert(args[i].(float64), inFnField)))
			continue
		}
		if arg.Kind() == reflect.String { // json
			var ptr int
			for inFnField.Kind() == reflect.Ptr {
				inFnField = inFnField.Elem()
				ptr++
			}
			switch inFnField.Kind() {
			case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map:
				v := reflect.New(inFnField)
				if err := json.Unmarshal([]byte(args[i].(string)), v.Interface()); err != nil {
					return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("go call json.Unmarshal error: %s",
						err))
				}
				if ptr == 0 {
					v = v.Elem()
				} else if ptr != 1 {
					for i := ptr - 1; i > 0; i-- {
						temp := reflect.New(v.Type())
						temp.Elem().Set(v)
						v = temp
					}
				}
				ins = append(ins, v)
				continue
			}
		}
		return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("go code error: fn in args type is not match"))
	}
	outs := fnv.Call(ins)
	if len(outs) == 0 {
		return "", nil
	}
	if fnt.Out(len(outs) - 1).Implements(reflect.ValueOf(new(error)).Elem().Type()) {
		if errValueOf := outs[len(outs)-1]; !errValueOf.IsNil() {
			log.ZError(ctx, "fn call error", errValueOf.Interface().(error), "function name",
				funcName, "cost time", time.Since(t))
			return nil, errValueOf.Interface().(error)
		}
		if len(outs) == 1 {
			return "", nil
		}
		outs = outs[:len(outs)-1]
	}
	for i := 0; i < len(outs); i++ {
		out := outs[i]
		switch out.Kind() {
		case reflect.Map:
			if out.IsNil() {
				outs[i] = reflect.MakeMap(out.Type())
			}
		case reflect.Slice:
			if out.IsNil() {
				outs[i] = reflect.MakeSlice(out.Type(), 0, 0)
			}
		}
	}
	if len(outs) == 1 {
		log.ZInfo(ctx, "output resp", "function name", funcName, "resp", outs[0].Interface(),
			"cost time", time.Since(t))
		return outs[0].Interface(), nil
	}
	val := make([]any, 0, len(outs))
	for i := range outs {
		val = append(val, outs[i].Interface())
	}
	log.ZInfo(ctx, "output resp", "function name", funcName, "resp", val, "cost time", time.Since(t))
	return val, nil
}
func isInteger(arg reflect.Type) bool {
	switch arg.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return true
	default:
		return false

	}
}
func convert(arg float64, p reflect.Type) any {
	switch p.Kind() {
	case reflect.Int:
		return int(arg)
	case reflect.Int8:
		return int8(arg)
	case reflect.Int16:
		return int16(arg)
	case reflect.Int32:
		return int32(arg)
	case reflect.Int64:
		return int64(arg)
	default:
		return arg

	}
}
func (f *FuncRouter) messageCall(operationID string, fn any, args ...any) {
	go func() {
		funcPtr := reflect.ValueOf(fn).Pointer()
		funcName := runtime.FuncForPC(funcPtr).Name()
		parts := strings.Split(funcName, ".")
		var trimFuncName string
		if trimFuncNameList := strings.Split(parts[len(parts)-1], "-"); len(trimFuncNameList) == 0 {
			f.respMessage.sendOnErrorResp(operationID, "FuncError",
				errors.New("call function trimFuncNameList is empty"))
			return
		} else {
			trimFuncName = trimFuncNameList[0]
		}
		sendMessageCallback := NewSendMessageCallback(trimFuncName, f.respMessage)
		res, err := f.messageCall_(sendMessageCallback, operationID, fn, funcName, args...)
		if err != nil {
			f.respMessage.sendOnErrorResp(operationID, trimFuncName, err)
			return
		}
		data, err := json.Marshal(res)
		if err != nil {
			f.respMessage.sendOnErrorResp(operationID, trimFuncName, err)
			return
		} else {
			f.respMessage.sendOnSuccessResp(operationID, trimFuncName, string(data))
		}
	}()
}
func (f *FuncRouter) messageCall_(callback open_im_sdk_callback.SendMsgCallBack, operationID string,
	fn any, funcName string, args ...any) (res any, err error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %+v\n%s", r, debug.Stack())
			err = fmt.Errorf("call panic: %+v", r)
		}
	}()
	if operationID == "" {
		return nil, sdkerrs.ErrArgs.Wrap("call function operationID is empty")
	}
	if err := CheckResourceLoad(f.userForSDK, ""); err != nil {
		return nil, sdkerrs.ErrResourceLoad.Wrap("not load resource")
	}

	ctx := ccontext.WithOperationID(f.userForSDK.BaseCtx(), operationID)
	ctx = ccontext.WithSendMessageCallback(ctx, callback)

	fnv := reflect.ValueOf(fn)
	if fnv.Kind() != reflect.Func {
		return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("call function fn is not function, is %T", fn))
	}
	log.ZInfo(ctx, "input req", "function name", funcName, "args", args)
	fnt := fnv.Type()
	nin := fnt.NumIn()
	if len(args)+1 != nin {
		return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("go code error: fn in args num is not match"))
	}
	t := time.Now()
	ins := make([]reflect.Value, 0, nin)
	ins = append(ins, reflect.ValueOf(ctx))
	for i := 0; i < len(args); i++ {
		inFnField := fnt.In(i + 1)
		arg := reflect.TypeOf(args[i])
		if arg.String() == inFnField.String() || inFnField.Kind() == reflect.Interface {
			ins = append(ins, reflect.ValueOf(args[i]))
			continue
		}
		if arg.String() == "float64" && isInteger(inFnField) {
			ins = append(ins, reflect.ValueOf(convert(args[i].(float64), inFnField)))
			continue
		}
		if arg.Kind() == reflect.String { // json
			var ptr int
			for inFnField.Kind() == reflect.Ptr {
				inFnField = inFnField.Elem()
				ptr++
			}
			switch inFnField.Kind() {
			case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map:
				v := reflect.New(inFnField)
				if err := json.Unmarshal([]byte(args[i].(string)), v.Interface()); err != nil {
					return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("go call json.Unmarshal error: %s",
						err))
				}
				if ptr == 0 {
					v = v.Elem()
				} else if ptr != 1 {
					for i := ptr - 1; i > 0; i-- {
						temp := reflect.New(v.Type())
						temp.Elem().Set(v)
						v = temp
					}
				}
				ins = append(ins, v)
				continue
			}
		}
		return nil, sdkerrs.ErrSdkInternal.Wrap(fmt.Sprintf("go code error: fn in args type is not match"))
	}
	outs := fnv.Call(ins)
	if len(outs) == 0 {
		return "", nil
	}
	if fnt.Out(len(outs) - 1).Implements(reflect.ValueOf(new(error)).Elem().Type()) {
		if errValueOf := outs[len(outs)-1]; !errValueOf.IsNil() {
			log.ZError(ctx, "fn call error", errValueOf.Interface().(error), "function name",
				funcName, "cost time", time.Since(t))
			return nil, errValueOf.Interface().(error)
		}
		if len(outs) == 1 {
			return "", nil
		}
		outs = outs[:len(outs)-1]
	}
	for i := 0; i < len(outs); i++ {
		out := outs[i]
		switch out.Kind() {
		case reflect.Map:
			if out.IsNil() {
				outs[i] = reflect.MakeMap(out.Type())
			}
		case reflect.Slice:
			if out.IsNil() {
				outs[i] = reflect.MakeSlice(out.Type(), 0, 0)
			}
		}
	}
	if len(outs) == 1 {
		log.ZInfo(ctx, "output resp", "function name", funcName, "resp", outs[0].Interface(),
			"cost time", time.Since(t))
		return outs[0].Interface(), nil
	}
	val := make([]any, 0, len(outs))
	for i := range outs {
		val = append(val, outs[i].Interface())
	}
	log.ZInfo(ctx, "output resp", "function name", funcName, "resp", val, "cost time", time.Since(t))
	return val, nil
}
