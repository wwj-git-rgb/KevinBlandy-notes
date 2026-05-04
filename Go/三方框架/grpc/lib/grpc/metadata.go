-------------------------
元数据
-------------------------


-------------------------
var
-------------------------


-------------------------
type
-------------------------
	# type MD map[string][]string

		* MetaData
		* 所有的 key 操作的时候都会被转换为小写。

		func FromIncomingContext(ctx context.Context) (MD, bool)
		func FromOutgoingContext(ctx context.Context) (MD, bool)
			* 根 In 或 Out 获取到全局上下文，Key 分别是：
				type mdIncomingKey struct{}
				type mdOutgoingKey struct{}

			* 返回的结果是复制的

		func Join(mds ...MD) MD
			* 拼接所有的元数据

		func New(m map[string]string) MD
			* 根据 m 构建新的 MD，完整复制

		func Pairs(kv ...string) MD
			* 根据键值对构建 MD，kv 必须是 2 的整数倍，否则 panic

		func (md MD) Append(k string, vals ...string)
		func (md MD) Copy() MD
		func (md MD) Delete(k string)
		func (md MD) Get(k string) []string
		func (md MD) Len() int
		func (md MD) Set(k string, vals ...string)

-------------------------
func
-------------------------

	func AppendToOutgoingContext(ctx context.Context, kv ...string) context.Context
	func DecodeKeyValue(k, v string) (string, string, error)

	func NewIncomingContext(ctx context.Context, md MD) context.Context
	func NewOutgoingContext(ctx context.Context, md MD) context.Context
		* 构建新的 Context
		* In 表示入， Out 表示出

	func ValueFromIncomingContext(ctx context.Context, key string) []string
