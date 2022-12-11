package server

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/xhttp"
	"net/http"
	"testing"
	"time"
)

func TestServer_StartDefault(t *testing.T) {
	server := Server{}
	server.StartDefault()
}

func Test_HTTP(t *testing.T) {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello, go gateway!"))
		if err != nil {
			return
		}
	})

	// channel pipeline initializer.
	setupCodec := func(channel netty.Channel) {
		channel.Pipeline().
			// decode http request from channel
			AddLast(xhttp.ServerCodec()).
			// print http access log
			AddLast(new(httpStateHandler)).
			// compatible with http.Handler
			AddLast(xhttp.Handler(httpMux))
	}

	// setup bootstrap & startup server.
	netty.NewBootstrap(netty.WithChildInitializer(setupCodec)).
		Listen("0.0.0.0:7380").Sync()
}

type httpStateHandler struct{}

func (*httpStateHandler) HandleActive(ctx netty.ActiveContext) {
	fmt.Printf("http client active: %s\n", ctx.Channel().RemoteAddr())
	ctx.HandleActive()
}

func (*httpStateHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	if request, ok := message.(*http.Request); ok {
		fmt.Printf("[%d]%s: %s %s\n", ctx.Channel().ID(), ctx.Channel().RemoteAddr(), request.Method, request.URL.Path)
	}
	ctx.HandleRead(message)
}

func (*httpStateHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	if responseWriter, ok := message.(http.ResponseWriter); ok {
		// set response header.
		responseWriter.Header().Add("x-time", time.Now().String())
	}
	ctx.HandleWrite(message)
}

//func (*httpStateHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
//	fmt.Printf("http client except: %s %v\n", ctx.Channel().RemoteAddr(), ex)
//}

func (*httpStateHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	fmt.Printf("http client inactive: %s %v\n", ctx.Channel().RemoteAddr(), ex)
	ctx.HandleInactive(ex)
}
