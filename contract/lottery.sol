pragma solidity ^0.4.25;

contract lotteryRecord {
    struct TRC20Data {
        address sender;
        address receiver;
        string orderHash; // txhash
        string data;
        bool used;
    }

    mapping(string => TRC20Data) transactionByTx;

    /**
     * 保存交易数据  输入数据, 返回bcos链hash
     * @param _sender 公链交易发送方
     * @param _receiver 公链交易接收方
     * @param _orderHash  公链交易id
     * @param _data 交易数据
     */
    function saveTransactionData(
        address _sender,
        address _receiver,
        string memory _orderHash,
        string memory _data
    ) public returns (bytes32) {
        bytes32 txHash =
        keccak256(abi.encodePacked(_sender, _receiver, _orderHash, _data, block.timestamp));
        require(!isTxExists(_orderHash), "TRC20Data already exists");
        TRC20Data storage transaction = transactionByTx[_orderHash];
        transaction.sender = _sender;
        transaction.receiver = _receiver;
        transaction.orderHash = _orderHash;
        transaction.data = _data;
        transaction.used = true;
        return txHash;
    }


    /**
     * 获取公链交易数据： 输入公链交易hash, 返回数值
     * @param _txStr  公链交易id
     */
    function getTransactionData(string _txStr)
    public view returns (
        address sender,
        address receiver,
        string memory orderHash,
        string memory data
    )
    {
        require(isTxExists(_txStr), "Transaction does not exist");
        TRC20Data storage transaction = transactionByTx[_txStr];
        return (transaction.sender, transaction.receiver, transaction.orderHash, transaction.data);
    }

    function isTxExists(string memory key) public view returns (bool) {
        return transactionByTx[key].used;
    }

    function destroy() public{
        selfdestruct(msg.sender);
    }
}