package server

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/xhttp"
	"github.com/kwseeker/gateway/gateway-core/handler"
)

const DefaultServerUrl = "0.0.0.0:7380"

type Server struct {
	url string
}

func (s *Server) StartDefault() {
	s.Start(DefaultServerUrl)
}

func (s *Server) Start(url string) {
	// child pipeline initializer.
	setupCodec := func(channel netty.Channel) {
		channel.Pipeline().
			// decode http request from channel
			AddLast(xhttp.ServerCodec()).
			// print http access log
			AddLast(new(handler.StateHandler)).
			// compatible with http.Handler
			//AddLast(xhttp.Handler(httpMux))
			AddLast(xhttp.Handler(handler.NewSession())) //TODO 这里后面参考xhttp.Handler重新实现一个Handler
	}

	// setup bootstrap & startup server.
	// TODO SO_BACKLOG等选项
	err := netty.NewBootstrap(netty.WithChildInitializer(setupCodec)).
		Listen(url).Sync()
	if err != nil {
		fmt.Println("Session server start failed!")
		return
	}
}

//	  private final EventLoopGroup boss = new NioEventLoopGroup(1);
//    private final EventLoopGroup work = new NioEventLoopGroup();
//    private Channel channel;
//
//    @Override
//    public Channel call() throws Exception {
//        ChannelFuture channelFuture = null;
//        try {
//            ServerBootstrap b = new ServerBootstrap();
//            b.group(boss, work)
//                    .channel(NioServerSocketChannel.class)
//                    .option(ChannelOption.SO_BACKLOG, 128)
//                    .childHandler(new SessionChannelInitializer());
//
//            channelFuture = b.bind(new InetSocketAddress(7397)).syncUninterruptibly();
//            this.channel = channelFuture.channel();
//        } catch (Exception e) {
//            logger.error("socket server start error.", e);
//        } finally {
//            if (null != channelFuture && channelFuture.isSuccess()) {
//                logger.info("socket server start done.");
//            } else {
//                logger.error("socket server start error.");
//            }
//        }
//        return channel;
//    }
