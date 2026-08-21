----------------------------
flate
----------------------------
	# DEFLATE 压缩实现
		* 实现了 RFC 1951 中描述的 DEFLATE 压缩数据格式。
		* compress/gzip 和 compress/zlib 软件包实现了对基于 DEFLATE 的文件格式的访问。
	
----------------------------
var
----------------------------

	const (
		NoCompression      = 0
		BestSpeed          = 1
		BestCompression    = 9
		DefaultCompression = -1

		// HuffmanOnly disables Lempel-Ziv match searching and only performs Huffman
		// entropy encoding. This mode is useful in compressing data that has
		// already been compressed with an LZ style algorithm (e.g. Snappy or LZ4)
		// that lacks an entropy encoder. Compression gains are achieved when
		// certain bytes in the input stream occur more frequently than others.
		//
		// Note that HuffmanOnly produces a compressed output that is
		// RFC 1951 compliant. That is, any valid DEFLATE decompressor will
		// continue to be able to decompress this output.
		HuffmanOnly = -2
	)

----------------------------
type
----------------------------
	# type CorruptInputError int64

		func (e CorruptInputError) Error() string
	
	# type InternalError string
		func (e InternalError) Error() string
	
	# type ReadError struct {
			Offset int64 // byte offset where error occurred
			Err    error // error returned by underlying Read
		}
		func (e *ReadError) Error() string
	
	# type Reader interface {
			io.Reader
			io.ByteReader
		}
	
	# type Resetter interface {
			Reset(r io.Reader, dict []byte) error
		}
	
	# type WriteError struct {
			Offset int64 // byte offset where error occurred
			Err    error // error returned by underlying Write
		}
		func (e *WriteError) Error() string
	
	# type Writer struct {
		}
		func NewWriter(w io.Writer, level int) (*Writer, error)
		func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)
		func (w *Writer) Close() error
		func (w *Writer) Flush() error
		func (w *Writer) Reset(dst io.Writer)
		func (w *Writer) Write(data []byte) (n int, err error)

----------------------------
func
----------------------------
	func NewReader(r io.Reader) io.ReadCloser
	func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
