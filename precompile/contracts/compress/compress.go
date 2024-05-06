package compress

import (
	"bytes"
	"io"

	"github.com/andybalholm/brotli"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/abi"
)

type Compress struct {
	precompile.StatefulPrecompiledContract

	reader *brotli.Reader
	writer *brotli.Writer
}

func NewCompress() *Compress {
	return &Compress{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			abi.CompressABI,
		),
		reader: brotli.NewReader(nil),
		writer: brotli.NewWriterLevel(nil, brotli.DefaultCompression),
	}
}

func (c *Compress) Compress(ctx precompile.StatefulContext, input []byte) ([]byte, error) {
	var buf bytes.Buffer
	c.writer.Reset(&buf)
	_, err := c.writer.Write(input)
	if err != nil {
		return nil, err
	}
	if err := c.writer.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Compress) Decompress(ctx precompile.StatefulContext, input []byte) ([]byte, error) {
	var buf bytes.Buffer
	c.reader.Reset(bytes.NewReader(input))
	_, err := io.Copy(&buf, c.reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// gas calcs
func (c *Compress) CompressRequiredGas(ctx precompile.StatefulContext, input []byte) uint64 {
	return precompile.WordLength(input, 256) * precompile.GasQuickStep
}

func (c *Compress) DecompressRequiredGas(ctx precompile.StatefulContext, input []byte) uint64 {
	return precompile.WordLength(input, 256) * precompile.GasQuickStep
}
