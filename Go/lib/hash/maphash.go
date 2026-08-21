-----------------
var
-----------------

-----------------
type
-----------------
	# type ComparableHasher[T comparable] struct {
			// contains filtered or unexported fields
		}
		func (ComparableHasher[T]) Equal(x, y T) bool
		func (ComparableHasher[T]) Hash(h *Hash, v T)

	# type Hash struct {
			// contains filtered or unexported fields
		}
		func (h *Hash) BlockSize() int
		func (h *Hash) Clone() (hash.Cloner, error)
		func (h *Hash) Reset()
		func (h *Hash) Seed() Seed
		func (h *Hash) SetSeed(seed Seed)
		func (h *Hash) Size() int
		func (h *Hash) Sum(b []byte) []byte
		func (h *Hash) Sum64() uint64
		func (h *Hash) Write(b []byte) (int, error)
		func (h *Hash) WriteByte(b byte) error
		func (h *Hash) WriteString(s string) (int, error)

	# type Hasher[T any] interface {
			Hash(*Hash, T)
			Equal(x, y T) bool
		}

		func (h *Hash) BlockSize() int
		func (h *Hash) Clone() (hash.Cloner, error)
		func (h *Hash) Reset()
		func (h *Hash) Seed() Seed
		func (h *Hash) SetSeed(seed Seed)
		func (h *Hash) Size() int
		func (h *Hash) Sum(b []byte) []byte
		func (h *Hash) Sum64() uint64
		func (h *Hash) Write(b []byte) (int, error)
		func (h *Hash) WriteByte(b byte) error
		func (h *Hash) WriteString(s string) (int, error)

	# type Seed struct {
			// contains filtered or unexported fields
		}
		func MakeSeed() Seed
	
-----------------
func
-----------------
	func Bytes(seed Seed, b []byte) uint64
	func Comparable[T comparable](seed Seed, v T) uint64
	func String(seed Seed, s string) uint64
	func WriteComparable[T comparable](h *Hash, x T)
