// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

/**
 * @title Small library for checking Merkle proofs.
 * @author Original authorship of this code is unclear. This implementation is adapted from Polygon's.
 * See https://github.com/maticnetwork/contracts/commits/main/contracts/common/lib/Merkle.sol
 * with added functions for sha256 support for Ethereum Consensus Layer Merkleization.
 */
library Merkle {
    /**
     @notice this function checks whether the given @param leaf is actually a member (leaf) of the 
             merkle tree with @param rootHash being the Merkle root or not.   
     @param leaf is the element whose membership in the merkle tree is being checked,
     @param index is the leaf index
     @param rootHash is the Merkle root of the Merkle tree,
     @param proof is the Merkle proof associated with the @param leaf and @param rootHash.
     */ 
    function checkMembership(
        bytes32 leaf,
        uint256 index,
        bytes32 rootHash,
        bytes memory proof
    ) internal pure returns (bool) {

        require(proof.length % 32 == 0, "Invalid proof length");

        /**
         Merkle proof consists of all siblings along the path to the Merkle root, each of 32 bytes
         */
        uint256 proofHeight = proof.length / 32;

        /**
          Proof of size n means, height of the tree is n+1.
          In a tree of height n+1, max #leafs possible is 2**n.
         */
        require(index < 2**proofHeight, "Leaf index is too big");

        bytes32 proofElement;

        // starting from the leaf
        bytes32 computedHash = leaf;

        // going up the Merkle tree
        for (uint256 i = 32; i <= proof.length; i += 32) {

            // retrieve the sibling along the merkle proof
            assembly {
                proofElement := mload(add(proof, i))
            }


            /**
             check whether the association with the parent is of the format:

                computedHash of Parent                    computedHash of Parent        
                             *                                      *
                           *   *                or                *   *
                         *       *                              *       *
                       *           *                          *           * 
                computedHash    proofElement            proofElement   computedHash
                             
             */
            // association is of first type
            if (index % 2 == 0) {
                computedHash = keccak256(
                    abi.encodePacked(computedHash, proofElement)
                );

            // association is of second type
            } else {
                computedHash = keccak256(
                    abi.encodePacked(proofElement, computedHash)
                );
            }

            index = index / 2;
        }

        // check whether computed root is same as the Merkle root
        return computedHash == rootHash;
    }

    /**
     @notice this function returns the merkle root of a tree created from a set of leaves using sha256 as its hash function
     @param leaves the leaves of the merkle tree

     @notice requires the leaves.length is a power of 2
     */ 
    function merkleizeSha256(
        bytes32[] memory leaves
    ) internal pure returns (bytes32) {
        //there are half as many nodes in the layer above the leaves
        uint256 numNodesInLayer = leaves.length / 2;
        //create a layer to store the internal nodes
        bytes32[] memory layer = new bytes32[](numNodesInLayer);
        //fill the layer with the pairwise hashes of the leaves
        for (uint i = 0; i < numNodesInLayer; i++) {
            layer[i] = sha256(abi.encodePacked(leaves[2*i], leaves[2*i+1]));
        }
        //the next layer above has half as many nodes
        numNodesInLayer /= 2;
        //while we haven't computed the root
        while (numNodesInLayer != 0) {
            //overwrite the first numNodesInLayer nodes in layer with the pairwise hashes of their children
            for (uint i = 0; i < numNodesInLayer; i++) {
                layer[i] = sha256(abi.encodePacked(layer[2*i], layer[2*i+1]));
            }
            //the next layer above has half as many nodes
            numNodesInLayer /= 2;
        }
        //the first node in the layer is the root
        return layer[0];
    }

    function checkMembershipSha256(
        bytes32 leaf,
        uint256 index,
        bytes32 rootHash,
        bytes memory proof
    ) internal pure returns (bool) {
        require(proof.length % 32 == 0, "Invalid proof length");
        uint256 proofHeight = proof.length / 32;
        // Proof of size n means, height of the tree is n+1.
        // In a tree of height n+1, max #leafs possible is 2 ^ n
        require(index < 2 ** proofHeight, "Leaf index is too big");

        bytes32 proofElement;
        bytes32 computedHash = leaf;
        for (uint256 i = 32; i <= proof.length; i += 32) {
            assembly {
                proofElement := mload(add(proof, i))
            }

            if (index % 2 == 0) {
                computedHash = sha256(
                    abi.encodePacked(computedHash, proofElement)
                );
            } else {
                computedHash = sha256(
                    abi.encodePacked(proofElement, computedHash)
                );
            }

            index = index / 2;
        }
        return computedHash == rootHash;
    }
}