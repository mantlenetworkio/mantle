package genesis

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mantlenetworkio/mantle/mt-chain-ops/state"
)

func PrintAlloc(genesis *core.Genesis) {

	fmt.Printf("========================================================\n")
	for address, account := range genesis.Alloc {
		str, err := json.Marshal(account)
		if err == nil {
			fmt.Printf("%s : %s \n", address, str)
		}
	}
	fmt.Printf("========================================================\n")

}

// BuildL2DeveloperGenesis will build the developer Mantle Genesis
// Block. Suitable for devnets.
func BuildL2DeveloperGenesis(config *DeployConfig, l1StartBlock *types.Block) (*core.Genesis, error) {
	genspec, err := NewL2Genesis(config, l1StartBlock)

	if err != nil {
		return nil, err
	}
	db := state.NewMemoryStateDB(genspec)

	if config.FundDevAccounts {
		FundDevAccounts(db)
	}
	SetPrecompileBalances(db)

	storage, err := NewL2StorageConfig(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	immutable, err := NewL2ImmutableConfig(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	if err := SetL2Proxies(db); err != nil {
		return nil, err
	}

	if err := SetImplementations(db, storage, immutable); err != nil {
		return nil, err
	}

	if err := SetDevOnlyL2Implementations(db, storage, immutable); err != nil {
		return nil, err
	}

	return db.Genesis(), nil
}
