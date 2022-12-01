// SPDX-License-Identifier: UNLICENSED AND MIT
// several functions are from https://github.com/ChihChengLiang/bls_solidity_python/blob/master/contracts/BLS.sol (MIT license)
// remainder is UNLICENSED
pragma solidity ^0.8.9;

/**
 * @title Library for operations related to BLS Signatures used in EigenLayer middleware.
 * @author Layr Labs, Inc. with credit to Chih Cheng Liang
 * @notice Uses the BN254 curve.
 */
library BLS {
    // BN 254 CONSTANTS
    // modulus for the underlying field F_q of the elliptic curve
    uint256 internal constant MODULUS = 21888242871839275222246405745257275088696311157297823662689037894645226208583;

    // negation of the generator of group G2
    /**
     * @dev Generator point lies in F_q2 is of the form: (x0 + ix1, y0 + iy1).
     */
    uint256 internal constant nG2x1 = 11559732032986387107991004021392285783925812861821192530917403151452391805634;
    uint256 internal constant nG2x0 = 10857046999023057135944570762232829481370756359578518086990519993285655852781;
    uint256 internal constant nG2y1 = 17805874995975841540914202342111839520379459829704422454583296818431106115052;
    uint256 internal constant nG2y0 = 13392588948715843804641432497768002650278120570034223513918757245338268106653;

    // generator of group G2
    /**
     * @dev Generator point lies in F_q2 is of the form: (x0 + ix1, y0 + iy1).
     */
    uint256 internal constant G2x1 = 11559732032986387107991004021392285783925812861821192530917403151452391805634;
    uint256 internal constant G2x0 = 10857046999023057135944570762232829481370756359578518086990519993285655852781;
    uint256 internal constant G2y1 = 4082367875863433681332203403145435568316851327593401208105741076214120093531;
    uint256 internal constant G2y0 = 8495653923123431417604973247489272438418190587263600148770280649306958101930;

    bytes32 internal constant powersOfTauMerkleRoot = 0x22c998e49752bbb1918ba87d6d59dd0e83620a311ba91dd4b2cc84990b31b56f;

    function hashPubkey(uint256[4] memory pk) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(pk[0], pk[1], pk[2], pk[3]));
    }

    function hashPubkey(uint256[6] memory pk) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(pk[0], pk[1], pk[2], pk[3]));
    }

    /**
     * @notice verification of BLS signature with the message being pubkey hash
     */
    /**
     * @dev first paramater, data, is the calldata that contains the coordinates for pubkey on G2 and signature on G1
     * @return pubkey is the pubkey and is of the format [x1, x0, y1, y0]
     */
    function verifyBLSSigOfPubKeyHash(bytes calldata data, address operator)
        internal
        view
        returns (uint256, uint256, uint256, uint256)
    {
        // uint256 offset = 68;
        // e(H(m), pk)e(sigma, -g2) = e(H(m), pk)(e(sigma, g2)^-1) == 1?
        // is the same as
        // e(H(m), pk) == e(sigma, g2)?
        uint256[12] memory input;

        assembly {
            //store pk in indexes 2-5, it is a G2 point
            mstore(add(input, 0x40), calldataload(data.offset))
            mstore(add(input, 0x60), calldataload(add(data.offset, 32)))
            mstore(add(input, 0x80), calldataload(add(data.offset, 64)))
            mstore(add(input, 0xA0), calldataload(add(data.offset, 96)))
            //store sigma (signature) in indexes 6-7, it is a G1 point
            mstore(add(input, 0xC0), calldataload(add(data.offset, 128)))
            mstore(add(input, 0xE0), calldataload(add(data.offset, 160)))
            //store the negated G2 generator in indexes 8-11
            mstore(add(input, 0x100), nG2x1)
            mstore(add(input, 0x120), nG2x0)
            mstore(add(input, 0x140), nG2y1)
            mstore(add(input, 0x160), nG2y0)
        }

        // calculate H(m) = H(pk)
        (input[0], input[1]) = hashToG1(keccak256(abi.encodePacked(input[2], input[3], input[4], input[5], operator)));

        assembly {
            // check the pairing
            if iszero(
                // call ecPairing precompile with 384 bytes of data,
                // i.e. input[0] through (including) input[11], and get 32 bytes of return data
                staticcall(not(0), 0x08, input, 0x180, add(input, 0x20), 0x20)
            ) { revert(0, 0) }
        }

        require(input[1] == 1, "Pairing was unsuccessful");

        // return pubkey, the format being [x1, x0, y1, y0]
        return (input[3], input[2], input[5], input[4]);
    }

    /**
     * @notice same function as AddAssign in https://github.com/ConsenSys/gnark-crypto/blob/master/ecc/bn254/g2.go
     */
    function addJac(uint256[6] memory jac1, uint256[6] memory jac2) internal pure returns (uint256[6] memory) {
        //NOTE: JAC IS REFERRED TO AS X, Y, Z
        //ALL 2 ELEMENTS EACH
        // var XX, YY, YYYY, ZZ, S, M, T fptower.E2

        if (jac1[4] == 0 && jac1[5] == 0) {
            // on point 1 being a point at infinity
            return jac2;
        } else if (jac2[4] == 0 && jac2[5] == 0) {
            // on point 2 being a point at infinity
            return jac1;
        }

        // var Z1Z1, Z2Z2, U1, U2, S1, S2, H, I, J, r, V fptower.E2
        //z1z1 = a.z^2
        uint256[4] memory z1z1z2z2;
        (z1z1z2z2[0], z1z1z2z2[1]) = square(jac2[4], jac2[5]);
        //z2z2 = p.z^2
        // uint256[2] memory z2z2;
        (z1z1z2z2[2], z1z1z2z2[3]) = square(jac1[4], jac1[5]);
        //u1 = a.x*z2z2
        uint256[4] memory u1u2;
        (u1u2[0], u1u2[1]) = mul(jac2[0], jac2[1], z1z1z2z2[2], z1z1z2z2[3]);
        //u2 = p.x*z1z1
        // uint256[2] memory u2;
        (u1u2[2], u1u2[3]) = mul(jac1[0], jac1[1], z1z1z2z2[0], z1z1z2z2[1]);
        //s1 = a.y*p.z*z2z2
        uint256[2] memory s1;
        (s1[0], s1[1]) = mul(jac2[2], jac2[3], jac1[4], jac1[5]);
        (s1[0], s1[1]) = mul(s1[0], s1[1], z1z1z2z2[2], z1z1z2z2[3]);

        //s2 = p.y*a.z*z1z1
        uint256[2] memory s2;
        (s2[0], s2[1]) = mul(jac1[2], jac1[3], jac2[4], jac2[5]);
        (s2[0], s2[1]) = mul(s2[0], s2[1], z1z1z2z2[0], z1z1z2z2[1]);

        // // if p == a, we double instead, is this too inefficient?
        // // if (u1[0] == 0 && u1[1] == 0 && u2[0] == 0 && u2[1] == 0) {
        // //     return p.DoubleAssign()
        // // } else {

        // // }

        uint256[2] memory h;
        uint256[2] memory i;

        assembly {
            //h = u2 - u1
            mstore(h, addmod(mload(add(u1u2, 0x040)), sub(MODULUS, mload(u1u2)), MODULUS))
            mstore(add(h, 0x20), addmod(mload(add(u1u2, 0x60)), sub(MODULUS, mload(add(u1u2, 0x20))), MODULUS))

            //i = 2h
            mstore(i, mulmod(mload(h), 2, MODULUS))
            mstore(add(i, 0x20), mulmod(mload(add(h, 0x20)), 2, MODULUS))
        }

        (i[0], i[1]) = square(i[0], i[1]);

        uint256[2] memory j;
        (j[0], j[1]) = mul(h[0], h[1], i[0], i[1]);

        uint256[2] memory r;
        assembly {
            //r = s2 - s1
            mstore(r, addmod(mload(s2), sub(MODULUS, mload(s1)), MODULUS))
            mstore(add(r, 0x20), addmod(mload(add(s2, 0x20)), sub(MODULUS, mload(add(s1, 0x20))), MODULUS))

            //r *= 2
            mstore(r, mulmod(mload(r), 2, MODULUS))
            mstore(add(r, 0x20), mulmod(mload(add(r, 0x20)), 2, MODULUS))
        }

        uint256[2] memory v;
        (v[0], v[1]) = mul(u1u2[0], u1u2[1], i[0], i[1]);

        (jac1[0], jac1[1]) = square(r[0], r[1]);

        assembly {
            //x -= j
            mstore(jac1, addmod(mload(jac1), sub(MODULUS, mload(j)), MODULUS))
            mstore(add(jac1, 0x20), addmod(mload(add(jac1, 0x20)), sub(MODULUS, mload(add(j, 0x20))), MODULUS))
            //x -= v
            mstore(jac1, addmod(mload(jac1), sub(MODULUS, mload(v)), MODULUS))
            mstore(add(jac1, 0x20), addmod(mload(add(jac1, 0x20)), sub(MODULUS, mload(add(v, 0x20))), MODULUS))
            //x -= v
            mstore(jac1, addmod(mload(jac1), sub(MODULUS, mload(v)), MODULUS))
            mstore(add(jac1, 0x20), addmod(mload(add(jac1, 0x20)), sub(MODULUS, mload(add(v, 0x20))), MODULUS))
            //y = v - x
            mstore(add(jac1, 0x40), addmod(mload(v), sub(MODULUS, mload(jac1)), MODULUS))
            mstore(add(jac1, 0x60), addmod(mload(add(v, 0x20)), sub(MODULUS, mload(add(jac1, 0x20))), MODULUS))
        }

        (jac1[2], jac1[3]) = mul(jac1[2], jac1[3], r[0], r[1]);
        (s1[0], s1[1]) = mul(s1[0], s1[1], j[0], j[1]);

        assembly {
            //s1 *= 2
            mstore(s1, mulmod(mload(s1), 2, MODULUS))
            mstore(add(s1, 0x20), mulmod(mload(add(s1, 0x20)), 2, MODULUS))
            //y -= s1
            mstore(add(jac1, 0x40), addmod(mload(add(jac1, 0x40)), sub(MODULUS, mload(s1)), MODULUS))
            mstore(add(jac1, 0x60), addmod(mload(add(jac1, 0x60)), sub(MODULUS, mload(add(s1, 0x20))), MODULUS))
            //z = a.z + p.z
            mstore(add(jac1, 0x80), addmod(mload(add(jac1, 0x80)), mload(add(jac2, 0x80)), MODULUS))
            mstore(add(jac1, 0xA0), addmod(mload(add(jac1, 0xA0)), mload(add(jac2, 0xA0)), MODULUS))
        }

        (jac1[4], jac1[5]) = square(jac1[4], jac1[5]);

        assembly {
            //z -= z1z1
            mstore(add(jac1, 0x80), addmod(mload(add(jac1, 0x80)), sub(MODULUS, mload(z1z1z2z2)), MODULUS))
            mstore(add(jac1, 0xA0), addmod(mload(add(jac1, 0xA0)), sub(MODULUS, mload(add(z1z1z2z2, 0x20))), MODULUS))
            //z -= z2z2
            mstore(add(jac1, 0x80), addmod(mload(add(jac1, 0x80)), sub(MODULUS, mload(add(z1z1z2z2, 0x40))), MODULUS))
            mstore(add(jac1, 0xA0), addmod(mload(add(jac1, 0xA0)), sub(MODULUS, mload(add(z1z1z2z2, 0x60))), MODULUS))
        }

        (jac1[4], jac1[5]) = mul(jac1[4], jac1[5], h[0], h[1]);

        return jac1;
    }

    /**
     * @notice used for squaring a Fq2 element - (x0 + ix1)
     */
    function square(uint256 x0, uint256 x1) internal pure returns (uint256, uint256) {
        uint256[4] memory z;
        assembly {
            //a = x0 + x1
            mstore(z, addmod(x0, x1, MODULUS))
            //b = x0 - x1
            mstore(add(z, 0x20), addmod(x0, sub(MODULUS, x1), MODULUS))
            //a = (x0 + x1)(x0 - x1)
            mstore(add(z, 0x40), mulmod(mload(z), mload(add(z, 0x20)), MODULUS))
            //b = 2x0y0
            mstore(add(z, 0x60), mulmod(2, mulmod(x0, x1, MODULUS), MODULUS))
        }
        return (z[2], z[3]);
    }

    function jacToAff(uint256[6] memory jac) internal view returns (uint256, uint256, uint256, uint256) {
        if (jac[4] == 0 && jac[5] == 0) {
            return (uint256(0), uint256(0), uint256(0), uint256(0));
        }

        (jac[4], jac[5]) = inverse(jac[4], jac[5]);
        (uint256 b0, uint256 b1) = square(jac[4], jac[5]);
        (jac[0], jac[1]) = mul(jac[0], jac[1], b0, b1);
        (jac[2], jac[3]) = mul(jac[2], jac[3], b0, b1);
        (jac[2], jac[3]) = mul(jac[2], jac[3], jac[4], jac[5]);

        return (jac[0], jac[1], jac[2], jac[3]);
    }

    /**
     * @notice same function as Inverse in https://github.com/ConsenSys/gnark-crypto/blob/528300a94e8717cb98d124ebf7de96dddca373ea/ecc/bn254/internal/fptower/e2_bn254.go#L73
     */
    function inverse(uint256 x0, uint256 x1) internal view returns (uint256, uint256) {
        uint256[2] memory t;
        assembly {
            mstore(t, mulmod(x0, x0, MODULUS))
            mstore(add(t, 0x20), mulmod(x1, x1, MODULUS))
            mstore(t, addmod(mload(t), mload(add(t, 0x20)), MODULUS))

            let freemem := mload(0x40)
            mstore(freemem, 0x20)
            mstore(add(freemem, 0x20), 0x20)
            mstore(add(freemem, 0x40), 0x20)
            mstore(add(freemem, 0x60), mload(t))
            // x^(n-2) = x^-1 mod q
            mstore(add(freemem, 0x80), sub(MODULUS, 2))
            // N = 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47
            mstore(add(freemem, 0xA0), MODULUS)
            if iszero(staticcall(sub(gas(), 2000), 5, freemem, 0xC0, add(t, 0x20), 0x20)) { revert(0, 0) }
            mstore(t, mulmod(x0, mload(add(t, 0x20)), MODULUS))
            mstore(add(t, 0x20), mulmod(x1, mload(add(t, 0x20)), MODULUS))
        }

        return (t[0], (MODULUS - t[1]) % MODULUS);
    }

    /**
     * @notice used for multiplying two Fq2 elements - (x0 + ix1) and (y0 + iy1).
     */
    function mul(uint256 x0, uint256 x1, uint256 y0, uint256 y1) internal pure returns (uint256, uint256) {
        uint256[5] memory z;
        assembly {
            //a = x0 + x1
            mstore(z, addmod(x0, x1, MODULUS))
            //b = y0 + y1
            mstore(add(z, 0x20), addmod(y0, y1, MODULUS))
            //a = (x0 + x1)(y0 + y1)
            mstore(z, mulmod(mload(z), mload(add(z, 0x20)), MODULUS))
            //b = x0y0
            mstore(add(z, 0x20), mulmod(x0, y0, MODULUS))
            //c = x1y1
            mstore(add(z, 0x40), mulmod(x1, y1, MODULUS))
            //c = -x1y1
            mstore(add(z, 0x40), sub(MODULUS, mload(add(z, 0x40))))
            //z0 = x0y0 - x1y1
            mstore(add(z, 0x60), addmod(mload(add(z, 0x20)), mload(add(z, 0x40)), MODULUS))
            //b = -x0y0
            mstore(add(z, 0x20), sub(MODULUS, mload(add(z, 0x20))))
            //z1 = x0y1 + x1y0
            mstore(add(z, 0x80), addmod(addmod(mload(z), mload(add(z, 0x20)), MODULUS), mload(add(z, 0x40)), MODULUS))
        }
        return (z[3], z[4]);
    }

    /**
     * @notice same as hashToPoint function in https://github.com/ChihChengLiang/bls_solidity_python/blob/master/contracts/BLS.sol
     */
    function hashToG1(bytes32 _x) internal view returns (uint256 x, uint256 y) {
        x = uint256(_x) % MODULUS;
        bool found = false;
        while (true) {
            y = mulmod(x, x, MODULUS);
            y = mulmod(y, x, MODULUS);
            y = addmod(y, 3, MODULUS);
            (y, found) = sqrt(y);
            if (found) {
                return (x, y);
            }
            x = addmod(x, 1, MODULUS);
        }
    }

    function sqrt(uint256 xx) internal view returns (uint256 x, bool hasRoot) {
        bool callSuccess;
        assembly {
            let freemem := mload(0x40)
            mstore(freemem, 0x20)
            mstore(add(freemem, 0x20), 0x20)
            mstore(add(freemem, 0x40), 0x20)
            mstore(add(freemem, 0x60), xx)
            // (N + 1) / 4 = 0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52
            mstore(add(freemem, 0x80), 0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52)
            // N = 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47
            mstore(add(freemem, 0xA0), 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47)
            callSuccess := staticcall(sub(gas(), 2000), 5, freemem, 0xC0, freemem, 0x20)
            x := mload(freemem)
            hasRoot := eq(xx, mulmod(x, x, MODULUS))
        }
        require(callSuccess, "BLS: sqrt modexp call failed");
    }

    /**
     * @notice This function is for removing a pubkey from aggregated pubkey. The thesis of this operation:
     * - conversion to Jacobian coordinates,
     * - do the subtraction of pubkey from aggregated pubkey,
     * - convert the updated aggregated pubkey back to affine coordinates.
     * @param pubkeyToRemoveAff is the pubkey that is to be removed,
     * @param existingAggPubkeyAff is the aggregated pubkey.
     * @dev Jacobian coordinates are stored in the form [x0, x1, y0, y1, z0, z1]
     */
    function removePubkeyFromAggregate(uint256[4] memory pubkeyToRemoveAff, uint256[4] memory existingAggPubkeyAff)
        internal
        view
        returns (uint256, uint256, uint256, uint256)
    {
        uint256[6] memory pubkeyToRemoveJac;
        uint256[6] memory existingAggPubkeyJac;

        // get x0, x1, y0, y1 from affine coordinates
        for (uint256 i = 0; i < 4;) {
            pubkeyToRemoveJac[i] = pubkeyToRemoveAff[i];
            existingAggPubkeyJac[i] = existingAggPubkeyAff[i];
            unchecked {
                ++i;
            }
        }
        // set z0 = 1
        pubkeyToRemoveJac[4] = 1;
        existingAggPubkeyJac[4] = 1;

        /// @notice subtract pubkeyToRemoveJac from the aggregate pubkey
        // negate pubkeyToRemoveJac
        pubkeyToRemoveJac[2] = (MODULUS - pubkeyToRemoveJac[2]) % MODULUS;
        pubkeyToRemoveJac[3] = (MODULUS - pubkeyToRemoveJac[3]) % MODULUS;
        // add the negation to existingAggPubkeyJac
        addJac(existingAggPubkeyJac, pubkeyToRemoveJac);

        // 'addJac' function above modifies the first input in memory, so now we can just return it (but first transform it back to affine)
        return (jacToAff(existingAggPubkeyJac));
    }
}
