// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2022, Specular contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.8.0;

/// @dev Thrown when unauthorized (!rollup) address calls an only-rollup function
/// @param sender Address of the caller
/// @param rollup The rollup address authorized to call this function
error NotRollup(address sender, address rollup);

/// @dev Thrown when unauthorized (!challenge) address calls an only-challenge function
/// @param sender Address of the caller
/// @param challenge The challenge address authorized to call this function
error NotChallenge(address sender, address challenge);

/// @dev Thrown when unauthorized (!sequencer) address calls an only-sequencer function
/// @param sender Address of the caller
/// @param sequencer The sequencer address authorized to call this function
error NotSequencer(address sender, address sequencer);

/// @dev Thrown when function is called with a zero address argument
error ZeroAddress();

/// @dev Thrown when function is called with a zero address argument
error RedundantInitialized();
