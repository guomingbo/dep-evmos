EXPECTED_GET_STORAGE_AT = (
    "0x00000000000000000000000000000000000000000000000000120a0b063499d4"
)

# curl https://mainnet.infura.io/v3/YOUR_KEY_HERE \
# -X POST \
# -H "Content-Type: application/json" \
# -d '{"jsonrpc": "2.0","method": "eth_getProof","params":
# ["0x7F0d15C7FAae65896648C8273B6d7E43f58Fa842",["0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"],
# "latest"],"id": 1}'
EXPECTED_GET_PROOF = {
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "accountProof": [
            "0xf90211a0d308f98795cbcc61d0520723e3361adc5513705272d90ee48ccd247246f7c349a010139d8b72e8086a732523b8905eabf1172c5202efc588a3c7cd22d86c01ccbda0faa34953c5e42eba596c0085082e7969a0fba937de89a482882f56c60eb61fbda00bb6a0aae1af594fabf4bb4be474fe7c99a3936b27b3ae838341d7ea31988df0a000fd250e61c6c9944d981d9ac2d1d3265239349bd82fd06ac5635fb3026ea3dea064625348caa521add14ff55a3efb5335d70c3a57c89ca74381b29c5e44bba22ca07c87aeff8d60769087900486cee0f400d4c73ac5f75e80ea7e419e3222493c61a0a85a1e93809071ab62ef124c05f71b33b5f3beb29a40d3362449589ad069f6c7a01e44cbe03b61463da0738947c35bfd6224f1a09c3c41da47f31d779b2c6a992fa0828a75c8da3f71e2df63ac7c0f19bb4b1c9dd861fc403823f189b588dd29eb75a0c0dfa1e4be056540e1a9524e3b060ab8f6d51bc688d84b7fcf7289cfaba1ceaba00865239fb25a3a83889326e8bc21f050b33c5181681c33b89e1d7ab883b3f3c5a0e148cdb3af3c738c1e53d1f85913b0accd46c45b34e3873abd6c98d5995d5f83a0044064693ce7d1a5888a1aca785ff1126c09cf83a00c4a3b386cb3da056eb432a0b5711def31905c703d958ae996321c9f8aefb12107acfa2dc20311801eb2558fa04005fcc01439852c004a84c6542d139b20ddac8c6e1e9684821d10e07cb4e33a80",  # noqa: E501
            "0xf90211a0051fd734a1442d85238b82a7fd9ab58cac3b56728e0d07141ca30ea09c502fd4a03537903e89b73f84f80346827ccb523c9658a478e4fac51269c1064d5b74dccea0fa27fc86a7025b9928cbc55d7fc4ad42fc4974f8f69e6e2c780360510398a28da0752a52276798a12b712b0edde30a8308da6618b0f58b3199c78d0654561541caa09e64d6fa19e6dadf9b107059d5e9e1ba0f8dade4605e31d36d84a8661a18b9c1a06f724081c76cdb3eeba6c164c16eecc7ca85ab6d38bc3bce4983feead3df67f9a02c66343b11b10ba25c84e2f75dc1e2a995c32929312ff9a347536138162be1d7a0994e90111ec6af30851974a047dd7a227d235b2c6d0643199a5c3f6a685bf77aa0f844ea047a6e348f4cd8d0845e5e48b06ea033c6fd4d82eb534a8bfac990be83a0db9d470d492623fd2157320e0d658e4a46ee37c12211db9356b977e09aa01551a062efa94455c86e0e45691054225aa509acf08d97ab1253c6589b9a9f7f210f0aa057a857a358c0f0b4363ee793dc7393efdc19271ffe1dfec2650aac0b867d9729a0fc714453737c12b3b20f3ca7be1a2c4e56ba75806677a3b5a75922d54637127aa00580bfa435f851327f608f21dabab8b74658fb80dfa4ae090de30e4134a652c0a080a7b5f559e6941f1e29a3dbbdc1bd41aac866845366e9435c2048233e185816a093d7ddf2501324e5f3c0b64b41829b9c929743f5db178809b01a0b513d66ab0380",  # noqa: E501
            "0xf90211a03ee5370675726aa9f895944a635a812819cd90853e34807e3763882b2f0dd978a03dcdfe595fb5b1c036e8c5618e75dc5dfbd20f4093b18f423d0c546a8841784fa0d181704ea7bcced784d9480abd6360e885657957e1eb596512b162cce3512433a0c75c61390552448a12958fd0fd51a88f7e5e6badd2c16b05f4c949f604b209e2a0bb24e0cc45c91c5d3f14c0a6eeb459db9447260696ba65e8a79f86170e2997e5a00870addc6060c3f528eacfabd6c52480f45f8e28ed3c7e58306daa37f7752205a0d9b904889003012fdafbfd6cae526bd627e6d6a8ccd0e1cecba8500649c9efcba0890f7bb42d79b87940acd735c61e6b8fe12b64c2aa7c162f62dfcb40b2a6bb26a09f8bb0d80d7f51b8803b12dd01bc77281f7d0746ae6d7331705d79be9ee676faa059f3d816f66dedd5ca93d54845d9fcbcd8a3b110c78c35629db79de7736d6181a0171aa28493b4a16cdaffb06ecdfc06f73bb5b2548426fdc0d1fc807fa734f219a0eedf3bd8db8007a360cb7ca38e7718e6c7de81244883caa3a276da03174e1697a086b32111c6000e06ae1d691588f96e589e00768aa3f8337633546ba3c5069694a06992a01de330340dba82add8897bf7bd9d1a4d89ee4f2c2e87407d5cd80b0a22a0a9aca4be841d84dfc3cfbcb492166d0a6a1bdf1bcd471c772eea60906da58da7a048439b7d9637e613a26f80ea269865dda9ceca2c5eea80e9a5acfb8ba83b1b7e80",  # noqa: E501
            "0xf90211a07e164dfd7f1ae0e9c7a5276c54f44dde92965a6c2f33a6ab70f15837e7334846a08d1f58f0b8113267e9ae0b00cc2814bafd7abdf64ef092da4022252723ec4009a08e4bd9b065fa8dab94bac469f47a04f2c6cd269c5345de50566a54ec34ef21dba0f9a8d5511d60abf59d66a80a252ea86e3f3b6c325eeac5d861fca7e16c4e4111a006ad844c311fe3730f814ed96a75dd268f609c30bd4476c7842d877cb2baa0aaa076e6eedf4846cb47b5343708159de812ab27f6ad350e68ed70a936bcad30401fa0aa867723f3f21a5eb867ed5c5b99609ea535b967f710e547945e3ec2a212c68fa09d3e7fd8cd22ab508d079b58b40f75c3242f7c7704a82ec1c573f682e2168824a0aba33ff032dc7e10fc3a228adbcff47480580b292c4afee754394c6017872d5ca0ec2cd81c2048e35a4bedee1b27e077129edab50a40972c0f6ccd25bd864bedf2a0371c03a000f78dff313551c7afbf52a1348127ab17b737732191dbd64e68b700a0df0aad31c4c0887c6109cadb0d2341a65f18bc35f6e5583ebd417853b9f56329a07c6437031b7543922a3e6a74ae237a844226db3f632a16dffbedd8679b03b018a0538592d2b6933ab7f0d276e7509151632fbca90c10079ee0c9cce87925b19abaa0d79a8e5a546626e60d5d5c7bb839612a42b5724e880deae4bc282f42496a1ef4a03d172648e5e01d32748c96362e197e37a13aeffa0ae213a23b8898e6508bb6a380",  # noqa: E501
            "0xf90211a07f77a654bce6a83ce5674a2936de6ea0768f4f11e19039c8b83ef62956d2ac15a0854518c708b71c9dfb97eaf4c300c0951108dc96ba9877e565cffced4224fbfca0d60083a3e2bdd537e16b3e2a6e62812fad9b7fbec0e56cb76676007fc1ac44eba0105b6a16896fc82d33f9c742890700b9a0312e37d15c7eb12dde2e015ce8920ea0f127c635dcde1f4898e8abcb45a1a8c07a32cc4d78f2f1fdc7386d26f41dc4aea06916ed562e2b84c6a4145a9a3c5716b5493db0637023347a883a3753522a57fea021b88faa7e7af0a40d03950f42abece8892452c3ce5d94266a36d444c8e84d73a0216e4a004f60aa058d227553b404ee93b5758f28d5bed24d6f3a34157b9a18bba0326b9ef2f85defffc61fc3561349969ff8aef23759ab4a4bd49b8dcd1453ee58a0649731f153d117e5ad56b07e3edfb911b3592009da413d952ff7fc2c1e0c8a25a077edfe281169e6425d0795670d9a33ea32ae11ee71896f9064d5a6189ed6693fa0f986846a0fff0ab469ae226de91e3f350757f49e21f48a57fc7c1e4142fa0c3aa0aae0a489f6e8142bce91638db238a1a2d32cc2c31d82eab86d20c89edca5ac3da04a789fc4610da58da2204e3c9f21e351b574b72c7a4194c996e2bc05d3086376a0184a75c2da01cbc27c48a0335318a60c9535ecd54a79bf96a9f821d8202a2080a0c22686c5baef46ef1364919c08a2cfe113cbc2ac0854d2a890df938538a5c36d80",  # noqa: E501
            "0xf90211a00b5b38caa0b8e500946e01b36441c04b3932af0af937866756aea01208e630e5a0347a4bd5d736dc5dd178c724de94ca9f3cf351d6cc697f547e115ac955c12762a0fc12087179e3debe3505a5da9cb133650322b5846644289211785ad3ec8c2400a001ef26f67f5ff67e0480d40e9822329c834f2683cc4973d11276c1d252c67d07a0dab8a89872ab8887433393a58c6838f56dd33598c17bc65b67b3975909a1d7d9a0d8e72a6583ec2542c5153d46ff671efcf0e86d3dd6594148be6c8f45b962372aa05d3dfc9d0271fa2e17f1af8a36a105755d7c8ffcd75b523bc9e508d1e4a7f9c5a0367b15490f03ee256758187c63bff678de58c716b94fcee3c68de9b88d4d6d8da0557dd25a49b199b78d194b2c418afcf654415b2db9354b9836490117ac9d768ea0a58e109994f8dccbc777bdf9af2b756c62a1102aba71513f3cbd9aab76fe5d99a034bfcb38403a1a9f3409ea13cb8139f820b702ee757cf87d229949825180a89ea0e3f8b04c275387c6222b6536967f2f1e6fc76dc6ee0e705d3e505f3d07e8bdd4a0deaf89533b13c20ff50ff3f03ff3a2daa8d1329a6a6d17a58c422cbe5ee2f80aa0d0754b6fa85d507289806f64b5ce973e4ded8ec2afb631c2ee13b007eba090cda0a762a467cda2fb000a163abbde837a5d22dc526c8101ad9f32d93dfde1c67765a0508c8133b154fa93189918d82e1048ee0e4094064ddbcee9cfbecf6867546c2380",  # noqa: E501
            "0xf90151a000595129cf1c67ce97ed615df7184fb47a0d018d9147a2a7952a639e464cdfd380a052081f67d885b439bee62a1c4e09c02a1ec729aef4d92be6584a29141a79c5c7a0d99cae2bc3b6fc3e8e2767b4d7a5884b9b2d9d20da7890318beefcdb9c7346d58080a0f63571735d99e763dafadb03d1abe3a93a98740ecddc8514d13df46a35a8d94680a081870f4697e57436ccdd8bdc1d624d7bc07861748ff8b3a352334df68f2c03ada0f0252504ee8753335ecf0ca1f73cf13d4505783ed5110bf4e9b384a68ad88f69a0d38a18ef7b993c9ad2d848453e7722c26998ff3d63e243b99cdab0dabf84157ea06e5ad479dbda7f1647ed613e35d837f5d2c697ef34160f3665173186553ee9c280a0bf8371cac505b6ea0cb0f67ef744713f306a2455f3f8844cfa527c62db59c69980a0d6d689405b1cf73786c5fdabb55ec22f9b5712d1ba9c0bfcd46d4b93721fcac780",  # noqa: E501
            "0xf8669d37f29e142d49f8824d4e7f1735ec3da219687387629b5fccd86812df84b846f8440180a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a01d93f60f105899172f7255c030301c3af4564edd4a48577dbdc448aec7ddb0ac",  # noqa: E501
        ],
        "address": "0x7f0d15c7faae65896648c8273b6d7e43f58fa842",
        "balance": "0x0",
        "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",  # noqa: E501
        "nonce": "0x0",
        "storageHash": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",  # noqa: E501
        "storageProof": [
            {
                "key": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",  # gitleaks:allow noqa: E501
                "proof": [],
                "value": "0x0",
            }
        ],
    },
}

# curl https://mainnet.infura.io/v3/YOUR_KEY_HERE \
#   -X POST \
#   -H "Content-Type: application/json" \
#   --data '{"method": "eth_getTransactionReceipt",
#   "params": ["0x88df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713944b"],
#   "id": 1,"jsonrpc": "2.0"}'
EXPECTED_ACCOUNT_PROOF = [
    "0x0ac1030a150157f96e6b86cdefdb3d412547816a82e3e0ebf9d212e8010a1e2f65746865726d696e742e74797065732e76312e4574684163636f756e7412c5010a7f0a2a63726331326c756b75367578656868616b303270793472637a36357a753073776837776a737277307070124f0a282f65746865726d696e742e63727970746f2e76312e657468736563703235366b312e5075624b657912230a21026e710a62a342de0ed4d7c4532dcbcbbafbf19652ed67b237efab70e8b207efac200112423078633564323436303138366637323333633932376537646232646363373033633065353030623635336361383232373362376266616438303435643835613437301a0b0801180120012a03000202222b08011204020402201a2120e2a580b805b8b4d1f80793092cefe57965d9582ba8e31505a72cf31a55fa173d222b08011204040802201a21206c0bd60f0d887a5f99dd023f133dfd12412b074c0c442ab8a25b17048ff34ae022290801122506100220a66ec49c8058aef1eabb21a9610e16227d95982482db0ab5b032f823d457c2b920222b080112040a2e02201a2120af3944b847407590f1b2a95edcd1c2c5408e27a404e0b22267bcc2d46df700ed",  # noqa: E501
    "0x0aff010a0361636312205632d1620b29d277324e7473898ac9e587a99a62022931b33ce6d0be3b19137b1a090801180120012a0100222708011201011a20f5b8da7f66d134e34242575499b1f125c07af21b36685c31f9a8999c71a21daf222708011201011a20b8ef619094b2b40ae87dd1d3413c5af70d9c43256dfea123ca835baa53ad54a4222708011201011a20b7d06a60784c017041761d99ed17a2d2d68ceb4eccf39bd9a527773f21e5b688222708011201011a20a54531b68a71accb459f478724bde077a82c9079ef2be7e0b8dc47764f1b96d2222708011201011a204f83dd3fed3f48e1dff2fe125c5d22f578b613ddb45bdc12d4bc1841232d9a8b",  # noqa: E501
]

EXPECTED_STORAGE_PROOF = [
    "0x12370a350257f96e6b86cdefdb3d412547816a82e3e0ebf9d20000000000000000000000000000000000000000000000000000000000000000",  # noqa: E501
    "0x0af9010a0365766d1220e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b8551a090801180120012a0100222508011221010842b5561422ad68d28682baffad7f99cd8a64a339d4bd62b72916894d3d9a19222508011221010b1d3ff3c723ae07cea430234113b9d2b3b4218b7661f596b39b0592aedf9b60222508011221019975643a5ffdd52f7ef9a9cc57244008e5e7269e55587af4e4f137e5ef8ed6d3222708011201011a20a54531b68a71accb459f478724bde077a82c9079ef2be7e0b8dc47764f1b96d2222708011201011a204f83dd3fed3f48e1dff2fe125c5d22f578b613ddb45bdc12d4bc1841232d9a8b",  # noqa: E501
]

EXPECTED_GET_TRANSACTION = {
    "jsonrpc": "2.0",
    "id": 0,
    "result": {
        "blockHash": "0x1d59ff54b1eb26b013ce3cb5fc9dab3705b415a67127a003c3e61eb445bb8df2",  # noqa: E501
        "blockNumber": "0x5daf3b",
        "chainId": "0x1",
        "from": "0xa7d9ddbe1f17865597fbd27ec712455208b6b76d",
        "gas": "0xc350",
        "gasPrice": "0x4a817c800",
        "hash": "0x88df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713944b",  # noqa: E501
        "input": "0x68656c6c6f21",
        "nonce": "0x15",
        "r": "0x1b5e176d927f8e9ab405058b2d2457392da3e20f328b16ddabcebc33eaac5fea",
        "s": "0x4ba69724e8f69de52f0125ad8b3c5c2cef33019bac3249e2c0a2192766d1721c",
        "to": "0xf02c1c8e6114b1dbe8937a39260b5b0a374432bb",
        "transactionIndex": "0x41",
        "type": "0x0",
        "v": "0x25",
        "value": "0xf3dbb76162000",
    },
}

# curl https://mainnet.infura.io/v3/YOUR_KEY_HERE \
# -X POST \
# -H "Content-Type: application/json" \
# -d '{"jsonrpc": "2.0","method": "eth_feeHistory",
# "params": [4, "latest", [25, 75]],"id": 1}'

EXPECTED_FEE_HISTORY = {
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "baseFeePerGas": [
            "0x3043e3967",
            "0x2ffe341bf",
            "0x309363803",
            "0x2b09095a4",
            "0x2994cf12a",
        ],
        "gasUsedRatio": [0.4774405666666667, 0.5485699666666667, 0.0437691, 0.3648538],
        "oldestBlock": "0xf15b41",
        "reward": [
            ["0x3b9aca00", "0x77359400"],
            ["0x59682f00", "0x9502f900"],
            ["0x59682f00", "0x59682f00"],
            ["0x21797ced", "0x59682f00"],
        ],
    },
}

# curl https://mainnet.infura.io/v3/YOUR_KEY_HERE \
#   -X POST \
#   -H "Content-Type: application/json" \
#   --data '{"method": "eth_getTransactionReceipt",
#   "params": ["0xe25a4c707d3f981afa8ec103b1bc0bee9c6b7bea75e76ffa0a221d5239a7066a"],"
#    id": 1,"jsonrpc": "2.0"}'

EXPECTED_GET_TRANSACTION_RECEIPT = {
    "jsonrpc": "2.0",
    "result": {
        "blockHash": "0x42aa557e22141e0eba2d42258c010820ef080f2c7ecd0ad18cb53047b7e0421f",  # noqa: E501
        "blockNumber": "0xa50131",
        "contractAddress": None,
        "cumulativeGasUsed": "0xb72e24",
        # "effectiveGasPrice": "0x147d357000",
        "from": "0x56ac35db407fe1fb4977edd41f712aa5d8e7eb08",
        "gasUsed": "0x5208",
        "logs": [],
        "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",  # noqa: E501
        "status": "0x1",
        "to": "0x56ac35db407fe1fb4977edd41f712aa5d8e7eb08",
        "transactionHash": "0xe25a4c707d3f981afa8ec103b1bc0bee9c6b7bea75e76ffa0a221d5239a7066a",  # noqa: E501
        "transactionIndex": "0x9a",
        "type": "0x0",
    },
    "id": 1,
}

EXPECTED_CALLTRACERS = {
    "from": "0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2",
    "gas": "0x0",
    "gasUsed": "0x0",
    "input": "0x",
    "output": "0x",
    "to": "0x378c50d9264c63f3f92b806d4ee56e9d86ffb3ec",
    "type": "CALL",
    "value": "0x64",
}

EXPECTED_STRUCT_TRACER = {
    "failed": False,
    "gas": 21000,
    "returnValue": "",
    "structLogs": [],
}

EXPECTED_CONTRACT_CREATE_TRACER = {
    "from": "0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2",
    "gas": "0x7ef9d",
    "gasUsed": "0x7ef9d",
    "input": "0x608060405234801561001057600080fd5b50604080518082018252600981526805465737445524332360bc1b60208083019182528351808501909452600484526315195cdd60e21b90840152815191929161005c91600391610156565b508051610070906004906020840190610156565b50505061008e336a52b7d2dcc80cd2e400000061009360201b60201c565b610250565b6001600160a01b0382166100ed5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546100ff91906101ef565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b82805461016290610215565b90600052602060002090601f01602090048101928261018457600085556101ca565b82601f1061019d57805160ff19168380011785556101ca565b828001600101855582156101ca579182015b828111156101ca5782518255916020019190600101906101af565b506101d69291506101da565b5090565b5b808211156101d657600081556001016101db565b6000821982111561021057634e487b7160e01b600052601160045260246000fd5b500190565b600181811c9082168061022957607f821691505b6020821081141561024a57634e487b7160e01b600052602260045260246000fd5b50919050565b61085d8061025f6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101a0565b6040516100c3919061069a565b60405180910390f35b6100df6100da36600461070b565b610232565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610735565b61024a565b604051601281526020016100c3565b6100df61013136600461070b565b61026e565b6100f3610144366004610771565b6001600160a01b031660009081526020819052604090205490565b6100b6610290565b6100df61017536600461070b565b61029f565b6100df61018836600461070b565b61031f565b6100f361019b366004610793565b61032d565b6060600380546101af906107c6565b80601f01602080910402602001604051908101604052809291908181526020018280546101db906107c6565b80156102285780601f106101fd57610100808354040283529160200191610228565b820191906000526020600020905b81548152906001019060200180831161020b57829003601f168201915b5050505050905090565b600033610240818585610358565b5060019392505050565b60003361025885828561047c565b6102638585856104f6565b506001949350505050565b600033610240818585610281838361032d565b61028b9190610801565b610358565b6060600480546101af906107c6565b600033816102ad828661032d565b9050838110156103125760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102638286868403610358565b6000336102408185856104f6565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103ba5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610309565b6001600160a01b03821661041b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610309565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6000610488848461032d565b905060001981146104f057818110156104e35760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610309565b6104f08484848403610358565b50505050565b6001600160a01b03831661055a5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610309565b6001600160a01b0382166105bc5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610309565b6001600160a01b038316600090815260208190526040902054818110156106345760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610309565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36104f0565b600060208083528351808285015260005b818110156106c7578581018301518582016040015282016106ab565b818111156106d9576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461070657600080fd5b919050565b6000806040838503121561071e57600080fd5b610727836106ef565b946020939093013593505050565b60008060006060848603121561074a57600080fd5b610753846106ef565b9250610761602085016106ef565b9150604084013590509250925092565b60006020828403121561078357600080fd5b61078c826106ef565b9392505050565b600080604083850312156107a657600080fd5b6107af836106ef565b91506107bd602084016106ef565b90509250929050565b600181811c908216806107da57607f821691505b602082108114156107fb57634e487b7160e01b600052602260045260246000fd5b50919050565b6000821982111561082257634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220fee840cc9e1cfac46e073588ade030be1401c580c5849dd4e63f659a75eb220c64736f6c634300080a0033",  # noqa: E501
    "output": "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101a0565b6040516100c3919061069a565b60405180910390f35b6100df6100da36600461070b565b610232565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610735565b61024a565b604051601281526020016100c3565b6100df61013136600461070b565b61026e565b6100f3610144366004610771565b6001600160a01b031660009081526020819052604090205490565b6100b6610290565b6100df61017536600461070b565b61029f565b6100df61018836600461070b565b61031f565b6100f361019b366004610793565b61032d565b6060600380546101af906107c6565b80601f01602080910402602001604051908101604052809291908181526020018280546101db906107c6565b80156102285780601f106101fd57610100808354040283529160200191610228565b820191906000526020600020905b81548152906001019060200180831161020b57829003601f168201915b5050505050905090565b600033610240818585610358565b5060019392505050565b60003361025885828561047c565b6102638585856104f6565b506001949350505050565b600033610240818585610281838361032d565b61028b9190610801565b610358565b6060600480546101af906107c6565b600033816102ad828661032d565b9050838110156103125760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102638286868403610358565b6000336102408185856104f6565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103ba5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610309565b6001600160a01b03821661041b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610309565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6000610488848461032d565b905060001981146104f057818110156104e35760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610309565b6104f08484848403610358565b50505050565b6001600160a01b03831661055a5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610309565b6001600160a01b0382166105bc5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610309565b6001600160a01b038316600090815260208190526040902054818110156106345760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610309565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36104f0565b600060208083528351808285015260005b818110156106c7578581018301518582016040015282016106ab565b818111156106d9576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461070657600080fd5b919050565b6000806040838503121561071e57600080fd5b610727836106ef565b946020939093013593505050565b60008060006060848603121561074a57600080fd5b610753846106ef565b9250610761602085016106ef565b9150604084013590509250925092565b60006020828403121561078357600080fd5b61078c826106ef565b9392505050565b600080604083850312156107a657600080fd5b6107af836106ef565b91506107bd602084016106ef565b90509250929050565b600181811c908216806107da57607f821691505b602082108114156107fb57634e487b7160e01b600052602260045260246000fd5b50919050565b6000821982111561082257634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220fee840cc9e1cfac46e073588ade030be1401c580c5849dd4e63f659a75eb220c64736f6c634300080a0033",  # noqa: E501
    "to": "0x8c76cfc1934d5120cc673b6e5ddf7b88feb1c18c",
    "type": "CREATE",
    "value": "0x0",
}
