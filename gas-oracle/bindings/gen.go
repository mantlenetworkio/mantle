// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bindings

//go:generate ./compile.sh
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/L1/rollup/StateCommitmentChain.sol/StateCommitmentChain.json --pkg bindings --type StateCommitmentChain --out StateCommitmentChain.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/L1/rollup/CanonicalTransactionChain.sol/CanonicalTransactionChain.json --pkg bindings --type CanonicalTransactionChain --out CanonicalTransactionChain.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/L2/predeploys/BVM_GasPriceOracle.sol/BVM_GasPriceOracle.json --pkg bindings --type BVM_GasPriceOracle --out BVM_GasPriceOracle.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/da/BVM_EigenDataLayrFee.sol/BVM_EigenDataLayrFee.json --pkg bindings --type BVM_EigenDataLayrFee --out BVM_EigenDataLayrFee.go
