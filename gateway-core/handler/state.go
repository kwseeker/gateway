package handler

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"net/http"
	"time"
)

// ISession golang不支持抽象类（golang推崇“组合优于继承”，面向对象的很多特性都是通过组合实现），
// 比如这里实现抽象类需要用 interface struct 组合的方式实现
//type ISession interface {
//	Session(ctx netty.InboundContext, message netty.Message)
//}
//

type StateHandler struct{}

func (*StateHandler) HandleActive(ctx netty.ActiveContext) {
	fmt.Printf("http client active: %s\n", ctx.Channel().RemoteAddr())
	ctx.HandleActive()
}

func (*StateHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	if request, ok := message.(*http.Request); ok {
		fmt.Printf("[%d]%s: %s %s\n", ctx.Channel().ID(), ctx.Channel().RemoteAddr(), request.Method, request.URL.Path)
	}
	ctx.HandleRead(message)
}

func (*StateHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	if responseWriter, ok := message.(http.ResponseWriter); ok {
		// set response header.
		responseWriter.Header().Add("x-time", time.Now().String())
	}
	ctx.HandleWrite(message)
}

//func (*StateHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
//	fmt.Printf("http client except: %s %v\n", ctx.Channel().RemoteAddr(), ex)
//}

func (*StateHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	fmt.Printf("http client inactive: %s %v\n", ctx.Channel().RemoteAddr(), ex)
	ctx.HandleInactive(ex)
}
