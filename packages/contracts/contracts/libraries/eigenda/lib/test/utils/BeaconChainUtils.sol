//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract BeaconChainProofUtils {

    bytes32[] beaconStateMerkleProof;
    bytes32[] validatorMerkleProof;
    bytes32[] validatorContainerFields;

    //this function generates a proof for validator 0x5e2c2b702b0af22301f7ae52886da3827ea100b3d2a52222e6a10ea82e718a7f 
    //with an initial deposit of 32 ETH
    function getInitialDepositProof() public returns(bytes32[] memory, bytes32[] memory, bytes32[] memory, bytes32, bytes32){
        beaconStateMerkleProof.push(0x0000000000000000000000000000000000000000000000000000000000000000);
        beaconStateMerkleProof.push(0x8a7c6aed738e0a0cf25ebb8c5b4da41173285b41451674890a0ca5a100c2d3c9);
        beaconStateMerkleProof.push(0xd63afe2579fd495c87de3b1dc9dded10c6e65b1a717a4efed9e826b969e2a6d9);
        beaconStateMerkleProof.push(0x086ef90e3db0073ad2f8b2e6b38653d726e850fde26859dd881da1ac523598f0);
        beaconStateMerkleProof.push(0x1260718cd540a187a9dcff9f4d39116cdc1a0aed8a94fbe7a69fb87eae747be5);

        validatorContainerFields.push(0x5e2c2b702b0af22301f7ae52886da3827ea100b3d2a52222e6a10ea82e718a7f);
        validatorContainerFields.push(0x010000000000000000000000a3e2f3de1b0a71ed49ec592caeb42322f89ebaea);
        validatorContainerFields.push(0x2000000000000000000000000000000000000000000000000000000000000000);
        validatorContainerFields.push(0x0000000000000000000000000000000000000000000000000000000000000000);
        validatorContainerFields.push(0x0200000000000000000000000000000000000000000000000000000000000000);
        validatorContainerFields.push(0x0300000000000000000000000000000000000000000000000000000000000000);
        validatorContainerFields.push(0x0600000000000000000000000000000000000000000000000000000000000000);
        validatorContainerFields.push(0x0900000000000000000000000000000000000000000000000000000000000000);

        validatorMerkleProof.push(0x0000000000000000000000000000000000000000000000000000000000000000);
        validatorMerkleProof.push(0xf5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b);
        validatorMerkleProof.push(0xdb56114e00fdd4c1f85c892bf35ac9a89289aaecb1ebd0a96cde606a748b5d71);
        validatorMerkleProof.push(0xc78009fdf07fc56a11f122370658a353aaa542ed63e44c4bc15ff4cd105ab33c);
        validatorMerkleProof.push(0x536d98837f2dd165a55d5eeae91485954472d56f246df256bf3cae19352a123c);
        validatorMerkleProof.push(0x9efde052aa15429fae05bad4d0b1d7c64da64d03d7a1854a588c2cb8430c0d30);
        validatorMerkleProof.push(0xd88ddfeed400a8755596b21942c1497e114c302e6118290f91e6772976041fa1);
        validatorMerkleProof.push(0x87eb0ddba57e35f6d286673802a4af5975e22506c7cf4c64bb6be5ee11527f2c);
        validatorMerkleProof.push(0x26846476fd5fc54a5d43385167c95144f2643f533cc85bb9d16b782f8d7db193);
        validatorMerkleProof.push(0x506d86582d252405b840018792cad2bf1259f1ef5aa5f887e13cb2f0094f51e1);
        validatorMerkleProof.push(0xffff0ad7e659772f9534c195c815efc4014ef1e1daed4404c06385d11192e92b);
        validatorMerkleProof.push(0x6cf04127db05441cd833107a52be852868890e4317e6a02ab47683aa75964220);
        validatorMerkleProof.push(0xb7d05f875f140027ef5118a2247bbb84ce8f2f0f1123623085daf7960c329f5f);
        validatorMerkleProof.push(0xdf6af5f5bbdb6be9ef8aa618e4bf8073960867171e29676f8b284dea6a08a85e);
        validatorMerkleProof.push(0xb58d900f5e182e3c50ef74969ea16c7726c549757cc23523c369587da7293784);
        validatorMerkleProof.push(0xd49a7502ffcfb0340b1d7885688500ca308161a7f96b62df9d083b71fcc8f2bb);
        validatorMerkleProof.push(0x8fe6b1689256c0d385f42f5bbe2027a22c1996e110ba97c171d3e5948de92beb);
        validatorMerkleProof.push(0x8d0d63c39ebade8509e0ae3c9c3876fb5fa112be18f905ecacfecb92057603ab);
        validatorMerkleProof.push(0x95eec8b2e541cad4e91de38385f2e046619f54496c2382cb6cacd5b98c26f5a4);
        validatorMerkleProof.push(0xf893e908917775b62bff23294dbbe3a1cd8e6cc1c35b4801887b646a6f81f17f);
        validatorMerkleProof.push(0xcddba7b592e3133393c16194fac7431abf2f5485ed711db282183c819e08ebaa);
        validatorMerkleProof.push(0x8a8d7fe3af8caa085a7639a832001457dfb9128a8061142ad0335629ff23ff9c);
        validatorMerkleProof.push(0xfeb3c337d7a51a6fbf00b9e34c52e1c9195c969bd4e7a0bfd51d5c5bed9c1167);
        validatorMerkleProof.push(0xe71f0aa83cc32edfbefa9f4d3e0174ca85182eec9f3a09f6a6c0df6377a510d7);
        validatorMerkleProof.push(0x31206fa80a50bb6abe29085058f16212212a60eec8f049fecb92d8c8e0a84bc0);
        validatorMerkleProof.push(0x21352bfecbeddde993839f614c3dac0a3ee37543f9b412b16199dc158e23b544);
        validatorMerkleProof.push(0x619e312724bb6d7c3153ed9de791d764a366b389af13c58bf8a8d90481a46765);
        validatorMerkleProof.push(0x7cdd2986268250628d0c10e385c58c6191e6fbe05191bcc04f133f2cea72c1c4);
        validatorMerkleProof.push(0x848930bd7ba8cac54661072113fb278869e07bb8587f91392933374d017bcbe1);
        validatorMerkleProof.push(0x8869ff2c22b28cc10510d9853292803328be4fb0e80495e8bb8d271f5b889636);
        validatorMerkleProof.push(0xb5fe28e79f1b850f8658246ce9b6a1e7b49fc06db7143e8fe0b4f2b0c5523a5c);
        validatorMerkleProof.push(0x985e929f70af28d0bdd1a90a808f977f597c7c778c489e98d3bd8910d31ac0f7);
        validatorMerkleProof.push(0xc6f67e02e6e4e1bdefb994c6098953f34636ba2b6ca20a4721d2b26a886722ff);
        validatorMerkleProof.push(0x1c9a7e5ff1cf48b4ad1582d3f4e4a1004f3b20d8c5a2b71387a4254ad933ebc5);
        validatorMerkleProof.push(0x2f075ae229646b6f6aed19a5e372cf295081401eb893ff599b3f9acc0c0d3e7d);
        validatorMerkleProof.push(0x328921deb59612076801e8cd61592107b5c67c79b846595cc6320c395b46362c);
        validatorMerkleProof.push(0xbfb909fdb236ad2411b4e4883810a074b840464689986c3f8a8091827e17c327);
        validatorMerkleProof.push(0x55d8fb3687ba3ba49f342c77f5a1f89bec83d811446e1a467139213d640b6a74);
        validatorMerkleProof.push(0xf7210d4f8e7e1039790e7bf4efa207555a10a6db1dd4b95da313aaa88b88fe76);
        validatorMerkleProof.push(0xad21b516cbc645ffe34ab5de1c8aef8cd4e7f8d2b51e8e1456adc7563cda206f);
        validatorMerkleProof.push(0x0100000000000000000000000000000000000000000000000000000000000000);

        //hash tree root of list of validators
        bytes32 validatorTreeRoot = 0x8fa2006e3402e8226353692fad2d35457e5b9c3bfe73130d488f1c220363baa9;

        //hash tree root of individual validator container
        bytes32 validatorRoot = 0x8a38196fcb74b5c932c486598e4bff65b4ae1dc3eacaa839f0fc6fe8171ee92b;

        return (beaconStateMerkleProof, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot);

    }
    //simulates a 16ETH slashing
    function getSlashedDepositProof() public returns(bytes32[] memory, bytes32[] memory, bytes32[] memory, bytes32, bytes32){
        
        beaconStateMerkleProof[0] = 0x0000000000000000000000000000000000000000000000000000000000000000;
        beaconStateMerkleProof[1] = 0x8a7c6aed738e0a0cf25ebb8c5b4da41173285b41451674890a0ca5a100c2d3c9;
        beaconStateMerkleProof[2] = 0xd63afe2579fd495c87de3b1dc9dded10c6e65b1a717a4efed9e826b969e2a6d9;
        beaconStateMerkleProof[3] = 0x086ef90e3db0073ad2f8b2e6b38653d726e850fde26859dd881da1ac523598f0;
        beaconStateMerkleProof[4] = 0x1260718cd540a187a9dcff9f4d39116cdc1a0aed8a94fbe7a69fb87eae747be5;

        validatorContainerFields[0] = 0x5e2c2b702b0af22301f7ae52886da3827ea100b3d2a52222e6a10ea82e718a7f;
        validatorContainerFields[1] = 0x010000000000000000000000a3e2f3de1b0a71ed49ec592caeb42322f89ebaea;
        validatorContainerFields[2] = 0x1000000000000000000000000000000000000000000000000000000000000000;
        validatorContainerFields[3] = 0x0100000000000000000000000000000000000000000000000000000000000000;
        validatorContainerFields[4] = 0x0200000000000000000000000000000000000000000000000000000000000000;
        validatorContainerFields[5] = 0x0300000000000000000000000000000000000000000000000000000000000000;
        validatorContainerFields[6] = 0x0600000000000000000000000000000000000000000000000000000000000000;
        validatorContainerFields[7] = 0x0900000000000000000000000000000000000000000000000000000000000000;


        validatorMerkleProof[0] = 0x0000000000000000000000000000000000000000000000000000000000000000;
        validatorMerkleProof[1] = 0xf5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b;
        validatorMerkleProof[2] = 0xdb56114e00fdd4c1f85c892bf35ac9a89289aaecb1ebd0a96cde606a748b5d71;
        validatorMerkleProof[3] = 0xc78009fdf07fc56a11f122370658a353aaa542ed63e44c4bc15ff4cd105ab33c;
        validatorMerkleProof[4] = 0x536d98837f2dd165a55d5eeae91485954472d56f246df256bf3cae19352a123c;
        validatorMerkleProof[5] = 0x9efde052aa15429fae05bad4d0b1d7c64da64d03d7a1854a588c2cb8430c0d30;
        validatorMerkleProof[6] = 0xd88ddfeed400a8755596b21942c1497e114c302e6118290f91e6772976041fa1;
        validatorMerkleProof[7] = 0x87eb0ddba57e35f6d286673802a4af5975e22506c7cf4c64bb6be5ee11527f2c;
        validatorMerkleProof[8] = 0x26846476fd5fc54a5d43385167c95144f2643f533cc85bb9d16b782f8d7db193;
        validatorMerkleProof[9] = 0x506d86582d252405b840018792cad2bf1259f1ef5aa5f887e13cb2f0094f51e1;
        validatorMerkleProof[10] = 0xffff0ad7e659772f9534c195c815efc4014ef1e1daed4404c06385d11192e92b;
        validatorMerkleProof[11] = 0x6cf04127db05441cd833107a52be852868890e4317e6a02ab47683aa75964220;
        validatorMerkleProof[12] = 0xb7d05f875f140027ef5118a2247bbb84ce8f2f0f1123623085daf7960c329f5f;
        validatorMerkleProof[13] = 0xdf6af5f5bbdb6be9ef8aa618e4bf8073960867171e29676f8b284dea6a08a85e;
        validatorMerkleProof[14] = 0xb58d900f5e182e3c50ef74969ea16c7726c549757cc23523c369587da7293784;
        validatorMerkleProof[15] = 0xd49a7502ffcfb0340b1d7885688500ca308161a7f96b62df9d083b71fcc8f2bb;
        validatorMerkleProof[16] = 0x8fe6b1689256c0d385f42f5bbe2027a22c1996e110ba97c171d3e5948de92beb;
        validatorMerkleProof[17] = 0x8d0d63c39ebade8509e0ae3c9c3876fb5fa112be18f905ecacfecb92057603ab;
        validatorMerkleProof[18] = 0x95eec8b2e541cad4e91de38385f2e046619f54496c2382cb6cacd5b98c26f5a4;
        validatorMerkleProof[19] = 0xf893e908917775b62bff23294dbbe3a1cd8e6cc1c35b4801887b646a6f81f17f;
        validatorMerkleProof[20] = 0xcddba7b592e3133393c16194fac7431abf2f5485ed711db282183c819e08ebaa;
        validatorMerkleProof[21] = 0x8a8d7fe3af8caa085a7639a832001457dfb9128a8061142ad0335629ff23ff9c;
        validatorMerkleProof[22] = 0xfeb3c337d7a51a6fbf00b9e34c52e1c9195c969bd4e7a0bfd51d5c5bed9c1167;
        validatorMerkleProof[23] = 0xe71f0aa83cc32edfbefa9f4d3e0174ca85182eec9f3a09f6a6c0df6377a510d7;
        validatorMerkleProof[24] = 0x31206fa80a50bb6abe29085058f16212212a60eec8f049fecb92d8c8e0a84bc0;
        validatorMerkleProof[25] = 0x21352bfecbeddde993839f614c3dac0a3ee37543f9b412b16199dc158e23b544;
        validatorMerkleProof[26] = 0x619e312724bb6d7c3153ed9de791d764a366b389af13c58bf8a8d90481a46765;
        validatorMerkleProof[27] = 0x7cdd2986268250628d0c10e385c58c6191e6fbe05191bcc04f133f2cea72c1c4;
        validatorMerkleProof[28] = 0x848930bd7ba8cac54661072113fb278869e07bb8587f91392933374d017bcbe1;
        validatorMerkleProof[29] = 0x8869ff2c22b28cc10510d9853292803328be4fb0e80495e8bb8d271f5b889636;
        validatorMerkleProof[30] = 0xb5fe28e79f1b850f8658246ce9b6a1e7b49fc06db7143e8fe0b4f2b0c5523a5c;
        validatorMerkleProof[31] = 0x985e929f70af28d0bdd1a90a808f977f597c7c778c489e98d3bd8910d31ac0f7;
        validatorMerkleProof[32] = 0xc6f67e02e6e4e1bdefb994c6098953f34636ba2b6ca20a4721d2b26a886722ff;
        validatorMerkleProof[33] = 0x1c9a7e5ff1cf48b4ad1582d3f4e4a1004f3b20d8c5a2b71387a4254ad933ebc5;
        validatorMerkleProof[34] = 0x2f075ae229646b6f6aed19a5e372cf295081401eb893ff599b3f9acc0c0d3e7d;
        validatorMerkleProof[35] = 0x328921deb59612076801e8cd61592107b5c67c79b846595cc6320c395b46362c;
        validatorMerkleProof[36] = 0xbfb909fdb236ad2411b4e4883810a074b840464689986c3f8a8091827e17c327;
        validatorMerkleProof[37] = 0x55d8fb3687ba3ba49f342c77f5a1f89bec83d811446e1a467139213d640b6a74;
        validatorMerkleProof[38] = 0xf7210d4f8e7e1039790e7bf4efa207555a10a6db1dd4b95da313aaa88b88fe76;
        validatorMerkleProof[39] = 0xad21b516cbc645ffe34ab5de1c8aef8cd4e7f8d2b51e8e1456adc7563cda206f;
        validatorMerkleProof[40] = 0x0100000000000000000000000000000000000000000000000000000000000000;


        //hash tree root of list of validators
        bytes32 validatorTreeRoot = 0x678b8500eaee69d975a9ba9cecf1d64bf102da3f10f65d6fb9331abd431e62bd;

        //hash tree root of individual validator container
        bytes32 validatorRoot = 0x24c79c612ac9d0ac4663b8b9b698ac8ddd4d1942ccb425d43f216f302d251b9d;

        return (beaconStateMerkleProof, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot);
    }



}