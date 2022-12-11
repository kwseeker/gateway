package handler

import (
	"fmt"
	"net/http"
)

type Session struct {
	serveMux *http.ServeMux
}

func NewSession() *Session {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", Handle)
	return &Session{serveMux: httpMux}
}

func (h *Session) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.serveMux.ServeHTTP(w, r)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("网关接收到请求 URI: %v\n", r.URL)

	//TODO net/http 格式、KEEP_ALIVE、跨域等控制

	_, err := w.Write([]byte("Hello, go gateway!"))
	if err != nil {
		return
	}
}

//		  logger.info("网关接收请求 uri：{} method：{}", request.uri(), request.method());
//
//        // 返回信息处理
//        DefaultFullHttpResponse response = new DefaultFullHttpResponse(HttpVersion.HTTP_1_1, HttpResponseStatus.OK);
//        // 返回信息控制
//        response.content().writeBytes(JSON.toJSONBytes("你访问路径被小傅哥的网关管理了 URI：" + request.uri(), SerializerFeature.PrettyFormat));
//        // 头部信息设置
//        HttpHeaders heads = response.headers();
//        // 返回内容类型
//        heads.add(HttpHeaderNames.CONTENT_TYPE, HttpHeaderValues.APPLICATION_JSON + "; charset=UTF-8");
//        // 响应体的长度
//        heads.add(HttpHeaderNames.CONTENT_LENGTH, response.content().readableBytes());
//        // 配置持久连接
//        heads.add(HttpHeaderNames.CONNECTION, HttpHeaderValues.KEEP_ALIVE);
//        // 配置跨域访问
//        heads.add(HttpHeaderNames.ACCESS_CONTROL_ALLOW_ORIGIN, "*");
//        heads.add(HttpHeaderNames.ACCESS_CONTROL_ALLOW_HEADERS, "*");
//        heads.add(HttpHeaderNames.ACCESS_CONTROL_ALLOW_METHODS, "GET, POST, PUT, DELETE");
//        heads.add(HttpHeaderNames.ACCESS_CONTROL_ALLOW_CREDENTIALS, "true");
//
//        channel.writeAndFlush(response);
