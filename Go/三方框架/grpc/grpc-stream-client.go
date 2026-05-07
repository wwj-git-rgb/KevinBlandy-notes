------------------
Stream Client
------------------


------------------
客户端单向流
------------------
	# 客户端只流式发送不流式读

		func ClientStream(client api.MemberServiceClient) error {

			// 传递 Header
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("h1", "v1"))

			// 调用服务端
			stream, err := client.Publisher(ctx)
			if err != nil {
				return err
			}

			// 发送 N 条消息
			for i := range 10 {
				if err := stream.Send(&api.Request{Member: &model.Member{Id: new(int64(i))}}); err != nil {
					return err
				}
			}

			//// 直接读取响应 & 关闭
			//resp, err := stream.CloseAndRecv()
			//if err != nil {
			//	return err
			//}
			//fmt.Println("recv", resp)

			// 客户端发送完毕
			if err := stream.CloseSend(); err != nil {
				return err
			}

			// 读取响应头
			header, err := stream.Header()
			if err != nil {
				return err
			}
			fmt.Println("header", header)

			// 读取响应的消息
			var response = &api.Response{}
			if err := stream.RecvMsg(response); err != nil {
				return err
			}
			fmt.Println("recv", response)

			// 读取 Tailer
			tailer := stream.Trailer()
			fmt.Println("tailer", tailer)

			return nil
		}

------------------
服务端单向流
------------------
	# 客户端只接收数据，不发送数据
		func ServerStream(client api.MemberServiceClient) error {
			// 设置 header
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("k1", "v1"))

			// 发起请求
			stream, err := client.Subscribe(ctx, &api.Request{Member: &model.Member{
				Id: new(int64(1)),
			}})
			if err != nil {
				return err
			}

			// 读取 Header
			header, err := stream.Header()
			if err != nil {
				return err
			}
			fmt.Println("header", header)

			// 轮询服务器消息
			for {
				msg, err := stream.Recv()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}
				fmt.Println("recv", msg)
			}

			// 读取 Tailer
			trailer := stream.Trailer()
			fmt.Println("trailer", trailer)

			// 服务端关闭了连接
			fmt.Println("server closed")

			return nil
		}

------------------
双向流
------------------
	# 双全工通信，可读可写

		func BidiStream(client api.MemberServiceClient) error {

			// 创建 Header
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("h1", "zoo"))

			// 创建连接
			stream, err := client.Stream(ctx)
			if err != nil {
				return err
			}

			wg := &sync.WaitGroup{}
			// 并发读
			wg.Go(func() {
				// 读取 Header
				header, err := stream.Header()
				if err != nil {
					fmt.Println("header 读取异常", err.Error())
					return
				}
				fmt.Println("header", header)

				// 迭代消息
				for {
					msg, err := stream.Recv()
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						fmt.Println("消息读取异常", err.Error())
						return
					}
					fmt.Println("revc", msg)
				}

				// 服务端关闭写入流
				fmt.Println("读取完毕")

				// 读取 Trailer
				trailer := stream.Trailer()

				fmt.Println("trailer", trailer)
			})

			// 并发写
			wg.Go(func() {

				msgCount := rand.Intn(10) + 1

				// 写入 N 条数据
				for i := range msgCount {
					if err := stream.Send(&api.Request{Member: &model.Member{Id: new(int64(i)), Title: new("来自于客户端的消息")}}); err != nil {
						fmt.Println("写入异常", err.Error())
						return
					}
				}
				// 写入完毕
				if err := stream.CloseSend(); err != nil {
					fmt.Println("写入流关闭异常", err.Error())
					return
				}
				// 服务端关闭写入流
				fmt.Println("写入完毕", msgCount)
			})

			wg.Wait()

			return nil
		}
