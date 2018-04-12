// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 钱包模块存贮的tx交易详细信息
// 	 tx : tx交易信息
// 	 receipt :交易收据信息
// 	 height :交易所在的区块高度
// 	 index :交易所在区块中的索引
// 	 blocktime :交易所在区块的时标
// 	 amount :交易量
// 	 fromaddr :交易打出地址
// 	 txhash : 交易对应的哈希值
// 	 actionName  :交易对应的函数调用
type WalletTxDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Height     int64        `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,6,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,7,opt,name=fromaddr" json:"fromaddr,omitempty"`
	Txhash     []byte       `protobuf:"bytes,8,opt,name=txhash,proto3" json:"txhash,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *WalletTxDetail) Reset()                    { *m = WalletTxDetail{} }
func (m *WalletTxDetail) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetail) ProtoMessage()               {}
func (*WalletTxDetail) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *WalletTxDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *WalletTxDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *WalletTxDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *WalletTxDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WalletTxDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *WalletTxDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WalletTxDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *WalletTxDetail) GetTxhash() []byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

func (m *WalletTxDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type WalletTxDetails struct {
	TxDetails []*WalletTxDetail `protobuf:"bytes,1,rep,name=txDetails" json:"txDetails,omitempty"`
}

func (m *WalletTxDetails) Reset()                    { *m = WalletTxDetails{} }
func (m *WalletTxDetails) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetails) ProtoMessage()               {}
func (*WalletTxDetails) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *WalletTxDetails) GetTxDetails() []*WalletTxDetail {
	if m != nil {
		return m.TxDetails
	}
	return nil
}

// 钱包模块存贮的账户信息
// 	 privkey : 账户地址对应的私钥
// 	 label :账户地址对应的标签
// 	 addr :账户地址
// 	 timeStamp :创建账户时的时标
type WalletAccountStore struct {
	Privkey   string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label     string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Addr      string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	TimeStamp string `protobuf:"bytes,4,opt,name=timeStamp" json:"timeStamp,omitempty"`
}

func (m *WalletAccountStore) Reset()                    { *m = WalletAccountStore{} }
func (m *WalletAccountStore) String() string            { return proto.CompactTextString(m) }
func (*WalletAccountStore) ProtoMessage()               {}
func (*WalletAccountStore) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *WalletAccountStore) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *WalletAccountStore) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WalletAccountStore) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *WalletAccountStore) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

// 钱包模块通过一个随机值对钱包密码加密
// 	 pwHash : 对钱包密码和一个随机值组合进行哈希计算
// 	 randstr :对钱包密码加密的一个随机值
type WalletPwHash struct {
	PwHash  []byte `protobuf:"bytes,1,opt,name=pwHash,proto3" json:"pwHash,omitempty"`
	Randstr string `protobuf:"bytes,2,opt,name=randstr" json:"randstr,omitempty"`
}

func (m *WalletPwHash) Reset()                    { *m = WalletPwHash{} }
func (m *WalletPwHash) String() string            { return proto.CompactTextString(m) }
func (*WalletPwHash) ProtoMessage()               {}
func (*WalletPwHash) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *WalletPwHash) GetPwHash() []byte {
	if m != nil {
		return m.PwHash
	}
	return nil
}

func (m *WalletPwHash) GetRandstr() string {
	if m != nil {
		return m.Randstr
	}
	return ""
}

// 钱包当前的状态
// 	 isWalletLock : 钱包是否锁状态，true锁定，false解锁
// 	 isAutoMining :钱包是否开启挖矿功能，true开启挖矿，false关闭挖矿
// 	 isHasSeed : 钱包是否有种子，true已有，false没有
// 	 isTicketLock :钱包挖矿买票锁状态，true锁定，false解锁，只能用于挖矿转账
type WalletStatus struct {
	IsWalletLock bool `protobuf:"varint,1,opt,name=isWalletLock" json:"isWalletLock,omitempty"`
	IsAutoMining bool `protobuf:"varint,2,opt,name=isAutoMining" json:"isAutoMining,omitempty"`
	IsHasSeed    bool `protobuf:"varint,3,opt,name=isHasSeed" json:"isHasSeed,omitempty"`
	IsTicketLock bool `protobuf:"varint,4,opt,name=isTicketLock" json:"isTicketLock,omitempty"`
}

func (m *WalletStatus) Reset()                    { *m = WalletStatus{} }
func (m *WalletStatus) String() string            { return proto.CompactTextString(m) }
func (*WalletStatus) ProtoMessage()               {}
func (*WalletStatus) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *WalletStatus) GetIsWalletLock() bool {
	if m != nil {
		return m.IsWalletLock
	}
	return false
}

func (m *WalletStatus) GetIsAutoMining() bool {
	if m != nil {
		return m.IsAutoMining
	}
	return false
}

func (m *WalletStatus) GetIsHasSeed() bool {
	if m != nil {
		return m.IsHasSeed
	}
	return false
}

func (m *WalletStatus) GetIsTicketLock() bool {
	if m != nil {
		return m.IsTicketLock
	}
	return false
}

type WalletAccounts struct {
	Wallets []*WalletAccount `protobuf:"bytes,1,rep,name=wallets" json:"wallets,omitempty"`
}

func (m *WalletAccounts) Reset()                    { *m = WalletAccounts{} }
func (m *WalletAccounts) String() string            { return proto.CompactTextString(m) }
func (*WalletAccounts) ProtoMessage()               {}
func (*WalletAccounts) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *WalletAccounts) GetWallets() []*WalletAccount {
	if m != nil {
		return m.Wallets
	}
	return nil
}

type WalletAccount struct {
	Acc   *Account `protobuf:"bytes,1,opt,name=acc" json:"acc,omitempty"`
	Label string   `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *WalletAccount) Reset()                    { *m = WalletAccount{} }
func (m *WalletAccount) String() string            { return proto.CompactTextString(m) }
func (*WalletAccount) ProtoMessage()               {}
func (*WalletAccount) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *WalletAccount) GetAcc() *Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *WalletAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 钱包解锁
// 	 passwd : 钱包密码
// 	 timeout :钱包解锁时间，0，一直解锁，非0值，超时之后继续锁定
// 	 walletOrTicket :解锁整个钱包还是只解锁挖矿买票功能，1只解锁挖矿买票，0解锁整个钱包
type WalletUnLock struct {
	Passwd         string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
	Timeout        int64  `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	WalletOrTicket bool   `protobuf:"varint,3,opt,name=walletOrTicket" json:"walletOrTicket,omitempty"`
}

func (m *WalletUnLock) Reset()                    { *m = WalletUnLock{} }
func (m *WalletUnLock) String() string            { return proto.CompactTextString(m) }
func (*WalletUnLock) ProtoMessage()               {}
func (*WalletUnLock) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *WalletUnLock) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *WalletUnLock) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *WalletUnLock) GetWalletOrTicket() bool {
	if m != nil {
		return m.WalletOrTicket
	}
	return false
}

type GenSeedLang struct {
	Lang int32 `protobuf:"varint,1,opt,name=lang" json:"lang,omitempty"`
}

func (m *GenSeedLang) Reset()                    { *m = GenSeedLang{} }
func (m *GenSeedLang) String() string            { return proto.CompactTextString(m) }
func (*GenSeedLang) ProtoMessage()               {}
func (*GenSeedLang) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *GenSeedLang) GetLang() int32 {
	if m != nil {
		return m.Lang
	}
	return 0
}

type GetSeedByPw struct {
	Passwd string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *GetSeedByPw) Reset()                    { *m = GetSeedByPw{} }
func (m *GetSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*GetSeedByPw) ProtoMessage()               {}
func (*GetSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *GetSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

// 存储钱包的种子
// 	 seed : 钱包种子
// 	 passwd :钱包密码
type SaveSeedByPw struct {
	Seed   string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *SaveSeedByPw) Reset()                    { *m = SaveSeedByPw{} }
func (m *SaveSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*SaveSeedByPw) ProtoMessage()               {}
func (*SaveSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *SaveSeedByPw) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *SaveSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type ReplySeed struct {
	Seed string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
}

func (m *ReplySeed) Reset()                    { *m = ReplySeed{} }
func (m *ReplySeed) String() string            { return proto.CompactTextString(m) }
func (*ReplySeed) ProtoMessage()               {}
func (*ReplySeed) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *ReplySeed) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type ReqWalletSetPasswd struct {
	Oldpass string `protobuf:"bytes,1,opt,name=oldpass" json:"oldpass,omitempty"`
	Newpass string `protobuf:"bytes,2,opt,name=newpass" json:"newpass,omitempty"`
}

func (m *ReqWalletSetPasswd) Reset()                    { *m = ReqWalletSetPasswd{} }
func (m *ReqWalletSetPasswd) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetPasswd) ProtoMessage()               {}
func (*ReqWalletSetPasswd) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *ReqWalletSetPasswd) GetOldpass() string {
	if m != nil {
		return m.Oldpass
	}
	return ""
}

func (m *ReqWalletSetPasswd) GetNewpass() string {
	if m != nil {
		return m.Newpass
	}
	return ""
}

type ReqNewAccount struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
}

func (m *ReqNewAccount) Reset()                    { *m = ReqNewAccount{} }
func (m *ReqNewAccount) String() string            { return proto.CompactTextString(m) }
func (*ReqNewAccount) ProtoMessage()               {}
func (*ReqNewAccount) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *ReqNewAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type MinerFlag struct {
	Flag    int32 `protobuf:"varint,1,opt,name=flag" json:"flag,omitempty"`
	Reserve int64 `protobuf:"varint,2,opt,name=reserve" json:"reserve,omitempty"`
}

func (m *MinerFlag) Reset()                    { *m = MinerFlag{} }
func (m *MinerFlag) String() string            { return proto.CompactTextString(m) }
func (*MinerFlag) ProtoMessage()               {}
func (*MinerFlag) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{14} }

func (m *MinerFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *MinerFlag) GetReserve() int64 {
	if m != nil {
		return m.Reserve
	}
	return 0
}

// 获取钱包交易的详细信息
// 	 fromTx : []byte( Sprintf("%018d", height*100000 + index)，
// 				表示从高度 height 中的 index 开始获取交易列表；
// 			    第一次传参为空，获取最新的交易。)
// 	 count :获取交易列表的个数。
// 	 direction :查找方式；0，上一页；1，下一页。
type ReqWalletTransactionList struct {
	FromTx    []byte `protobuf:"bytes,1,opt,name=fromTx,proto3" json:"fromTx,omitempty"`
	Count     int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
}

func (m *ReqWalletTransactionList) Reset()                    { *m = ReqWalletTransactionList{} }
func (m *ReqWalletTransactionList) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletTransactionList) ProtoMessage()               {}
func (*ReqWalletTransactionList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{15} }

func (m *ReqWalletTransactionList) GetFromTx() []byte {
	if m != nil {
		return m.FromTx
	}
	return nil
}

func (m *ReqWalletTransactionList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqWalletTransactionList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ReqWalletImportPrivKey struct {
	Privkey string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label   string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletImportPrivKey) Reset()                    { *m = ReqWalletImportPrivKey{} }
func (m *ReqWalletImportPrivKey) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletImportPrivKey) ProtoMessage()               {}
func (*ReqWalletImportPrivKey) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{16} }

func (m *ReqWalletImportPrivKey) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *ReqWalletImportPrivKey) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 发送交易
// 	 from : 打出地址
// 	 to :接受地址
// 	 amount : 转账额度
// 	 note :转账备注
type ReqWalletSendToAddress struct {
	From        string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To          string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Note        string `protobuf:"bytes,4,opt,name=note" json:"note,omitempty"`
	Istoken     bool   `protobuf:"varint,5,opt,name=istoken" json:"istoken,omitempty"`
	TokenSymbol string `protobuf:"bytes,6,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
}

func (m *ReqWalletSendToAddress) Reset()                    { *m = ReqWalletSendToAddress{} }
func (m *ReqWalletSendToAddress) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSendToAddress) ProtoMessage()               {}
func (*ReqWalletSendToAddress) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{17} }

func (m *ReqWalletSendToAddress) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqWalletSendToAddress) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetIstoken() bool {
	if m != nil {
		return m.Istoken
	}
	return false
}

func (m *ReqWalletSendToAddress) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

type ReqWalletSetFee struct {
	Amount int64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *ReqWalletSetFee) Reset()                    { *m = ReqWalletSetFee{} }
func (m *ReqWalletSetFee) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetFee) ProtoMessage()               {}
func (*ReqWalletSetFee) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{18} }

func (m *ReqWalletSetFee) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ReqWalletSetLabel struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletSetLabel) Reset()                    { *m = ReqWalletSetLabel{} }
func (m *ReqWalletSetLabel) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetLabel) ProtoMessage()               {}
func (*ReqWalletSetLabel) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{19} }

func (m *ReqWalletSetLabel) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqWalletSetLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ReqWalletMergeBalance struct {
	To string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
}

func (m *ReqWalletMergeBalance) Reset()                    { *m = ReqWalletMergeBalance{} }
func (m *ReqWalletMergeBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletMergeBalance) ProtoMessage()               {}
func (*ReqWalletMergeBalance) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{20} }

func (m *ReqWalletMergeBalance) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type ReqStr struct {
	Reqstr string `protobuf:"bytes,1,opt,name=reqstr" json:"reqstr,omitempty"`
}

func (m *ReqStr) Reset()                    { *m = ReqStr{} }
func (m *ReqStr) String() string            { return proto.CompactTextString(m) }
func (*ReqStr) ProtoMessage()               {}
func (*ReqStr) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{21} }

func (m *ReqStr) GetReqstr() string {
	if m != nil {
		return m.Reqstr
	}
	return ""
}

type ReplyStr struct {
	Replystr string `protobuf:"bytes,1,opt,name=replystr" json:"replystr,omitempty"`
}

func (m *ReplyStr) Reset()                    { *m = ReplyStr{} }
func (m *ReplyStr) String() string            { return proto.CompactTextString(m) }
func (*ReplyStr) ProtoMessage()               {}
func (*ReplyStr) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{22} }

func (m *ReplyStr) GetReplystr() string {
	if m != nil {
		return m.Replystr
	}
	return ""
}

type ReqTokenPreCreate struct {
	CreatorAddr  string `protobuf:"bytes,1,opt,name=creator_addr,json=creatorAddr" json:"creator_addr,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Symbol       string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
	Introduction string `protobuf:"bytes,4,opt,name=introduction" json:"introduction,omitempty"`
	OwnerAddr    string `protobuf:"bytes,5,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
	Total        int64  `protobuf:"varint,6,opt,name=total" json:"total,omitempty"`
	Price        int64  `protobuf:"varint,7,opt,name=price" json:"price,omitempty"`
}

func (m *ReqTokenPreCreate) Reset()                    { *m = ReqTokenPreCreate{} }
func (m *ReqTokenPreCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenPreCreate) ProtoMessage()               {}
func (*ReqTokenPreCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{23} }

func (m *ReqTokenPreCreate) GetCreatorAddr() string {
	if m != nil {
		return m.CreatorAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTokenPreCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenPreCreate) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *ReqTokenPreCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReqTokenPreCreate) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ReqTokenFinishCreate struct {
	FinisherAddr string `protobuf:"bytes,1,opt,name=finisher_addr,json=finisherAddr" json:"finisher_addr,omitempty"`
	Symbol       string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	OwnerAddr    string `protobuf:"bytes,3,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
}

func (m *ReqTokenFinishCreate) Reset()                    { *m = ReqTokenFinishCreate{} }
func (m *ReqTokenFinishCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenFinishCreate) ProtoMessage()               {}
func (*ReqTokenFinishCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{24} }

func (m *ReqTokenFinishCreate) GetFinisherAddr() string {
	if m != nil {
		return m.FinisherAddr
	}
	return ""
}

func (m *ReqTokenFinishCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenFinishCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

type ReqTokenRevokeCreate struct {
	RevokerAddr string `protobuf:"bytes,1,opt,name=revoker_addr,json=revokerAddr" json:"revoker_addr,omitempty"`
	Symbol      string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	OwnerAddr   string `protobuf:"bytes,3,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
}

func (m *ReqTokenRevokeCreate) Reset()                    { *m = ReqTokenRevokeCreate{} }
func (m *ReqTokenRevokeCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenRevokeCreate) ProtoMessage()               {}
func (*ReqTokenRevokeCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{25} }

func (m *ReqTokenRevokeCreate) GetRevokerAddr() string {
	if m != nil {
		return m.RevokerAddr
	}
	return ""
}

func (m *ReqTokenRevokeCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenRevokeCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

type ReqSellToken struct {
	Sell  *TradeForSell `protobuf:"bytes,1,opt,name=sell" json:"sell,omitempty"`
	Owner string        `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
}

func (m *ReqSellToken) Reset()                    { *m = ReqSellToken{} }
func (m *ReqSellToken) String() string            { return proto.CompactTextString(m) }
func (*ReqSellToken) ProtoMessage()               {}
func (*ReqSellToken) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{26} }

func (m *ReqSellToken) GetSell() *TradeForSell {
	if m != nil {
		return m.Sell
	}
	return nil
}

func (m *ReqSellToken) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReqBuyToken struct {
	Buy   *TradeForBuy `protobuf:"bytes,1,opt,name=buy" json:"buy,omitempty"`
	Buyer string       `protobuf:"bytes,2,opt,name=buyer" json:"buyer,omitempty"`
}

func (m *ReqBuyToken) Reset()                    { *m = ReqBuyToken{} }
func (m *ReqBuyToken) String() string            { return proto.CompactTextString(m) }
func (*ReqBuyToken) ProtoMessage()               {}
func (*ReqBuyToken) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{27} }

func (m *ReqBuyToken) GetBuy() *TradeForBuy {
	if m != nil {
		return m.Buy
	}
	return nil
}

func (m *ReqBuyToken) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

type ReqRevokeSell struct {
	Revoke *TradeForRevokeSell `protobuf:"bytes,1,opt,name=revoke" json:"revoke,omitempty"`
	Owner  string              `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
}

func (m *ReqRevokeSell) Reset()                    { *m = ReqRevokeSell{} }
func (m *ReqRevokeSell) String() string            { return proto.CompactTextString(m) }
func (*ReqRevokeSell) ProtoMessage()               {}
func (*ReqRevokeSell) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{28} }

func (m *ReqRevokeSell) GetRevoke() *TradeForRevokeSell {
	if m != nil {
		return m.Revoke
	}
	return nil
}

func (m *ReqRevokeSell) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReqModifyConfig struct {
	Key      string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Op       string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Modifier string `protobuf:"bytes,4,opt,name=modifier" json:"modifier,omitempty"`
}

func (m *ReqModifyConfig) Reset()                    { *m = ReqModifyConfig{} }
func (m *ReqModifyConfig) String() string            { return proto.CompactTextString(m) }
func (*ReqModifyConfig) ProtoMessage()               {}
func (*ReqModifyConfig) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{29} }

func (m *ReqModifyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqModifyConfig) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *ReqModifyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReqModifyConfig) GetModifier() string {
	if m != nil {
		return m.Modifier
	}
	return ""
}

func init() {
	proto.RegisterType((*WalletTxDetail)(nil), "types.WalletTxDetail")
	proto.RegisterType((*WalletTxDetails)(nil), "types.WalletTxDetails")
	proto.RegisterType((*WalletAccountStore)(nil), "types.WalletAccountStore")
	proto.RegisterType((*WalletPwHash)(nil), "types.WalletPwHash")
	proto.RegisterType((*WalletStatus)(nil), "types.WalletStatus")
	proto.RegisterType((*WalletAccounts)(nil), "types.WalletAccounts")
	proto.RegisterType((*WalletAccount)(nil), "types.WalletAccount")
	proto.RegisterType((*WalletUnLock)(nil), "types.WalletUnLock")
	proto.RegisterType((*GenSeedLang)(nil), "types.GenSeedLang")
	proto.RegisterType((*GetSeedByPw)(nil), "types.GetSeedByPw")
	proto.RegisterType((*SaveSeedByPw)(nil), "types.SaveSeedByPw")
	proto.RegisterType((*ReplySeed)(nil), "types.ReplySeed")
	proto.RegisterType((*ReqWalletSetPasswd)(nil), "types.ReqWalletSetPasswd")
	proto.RegisterType((*ReqNewAccount)(nil), "types.ReqNewAccount")
	proto.RegisterType((*MinerFlag)(nil), "types.MinerFlag")
	proto.RegisterType((*ReqWalletTransactionList)(nil), "types.ReqWalletTransactionList")
	proto.RegisterType((*ReqWalletImportPrivKey)(nil), "types.ReqWalletImportPrivKey")
	proto.RegisterType((*ReqWalletSendToAddress)(nil), "types.ReqWalletSendToAddress")
	proto.RegisterType((*ReqWalletSetFee)(nil), "types.ReqWalletSetFee")
	proto.RegisterType((*ReqWalletSetLabel)(nil), "types.ReqWalletSetLabel")
	proto.RegisterType((*ReqWalletMergeBalance)(nil), "types.ReqWalletMergeBalance")
	proto.RegisterType((*ReqStr)(nil), "types.ReqStr")
	proto.RegisterType((*ReplyStr)(nil), "types.ReplyStr")
	proto.RegisterType((*ReqTokenPreCreate)(nil), "types.ReqTokenPreCreate")
	proto.RegisterType((*ReqTokenFinishCreate)(nil), "types.ReqTokenFinishCreate")
	proto.RegisterType((*ReqTokenRevokeCreate)(nil), "types.ReqTokenRevokeCreate")
	proto.RegisterType((*ReqSellToken)(nil), "types.ReqSellToken")
	proto.RegisterType((*ReqBuyToken)(nil), "types.ReqBuyToken")
	proto.RegisterType((*ReqRevokeSell)(nil), "types.ReqRevokeSell")
	proto.RegisterType((*ReqModifyConfig)(nil), "types.ReqModifyConfig")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 1164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x25, 0xcb, 0xb6, 0xc6, 0xb2, 0x93, 0x6c, 0x93, 0x40, 0x35, 0xfa, 0xa3, 0x6c, 0x9b,
	0xc4, 0x05, 0x0a, 0x03, 0x4d, 0x4e, 0x2d, 0x50, 0x20, 0x76, 0x02, 0xc7, 0x41, 0xed, 0xd4, 0x58,
	0xb9, 0x68, 0x6f, 0xc5, 0x9a, 0x1c, 0x5b, 0x84, 0x28, 0xae, 0xbc, 0x5c, 0xfd, 0xbd, 0x49, 0xd1,
	0x73, 0x5f, 0xa8, 0x6f, 0x54, 0xcc, 0xec, 0x52, 0x24, 0x53, 0xe7, 0x50, 0xf4, 0x36, 0xdf, 0xec,
	0xec, 0xfc, 0x7c, 0x33, 0x9c, 0x25, 0xf4, 0x16, 0x3a, 0xcb, 0xd0, 0x1d, 0x4e, 0xad, 0x71, 0x46,
	0x74, 0xdc, 0x6a, 0x8a, 0xc5, 0xfe, 0x03, 0x67, 0x75, 0x5e, 0xe8, 0xd8, 0xa5, 0x26, 0xf7, 0x27,
	0xfb, 0xf7, 0xaf, 0x32, 0x13, 0x8f, 0xe3, 0x91, 0x4e, 0x4b, 0xcd, 0xae, 0x8e, 0x63, 0x33, 0xcb,
	0xc3, 0xd5, 0xfd, 0x3d, 0x5c, 0x62, 0x3c, 0x73, 0xc6, 0x7a, 0x2c, 0xff, 0x6c, 0xc1, 0xde, 0xaf,
	0xec, 0xfb, 0x72, 0xf9, 0x06, 0x9d, 0x4e, 0x33, 0x21, 0xa1, 0xe5, 0x96, 0xfd, 0x68, 0x10, 0x1d,
	0xec, 0xbc, 0x10, 0x87, 0x1c, 0xea, 0xf0, 0xb2, 0x8a, 0xa4, 0x5a, 0x6e, 0x29, 0xbe, 0x85, 0x2d,
	0x8b, 0x31, 0xa6, 0x53, 0xd7, 0x6f, 0x35, 0x0c, 0x95, 0xd7, 0xbe, 0xd1, 0x4e, 0xab, 0xd2, 0x44,
	0x3c, 0x86, 0xcd, 0x11, 0xa6, 0x37, 0x23, 0xd7, 0x6f, 0x0f, 0xa2, 0x83, 0xb6, 0x0a, 0x48, 0x3c,
	0x84, 0x4e, 0x9a, 0x27, 0xb8, 0xec, 0x6f, 0xb0, 0xda, 0x03, 0xf1, 0x19, 0x74, 0xb9, 0x0a, 0x97,
	0x4e, 0xb0, 0xdf, 0xe1, 0x93, 0x4a, 0x41, 0xbe, 0xf4, 0x84, 0x0a, 0xea, 0x6f, 0x7a, 0x5f, 0x1e,
	0x89, 0x7d, 0xd8, 0xbe, 0xb6, 0x66, 0xa2, 0x93, 0xc4, 0xf6, 0xb7, 0x06, 0xd1, 0x41, 0x57, 0xad,
	0x31, 0xdd, 0x71, 0xcb, 0x91, 0x2e, 0x46, 0xfd, 0xed, 0x41, 0x74, 0xd0, 0x53, 0x01, 0x89, 0x2f,
	0x00, 0x7c, 0x4d, 0xef, 0xf5, 0x04, 0xfb, 0x5d, 0xbe, 0x55, 0xd3, 0xc8, 0x13, 0xb8, 0xd7, 0xe4,
	0xa6, 0x10, 0x2f, 0xa1, 0xeb, 0x4a, 0xd0, 0x8f, 0x06, 0xed, 0x83, 0x9d, 0x17, 0x8f, 0x42, 0xe9,
	0x4d, 0x53, 0x55, 0xd9, 0xc9, 0x39, 0x08, 0x7f, 0x78, 0xe4, 0x7b, 0x31, 0x74, 0xc6, 0xa2, 0xe8,
	0xc3, 0xd6, 0xd4, 0xa6, 0xf3, 0x31, 0xae, 0x98, 0xec, 0xae, 0x2a, 0x21, 0xf1, 0x92, 0xe9, 0x2b,
	0xcc, 0x98, 0xdb, 0xae, 0xf2, 0x40, 0x08, 0xd8, 0xe0, 0xea, 0xda, 0xac, 0x64, 0x99, 0xb8, 0x22,
	0x56, 0x86, 0x4e, 0x4f, 0xa6, 0xcc, 0x62, 0x57, 0x55, 0x0a, 0xf9, 0x0a, 0x7a, 0x3e, 0xee, 0xc5,
	0xe2, 0x94, 0xea, 0x7d, 0x0c, 0x9b, 0x53, 0x96, 0x38, 0x60, 0x4f, 0x05, 0x44, 0x99, 0x58, 0x9d,
	0x27, 0x85, 0xb3, 0x21, 0x62, 0x09, 0xe5, 0x1f, 0x51, 0xe9, 0x62, 0xe8, 0xb4, 0x9b, 0x15, 0x42,
	0x42, 0x2f, 0x2d, 0xbc, 0xe6, 0xcc, 0xc4, 0x63, 0x76, 0xb4, 0xad, 0x1a, 0x3a, 0x6f, 0x73, 0x34,
	0x73, 0xe6, 0x3c, 0xcd, 0xd3, 0xfc, 0x86, 0x7d, 0xb2, 0x4d, 0xa5, 0xa3, 0xc4, 0xd3, 0xe2, 0x54,
	0x17, 0x43, 0xc4, 0x84, 0x2b, 0xda, 0x56, 0x95, 0xc2, 0x7b, 0xb8, 0x4c, 0xe3, 0x71, 0x88, 0xb2,
	0x51, 0x7a, 0xa8, 0x74, 0xf2, 0x55, 0x39, 0xb8, 0x81, 0xd4, 0x42, 0x1c, 0xc2, 0x96, 0xff, 0x4c,
	0xca, 0xce, 0x3c, 0x6c, 0x74, 0x26, 0xd8, 0xa9, 0xd2, 0x48, 0xbe, 0x85, 0xdd, 0xc6, 0x89, 0x18,
	0x40, 0x5b, 0xc7, 0x71, 0x18, 0xfd, 0xbd, 0x70, 0xb9, 0xbc, 0x46, 0x47, 0x77, 0x77, 0x46, 0x8e,
	0x4a, 0x92, 0x7e, 0xc9, 0x99, 0x00, 0xe2, 0x59, 0x17, 0xc5, 0x22, 0x09, 0x8d, 0x0d, 0x88, 0x78,
	0xa6, 0xe6, 0x98, 0x99, 0xff, 0x6a, 0xda, 0xaa, 0x84, 0xe2, 0x19, 0xec, 0xf9, 0xac, 0x7e, 0xb6,
	0xbe, 0xc4, 0xc0, 0xc9, 0x07, 0x5a, 0xf9, 0x04, 0x76, 0xde, 0x62, 0x4e, 0x1c, 0x9d, 0xe9, 0xfc,
	0x86, 0x46, 0x22, 0xd3, 0xf9, 0x0d, 0x87, 0xe9, 0x28, 0x96, 0xe5, 0x53, 0x32, 0x71, 0x64, 0x72,
	0xbc, 0xba, 0x58, 0x7c, 0x2c, 0x17, 0xf9, 0x03, 0xf4, 0x86, 0x7a, 0x8e, 0x6b, 0x3b, 0x01, 0x1b,
	0x05, 0xf5, 0xc2, 0x5b, 0xb1, 0x5c, 0xbb, 0xdb, 0x6a, 0xdc, 0xfd, 0x12, 0xba, 0x0a, 0xa7, 0xd9,
	0x8a, 0x7b, 0x75, 0xc7, 0x45, 0x79, 0x0a, 0x42, 0xe1, 0x6d, 0x18, 0x1c, 0x74, 0x17, 0xeb, 0xf2,
	0x4d, 0x96, 0x90, 0x8f, 0x72, 0xe0, 0x03, 0xa4, 0x93, 0x1c, 0x17, 0x7c, 0x12, 0x06, 0x30, 0x40,
	0xf9, 0x14, 0x76, 0x15, 0xde, 0xbe, 0xc7, 0x45, 0xd9, 0xa3, 0x75, 0x07, 0xa2, 0x7a, 0x07, 0xbe,
	0x87, 0xee, 0x79, 0x9a, 0xa3, 0x3d, 0xc9, 0x34, 0xb3, 0x72, 0x9d, 0xe9, 0x35, 0x2b, 0x24, 0xf3,
	0x88, 0x63, 0x81, 0x76, 0x8e, 0x25, 0xf5, 0x01, 0xca, 0x6b, 0xe8, 0xaf, 0x73, 0xad, 0xad, 0xb9,
	0xb3, 0xb4, 0xe0, 0xc5, 0x45, 0x4b, 0xe4, 0x72, 0x59, 0x7e, 0x30, 0x1e, 0x51, 0x12, 0x9c, 0x0d,
	0xfb, 0xea, 0x28, 0x0f, 0x68, 0xa6, 0x93, 0xd4, 0x22, 0x5f, 0xe7, 0xfe, 0x75, 0x54, 0xa5, 0x90,
	0xa7, 0xf0, 0x78, 0x1d, 0xe7, 0xdd, 0x64, 0x6a, 0xac, 0xbb, 0xb0, 0xe9, 0xfc, 0x27, 0x5c, 0xfd,
	0xd7, 0x45, 0x20, 0xff, 0x8a, 0x6a, 0xae, 0x86, 0x98, 0x27, 0x97, 0xe6, 0x28, 0x49, 0x2c, 0x16,
	0x05, 0x97, 0x6e, 0xcd, 0xa4, 0x6c, 0x06, 0xc9, 0x62, 0x0f, 0x5a, 0xce, 0x04, 0x0f, 0x2d, 0x67,
	0x6a, 0x1b, 0xb4, 0xdd, 0xd8, 0xa0, 0x02, 0x36, 0x72, 0xe3, 0x30, 0xac, 0x11, 0x96, 0x29, 0xb5,
	0xb4, 0x70, 0x66, 0x8c, 0x39, 0x6f, 0xe2, 0x6d, 0x55, 0x42, 0x31, 0x80, 0x1d, 0x16, 0x86, 0xab,
	0xc9, 0x95, 0xc9, 0x78, 0x19, 0x77, 0x55, 0x5d, 0x25, 0xbf, 0x81, 0x7b, 0xf5, 0x21, 0x38, 0xc1,
	0xfa, 0xf2, 0x8e, 0xea, 0xa1, 0xe5, 0x8f, 0xf0, 0xa0, 0x6e, 0x7a, 0xd6, 0xd8, 0x77, 0x51, 0x6d,
	0xdf, 0xdd, 0x4d, 0xc8, 0x73, 0x78, 0xb4, 0xbe, 0x7e, 0x8e, 0xf6, 0x06, 0x8f, 0x75, 0xa6, 0xf3,
	0x18, 0x43, 0xe9, 0x51, 0x59, 0xba, 0x1c, 0xc0, 0xa6, 0xc2, 0xdb, 0xa1, 0xe3, 0x27, 0xc1, 0xe2,
	0x2d, 0x6d, 0xbc, 0xf0, 0x59, 0x78, 0x24, 0x9f, 0xc1, 0xb6, 0x1f, 0x6d, 0x67, 0xe9, 0x49, 0xb1,
	0x24, 0x57, 0x56, 0x6b, 0x2c, 0xff, 0x8e, 0x38, 0xe5, 0x4b, 0xaa, 0xf7, 0xc2, 0xe2, 0x6b, 0x8b,
	0xda, 0xa1, 0x78, 0x02, 0xbd, 0x98, 0x24, 0x63, 0x7f, 0xaf, 0xa5, 0xbe, 0x13, 0x74, 0xd4, 0x24,
	0x66, 0x99, 0x5e, 0x9b, 0x56, 0x60, 0x59, 0xfb, 0x37, 0xad, 0xf0, 0x34, 0xfa, 0xdd, 0x1e, 0x10,
	0xaf, 0xc1, 0xdc, 0x59, 0x93, 0xcc, 0xfc, 0x4c, 0xf9, 0xce, 0x34, 0x74, 0xe2, 0x73, 0x00, 0xb3,
	0xc8, 0x31, 0x04, 0xec, 0xf8, 0x27, 0x80, 0x35, 0x47, 0x81, 0x30, 0x67, 0x9c, 0xce, 0xc2, 0x6b,
	0xe9, 0x01, 0x69, 0xa7, 0x36, 0x8d, 0x91, 0x5f, 0xca, 0xb6, 0xf2, 0x40, 0x5a, 0x78, 0x58, 0x96,
	0x74, 0x92, 0xe6, 0x69, 0x31, 0x0a, 0x55, 0x7d, 0x05, 0xbb, 0xd7, 0x8c, 0xb1, 0x51, 0x56, 0xaf,
	0x54, 0x1e, 0x85, 0x37, 0x36, 0xd4, 0xd0, 0x6a, 0xd4, 0xd0, 0xcc, 0xaf, 0xfd, 0x41, 0x7e, 0x72,
	0x5a, 0xc5, 0x54, 0x38, 0x37, 0xe3, 0x1a, 0x93, 0x96, 0x71, 0x93, 0xc9, 0xa0, 0xfb, 0x3f, 0x11,
	0xcf, 0xa1, 0x47, 0x33, 0x80, 0x59, 0xc6, 0x51, 0xc5, 0x73, 0xda, 0x5f, 0x59, 0x16, 0xb6, 0xfe,
	0x27, 0xd5, 0x0f, 0x4f, 0x82, 0x27, 0xc6, 0x92, 0x9d, 0x62, 0x03, 0x22, 0x8d, 0xbd, 0x94, 0xb3,
	0xc7, 0x40, 0xbe, 0x83, 0x1d, 0x85, 0xb7, 0xc7, 0xb3, 0x95, 0xf7, 0xf6, 0x35, 0xb4, 0xaf, 0x66,
	0xab, 0x7f, 0xff, 0x3d, 0xb1, 0xb3, 0xe3, 0xd9, 0x4a, 0xd1, 0x31, 0xb9, 0xba, 0x9a, 0xad, 0x2a,
	0x57, 0x0c, 0xe4, 0x6f, 0xbc, 0xeb, 0x3c, 0x0d, 0x14, 0x57, 0x7c, 0x47, 0x43, 0x4a, 0x28, 0xf8,
	0xfb, 0xf4, 0x03, 0x7f, 0x95, 0xa9, 0x0a, 0x86, 0x1f, 0x49, 0x12, 0xf9, 0x53, 0x3c, 0x37, 0x49,
	0x7a, 0xbd, 0x7a, 0x6d, 0xf2, 0xeb, 0xf4, 0x46, 0xdc, 0x87, 0x76, 0xb5, 0x70, 0x48, 0xa4, 0x8f,
	0xc5, 0x4c, 0xcb, 0x3d, 0x61, 0xa6, 0xe4, 0x6a, 0xae, 0xb3, 0x19, 0x06, 0x0a, 0x3d, 0xa0, 0x8f,
	0x62, 0x42, 0x7e, 0x52, 0xb4, 0x61, 0x1e, 0xd7, 0xf8, 0x6a, 0x93, 0xff, 0x29, 0x5f, 0xfe, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x0a, 0x15, 0xc7, 0x70, 0xae, 0x0a, 0x00, 0x00,
}
