package client

import (
	"gitlab.33.cn/chain33/chain33/types"
)

// 消息通道交互API接口定义
type QueueProtocolAPI interface {
	Version() (*types.Reply, error)
	Close()
	// +++++++++++++++ mempool interfaces begin
	// 同步发送交易信息到指定模块，获取应答消息 types.EventTx
	SendTx(param *types.Transaction) (*types.Reply, error)
	// types.EventTxList
	GetTxList(param *types.TxHashList) (*types.ReplyTxList, error)
	// types.EventGetMempool
	GetMempool() (*types.ReplyTxList, error)
	// types.EventGetLastMempool
	GetLastMempool() (*types.ReplyTxList, error)
	// types.EventQuery
	Query(param *types.Query) (*types.Message, error)
	// --------------- mempool interfaces end

	// +++++++++++++++ p2p interfaces begin
	// types.EventPeerInfo
	PeerInfo() (*types.PeerList, error)
	// types.EventGetNetInfo
	GetNetInfo() (*types.NodeNetInfo, error)
	// --------------- p2p interfaces end

	// +++++++++++++++ consensus interfaces begin
	// types.EventGetTicketCount
	GetTicketCount() (*types.Int64, error)
	// --------------- consensus interfaces end

	// +++++++++++++++ wallet interfaces begin
	// types.EventLocalGet
	LocalGet(param *types.ReqHash) (*types.LocalReplyValue, error)
	// types.EventWalletGetAccountList
	WalletGetAccountList() (*types.WalletAccounts, error)
	// types.EventNewAccount
	NewAccount(param *types.ReqNewAccount) (*types.WalletAccount, error)
	// types.EventWalletTransactionList
	WalletTransactionList(param *types.ReqWalletTransactionList) (*types.WalletTxDetails, error)
	// types.EventWalletImportprivkey
	WalletImportprivkey(param *types.ReqWalletImportPrivKey) (*types.WalletAccount, error)
	// types.EventWalletSendToAddress
	WalletSendToAddress(param *types.ReqWalletSendToAddress) (*types.ReplyHash, error)
	// types.EventWalletSetFee
	WalletSetFee(param *types.ReqWalletSetFee) (*types.Reply, error)
	// types.EventWalletSetLabel
	WalletSetLabel(param *types.ReqWalletSetLabel) (*types.WalletAccount, error)
	// types.EventWalletMergeBalance
	WalletMergeBalance(param *types.ReqWalletMergeBalance) (*types.ReplyHashes, error)
	// types.EventWalletSetPasswd
	WalletSetPasswd(param *types.ReqWalletSetPasswd) (*types.Reply, error)
	// types.EventWalletLock
	WalletLock() (*types.Reply, error)
	// types.EventWalletUnLock
	WalletUnLock(param *types.WalletUnLock) (*types.Reply, error)
	// types.EventGenSeed
	GenSeed(param *types.GenSeedLang) (*types.ReplySeed, error)
	// types.EventSaveSeed
	SaveSeed(param *types.SaveSeedByPw) (*types.Reply, error)
	// types.EventGetSeed
	GetSeed(param *types.GetSeedByPw) (*types.ReplySeed, error)
	// types.EventGetWalletStatus
	GetWalletStatus() (*types.WalletStatus, error)
	// types.EventWalletAutoMiner
	WalletAutoMiner(param *types.MinerFlag) (*types.Reply, error)
	// types.EventDumpPrivkey
	DumpPrivkey(param *types.ReqStr) (*types.ReplyStr, error)
	// types.EventCloseTickets
	CloseTickets() (*types.ReplyHashes, error)
	// types.EventSignRawTx
	SignRawTx(param *types.ReqSignRawTx) (*types.ReplySignRawTx, error)
	// --------------- wallet interfaces end

	// +++++++++++++++ blockchain interfaces begin
	// types.EventGetBlocks
	GetBlocks(param *types.ReqBlocks) (*types.BlockDetails, error)
	// types.EventQueryTx
	QueryTx(param *types.ReqHash) (*types.TransactionDetail, error)
	// types.EventGetTransactionByAddr
	GetTransactionByAddr(param *types.ReqAddr) (*types.ReplyTxInfos, error)
	// types.EventGetTransactionByHash
	GetTransactionByHash(param *types.ReqHashes) (*types.TransactionDetails, error)
	// types.EventGetHeaders
	GetHeaders(param *types.ReqBlocks) (*types.Headers, error)
	// types.EventGetBlockOverview
	GetBlockOverview(param *types.ReqHash) (*types.BlockOverview, error)
	// types.EventGetAddrOverview
	GetAddrOverview(param *types.ReqAddr) (*types.AddrOverview, error)
	// types.EventGetBlockHash
	GetBlockHash(param *types.ReqInt) (*types.ReplyHash, error)
	// types.EventIsSync
	IsSync() (*types.Reply, error)
	// types.EventIsNtpClockSync
	IsNtpClockSync() (*types.Reply, error)
	// types.EventGetLastHeader
	GetLastHeader() (*types.Header, error)
	// --------------- blockchain interfaces end

	// +++++++++++++++ store interfaces begin
	StoreGet(*types.StoreGet) (*types.StoreReplyValue, error)
	StoreGetTotalCoins(*types.IterateRangeByStateHash) (*types.ReplyGetTotalCoins, error)
	// --------------- store interfaces end

	// +++++++++++++++ other interfaces begin
	// --------------- other interfaces end
}