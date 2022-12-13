pragma solidity ^0.8.9;


contract Parser  {
    function parse(bytes[] calldata polys, uint256 startIndex, uint256 length) public returns(bytes memory) {
        bytes memory provenString;
        uint256 polyIndex = 0;
        uint256 index = startIndex;
        while(provenString.length < length) {
            //read till the beginning of the next 32 byte segment
            //or however many more bytes there are left in the fraud string
            //or till the end of the polynomial
            uint256 bytesToRead = min(
                32*(index/32 + 1) - index,
                length - provenString.length,
                polys[polyIndex].length - index
            );
            //append read bytes to the end of the proven string
            provenString = abi.encodePacked(provenString, polys[polyIndex][index:index+bytesToRead]);
            if(length <= provenString.length) {
                //if we have read same number of bytes in the fraud string, break
                break;
            } else if (index+bytesToRead == polys[polyIndex].length) {
                //if we finished reading the current polynomial so we move onto the next
                polyIndex++;
                index = 1;
            } else {
                //we have read readUntil bytes and add 1 more to skip the zero byte at the beginning of every 32 bytes
                index += bytesToRead + 1;
            }
        }
        return provenString;
    }

    function min(uint256 a, uint256 b, uint256 c) internal pure returns(uint256) {
        return (a < b) ? ((a < c) ? a : c) : ((b < c) ? b : c);
    }
}
