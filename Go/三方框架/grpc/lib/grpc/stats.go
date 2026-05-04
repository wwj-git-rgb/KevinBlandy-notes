--------------------------
stats
--------------------------

--------------------------
var
--------------------------

--------------------------
type
--------------------------
	# type Begin struct {
			Client bool
			BeginTime time.Time
			FailFast bool
			IsClientStream bool
			IsServerStream bool
			IsTransparentRetryAttempt bool
		}

		func (s *Begin) IsClient() bool
	
	# type ConnBegin struct {
			Client bool
		}

		func (s *ConnBegin) IsClient() bool
	
	# type ConnEnd struct {
			Client bool
		}
		
		func (s *ConnEnd) IsClient() bool
	
	# type ConnStats interface {
			IsClient() bool
		}
	
	# type ConnTagInfo struct {
			RemoteAddr net.Addr
			LocalAddr net.Addr
		}
	
	# type DelayedPickComplete struct{}
		func (*DelayedPickComplete) IsClient() bool
	
	# type End struct {
			Client bool
			BeginTime time.Time
			EndTime time.Time
			Trailer metadata.MD
			Error error
		}

		func (s *End) IsClient() bool
	
	# type Handler interface {
			TagRPC(context.Context, *RPCTagInfo) context.Context
			HandleRPC(context.Context, RPCStats)
			TagConn(context.Context, *ConnTagInfo) context.Context
			HandleConn(context.Context, ConnStats)
		}
	
	# type InHeader struct {
			Client bool
			WireLength int
			Compression string
			Header metadata.MD
			FullMethod string
			RemoteAddr net.Addr
			LocalAddr net.Addr
		}
		
		func (s *InHeader) IsClient() bool
	
	# type InPayload struct {
			Client bool
			Payload any
			Length int
			CompressedLength int
			WireLength int
			RecvTime time.Time
		}

		func (s *InPayload) IsClient() bool
	
	# type InTrailer struct {
			Client bool
			WireLength int
			Trailer metadata.MD
		}

		func (s *InTrailer) IsClient() bool
	
	# type MetricSet struct {
		}
		

		func NewMetricSet(metricNames ...string) *MetricSet
		func (m *MetricSet) Add(metricNames ...string) *MetricSet
		func (m *MetricSet) Join(metrics *MetricSet) *MetricSet
		func (m *MetricSet) Metrics() map[string]bool
		func (m *MetricSet) Remove(metricNames ...string) *MetricSet
	
	# type OutHeader struct {
			Client bool
			Compression string
			Header metadata.MD
			FullMethod string
			RemoteAddr net.Addr
			LocalAddr net.Addr
		}

		func (s *OutHeader) IsClient() bool
	
	# type OutPayload struct {
			Client bool
			Payload any
			Length int
			CompressedLength int
			WireLength int
			SentTime time.Time
		}

		func (s *OutPayload) IsClient() bool
	
	# type OutTrailer struct {
			Client bool
			WireLength int
			Trailer metadata.MD
		}

		func (s *OutTrailer) IsClient() bool
	
	# type PickerUpdated = DelayedPickComplete

	# type RPCStats interface {
			IsClient() bool
		}
	
	# type RPCTagInfo struct {
			FullMethodName string
			FailFast bool
			NameResolutionDelay bool
		}
	

--------------------------
func
--------------------------
	func SetTags(ctx context.Context, b []byte) context.Context
	func SetTrace(ctx context.Context, b []byte) context.Context
	func Tags(ctx context.Context) []byte
	func Trace(ctx context.Context) []byte

