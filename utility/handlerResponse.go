package utility

import (
	"gf-demo-api/internal/model"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	ghttp "github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

type DefaultResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.BufferLength() > 0 && r.Response.Status == http.StatusOK {
		return
	}
	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	defaultResponse := DefaultResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    res,
	}
	if err != nil {
		if code == gcode.CodeNil || code == gcode.CodeInternalError {
			defaultResponse.Code = model.CodeInternalError.Code()
			defaultResponse.Message = model.CodeInternalError.Desc()
		}
		r.Response.ClearBuffer()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		switch r.Response.Status {
		case http.StatusNotFound:
			{
				defaultResponse.Code = model.CodeNotFound.Code()
				defaultResponse.Message = model.CodeNotFound.Desc()
			}
		case http.StatusForbidden:
			{
				defaultResponse.Code = model.UnAuthorized.Code()
				defaultResponse.Message = model.UnAuthorized.Desc()
			}
		default:
			{
				defaultResponse.Code = model.CodeUnknown.Code()
				defaultResponse.Message = model.CodeUnknown.Desc()
			}
		}
	} else {
		defaultResponse.Code = model.CodeOK.Code()
		defaultResponse.Message = model.CodeOK.Desc()
	}
	r.Response.WriteJson(defaultResponse)
}
