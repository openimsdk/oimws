package core_func

import (
	"fmt"
	log "github.com/xuexihuang/new_log15"

	"github.com/OpenIMSDK/tools/errs"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
)

type RespMessage struct {
	respMessagesChan chan *EventData
}

// NewRespMessage creates a new instance of RespMessage with the given channel.
func NewRespMessage(respMessagesChan chan *EventData) *RespMessage {
	return &RespMessage{respMessagesChan: respMessagesChan}
}

// sendOnSuccessResp sends a success response message for a given operation.
func (r *RespMessage) sendOnSuccessResp(operationID, event string, data string) {
	r.respMessagesChan <- &EventData{
		Event:       event,
		OperationID: operationID,
		Data:        data,
	}
}

// sendOnErrorResp sends an error response message for a given operation.
func (r *RespMessage) sendOnErrorResp(operationID, event string, err error) {
	log.Error("sendOnErrorResp", "operationID", operationID, "event", event, "err", err)
	resp := &EventData{
		Event:       event,
		OperationID: operationID,
	}
	if code, ok := err.(errs.CodeError); ok {
		resp.ErrCode = int32(code.Code())
		resp.ErrMsg = code.Error()
	} else {
		resp.ErrCode = sdkerrs.UnknownCode
		resp.ErrMsg = fmt.Sprintf("error %T not implement CodeError: %s", err, err)
	}
	r.respMessagesChan <- resp
}

// sendEventFailedRespNoErr sends a failed event response without error details.
// event: Name of the event.
func (r *RespMessage) sendEventFailedRespNoErr(event string) {
	r.respMessagesChan <- &EventData{
		Event: event,
	}
}

// sendEventSuccessRespWithData sends a successful event response with associated data.
func (r *RespMessage) sendEventSuccessRespWithData(event string, data string) {
	r.respMessagesChan <- &EventData{
		Event: event,
		Data:  data,
	}
}

// sendEventSuccessRespNoData sends a successful event response without any associated data.
// This is included for completeness but not used in the above callback methods.
func (r *RespMessage) sendEventSuccessRespNoData(event string) {
	r.respMessagesChan <- &EventData{
		Event: event,
	}
}

// sendEventFailedRespNoData sends a failed event response with error code and message, without any associated data.
// This function may be used if there are any future error handling requirements in the SignalingCallback.
func (r *RespMessage) sendEventFailedRespNoData(event string, errCode int32, errMsg string) {
	r.respMessagesChan <- &EventData{
		Event:   event,
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}
