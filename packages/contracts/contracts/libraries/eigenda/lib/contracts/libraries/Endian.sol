// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.9;

library Endian {
    //copied from https://etherscan.io/address/0x3FEFc5A4B1c02f21cBc8D3613643ba0635b9a873#code, thanks
    function fromLittleEndianUint64(bytes32 num) internal pure returns (uint64) {
        uint64 v = uint64(uint256(num >> 192));
        //if we number the bytes (1, 2, 3, 4, 5, 6, 7, 8)
        v = ((v & 0x00ff00ff00ff00ff) << 8) | ((v & 0xff00ff00ff00ff00) >> 8);
        // (2, 0, 4, 0, 6, 0, 8, 0) | (0, 1, 0, 3, 0, 5, 0, 7)
        // = (2, 1, 4, 3, 6, 5, 8, 7)
        v = ((v & 0x0000ffff0000ffff) << 16) | ((v & 0xffff0000ffff0000) >> 16);
        // (4, 3, 0, 0, 8, 7, 0, 0) | (0, 0, 2, 1, 0, 0, 6, 5)
        // = (4, 3, 2, 1, 8, 7, 6, 5)
        // then
        // = (8, 7, 6, 5, 4, 3, 2, 1)
        return (v << 32) | (v >> 32);
    }
}
