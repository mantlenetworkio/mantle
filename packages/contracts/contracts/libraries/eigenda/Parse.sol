pragma solidity ^0.8.9;

contract Parser {
    /**
     * @notice Parses data from non-interactive polynomial proofs.
     * @param polys The non-interactive polynomial proofs themselves
     * @param startIndex The byte index from which to begin reading data.
     * @param length The length of data to parse, in bytes.
     * @return provenString The parsed data.
     */
    function parse(bytes[] calldata polys, uint256 startIndex, uint256 length) public pure returns(bytes memory provenString) {
        // each symbol encodes 31 bytes, and is padded to 32 bytes -- this verifies that we are beginning to parse the data from a non-padded byte
        require(startIndex % 32 != 0, "Cannot start reading from a padded byte");
        // index of the `polys` array from which we are currently reading
        uint256 polyIndex = 0;
        // keeps track of the index to read inside of the current polynomial
        uint256 index = startIndex;
        // continue reading until we reach the desired length
        while(provenString.length < length) {
            /**
             * Read:
             * 1) until the beginning of the next 32 byte segment OR
             * 2) however many more bytes there are left in the fraud string
             * -- whichever amount is the *smallest*
             */
            uint256 bytesToRead = min(
            // the amount of bytes until the end of the current 32 byte segment
                (32 * ((index / 32) + 1)) - index,
            // the remaining total bytes to parse
                length - provenString.length
            );
            /**
             * Append the read bytes to the end of the proven string.
             * Note that indexing of bytes is inclusive of the first index and exclusive of the second, meaning
             * that, for example, polys[0][x:x+1] specifies the *single byte* at position x of `polys[0]`, and
             * polys[0][x:x] will specify an empty byte string.
             */
            provenString = abi.encodePacked(provenString, polys[polyIndex][index:index+bytesToRead]);
            // if we finished reading the current polynomial, then we move onto the next one
            if (index + bytesToRead == polys[polyIndex].length) {
                polyIndex++;
                // skip the first byte of the polynomial since this is zero padding
                index = 1;
                // we have read `index + bytesToRead` bytes, and add 1 more to skip the zero-padding byte at the beginning of every 32 bytes
            } else {
                index += bytesToRead + 1;
            }
        }
        return provenString;
    }

    /// @notice Calculates the minimum of 2 numbers
    function min(uint256 a, uint256 b) internal pure returns(uint256) {
        return (a < b) ? a : b;
    }
}
