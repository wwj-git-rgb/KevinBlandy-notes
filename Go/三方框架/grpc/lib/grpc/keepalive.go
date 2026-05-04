-------------------------
keepalive
-------------------------

-------------------------
var
-------------------------

-------------------------
type
-------------------------
	# type ClientParameters struct {
			Time time.Duration
			Timeout time.Duration
			PermitWithoutStream bool
		}

		* 客户端连接超时配置
	
	# type EnforcementPolicy struct {
			MinTime time.Duration // The current default value is 5 minutes.
			PermitWithoutStream bool // false by default.
		}

		* 用于在服务器端设置保持活动连接的强制策略。
		* 如果客户端违反此策略，服务器将关闭与该客户端的连接。

	
	# type ServerParameters struct {
			MaxConnectionIdle time.Duration // The current default value is infinity.
			MaxConnectionAge time.Duration // The current default value is infinity.
			MaxConnectionAgeGrace time.Duration // The current default value is infinity.
			Time time.Duration // The current default value is 2 hours.
			Timeout time.Duration // The current default value is 20 seconds.
		}

		* 用于在服务器端设置 keepalive 和 max-age 参数。

-------------------------
func
-------------------------
