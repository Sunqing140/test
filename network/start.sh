#!/bin/bash

if [[ `uname` == 'Darwin' ]]; then
    echo "Mac OS"
    export PATH=${PWD}/hyperledger-fabric-darwin-amd64-1.4.12/bin:$PATH
fi
if [[ `uname` == 'Linux' ]]; then
    echo "Linux"
    export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH
fi

echo "一、清理环境"
./stop.sh

echo "二、生成证书和秘钥（ MSP 材料），生成结果将保存在 crypto-config 文件夹中"
cryptogen generate --config=./crypto-config.yaml

echo "三、创建排序通道创世区块"
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block -channelID firstchannel

echo "四、生成通道配置事务'appchannel.tx'"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/appchannel.tx -channelID appchannel

echo "五、为 EN1 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/EN1Anchor.tx -channelID appchannel -asOrg EN1

echo "六、为 EN2 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/EN2Anchor.tx -channelID appchannel -asOrg EN2

echo "区块链 ： 启动"
docker-compose up -d
echo "正在等待节点的启动完成，等待10秒"
sleep 10

EN1Peer0Cli="CORE_PEER_ADDRESS=peer0.enterprise1.com:7051 CORE_PEER_LOCALMSPID=EN1MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/enterprise1.com/users/Admin@enterprise1.com/msp"
EN1Peer1Cli="CORE_PEER_ADDRESS=peer1.enterprise1.com:7051 CORE_PEER_LOCALMSPID=EN1MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/enterprise1.com/users/Admin@enterprise1.com/msp"
EN2Peer0Cli="CORE_PEER_ADDRESS=peer0.enterprise2.com:7051 CORE_PEER_LOCALMSPID=EN2MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/enterprise2.com/users/Admin@enterprise2.com/msp"
EN2Peer1Cli="CORE_PEER_ADDRESS=peer1.enterprise2.com:7051 CORE_PEER_LOCALMSPID=EN2MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/enterprise2.com/users/Admin@enterprise2.com/msp"

echo "七、创建通道"
docker exec cli bash -c "$EN1Peer0Cli peer channel create -o orderer.cowunion.com:7050 -c appchannel -f /etc/hyperledger/config/appchannel.tx"

echo "八、将所有节点加入通道"
docker exec cli bash -c "$EN1Peer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$EN1Peer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$EN2Peer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$EN2Peer1Cli peer channel join -b appchannel.block"

echo "九、更新锚节点"
docker exec cli bash -c "$EN1Peer0Cli peer channel update -o orderer.cowunion.com:7050 -c appchannel -f /etc/hyperledger/config/EN1Anchor.tx"
docker exec cli bash -c "$EN2Peer0Cli peer channel update -o orderer.cowunion.com:7050 -c appchannel -f /etc/hyperledger/config/EN2Anchor.tx"

# -n 链码名，可以自己随便设置
# -v 版本号
# -p 链码目录，在 /opt/gopath/src/ 目录下
echo "十、安装链码"
docker exec cli bash -c "$EN1Peer0Cli peer chaincode install -n cow-trace -v 1.0.0 -l golang -p chaincode"
docker exec cli bash -c "$EN2Peer0Cli peer chaincode install -n cow-trace -v 1.0.0 -l golang -p chaincode"

# 只需要其中一个节点实例化
# -n 对应上一步安装链码的名字
# -v 版本号
# -C 是通道，在fabric的世界，一个通道就是一条不同的链
# -c 为传参，传入init参数
echo "十一、实例化链码"
docker exec cli bash -c "$EN1Peer0Cli peer chaincode instantiate -o orderer.cowunion.com:7050 -C appchannel -n cow-trace -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"AND ('EN1MSP.member','EN2MSP.member')\""

echo "正在等待链码实例化完成，等待5秒"
sleep 5

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
echo "十二、验证链码"
docker exec cli bash -c "$EN1Peer0Cli peer chaincode invoke -C appchannel -n cow-trace -c '{\"Args\":[\"hello\"]}'"
docker exec cli bash -c "$EN2Peer0Cli peer chaincode invoke -C appchannel -n cow-trace -c '{\"Args\":[\"hello\"]}'"
