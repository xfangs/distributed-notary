package api

import (
	"fmt"
	"time"

	"net/http"

	"os"

	"github.com/SmartMeshFoundation/distributed-notary/utils"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/nkbai/log"
)

// defaultAPITimeout : 默认api请求超时时间
var defaultAPITimeout = 30 * time.Second

/*
BaseAPI : 提供一些公共方法
*/
type BaseAPI struct {
	serverName  string
	host        string
	router      rest.App
	middleWares []rest.Middleware
	api         *rest.Api
	timeout     time.Duration // 调用service层的超时时间
	requestChan chan Request
}

// NewBaseAPI :
func NewBaseAPI(serverName string, host string, router rest.App, middleWares ...rest.Middleware) BaseAPI {
	return BaseAPI{
		serverName:  serverName,
		host:        host,
		router:      router,
		timeout:     defaultAPITimeout,
		middleWares: middleWares,
		requestChan: make(chan Request, 10),
	}
}

// Start 启动监听线程
func (ba *BaseAPI) Start(sync bool) {
	ba.api = rest.NewApi()
	ba.api.Use(rest.DefaultCommonStack...)
	if len(ba.middleWares) > 0 {
		ba.api.Use(ba.middleWares...)
	}
	ba.api.SetApp(ba.router)
	log.Info("%s listen at %s", ba.serverName, ba.host)
	if sync {
		err := http.ListenAndServe(ba.host, ba.api.MakeHandler())
		if err != nil {
			log.Error("http server start err : %s", err.Error())
			os.Exit(-1)
		}
	} else {
		go func() {
			err := http.ListenAndServe(ba.host, ba.api.MakeHandler())
			if err != nil {
				log.Error("http server start err : %s", err.Error())
				os.Exit(-1)
			}
		}()
	}
}

// GetRequestChan :
func (ba *BaseAPI) GetRequestChan() <-chan Request {
	return ba.requestChan
}

// SetTimeout :
func (ba *BaseAPI) SetTimeout(timeout time.Duration) {
	ba.timeout = timeout
}

// SendToServiceAndWaitResponse :
func (ba *BaseAPI) SendToServiceAndWaitResponse(req Request, timeout ...time.Duration) *BaseResponse {
	if r, ok := req.(NotaryRequest); ok {
		if !VerifyNotarySignature(r) {
			return NewFailResponse(req.GetRequestID(), ErrorCodePermissionDenied)
		}
	}
	var resp *BaseResponse
	requestTimeout := ba.timeout
	if len(timeout) > 0 && timeout[0] > 0 {
		requestTimeout = timeout[0]
	}
	ba.requestChan <- req
	if requestTimeout > 0 {
		select {
		case resp = <-req.GetResponseChan():
		case <-time.After(requestTimeout):
			resp = NewFailResponse(req.GetRequestID(), ErrorCodeTimeout)
		}
	} else {
		resp = <-req.GetResponseChan()
	}
	apiLog(req, resp)
	return resp
}

/*
tool functions
*/

// Return :
func Return(w rest.ResponseWriter, response *BaseResponse) {
	if w == nil {
		return
	}
	err := w.WriteJson(response)
	if err != nil {
		log.Warn(fmt.Sprintf("writejson err %s", err))
	}
}

func apiLog(req Request, resp *BaseResponse) {
	prefix := ""
	body := utils.ToJSONStringFormat(req)
	switch r := req.(type) {
	case CrossChainRequest:
		prefix = fmt.Sprintf("==> API [RequestID=%s Name=%s SCToken=%s]", req.GetRequestID(), req.GetRequestName(), utils.APex(r.GetSCTokenAddress()))
	case NotaryRequest:
		type requestToLog struct {
			BaseRequest
			BaseNotaryRequest
			BaseCrossChainRequest
		}
		var l requestToLog
		l.BaseRequest.RequestID = req.GetRequestID()
		l.BaseRequest.Name = req.GetRequestName()
		l.BaseNotaryRequest.SessionID = r.GetSessionID()
		l.BaseNotaryRequest.Sender = r.GetSender()
		l.BaseNotaryRequest.Signature = r.getSignature()
		body = utils.ToJSONStringFormat(l)
		prefix = fmt.Sprintf("[SessionID=%s] ==> API [RequestID=%s Name=%s SenderID=%d]", utils.HPex(r.GetSessionID()), req.GetRequestID(), req.GetRequestName(), r.GetSenderID())
	default:
		prefix = fmt.Sprintf("==> API [RequestID=%s Name=%s]", req.GetRequestID(), req.GetRequestName())
	}
	if resp.GetErrorCode() == ErrorCodeSuccess {
		log.Trace(fmt.Sprintf("%s deal SUCCESS", prefix))
	} else {
		log.Error(fmt.Sprintf("%s deal FAIL: \nRequest :\n%s\nResponse :\n%s", prefix, body, utils.ToJSONStringFormat(resp)))
	}
}
