-------------------------
UUID
-------------------------
	# UUID 的官方版本

-------------------------
var
-------------------------

-------------------------
type
-------------------------
	# type UUID [16]byte

		func Max() UUID
		func MustParse(s string) UUID
		func New() UUID
		func NewV4() UUID
		func NewV7() UUID
			* 有序的
		
		func Nil() UUID
			* 返回空的 UUID{}

		func Parse(s string) (UUID, error)

		func (u UUID) AppendText(b []byte) ([]byte, error)
		func (u UUID) Compare(v UUID) int
		func (u UUID) MarshalText() ([]byte, error)
		func (u UUID) String() string
		func (u *UUID) UnmarshalText(b []byte) error

-------------------------
func
-------------------------

