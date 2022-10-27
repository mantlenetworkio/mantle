// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract Test {
    int32 public versionNum;

    function setVersion() public virtual {
        versionNum = 1;
    }

    function version() public view virtual returns (int32) {
        return versionNum;
    }
}

contract TestUpgrade is Test {
    function setVersion() public override {
        versionNum = 2;
    }

    function version() public view override returns (int32) {
        return versionNum;
    }
}
