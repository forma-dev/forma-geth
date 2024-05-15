package jsonstore

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/abi"
	"github.com/ethereum/go-ethereum/precompile/contracts/compress"
	"github.com/ethereum/go-ethereum/precompile/contracts/jsonutil"
)

type JsonStore struct {
	precompile.StatefulPrecompiledContract

	compress *compress.Compress
	jsonutil *jsonutil.JsonUtil
}

const (
	PrepaidStorageGas   uint64 = 7500 // per 32 bytes
	PrepaidStorageDAGas uint64 = 2500
)

// keccak256(abi.encode(uint256(keccak256("forma.storage.jsonstore.PrepaidStorage")) - 1)) & ~bytes32(uint256(0xff))
var PrepaidStorageKey = common.HexToHash("0xf05e12b1d81b8c3dbb2d4eed32a4c02bb640b0f2adc116ef1449d9fcd8d33400")

func NewJsonStore() *JsonStore {
	return &JsonStore{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			abi.JsonStoreABI,
		),
		compress: compress.NewCompress(),
		jsonutil: jsonutil.NewJsonUtil(),
	}
}

// Solidity: function exists(bytes32 _slot) view returns(bool)
func (c *JsonStore) Exists(ctx precompile.StatefulContext, slot [32]byte) (bool, error) {
	storageVal := ctx.GetState(c.getStorageKey(ctx.MsgSender(), slot))
	return storageVal != (common.Hash{}), nil
}

// Solidity: function exists(address _key, bytes32 _slot) view returns(bool)
func (c *JsonStore) Exists0(ctx precompile.StatefulContext, key common.Address, slot [32]byte) (bool, error) {
	storageVal := ctx.GetState(c.getStorageKey(key, slot))
	return storageVal != (common.Hash{}), nil
}

// Solidity: function get(address _key, bytes32 _slot) view returns(string)
func (c *JsonStore) Get(ctx precompile.StatefulContext, key common.Address, slot [32]byte) (string, error) {
	return c.retrieveStorage(ctx, key, slot)
}

// Solidity: function get(bytes32 _slot) view returns(string)
func (c *JsonStore) Get0(ctx precompile.StatefulContext, slot [32]byte) (string, error) {
	return c.retrieveStorage(ctx, ctx.MsgSender(), slot)
}

// Solidity: function uri(bytes32 _slot) view returns(string)
func (c *JsonStore) Uri(ctx precompile.StatefulContext, slot [32]byte) (string, error) {
	return c.Uri0(ctx, ctx.MsgSender(), slot)
}

// Solidity: function uri(address _key, bytes32 _slot) view returns(string)
func (c *JsonStore) Uri0(ctx precompile.StatefulContext, key common.Address, slot [32]byte) (string, error) {
	jsonBlob, err := c.retrieveStorage(ctx, key, slot)
	if err != nil {
		return "", err
	}
	if jsonBlob == "" {
		return "", nil
	}
	return c.jsonutil.DataURI(ctx, jsonBlob)
}

// Solidity: function prepaid() view returns(uint64)
func (c *JsonStore) Prepaid(ctx precompile.StatefulContext) (uint64, error) {
	return c.getPrepaidStorage(ctx, ctx.MsgSender()), nil
}

// Solidity: function prepaid(address _key) view returns(uint64)
func (c *JsonStore) Prepaid0(ctx precompile.StatefulContext, _key common.Address) (uint64, error) {
	return c.getPrepaidStorage(ctx, _key), nil
}

// Solidity: function clear(bytes32 _slot) returns(bool)
func (c *JsonStore) Clear(ctx precompile.StatefulContext, slot [32]byte) (bool, error) {
	_, existingKeys := c.getStorageKeys(ctx, ctx.MsgSender(), slot)
	for k := range existingKeys {
		ctx.SetState(k, common.Hash{})
	}

	c.setPrepaidStorage(ctx, ctx.MsgSender(), c.getPrepaidStorage(ctx, ctx.MsgSender())+uint64(len(existingKeys)))
	return true, nil
}

// Solidity: function set(bytes32 _slot, string _jsonBlob) returns(bool)
func (c *JsonStore) Set(ctx precompile.StatefulContext, slot [32]byte, jsonBlob string) (bool, error) {
	return c.updateStorage(ctx, ctx.MsgSender(), slot, jsonBlob)
}

// Solidity: function prepay(uint64 _numSlots) returns(bool)
func (c *JsonStore) Prepay(ctx precompile.StatefulContext, numSlots uint64) (bool, error) {
	c.setPrepaidStorage(ctx, ctx.MsgSender(), c.getPrepaidStorage(ctx, ctx.MsgSender())+numSlots)
	return true, nil
}

// gas calcs
func (c *JsonStore) ExistsRequiredGas(ctx precompile.StatefulContext, slot [32]byte) uint64 {
	return precompile.GasQuickStep
}

func (c *JsonStore) Exists0RequiredGas(ctx precompile.StatefulContext, key common.Address, slot [32]byte) uint64 {
	return precompile.GasQuickStep
}

func (c *JsonStore) GetRequiredGas(ctx precompile.StatefulContext, key common.Address, slot [32]byte) uint64 {
	return precompile.GasSlowStep
}

func (c *JsonStore) Get0RequiredGas(ctx precompile.StatefulContext, slot [32]byte) uint64 {
	return precompile.GasSlowStep
}

func (c *JsonStore) UriRequiredGas(ctx precompile.StatefulContext, slot [32]byte) uint64 {
	return precompile.GasSlowStep
}

func (c *JsonStore) Uri0RequiredGas(ctx precompile.StatefulContext, key common.Address, slot [32]byte) uint64 {
	return precompile.GasSlowStep
}

func (c *JsonStore) PrepaidRequiredGas(ctx precompile.StatefulContext) uint64 {
	return precompile.GasQuickStep
}

func (c *JsonStore) Prepaid0RequiredGas(ctx precompile.StatefulContext, _key common.Address) uint64 {
	return precompile.GasQuickStep
}

func (c *JsonStore) ClearRequiredGas(ctx precompile.StatefulContext, slot [32]byte) uint64 {
	return precompile.GasSlowStep
}

func (c *JsonStore) SetRequiredGas(ctx precompile.StatefulContext, slot [32]byte, jsonBlob string) uint64 {
	prepaidStorageSlots := c.getPrepaidStorage(ctx, ctx.MsgSender())

	// compact json
	compactJson, _ := c.jsonutil.Compact(ctx, jsonBlob)

	// compress for storage
	jsonBytes, _ := c.compress.Compress(ctx, []byte(compactJson))

	existingKeys, _ := c.getStorageKeys(ctx, ctx.MsgSender(), slot)
	currentSlots := uint64(len(existingKeys))
	reqSlots := uint64(len(c.buildStorageMap(ctx.MsgSender(), slot, jsonBytes)))

	// no new slots required
	if currentSlots == reqSlots {
		return precompile.GasFree
	}

	daGas := PrepaidStorageDAGas * reqSlots

	if currentSlots < reqSlots {
		unpaidSlots := reqSlots - currentSlots
		if prepaidStorageSlots > 0 {
			if unpaidSlots > prepaidStorageSlots {
				unpaidSlots -= prepaidStorageSlots
				c.setPrepaidStorage(ctx, ctx.MsgSender(), 0)
			} else {
				c.setPrepaidStorage(ctx, ctx.MsgSender(), prepaidStorageSlots-unpaidSlots)
				return daGas
			}
		}
		return (unpaidSlots * PrepaidStorageGas) + daGas
	} else {
		c.setPrepaidStorage(ctx, ctx.MsgSender(), currentSlots-reqSlots)
		return daGas
	}
}

func (c *JsonStore) PrepayRequiredGas(ctx precompile.StatefulContext, numSlots uint64) uint64 {
	return numSlots * PrepaidStorageGas
}

func (c *JsonStore) getPrepaidStorage(ctx precompile.StatefulContext, key common.Address) uint64 {
	return new(big.Int).SetBytes(ctx.GetState(c.getPrepaidStorageKey(key)).Bytes()).Uint64()
}

func (c *JsonStore) setPrepaidStorage(ctx precompile.StatefulContext, key common.Address, val uint64) {
	ctx.SetState(c.getPrepaidStorageKey(key), common.BigToHash(new(big.Int).SetUint64(val)))
}

func (c *JsonStore) getPrepaidStorageKey(key common.Address) common.Hash {
	return crypto.Keccak256Hash(
		PrepaidStorageKey.Bytes(),
		common.BytesToHash(key.Bytes()).Bytes(),
	)
}

func (c *JsonStore) getStorageKey(key common.Address, slot [32]byte) common.Hash {
	return crypto.Keccak256Hash(
		common.BytesToHash(key.Bytes()).Bytes(),
		slot[:],
	)
}

func (c *JsonStore) retrieveStorage(ctx precompile.StatefulContext, key common.Address, slot [32]byte) (string, error) {
	keys, readLenths := c.getStorageKeys(ctx, key, slot)
	if len(keys) == 0 {
		return "", nil
	}

	var storeBytes []byte
	for _, key := range keys {
		if readLenths[key] > 0 {
			keyBytes := ctx.GetState(key).Bytes()
			storeBytes = append(storeBytes, keyBytes[:readLenths[key]]...)
		}
	}

	res, err := c.compress.Decompress(ctx, storeBytes)
	if err != nil {
		return "", err
	}

	strRes := string(res)
	if strRes == "" {
		strRes = `{}`
	}

	return strRes, nil
}

func (c *JsonStore) updateStorage(ctx precompile.StatefulContext, key common.Address, slot [32]byte, jsonBlob string) (bool, error) {
	_, validationErr := c.jsonutil.Validate(ctx, jsonBlob)
	if validationErr != nil {
		return false, errors.New("invalid JSON")
	}

	// compact json
	compactJson, err := c.jsonutil.Compact(ctx, jsonBlob)
	if err != nil {
		return false, err
	}

	// compress for storage
	jsonBytes, err := c.compress.Compress(ctx, []byte(compactJson))
	if err != nil {
		return false, err
	}

	_, existingKeys := c.getStorageKeys(ctx, key, slot)
	newStorageMap := c.buildStorageMap(key, slot, jsonBytes)

	for k, v := range newStorageMap {
		ctx.SetState(k, v)
		delete(existingKeys, k)
	}

	// purge data from remaining existingKeys
	for k := range existingKeys {
		ctx.SetState(k, common.Hash{})
	}

	return true, nil
}

func (c *JsonStore) getStorageKeys(ctx precompile.StatefulContext, key common.Address, slot [32]byte) ([]common.Hash, map[common.Hash]int) {
	var order []common.Hash
	storage := make(map[common.Hash]int)
	storageKey := c.getStorageKey(key, slot)
	storageVal := ctx.GetState(storageKey).Bytes()

	if storageVal[31]&1 == 0 {
		length := int(storageVal[31] / 2)
		if length > 0 {
			order = append(order, storageKey)
			storage[storageKey] = length
		}
	} else {
		length := (new(big.Int).SetBytes(storageVal).Uint64() - 1) / 2
		order = append(order, storageKey)
		storage[storageKey] = 0
		chunkKeyInt := crypto.Keccak256Hash(storageKey.Bytes()).Big()
		for i := uint64(0); i < length; i += 32 {
			chunkKey := common.BigToHash(chunkKeyInt)
			storage[chunkKey] = 32

			end := i + 32
			if end > length {
				storage[chunkKey] = 32 - int(end-length)
			}

			order = append(order, chunkKey)

			chunkKeyInt.Add(chunkKeyInt, big.NewInt(1))
		}
	}

	return order, storage
}

func (c *JsonStore) buildStorageMap(key common.Address, slot [32]byte, data []byte) map[common.Hash]common.Hash {
	storage := make(map[common.Hash]common.Hash)
	storageKey := c.getStorageKey(key, slot)

	length := len(data)
	if length < 32 {
		store := make([]byte, 32)
		copy(store[0:length], data[:])
		store[31] = byte(length * 2)
		storage[storageKey] = common.BytesToHash(store)
	} else {
		storage[storageKey] = common.BigToHash(new(big.Int).SetUint64(uint64(length*2) + 1))
		chunkKeyInt := crypto.Keccak256Hash(storageKey.Bytes()).Big()
		for i := 0; i < length; i += 32 {
			end := i + 32
			if end > length {
				end = length
			}

			chunk := make([]byte, 32)
			copy(chunk, data[i:end])

			chunkKey := common.BigToHash(chunkKeyInt)
			storage[chunkKey] = common.BytesToHash(chunk)

			chunkKeyInt.Add(chunkKeyInt, big.NewInt(1))
		}
	}

	return storage
}
