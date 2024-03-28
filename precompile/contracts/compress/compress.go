package compress

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/bindings"
)

type Compress struct {
	precompile.StatefulPrecompiledContract
}

func NewCompress() *Compress {
	return &Compress{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			bindings.CompressABI,
		),
	}
}

func (c *Compress) Deflate(ctx precompile.StatefulContext, input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(input)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Compress) Inflate(ctx precompile.StatefulContext, input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz, err := gzip.NewReader(bytes.NewBuffer(input))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&buf, gz)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Compress) RequiredGas(input []byte) uint64 {
	return precompile.GasQuickStep
}
