// Copyright 2022, Specular contributors
//
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
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/L1/fraud-proof/AssertionMap.sol/AssertionMap.json --pkg bindings --type AssertionMap --out AssertionMap.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/L1/fraud-proof/challenge/IChallenge.sol/IChallenge.json --pkg bindings --type IChallenge --out IChallenge.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/L1/fraud-proof/IRollup.sol/IRollup.json --pkg bindings --type IRollup --out IRollup.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../packages/contracts/abi/contracts/L1/fraud-proof/Rollup.sol/Rollup.json --pkg bindings --type Rollup --out Rollup.go
