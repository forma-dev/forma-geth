package base64

import (
	b64 "encoding/base64"

	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/bindings"
)

type Base64 struct {
	precompile.StatefulPrecompiledContract
}

func NewBase64() *Base64 {
	return &Base64{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			bindings.Base64ABI,
		),
	}
}

func (c *Base64) Encode(ctx precompile.StatefulContext, data []byte) (string, error) {
	return b64.StdEncoding.EncodeToString(data), nil
}

func (c *Base64) EncodeURL(ctx precompile.StatefulContext, data []byte) (string, error) {
	return b64.URLEncoding.EncodeToString(data), nil
}

func (c *Base64) RequiredGas(input []byte) uint64 {
	return precompile.GasQuickStep
}
