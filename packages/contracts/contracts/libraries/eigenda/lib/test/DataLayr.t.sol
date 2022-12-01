// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./DataLayrTestHelper.t.sol";
import "forge-std/Test.sol";

contract DataLayrTests is DSTest, DataLayrTestHelper {
    //checks that it is possible to init a data store
    function testInitDataStore() public returns (bytes32) {
        uint256 numSigners = 15;


        //register all the operators
        _registerNumSigners(numSigners);

        //change the current timestamp to be in the future 100 seconds and init
        return _testInitDataStore(block.timestamp + 100, address(this), header).metadata.headerHash;
    }

    function testLoopInitDataStore() public {
        uint256 g = gasleft();
        uint256 numSigners = 15;

        for (uint256 i = 0; i < 20; i++) {
            if(i==0){
                _registerNumSigners(numSigners);
            }
            _testInitDataStore(block.timestamp + 100, address(this), header).metadata.headerHash;
        }
        emit log_named_uint("gas", g - gasleft());
    }

    //verifies that it is possible to confirm a data store
    //checks that the store is marked as committed
    function testConfirmDataStore() public {
        _testConfirmDataStoreSelfOperators(15);
    }

    function testLoopConfirmDataStore() public {
        _testConfirmDataStoreSelfOperators(15);
        uint256 g = gasleft();
        for (uint256 i = 1; i < 5; i++) {
            _testConfirmDataStoreWithoutRegister(block.timestamp, i, 15);
        }
        emit log_named_uint("gas", g - gasleft());
    }

    function testConfirmDataStoreTwelveOperators() public {
        _testConfirmDataStoreSelfOperators(12);
    }

    function testCodingRatio() public {

        uint256 numSigners = 15;
        //register all the operators
        _registerNumSigners(numSigners);

        /// @notice this header has numSys set to 9.  Thus coding ration = 9/15, which is greater than the set adversary threshold in DataLayrServiceManager.
        bytes memory header = hex"0e75f28b7a90f89995e522d0cd3a340345e60e249099d4cd96daef320a3abfc31df7f4c8f6f8bc5dc1de03f56202933ec2cc40acad1199f40c7b42aefd45bfb10000000800000009000000020000014000000000000000000000000000000000000000002b4982b07d4e522c2a94b3e7c5ab68bfeecc33c5fa355bc968491c62c12cf93f0cd04099c3d9742620bf0898cf3843116efc02e6f7d408ba443aa472f950e4f3";
        
        uint256 initTimestamp = block.timestamp + 100;
        address confirmer  = address(this);

        _testInitDataStoreExpectRevert(initTimestamp, confirmer, header, bytes("DataLayrServiceManager.initDataStore: Coding ratio is too high"));
    }

    function testZeroTotalBytes() public {        
        uint256 initTimestamp = block.timestamp + 100;
        address confirmer  = address(this);

        _testInitDataStoreExpectRevert(initTimestamp, confirmer, header, bytes("DataLayrServiceManager.initDataStore: totalBytes < MIN_STORE_SIZE"));

    }

    function testTotalOperatorIndex(uint32 wrongTotalOperatorsIndex) external {
        uint256 numSigners = 15;
        //register all the operators
        _registerNumSigners(numSigners);        
        uint256 initTimestamp = block.timestamp + 100;

        cheats.assume(wrongTotalOperatorsIndex > uint32(dlReg.getLengthOfTotalOperatorsHistory()));

        // weth is set as the paymentToken of dlsm, so we must approve dlsm to transfer weth
        weth.transfer(storer, 1e11);
        cheats.startPrank(storer);
        weth.approve(address(dataLayrPaymentManager), type(uint256).max);

        dataLayrPaymentManager.depositFutureFees(storer, 1e11);

        uint32 blockNumber = uint32(block.number);

        require(initTimestamp >= block.timestamp, "_testInitDataStore: warping back in time!");
        cheats.warp(initTimestamp);

        cheats.expectRevert();
        dlsm.initDataStore(
            storer,
            address(this),
            durationToInit,
            blockNumber,
            wrongTotalOperatorsIndex,
            header
        );
    }


     //testing inclusion of nonsigners in DLN quorum, ensuring that nonsigner inclusion proof is working correctly.
    function testInadequateQuorumStake(uint256 ethAmount, uint256 eigenAmount) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e10);

        {
            // address operator = signers[0];
            uint8 operatorType = 3;
            _testInitiateDelegation(0, eigenAmount, ethAmount);
            _testRegisterBLSPubKey(0);
            _testRegisterOperatorWithDataLayr(0, operatorType, testEphemeralKey, testSocket);
        }

        NonSignerPK memory nonsignerPK1;
        NonSignerPK memory nonsignerPK2;
        RegistrantAPK memory registrantAPK;
        SignerAggSig memory signerAggSig;

        //bytes memory header = hex"0e75f28b7a90f89995e522d0cd3a340345e60e249099d4cd96daef320a3abfc31df7f4c8f6f8bc5dc1de03f56202933ec2cc40acad1199f40c7b42aefd45bfb10000000800000002000000020000014000000000000000000000000000000000000000002b4982b07d4e522c2a94b3e7c5ab68bfeecc33c5fa355bc968491c62c12cf93f0cd04099c3d9742620bf0898cf3843116efc02e6f7d408ba443aa472f950e4f3";


        nonsignerPK1.xA0 = (uint256(9391974691841703379432258354827183968448857856995465041611595190399280871636));
        nonsignerPK1.xA1 = (uint256(13443635970046784780120024980077142239453332379977521244676409699881477679792));
        nonsignerPK1.yA0 = (uint256(18537770077305880837445921613844169212065436707683905058015983155555211983585));
        nonsignerPK1.yA1 = (uint256(2367997946501567411511477079128155943372633684242183349427117627773599205044));

        nonsignerPK2.xA0 = (uint256(10245738255635135293623161230197183222740738674756428343303263476182774511624));
        nonsignerPK2.xA1 = (uint256(10281853605827367652226404263211738087634374304916354347419537904612128636245));
        nonsignerPK2.yA0 = (uint256(3091447672609454381783218377241231503703729871039021245809464784750860882084));
        nonsignerPK2.yA1 = (uint256(18210007982945446441276599406248966847525243540006051743069767984995839204266));

        //aggreate public key of all registrants, including nonsigners
        registrantAPK.apk0 = uint256(20820493588973199354272631301248587752629863429201347184003644368113679196121);
        registrantAPK.apk1 = uint256(18507428821816114421698399069438744284866101909563082454551586195885282320634);
        registrantAPK.apk2 = uint256(1263326262781780932600377484793962587101562728383804037421955407439695092960);
        registrantAPK.apk3 = uint256(3512517006108887301063578607317108977425754510174956792003926207778790018672);
        
        signerAggSig.sigma0 = uint256(21866930911187421380583942535436041239527001942513358568552618703117179737517);
        signerAggSig.sigma1 = uint256(11815084818309926263646220976708780882220448121684663825366455811834597006341);

        {
            uint32 numberOfSigners = 15;
            _testRegisterSigners(numberOfSigners, false);
        }

        bytes memory data;
        uint256 initTime = 1000000001;
        IDataLayrServiceManager.DataStoreSearchData memory searchData = _testInitDataStore(initTime, address(this), header);

        // multiple scoped blocks helps fix 'stack too deep' errors
        {
            
            uint32 numberOfNonSigners = 2;
            uint32 dataStoreId = dlsm.taskNumber() - 1;
            data = abi.encodePacked(
                keccak256(
                    abi.encodePacked(
                        searchData.metadata.globalDataStoreId,
                        searchData.metadata.headerHash,
                        searchData.duration,
                        initTime,
                        uint32(0)
                    )
                ),
                uint48(dlReg.getLengthOfTotalStakeHistory() - 1),
                searchData.metadata.blockNumber,
                dataStoreId,
                numberOfNonSigners
            );
        }
        {
            data = abi.encodePacked(
                data,
                nonsignerPK1.xA0,
                nonsignerPK1.xA1,
                nonsignerPK1.yA0,
                nonsignerPK1.yA1,
                uint32(0)
            );

        }
        {
            data = abi.encodePacked(
                data,
                nonsignerPK2.xA0,
                nonsignerPK2.xA1,
                nonsignerPK2.yA0,
                nonsignerPK2.yA1,
                uint32(0),
                uint32(dlReg.getApkUpdatesLength() - 1)
            );
        }
        {
            data = abi.encodePacked(
                data,
                registrantAPK.apk0,
                registrantAPK.apk1,
                registrantAPK.apk2,
                registrantAPK.apk3,
                signerAggSig.sigma0,
                signerAggSig.sigma1

            );
        }

        cheats.expectRevert(bytes("DataLayrServiceManager.confirmDataStore: signatories do not own at least threshold percentage of both quorums"));
        dlsm.confirmDataStore(data, searchData);

    }

    function _registerNumSigners(uint256 numSigners) internal {
        for (uint256 i = 0; i < numSigners; ++i) {
            _testRegisterAdditionalSelfOperator(signers[i], registrationData[i], ephemeralKeyHashes[i]);
        }
    }
}
