package strings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/bindings"
)

type Strings struct {
	precompile.StatefulPrecompiledContract
}

func NewStrings() *Strings {
	return &Strings{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			bindings.StringsABI,
		),
	}
}

// Solidity: function equal(string calldata _a, string calldata _b) view external returns (bool)
func (c *Strings) Equal(ctx precompile.StatefulContext, a string, b string) (bool, error) {
	return a == b, nil
}

// Solidity: function equalCaseFold(string calldata _a, string calldata _b) view external returns (bool)
func (c *Strings) EqualCaseFold(ctx precompile.StatefulContext, a string, b string) (bool, error) {
	return strings.EqualFold(a, b), nil
}

// Solidity: function contains(string calldata _str, string calldata _substr) view external returns (bool)
func (c *Strings) Contains(ctx precompile.StatefulContext, str string, substr string) (bool, error) {
	return strings.Contains(str, substr), nil
}

// Solidity: function startsWith(string calldata _str, string calldata _substr) view external returns (bool)
func (c *Strings) StartsWith(ctx precompile.StatefulContext, str string, substr string) (bool, error) {
	return strings.HasPrefix(str, substr), nil
}

// Solidity: function endsWith(string calldata _str, string calldata _substr) view external returns (bool)
func (c *Strings) EndsWith(ctx precompile.StatefulContext, str string, substr string) (bool, error) {
	return strings.HasSuffix(str, substr), nil
}

// Solidity: function indexOf(string calldata _str, string calldata _substr) view external returns (uint256)
func (c *Strings) IndexOf(ctx precompile.StatefulContext, str string, substr string) (*big.Int, error) {
	return big.NewInt(int64(strings.Index(str, substr))), nil
}

// Solidity: function toUpperCase(string calldata _str) view external returns (string memory)
func (c *Strings) ToUpperCase(ctx precompile.StatefulContext, str string) (string, error) {
	return strings.ToUpper(str), nil
}

// Solidity: function toLowerCase(string calldata _str) view external returns (string memory)
func (c *Strings) ToLowerCase(ctx precompile.StatefulContext, str string) (string, error) {
	return strings.ToLower(str), nil
}

// Solidity: function padStart(string calldata _str, uint16 _len, string calldata _pad) view external returns (string memory)
func (c *Strings) PadStart(ctx precompile.StatefulContext, str string, _len uint16, pad string) (string, error) {
	var overallLen = int(_len)
	var padCountInt = 1 + ((overallLen - len(pad)) / len(pad))
	var retStr = strings.Repeat(pad, padCountInt) + str
	return retStr[(len(retStr) - overallLen):], nil
}

// Solidity: function padEnd(string calldata _str, uint16 _len, string calldata _pad) view external returns (string memory)
func (c *Strings) PadEnd(ctx precompile.StatefulContext, str string, _len uint16, pad string) (string, error) {
	var overallLen = int(_len)
	var padCountInt = 1 + ((overallLen - len(pad)) / len(pad))
	var retStr = str + strings.Repeat(pad, padCountInt)
	return retStr[:overallLen], nil
}

// Solidity: function repeat(string calldata _str, uint16 _count) view external returns (string memory)
func (c *Strings) Repeat(ctx precompile.StatefulContext, str string, count uint16) (string, error) {
	return strings.Repeat(str, int(count)), nil
}

// Solidity: function replace(string calldata _str, string calldata _old, string calldata _new, uint16 _n) view external returns (string memory)
func (c *Strings) Replace(ctx precompile.StatefulContext, str string, old string, new string, n uint16) (string, error) {
	return strings.Replace(str, old, new, int(n)), nil
}

// Solidity: function replaceAll(string calldata _str, string calldata _old, string calldata _new) view external returns (string memory)
func (c *Strings) ReplaceAll(ctx precompile.StatefulContext, str string, old string, new string) (string, error) {
	return strings.ReplaceAll(str, old, new), nil
}

// Solidity: function split(string calldata _str, string calldata _delim) view external returns (string[] memory)
func (c *Strings) Split(ctx precompile.StatefulContext, str string, delim string) ([]string, error) {
	return strings.Split(str, delim), nil
}

// Solidity: function trim(string calldata _str) view external returns (string memory)
func (c *Strings) Trim(ctx precompile.StatefulContext, str string) (string, error) {
	return strings.TrimSpace(str), nil
}

func (c *Strings) RequiredGas(input []byte) uint64 {
	return precompile.GasFastStep
}
