package precompile

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Gas costs
const (
	GasFree      uint64 = 0
	GasQuickStep uint64 = 2
	GasFastStep  uint64 = 40
	GasMidStep   uint64 = 100
	GasSlowStep  uint64 = 400
	GasExtStep   uint64 = 2000
)

type statefulPrecompiledContract struct {
	abi abi.ABI
}

func NewStatefulPrecompiledContract(abiStr string) StatefulPrecompiledContract {
	abi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		panic(err)
	}
	return &statefulPrecompiledContract{
		abi: abi,
	}
}

func (spc *statefulPrecompiledContract) GetABI() abi.ABI {
	return spc.abi
}

func (spc *statefulPrecompiledContract) RequiredGas(input []byte) uint64 {
	// This is a placeholder implementation. The actual gas required would depend on the specific contract.
	// You should replace this with the actual implementation.
	return GasFree
}
