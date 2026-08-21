------------------------
mldsa
------------------------
	# 抗衡量子计算的加密算法

------------------------
var
------------------------

------------------------
type
------------------------
	# type Options struct {
			Context string
		}
		func (opts *Options) HashFunc() crypto.Hash
	
	
	# type Parameters struct {
		}
		func MLDSA44() Parameters
		func MLDSA65() Parameters
		func MLDSA87() Parameters
		func (params Parameters) PublicKeySize() int
		func (params Parameters) SignatureSize() int
		func (params Parameters) String() string
	
	# type PrivateKey struct {
		}
		func GenerateKey(params Parameters) (*PrivateKey, error)
		func NewPrivateKey(params Parameters, seed []byte) (*PrivateKey, error)
		func (sk *PrivateKey) Bytes() []byte
		func (sk *PrivateKey) Equal(x crypto.PrivateKey) bool
		func (sk *PrivateKey) Public() crypto.PublicKey
		func (sk *PrivateKey) PublicKey() *PublicKey
		func (sk *PrivateKey) Sign(_ io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)
		func (sk *PrivateKey) SignDeterministic(message []byte, opts crypto.SignerOpts) (signature []byte, err error)
	
	# type PublicKey struct {
		}
		func NewPublicKey(params Parameters, encoding []byte) (*PublicKey, error)
		func (pk *PublicKey) Bytes() []byte
		func (pk *PublicKey) Equal(x crypto.PublicKey) bool
		func (pk *PublicKey) Parameters() Parameters

------------------------
func
------------------------
	func Verify(pk *PublicKey, message []byte, signature []byte, opts *Options) error
