-----------------
Go
-----------------
	# Metadata 即元数据，类似于 “Header” ，分为 Header 和 Tailer
		* 客户端设置为 out
		* 服务端读取为 In
		* 客户端、服务端都可以设置 Header
		* Trailer 只能是服务端设置，客户端读取

	# 服务端设置 Metadata

		Server -> Client:  Header ->  Message(可以多条) -> Trailer (含 status)

		* Header 只能必须在消息之前发送，只能发送一次，如果没有手动发送，则会在第一条消息发送的时候，先发送 Header。
		* Header 可以在发送前，通过 SetHeader(metadata.MD) error 多次设置 Header。
		* 没有 SendTrailer，因为 Trailer 一定是 RPC 结束（handler return 或出错）时由框架自动发。
	
	# 服务端读取 Metadata
		* 通过 Context 进行读取
			// 从一元调用的 Context 参数重读取 md
			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				fmt.Println(md)
			}

			// 从流式参数的 Context 中读取 md
			md, ok := metadata.FromIncomingContext(serverStream.Context())
			if ok {
				fmt.Println(md)
			}

	
	# 客户端设置 Metadata

		* 通过 Context 进行设置
			// 新建整个 Header
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"foo": "bar"}))

			// 添加到已有 Header
			ctx = metadata.AppendToOutgoingContext(ctx, "zoo", "joo")

			// 使用 Ctx 发起请求                                                                                             
			stream, err := serviceClient.Subscribe(ctx, &api.Request{...})    
			
	
	# 客户端读取 Metadata

		* 通过 grpc.CallOption 进行读取

			header := &metadata.MD{} // Header 
			tailer := &metadata.MD{} // Tailer

			serviceClient.Call(context.Background(), &api.Request{Member: &model.Member{
				Id:    new(int64(1)),
				Title: new("hi"),
			}}, grpc.Header(header), grpc.Trailer(tailer)) // 通过 grpc.CallOption 进行读取

			fmt.Println("header", header)
			fmt.Println("tailer", tailer)

		* 流式客户端还支持 API 直接读取（也支持 grpc.CallOption）
				
			Header() (metadata.MD, error)
				* 在流消息之前读取。
				* 会阻塞调用，直到 Header 可用。

			Trailer() metadata.MD
				* 在流消息读取完毕后读取。





