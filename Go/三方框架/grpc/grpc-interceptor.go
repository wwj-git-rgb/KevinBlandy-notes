-------------------
interceptor
-------------------
	# 拦截器
		* 文档
			https://grpc.io/docs/guides/interceptors/
		
		* gRPC 一共有 4 种拦截器，按 “端 × 调用类型” 组合
			Unary
				UnaryClientInterceptor
				UnaryServerInterceptor
				
			Stream
				StreamServerInterceptor
				StreamClientInterceptor
		
--------------------------------------
UnaryClientInterceptor
--------------------------------------
	# 一元调用，客户端拦截器
		* 在连接级别添加，拦截发出去的请求，在请求发出前/后进行执行。
		* grpc.WithChainUnaryInterceptor 可以添加多个，依次执行

		type UnaryClientInterceptor func(ctx context.Context, method string, req, reply any, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
			ctx
				* 用户调用时传入的 ctx 

			method
				* 调用的方法，如 /api.MemberService/Call
			req
				* 即将发出的请求
				* 可以对其进行修改
			reply
				* 返回值的指针（gRPC 会把响应反序列化进去）
				* 可以对其进行修改
			cc
				*  连接对象   
			invoker
				* 继续下一个调用的句柄
					func(ctx context.Context, method string, req, reply any, cc *ClientConn, opts ...CallOption) error
			opts
				* 设置的选项  
			error
				*  错误信息
		
	# Demo
		grpc.NewClient("localhost:7788",
			grpc.WithChainUnaryInterceptor(func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				fmt.Println("[interceptor-1] 请求", req)
				err := invoker(ctx, method, req, reply, cc, opts...)
				fmt.Println("[interceptor-1] 响应", reply)
				return err
			}, func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				fmt.Println("[interceptor-2] 请求", req)
				err := invoker(ctx, method, req, reply, cc, opts...)
				fmt.Println("[interceptor-2] 响应", reply)
				return err
			}),
		)
		// [interceptor-1] 请求 member:{id:10028} 
		// [interceptor-2] 请求 member:{id:10028}
		// [interceptor-2] 响应 member:{id:10028}
		// [interceptor-1] 响应 member:{id:8888}

--------------------------------------
UnaryServerInterceptor
--------------------------------------
	# 一元调用，服务端拦截器

		* 主要作用是在服务端进行拦截，在 Handler/Interceptor 执行之前/后执行。
		* grpc.ChainUnaryInterceptor 可以添加多个，按照添加顺序依次执行，可以对

		type UnaryServerInterceptor func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (resp any, err error)

			ctx
				* 当前请求的 context，可以读取到 metadata
			
			req
				* 已经解码后的 protobuf 请求对象，需要自己进行转换
			
			info
				* 一元请求信息，目前就俩属性
					Server any
					FullMethod string
				
			handler
				* 继续执行调用链的句柄，并且可以获取到执行结果，其结果就是上个调用返回的结果。
					UnaryHandler func(ctx context.Context, req any) (any, error)
				
			resp
				* 返回给客户端的响应
			
			err
				* 返回给客户端的异常
	
	# Demo	
		grpc.NewServer(grpc.ChainUnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			fmt.Println("[filter-1] 执行前", req)
			ret, err := handler(ctx, req)
			fmt.Println("[filter-1] 执行后", ret)
			return ret, err
		}, func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			fmt.Println("[filter-2] 执行前", req)
			ret, err := handler(ctx, req)
			fmt.Println("[filter-2] 执行后", ret)
			return ret, err
		}))

		// [filter-1] 执行前 member:{id:10028}
		// [filter-2] 执行前 member:{id:10028}
		// [filter-2] 执行后 member:{id:10028}
		// [filter-1] 执行后 member:{id:10028}


--------------------------------------
StreamClientInterceptor
--------------------------------------
	type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)

--------------------------------------
StreamServerInterceptor
--------------------------------------
	type StreamServerInterceptor func(srv any, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
	

