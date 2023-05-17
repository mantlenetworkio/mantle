abstract contract Whitelist {
    modifier onlyOwner() {
        require(msg.sender == owner, "Ownable: caller is not the owner");
        _;
    }

    modifier stakerWhitelistOnly(address _checkAddress) {
        if (!stakerWhitelist[_checkAddress]) {
            revert("NOT_IN_STAKER_WHITELIST");
        }
        _;
    }

    modifier operatorWhitelistOnly(address _checkAddress) {
        if (!operatorWhitelist[_checkAddress]) {
            revert("NOT_IN_OPERATOR_WHITELIST");
        }
        _;
    }

    address public owner;
    mapping(address => bool) public stakerWhitelist;
    mapping(address => bool) public operatorWhitelist;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @notice Add to staker whitelist
     */
    function addToStakerWhitelist(address[] calldata toAddAddresses) external onlyOwner {
        for (uint i = 0; i < toAddAddresses.length; i++) {
            stakerWhitelist[toAddAddresses[i]] = true;
        }
    }

    /**
     * @notice Remove from whitelist
     */
    function removeFromStakerWhitelist(address[] calldata toRemoveAddresses) external onlyOwner {
        for (uint i = 0; i < toRemoveAddresses.length; i++) {
            delete stakerWhitelist[toRemoveAddresses[i]];
        }
    }

    /**
 * @notice Add to whitelist
     */
    function addToOperatorWhitelist(address[] calldata toAddAddresses) external onlyOwner {
        for (uint i = 0; i < toAddAddresses.length; i++) {
            operatorWhitelist[toAddAddresses[i]] = true;
        }
    }

    /**
     * @notice Remove from whitelist
     */
    function removeFromOperatorWhitelist(address[] calldata toRemoveAddresses) external onlyOwner {
        for (uint i = 0; i < toRemoveAddresses.length; i++) {
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
