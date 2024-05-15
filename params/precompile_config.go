package params

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile"

	pcbase64 "github.com/ethereum/go-ethereum/precompile/contracts/base64"
	pccompress "github.com/ethereum/go-ethereum/precompile/contracts/compress"
	pcintegers "github.com/ethereum/go-ethereum/precompile/contracts/integers"
	pcjsonstore "github.com/ethereum/go-ethereum/precompile/contracts/jsonstore"
	pcjsonutil "github.com/ethereum/go-ethereum/precompile/contracts/jsonutil"
	pcnativeminter "github.com/ethereum/go-ethereum/precompile/contracts/nativeminter"
	pcstrings "github.com/ethereum/go-ethereum/precompile/contracts/strings"
)

var NullPrecompiles = precompile.PrecompileMap{}

var FormaGenesisPrecompiles = precompile.PrecompileMap{
	common.HexToAddress("0x0F043A000001"): pcnativeminter.NewNativeMinter(),
	common.HexToAddress("0x0F043A000002"): pccompress.NewCompress(),
	common.HexToAddress("0x0F043A000003"): pcjsonutil.NewJsonUtil(),
	common.HexToAddress("0x0F043A000004"): pcbase64.NewBase64(),
	common.HexToAddress("0x0F043A000005"): pcstrings.NewStrings(),
	common.HexToAddress("0x0F043A000006"): pcintegers.NewIntegers(),
	common.HexToAddress("0x0F043A000007"): pcjsonstore.NewJsonStore(),
}

// return precompiles that are enabled at height
func (c *ChainConfig) Precompiles(height uint64, timestamp uint64) precompile.PrecompileMap {
	return FormaGenesisPrecompiles
}
