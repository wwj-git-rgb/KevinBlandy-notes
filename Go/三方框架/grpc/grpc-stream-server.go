
------------------
Stream Server
------------------


------------------
客户端单向流
------------------
	# 客户端只流式读，不流式写

		func (ServiceServer) Publisher(clientStream grpc.ClientStreamingServer[api.Request, api.Response]) error {

			// 读取客户端的 Header
			md, ok := metadata.FromIncomingContext(clientStream.Context())
			if ok {
				fmt.Println("header", md)
			}

			// 轮询读取客户端消息
			for {
				msg, err := clientStream.Recv()
				if err != nil {
					// 读取完毕
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}
				fmt.Println("recv", msg)
			}

			defer func() {
				// 设置 Tailer，在方法返回后，框架会自动发送
				clientStream.SetTrailer(metadata.New(map[string]string{
					"t1": "bar",
				}))
			}()

			// 设置 Header
			err := clientStream.SetHeader(metadata.New(map[string]string{
				"h1": "zoo",
			}))
			if err != nil {
				return err
			}

			// 响应客户端数据后关闭
			if err := clientStream.SendAndClose(&api.Response{Member: &model.Member{Id: new(int64(9999))}}); err != nil {
				return err
			}
			fmt.Println("server closed")
			return nil
		}

------------------
服务端单向流
------------------
	# 服务端只流式写入，不读

		func (ServiceServer) Subscribe(req *api.Request, serverStream grpc.ServerStreamingServer[api.Response]) error {

			// 读取客户端的 Header
			header, ok := metadata.FromIncomingContext(serverStream.Context())
			if ok {
				fmt.Println("header", header)
			}

			// 读取客户端参数
			fmt.Println("request", req.Member)

			// 设置响应 Header
			if err := serverStream.SendHeader(metadata.Pairs("h1", "h2")); err != nil {
				return err
			}

			// 设置响应 Tailer，框架自动在末尾发送
			serverStream.SetTrailer(metadata.Pairs("t1", "val"))

			// 发送多条消息到客户端
			for i := range 10 {
				if err := serverStream.SendMsg(&api.Response{Member: &model.Member{Id: new(int64(i))}}); err != nil {
					return err
				}
			}

			fmt.Println("server closed")

			// 方法返回，表示流结束
			return nil
		}

------------------
双向流
------------------
	# 双全工通信，可读可写

		func (ServiceServer) Stream(stream grpc.BidiStreamingServer[api.Request, api.Response]) error {

			wg := &sync.WaitGroup{}

			// 并发读
			wg.Go(func() {
				// 读取 Header
				header, ok := metadata.FromIncomingContext(stream.Context())
				if ok {
					fmt.Println("header", header)
				}

				// 读取消息
				for {
					msg, err := stream.Recv()
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						fmt.Println("读取异常", err.Error())
						return
					}
					fmt.Println("recv", msg)
				}
				fmt.Println("读取完毕")
			})

			// 并发写
			wg.Go(func() {
				// 设置 Header 和 Trailer
				err := stream.SetHeader(metadata.New(map[string]string{
					"h1": "zoo",
				}))
				if err != nil {
					fmt.Println("Header 设置异常", err.Error())
					return
				}
				stream.SetTrailer(metadata.New(map[string]string{
					"h2": "zoo1",
				}))

				msgCount := rand.Intn(20) + 1

				// 写入消息
				for i := range msgCount {
					if err := stream.Send(&api.Response{Member: &model.Member{Id: new(int64(i)), Title: new("Server")}}); err != nil {
						fmt.Println("写入异常", err.Error())
						return
					}
				}
				fmt.Println("写入完毕", msgCount)
			})

			wg.Wait()

			// 方法返回，流关闭
			fmt.Println("服务端关闭")

			return nil
		}