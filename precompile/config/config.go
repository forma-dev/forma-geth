package config

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
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

// var FormaGenesisPrecompiles = precompile.PrecompileMap{
// 	common.HexToAddress("0x0F043A000001"): pcnativeminter.NewNativeMinter(),
// 	common.HexToAddress("0x0F043A000002"): pccompress.NewCompress(),
// 	common.HexToAddress("0x0F043A000003"): pcjsonutil.NewJsonUtil(),
// 	common.HexToAddress("0x0F043A000004"): pcbase64.NewBase64(),
// 	common.HexToAddress("0x0F043A000005"): pcstrings.NewStrings(),
// 	common.HexToAddress("0x0F043A000006"): pcintegers.NewIntegers(),
// 	common.HexToAddress("0x0F043A000007"): pcjsonstore.NewJsonStore(),
// }

// Cache of initialized precompiles for each fork
var precompileCache = make(map[string]precompile.PrecompileMap)

// return precompiles that are enabled at height
func GetPrecompiles(chainConfig *params.ChainConfig, height uint64, timestamp uint64) precompile.PrecompileMap {
	fork := chainConfig.GetAstriaForks().GetForkAtHeight(height)

	// Return early if no precompiles configured
	if len(fork.Precompiles) == 0 {
		return NullPrecompiles
	}

	// Check if we've already initialized precompiles for this fork
	if cached, exists := precompileCache[fork.Name]; exists {
		return cached
	}

	// Initialize precompiles for this fork
	precompiles := make(precompile.PrecompileMap)
	for addr, precompileType := range fork.Precompiles {
		switch *precompileType {
		case params.PrecompileBase64:
			precompiles[addr] = pcbase64.NewBase64()
		case params.PrecompileCompress:
			precompiles[addr] = pccompress.NewCompress()
		case params.PrecompileJsonUtil:
			precompiles[addr] = pcjsonutil.NewJsonUtil()
		case params.PrecompileJsonStore:
			precompiles[addr] = pcjsonstore.NewJsonStore()
		case params.PrecompileNativeMinter:
			precompiles[addr] = pcnativeminter.NewNativeMinter()
		case params.PrecompileStrings:
			precompiles[addr] = pcstrings.NewStrings()
		case params.PrecompileIntegers:
			precompiles[addr] = pcintegers.NewIntegers()
		default:
			log.Error("Unknown precompile type", "type", precompileType, "address", addr)
		}
	}

	// Cache the initialized precompiles
	precompileCache[fork.Name] = precompiles
	return precompiles
}
