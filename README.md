## 代码结构
```md
├─app
│  └─controllers.go
│  └─models.go
│  └─routers.go
├─configs
│  └─sdk_cert
│  └─config.go
├─contract
└─pkg
│  └─http
├─main

```
### 前置条件
go1.18
bcos 2.9

### 功能
+ 交易数据保存
+ 交易数据查询

## 编译智能合约
### abi/bin 文件
+ 将.sol 利用webase/remix 在线ide编译，生成abi和bin文件
+ 利用abigen生成go文件
  ```./abigen -abi lottery.abi -bin lottery.bin -type lottery -pkg contract -out lottery.go```
### 部署
利用console进行部署生成合约地址，并配置在config.yml里面
```markdown
cp lottery.sol复制到 console/contracts/solidity下
cd ~/fisco/console && bash start.sh
deploy lottery
```
deploy成功后会展示合约信息
```
transaction hash: 0x0000000000000
contract address: 0xfe0000000
currentAccount: 0x000000000
```


## todo
+ 修改web 框架 为fasthttp或者httprouter
+ 查询添加缓存
+ 创建并发条件下设计




