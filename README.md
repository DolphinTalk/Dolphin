# Dolphin白皮书
## 概述
### 什么是Dolphin
Dolphin是基于区块链技术的聊天网络。
有别于传统IM，Dolphin利用区块链技术实现IM数据的公开存储，从而保证用户的隐私。
因为数据的加密和使用方式都是公开透明的，通过算法保证了对用户数据隐私保护的承诺。
另外，Dolphin本身是一个去中心化网络，存在在Dolphin网络里面的好友关系和用户数据，永远属于用户，用户可以授权给任何服务使用。
Dolphin不再以用户关系作为黏性要挟用户，而且Dolphin也支持与主流通讯协议的互联互通。

### Dolpin构建的社交网络
#### 对于个人数据的尊重
个人数据，无论好友关系还是文件、照片等，所有权永远属于个人，任何服务的提供方在未经授权的下无法查看。这个由区块链的智能合约的隐私保护技术保证。而且，通过去中心化的网络，这些数据不属于任何运营商，也不存在丢失的可能。
第三方应用无法分析用户数据，这让我们的生活更加安全可靠。
#### 对于个人使用习惯的尊重
现在社交网络是割裂的，我们用了钉钉就无法跟微信交流。我们希望Dolpin所构建的通讯网络里面，所有的IM能够实现互联互通，就像IPhone能给HUAWEI Mate20打电话一样，用户有权选择通讯的终端程序和体验，我们不应该被微信所绑架。
#### 对于细分场景的支持
我们在任何细分场合，无论是淘宝购物还是小红书看网红，我们都能通过Dophin理解世界，我们不在频繁切换IM聊天工具，因为通讯是一个基础能力，被所有App所支持

### Dolpin的特点
它是一个标准化的通讯区块链网络，它将一个业务的后台服务完全用区块链重构，并将架构和数据公开给全社会，受全社会监督。
因此Dolphin所承诺的数据安全、隐私保护，以及互联互通，完全是可信的。
它提供了标准通讯协议，兼容XMPP等主流及时通讯协议，允许任意遵循协议的IM接入，并在用户授权的前提下使用数据

## Dolphin网络
<img src=https://github.com/DolphinTalk/Dolphin/blob/master/rc/Dolphin-network.png width=75%></img>

### 用户系统
#### 账户系统
我们采用逻辑账号系统，并记录对应的公钥，持有公钥的私钥才能操作用户账号和数据。
如果私钥丢失，我们通过社区治理手段可以找回。
#### 用户数据
在整个区块链网络里面，核心存储好友关系数据。都通过公钥做非对称加密，只有持有私钥方能解开。
另外，我们对所有加密数据放入噪音，保证无法暴利破解。
### IM运营商
IM运营商定义：想接入Dolphin网络的IM服务提供方。我们为这些服务提供方提供四套标准解决方案Client、Server、GateWay、dolphin-BlockChain。
Dolphin本身就是一个IM网咯运营商，免费提供IM通讯软件。
#### Client
连接网络的客户端SDK，支持IOS、Android、Windows、MacOS、以及多种IOT设备
#### Server
与IM维持长连接的服务器，保证消息在IM与服务器之间的双向传输。一个IM运营商允许部署多个Server，支持大量的IM连接
#### Gateway
IM运营商之间传递信息的网关通道，优化网络拓扑
#### dolphin-Blockchain
存储用户登录Session和好友关系的网络
### 监管责任
所有监管责任由IM运营商负责，IM运营商有义务负责政府对IM的监管需求。

### 网络激励
#### 如何使用网络
数据在网络上是加密存储的，读网络数据是免费的。但是如果需要发生更新，用户需要支付一定的费用(GAS)。包括更改用户状态，更新用户好友关系。
用户需要通过IM运营商去修改网络数据，也可以执行通过Dolphin浏览器更改。更改状态需要消耗用户的GAS或者是运营商的GAS。
#### 如何搭建网络
Dolphin是符合监管的半公开网络，它是通过DPOS共识协议构建的区块链网络。任何IM运营商都可以作为结点加入，并通过大家选举参与记账。
#### GAS的分发和使用
GAS是Dolphin网络运转的分配凭证，参与提供Dolphin网络的节点都会获得Dolphin，包括手续费和挖矿所得。
为了保证前期网络有人参与，保证前期参与者的利益，GAS总量恒定，且不允许挖矿。
Dolphin IM的创造者发行10亿Dolphin，用于支持Dolphin网络运转。IM使用者充值获得Dolphin，并享受服务。
Dolphin网络将手续费全部分给记账结点，用来激励他们对网络的共享。
前期为了扶持记账结点，Dolphin的拥有着会拿20%的持续激励给矿工。
* *注意*：
	* GAS不能点对点转账，只能通过平台充值获得。
	* GAS不对普通用户可见，只用作网络服务结算。
	* GAS超级结点因为提供服务获取用户GAS，可到充值平台获取网络服务分成。

### 项目如何开展
不募资，完全靠开源贡献
