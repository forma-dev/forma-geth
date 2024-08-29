package compress

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/mocks"
	"github.com/holiman/uint256"
)

var (
	zero = uint256.NewInt(0)

	SelfAddrForTest  = common.HexToAddress("0xSelfAddr")
	MsgSenderForTest = common.BytesToAddress([]byte("0xMsgSender"))
)

func TestCompress(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, MsgSenderForTest, zero)

	c := NewCompress()

	input := []byte("Hello, Forma! Hello, Forma! Hello, Forma!")

	// Test for Non-Error
	output, err := c.Compress(ctx, input)
	if err != nil {
		t.Fatalf("Compress failed: %v", err)
	}

	// Print the compressed output
	t.Logf("Compressed output: %v", output)

	// Test that the compressed data is smaller
	if len(output) >= len(input) {
		t.Errorf("Compression test failed. Compressed data is not smaller than the input data.")
	}

	// Test for Reversibility
	decompressedOutput, err := c.Decompress(ctx, output)
	if err != nil {
		t.Fatalf("Decompress failed: %v", err)
	}
	if !bytes.Equal(input, decompressedOutput) {
		t.Errorf("Reversibility test failed. Decompressed output differs from the original input.")
	}
}

func TestDecompress(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, MsgSenderForTest, zero)

	c := NewCompress()

	input := []byte{11, 6, 128, 72, 101, 108, 108, 111, 44, 32, 70, 111, 114, 109, 97, 33, 3}
	expectedOutput := []byte("Hello, Forma!")

	output, err := c.Decompress(ctx, input)
	if err != nil {
		t.Fatalf("Decompress failed: %v", err)
	}

	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("Decompress output mismatch. Got %v, expected %v", output, expectedOutput)
	}
}

func TestCompressRequiredGas(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, MsgSenderForTest, zero)

	c := NewCompress()

	input := []byte("Hello, Forma!")
	expectedGas := uint64(2)

	gas := c.CompressRequiredGas(ctx, input)
	if gas != expectedGas {
		t.Errorf("CompressRequiredGas mismatch. Got %v, expected %v", gas, expectedGas)
	}
}

func TestDecompressRequiredGas(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, MsgSenderForTest, zero)

	c := NewCompress()

	input := []byte{11, 6, 128, 72, 101, 108, 108, 111, 44, 32, 70, 111, 114, 109, 97, 33, 3}
	expectedGas := uint64(2) // Example expected gas value

	gas := c.DecompressRequiredGas(ctx, input)
	if gas != expectedGas {
		t.Errorf("DecompressRequiredGas mismatch. Got %v, expected %v", gas, expectedGas)
	}
}
