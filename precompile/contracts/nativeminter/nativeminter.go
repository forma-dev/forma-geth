package nativeminter

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/bindings"
)

type NativeMinter struct {
	precompile.StatefulPrecompiledContract
	initialOwner common.Address
}

func NewNativeMinter(initialOwner common.Address) *NativeMinter {
	return &NativeMinter{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			bindings.NativeMinterABI,
		),
		initialOwner: initialOwner,
	}
}

// StorageSlots is a struct that holds the storage slots for the HypNative contract
type StorageSlots struct {
	Owner          common.Hash
	Minter         common.Hash
	OwnerRenounced common.Hash
}

// Slots is a global variable that holds the storage slots for the NativeMinter contract
var Slots = StorageSlots{
	Owner:          common.BytesToHash([]byte{0x00}), // slot 0
	Minter:         common.BytesToHash([]byte{0x01}), // slot 1
	OwnerRenounced: common.BytesToHash([]byte{0x02}), // slot 2
}

var ZeroAddress = common.Address{}

// Owner returns the owner of the NativeMinter contract
func (c *NativeMinter) Owner(ctx precompile.StatefulContext) (common.Address, error) {
	return common.BytesToAddress(ctx.GetState(Slots.Owner).Bytes()), nil
}

// Minter returns the minter of the NativeMinter contract
func (c *NativeMinter) Minter(ctx precompile.StatefulContext) (common.Address, error) {
	return common.BytesToAddress(ctx.GetState(Slots.Minter).Bytes()), nil
}

// SetMinter sets a new minter for the HypNative contract
func (c *NativeMinter) SetMinter(ctx precompile.StatefulContext, newMinter common.Address) (bool, error) {
	if err := c.senderIsOwner(ctx); err != nil {
		return false, err
	}

	ctx.SetState(Slots.Minter, common.BytesToHash(newMinter.Bytes()))

	return true, nil
}

// TransferOwnership transfers the ownership of the HypNative contract to a new owner
func (c *NativeMinter) TransferOwnership(ctx precompile.StatefulContext, newOwner common.Address) (bool, error) {
	if bytes.Equal(newOwner.Bytes(), ZeroAddress.Bytes()) {
		return false, errors.New("new owner is the zero address")
	}

	return c.internalTransferOwnership(ctx, newOwner)
}

// RenounceOwnership renounces the ownership of the HypNative contract
func (c *NativeMinter) RenounceOwnership(ctx precompile.StatefulContext) (bool, error) {
	ctx.SetState(Slots.OwnerRenounced, common.BytesToHash([]byte{0x01}))
	return c.internalTransferOwnership(ctx, ZeroAddress)
}

func (c *NativeMinter) internalTransferOwnership(ctx precompile.StatefulContext, newOwner common.Address) (bool, error) {
	if err := c.senderIsOwner(ctx); err != nil {
		return false, err
	}

	ctx.SetState(Slots.Owner, common.BytesToHash(newOwner.Bytes()))

	return true, nil
}

// Mint new tokens to the specified address
func (c *NativeMinter) Mint(ctx precompile.StatefulContext, addr common.Address, amount *big.Int) (bool, error) {
	if err := c.senderIsMinter(ctx); err != nil {
		return false, err
	}

	ctx.AddBalance(addr, amount)

	return true, nil
}

// Burn tokens from the specified address
func (c *NativeMinter) Burn(ctx precompile.StatefulContext, addr common.Address, amount *big.Int) (bool, error) {
	if err := c.senderIsMinter(ctx); err != nil {
		return false, err
	}

	// check balance is valid
	balance := ctx.GetBalance(addr)
	if balance.Cmp(amount) < 0 {
		return false, errors.New("insufficient balance for burn")
	}

	ctx.SubBalance(addr, amount)

	return true, nil
}

// Access control funcs
func (c *NativeMinter) senderIsOwner(ctx precompile.StatefulContext) error {
	owner := common.BytesToAddress(ctx.GetState(Slots.Owner).Bytes())
	ownerRenounced := ctx.GetState(Slots.OwnerRenounced)

	// owner not explicity set and not renounced; use initial owner
	if owner.Cmp(ZeroAddress) == 0 && bytes.Equal(ownerRenounced.Bytes(), common.BytesToHash([]byte{0x00}).Bytes()) {
		if owner.Cmp(c.initialOwner) != 0 {
			return errors.New("caller is not the owner")
		}
		return nil
	}

	if owner.Cmp(ctx.MsgSender()) != 0 {
		return errors.New("caller is not the owner")
	}
	return nil
}

func (c *NativeMinter) senderIsMinter(ctx precompile.StatefulContext) error {
	allowedMinter := common.BytesToAddress(ctx.GetState(Slots.Minter).Bytes())
	if allowedMinter.Cmp(ctx.MsgSender()) != 0 {
		return errors.New("caller is not the minter")
	}
	return nil
}
