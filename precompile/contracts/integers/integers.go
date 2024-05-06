package integers

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/abi"
)

type Integers struct {
	precompile.StatefulPrecompiledContract
}

func NewIntegers() *Integers {
	return &Integers{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			abi.IntegersABI,
		),
	}
}

// Solidity: function toString(uint256 _i) pure returns(string)
func (c *Integers) ToString(ctx precompile.StatefulContext, i *big.Int) (string, error) {
	return i.String(), nil
}

// Solidity: function toString(int256 _i) pure returns(string)
func (c *Integers) ToString0(ctx precompile.StatefulContext, i *big.Int) (string, error) {
	return i.String(), nil
}

// Solidity: function toHexString(uint256 _i, uint256 _length) pure returns(string)
func (c *Integers) ToHexString(ctx precompile.StatefulContext, i *big.Int, length *big.Int) (string, error) {
	hex := fmt.Sprintf("%x", i)

	requiredLength := int(length.Uint64()) * 2
	if len(hex) > requiredLength {
		return "", fmt.Errorf("the length of the hex string is greater than the required length")
	}

	hex = "0x" + strings.Repeat("0", requiredLength-len(hex)) + hex
	return hex, nil
}

// Solidity: function toHexString(uint256 _i) pure returns(string)
func (c *Integers) ToHexString0(ctx precompile.StatefulContext, i *big.Int) (string, error) {
	hex := fmt.Sprintf("%x", i)
	if len(hex)%2 != 0 {
		hex = "0" + hex
	}
	hex = "0x" + hex
	return hex, nil
}

// Solidity: function fromHexString(string calldata _str) pure returns (uint256)
func (c *Integers) FromHexString(ctx precompile.StatefulContext, str string) (*big.Int, error) {
	str = strings.TrimPrefix(strings.ToLower(str), "0x")
	i, success := new(big.Int).SetString(str, 16)
	if !success {
		return nil, fmt.Errorf("could not convert hex string to uint256")
	}
	return i, nil
}

func (c *Integers) DefaultGas(input []byte) uint64 {
	return precompile.GasQuickStep
}
