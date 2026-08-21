----------------------
jsontext
----------------------	
	# JSON 语法处理。

----------------------
var
----------------------
	var ErrDuplicateName = errors.New("duplicate object member name")
	var ErrNonStringName = errors.New("object member name must be a string")
	var Internal exporter

----------------------
type
----------------------
	# type Decoder struct {
			// contains filtered or unexported fields
		}
		func NewDecoder(r io.Reader, opts ...Options) *Decoder
		func (d *Decoder) InputOffset() int64
		func (d *Decoder) Options() Options
		func (d *Decoder) PeekKind() Kind
		func (d *Decoder) ReadToken() (Token, error)
		func (d *Decoder) ReadValue() (Value, error)
		func (d *Decoder) Reset(r io.Reader, opts ...Options)
		func (d *Decoder) SkipValue() error
		func (d *Decoder) StackDepth() int
		func (d *Decoder) StackIndex(i int) (Kind, int64)
		func (d *Decoder) StackPointer() Pointer
		func (d *Decoder) UnreadBuffer() []byte
	
	# type Encoder struct {
			// contains filtered or unexported fields
		}
		func NewEncoder(w io.Writer, opts ...Options) *Encoder
		func (e *Encoder) AvailableBuffer() []byte
		func (e *Encoder) Options() Options
		func (e *Encoder) OutputOffset() int64
		func (e *Encoder) Reset(w io.Writer, opts ...Options)
		func (e *Encoder) StackDepth() int
		func (e *Encoder) StackIndex(i int) (Kind, int64)
		func (e *Encoder) StackPointer() Pointer
		func (e *Encoder) WriteToken(t Token) error
		func (e *Encoder) WriteValue(v Value) error
	
	# type Kind byte
		const (
			KindInvalid     Kind = 0   // invalid kind
			KindNull        Kind = 'n' // null
			KindFalse       Kind = 'f' // false
			KindTrue        Kind = 't' // true
			KindString      Kind = '"' // string
			KindNumber      Kind = '0' // number
			KindBeginObject Kind = '{' // begin object
			KindEndObject   Kind = '}' // end object
			KindBeginArray  Kind = '[' // begin array
			KindEndArray    Kind = ']' // end array
		)
		func (k Kind) String() string
	
	# type Options = jsonopts.Options
		func AllowDuplicateNames(v bool) Options
		func AllowInvalidUTF8(v bool) Options
		func CanonicalizeRawFloats(v bool) Options
		func CanonicalizeRawInts(v bool) Options
		func EscapeForHTML(v bool) Options
		func EscapeForJS(v bool) Options
		func Multiline(v bool) Options
		func PreserveRawStrings(v bool) Options
		func ReorderRawObjects(v bool) Options
		func SpaceAfterColon(v bool) Options
		func SpaceAfterComma(v bool) Options
		func WithIndent(indent string) Options
		func WithIndentPrefix(prefix string) Options
	
	# type Pointer string
		func (p Pointer) AppendToken(tok string) Pointer
		func (p Pointer) Contains(pc Pointer) bool
		func (p Pointer) IsValid() bool
		func (p Pointer) LastToken() string
		func (p Pointer) Parent() Pointer
		func (p Pointer) Tokens() iter.Seq[string]
	
	# type SyntacticError struct {
			ByteOffset int64
			JSONPointer Pointer
			Err error
		}
		func (e *SyntacticError) Error() string
		func (e *SyntacticError) Unwrap() error
	
	# type Token struct {
		}
		func Bool(b bool) Token
		func Float(n float64) Token
		func Float32(n float32) Token
		func Int(n int64) Token
		func String(s string) Token
		func Uint(n uint64) Token
		func (t Token) Bool() bool
		func (t Token) Clone() Token
		func (t Token) Float() (float64, error)
		func (t Token) Float32() (float32, error)
		func (t Token) Int() (int64, error)
		func (t Token) Kind() Kind
		func (t Token) String() string
		func (t Token) Uint() (uint64, error)
	
	# type Value []byte
		func (v *Value) Canonicalize(opts ...Options) error
		func (v Value) Clone() Value
		func (v *Value) Compact(opts ...Options) error
		func (v *Value) Format(opts ...Options) error
		func (v *Value) Indent(opts ...Options) error
		func (v Value) IsValid(opts ...Options) bool
		func (v Value) Kind() Kind
		func (v Value) MarshalJSON() ([]byte, error)
		func (v Value) String() string
		func (v *Value) UnmarshalJSON(b []byte) error


----------------------
func
----------------------

	func AppendFloat(dst []byte, src float64, bits int) []byte
	func AppendFormat[Bytes ~[]byte | ~string](dst []byte, src Bytes, opts ...Options) ([]byte, error)
	func AppendQuote[Bytes ~[]byte | ~string](dst []byte, src Bytes) ([]byte, error)
	func AppendUnquote[Bytes ~[]byte | ~string](dst []byte, src Bytes) ([]byte, error)
