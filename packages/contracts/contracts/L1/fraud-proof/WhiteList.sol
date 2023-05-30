abstract contract Whitelist {
    modifier onlyOwner() {
        require(msg.sender == owner, "Ownable: caller is not the owner");
        _;
    }

    modifier stakerWhitelistOnly(address _checkAddress) {
        if (stakerslist[stakerWhitelist[_checkAddress]] == _checkAddress) {
            revert("NOT_IN_STAKER_WHITELIST");
        }
        _;
    }

    modifier operatorWhitelistOnly(address _checkAddress) {
        if (operatorslist[operatorWhitelist[_checkAddress]] == _checkAddress) {
            revert("NOT_IN_OPERATOR_WHITELIST");
        }
        _;
    }

    address public owner;
    mapping(address => uint256) public stakerWhitelist;
    address[] public stakerslist;
    mapping(address => uint256) public operatorWhitelist;
    address[] public operatorslist;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @notice Add to staker whitelist
     */
    function addToStakerWhitelist(address[] calldata toAddAddresses) public onlyOwner {
        for (uint i = 0; i < toAddAddresses.length; i++) {
            stakerWhitelist[toAddAddresses[i]] = stakerslist.length + i;
            stakerslist.push(toAddAddresses[i]);
        }
    }

    /**
     * @notice Remove from whitelist
     */
    function removeFromStakerWhitelist(address[] calldata toRemoveAddresses) public onlyOwner {
        for (uint i = 0; i < toRemoveAddresses.length; i++) {
            uint256 index = stakerWhitelist[toRemoveAddresses[i]];
            stakerslist[index] = stakerslist[stakerslist.length-1];
            stakerslist.pop();
            delete stakerWhitelist[toRemoveAddresses[i]];
        }
    }

    /**
 * @notice Add to whitelist
     */
    function addToOperatorWhitelist(address[] calldata toAddAddresses) public onlyOwner {
        for (uint i = 0; i < toAddAddresses.length; i++) {
            operatorWhitelist[toAddAddresses[i]] = operatorslist.length + i;
            operatorslist.push(toAddAddresses[i]);
        }
    }

    /**
     * @notice Remove from whitelist
     */
    function removeFromOperatorWhitelist(address[] calldata toRemoveAddresses) public onlyOwner {
        for (uint i = 0; i < toRemoveAddresses.length; i++) {
            uint256 index = operatorWhitelist[toRemoveAddresses[i]];
            operatorslist[index] = operatorslist[operatorslist.length-1];
            operatorslist.pop();
            delete operatorWhitelist[toRemoveAddresses[i]];
        }
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = owner;
        owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}
