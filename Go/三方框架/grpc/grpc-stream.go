
-------------------
Stream
-------------------
	# 接口体系
		Stream
			|-ClientStream
				|-BidiStreamingClient[Req any, Res any]
				|-ClientStreamingClient[Req any, Res any]
				|-ServerStreamingClient[Res any]
				|-GenericClientStream[Req any, Res any]
			|-ServerStream
				|-BidiStreamingServer[Req any, Res any]
				|-ClientStreamingServer[Req any, Res any]
				|-ServerStreamingServer[Res any]
				|-GenericServerStream[Req any, Res any]


