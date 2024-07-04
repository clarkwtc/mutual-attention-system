package exceptions

import (
    "github.com/gin-gonic/gin"
)

type IErrorHandler interface {
    Match(error error) bool
    SetNext(next IErrorHandler)
    Response(ctx *gin.Context)
    Handle(ctx *gin.Context, err error)
}

type ErrorHandler struct {
    Concrete IErrorHandler
}

func (handler *ErrorHandler) Handle(ctx *gin.Context, error error) {
    if handler.Concrete.Match(error) {
        handler.Concrete.Response(ctx)
    } else if handler.Concrete != nil {
        handler.Concrete.Handle(ctx, error)
    }
}

func (handler *ErrorHandler) SetNext(next IErrorHandler) {
    handler.Concrete = next
}
