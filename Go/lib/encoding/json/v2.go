-------------------
json
-------------------
	# JSON 的第二个版本


-------------------
var
-------------------
	var ErrUnknownName = errors.New("unknown object member name")


-------------------
type
-------------------
	# type Marshaler interface {
			MarshalJSON() ([]byte, error)
		}
	
	# type MarshalerTo interface {
			MarshalJSONTo(*jsontext.Encoder) error
		}
	
	# type Marshalers = typedMarshalers
		func JoinMarshalers(ms ...*Marshalers) *Marshalers
		func MarshalFunc[T any](fn func(T) ([]byte, error)) *Marshalers
		func MarshalToFunc[T any](fn func(*jsontext.Encoder, T) error) *Marshalers
	
	# type Options = jsonopts.Options
		func DefaultOptionsV2() Options
		func Deterministic(v bool) Options
		func FormatNilMapAsNull(v bool) Options
		func FormatNilSliceAsNull(v bool) Options
		func JoinOptions(srcs ...Options) Options
		func MatchCaseInsensitiveNames(v bool) Options
		func OmitZeroStructFields(v bool) Options
		func RejectUnknownMembers(v bool) Options
		func StringifyNumbers(v bool) Options
		func WithMarshalers(v *Marshalers) Options
		func WithUnmarshalers(v *Unmarshalers) Options
	
	# type SemanticError struct {
			ByteOffset int64
			JSONPointer jsontext.Pointer

			JSONKind jsontext.Kind // may be zero if unknown
			JSONValue jsontext.Value // may be nil if irrelevant or unknown
			GoType reflect.Type // may be nil if unknown

			Err error // may be nil
		}
		func (e *SemanticError) Error() string
		func (e *SemanticError) Unwrap() error
	
	# type Unmarshaler interface {
			UnmarshalJSON([]byte) error
		}
	
	# type UnmarshalerFrom interface {
			UnmarshalJSONFrom(*jsontext.Decoder) error
		}
	
	# type Unmarshalers = typedUnmarshalers
		func JoinUnmarshalers(us ...*Unmarshalers) *Unmarshalers
		func UnmarshalFromFunc[T any](fn func(*jsontext.Decoder, T) error) *Unmarshalers
		func UnmarshalFunc[T any](fn func([]byte, T) error) *Unmarshalers

-------------------
func
-------------------

	func GetOption[T any](opts Options, setter func(T) Options) (T, bool)
	func Marshal(in any, opts ...Options) (out []byte, err error)
	func MarshalEncode(out *jsontext.Encoder, in any, opts ...Options) (err error)
	func MarshalWrite(out io.Writer, in any, opts ...Options) (err error)
	func Unmarshal(in []byte, out any, opts ...Options) (err error)
	func UnmarshalDecode(in *jsontext.Decoder, out any, opts ...Options) (err error)
	func UnmarshalRead(in io.Reader, out any, opts ...Options) (err error)
