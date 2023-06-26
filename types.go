package flashbotsrpc

import (
	"bytes"
	"encoding/json"
	"math/big"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

// ErrRelayErrorResponse means it's a standard Flashbots relay error response - probably a user error rather than JSON or network error
var ErrRelayErrorResponse = errors.New("relay error response")

// Syncing - object with syncing data info
type Syncing struct {
	IsSyncing     bool
	StartingBlock int
	CurrentBlock  int
	HighestBlock  int
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Syncing) UnmarshalJSON(data []byte) error {
	proxy := new(proxySyncing)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	proxy.IsSyncing = true
	*s = *(*Syncing)(unsafe.Pointer(proxy))

	return nil
}

// T - input transaction object
type T struct {
	From     string
	To       string
	Gas      int
	GasPrice *big.Int
	Value    *big.Int
	Data     string
	Nonce    int
}

// MarshalJSON implements the json.Unmarshaler interface.
func (t T) MarshalJSON() ([]byte, error) {
	params := map[string]interface{}{
		"from": t.From,
	}
	if t.To != "" {
		params["to"] = t.To
	}
	if t.Gas > 0 {
		params["gas"] = IntToHex(t.Gas)
	}
	if t.GasPrice != nil {
		params["gasPrice"] = BigToHex(*t.GasPrice)
	}
	if t.Value != nil {
		params["value"] = BigToHex(*t.Value)
	}
	if t.Data != "" {
		params["data"] = t.Data
	}
	if t.Nonce > 0 {
		params["nonce"] = IntToHex(t.Nonce)
	}

	return json.Marshal(params)
}

// Transaction - transaction object
type Transaction struct {
	Hash             string
	Nonce            int
	BlockHash        string
	BlockNumber      *int
	TransactionIndex *int
	From             string
	To               string
	Value            big.Int
	Gas              int
	GasPrice         big.Int
	Input            string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransaction)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*Transaction)(unsafe.Pointer(proxy))

	return nil
}

// Log - log object
type Log struct {
	Removed          bool
	LogIndex         int
	TransactionIndex int
	TransactionHash  string
	BlockNumber      int
	BlockHash        string
	Address          string
	Data             string
	Topics           []string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (log *Log) UnmarshalJSON(data []byte) error {
	proxy := new(proxyLog)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*log = *(*Log)(unsafe.Pointer(proxy))

	return nil
}

// FilterParams - Filter parameters object
type FilterParams struct {
	FromBlock string     `json:"fromBlock,omitempty"`
	ToBlock   string     `json:"toBlock,omitempty"`
	Address   []string   `json:"address,omitempty"`
	Topics    [][]string `json:"topics,omitempty"`
}

// TransactionReceipt - transaction receipt object
type TransactionReceipt struct {
	TransactionHash   string
	TransactionIndex  int
	BlockHash         string
	BlockNumber       int
	CumulativeGasUsed int
	GasUsed           int
	ContractAddress   string
	Logs              []Log
	LogsBloom         string
	Root              string
	Status            string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TransactionReceipt) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransactionReceipt)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*TransactionReceipt)(unsafe.Pointer(proxy))

	return nil
}

// Block - block object
type Block struct {
	Number           int
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	Miner            string
	Difficulty       big.Int
	TotalDifficulty  big.Int
	ExtraData        string
	Size             int
	GasLimit         int
	GasUsed          int
	Timestamp        int
	Uncles           []string
	Transactions     []Transaction
}

type proxySyncing struct {
	IsSyncing     bool   `json:"-"`
	StartingBlock hexInt `json:"startingBlock"`
	CurrentBlock  hexInt `json:"currentBlock"`
	HighestBlock  hexInt `json:"highestBlock"`
}

type proxyTransaction struct {
	Hash             string  `json:"hash"`
	Nonce            hexInt  `json:"nonce"`
	BlockHash        string  `json:"blockHash"`
	BlockNumber      *hexInt `json:"blockNumber"`
	TransactionIndex *hexInt `json:"transactionIndex"`
	From             string  `json:"from"`
	To               string  `json:"to"`
	Value            hexBig  `json:"value"`
	Gas              hexInt  `json:"gas"`
	GasPrice         hexBig  `json:"gasPrice"`
	Input            string  `json:"input"`
}

type proxyLog struct {
	Removed          bool     `json:"removed"`
	LogIndex         hexInt   `json:"logIndex"`
	TransactionIndex hexInt   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockNumber      hexInt   `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

type proxyTransactionReceipt struct {
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  hexInt `json:"transactionIndex"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       hexInt `json:"blockNumber"`
	CumulativeGasUsed hexInt `json:"cumulativeGasUsed"`
	GasUsed           hexInt `json:"gasUsed"`
	ContractAddress   string `json:"contractAddress,omitempty"`
	Logs              []Log  `json:"logs"`
	LogsBloom         string `json:"logsBloom"`
	Root              string `json:"root"`
	Status            string `json:"status,omitempty"`
}

type hexInt int

func (i *hexInt) UnmarshalJSON(data []byte) error {
	result, err := ParseInt(string(bytes.Trim(data, `"`)))
	*i = hexInt(result)

	return err
}

type hexBig big.Int

func (i *hexBig) UnmarshalJSON(data []byte) error {
	result, err := ParseBigInt(string(bytes.Trim(data, `"`)))
	*i = hexBig(result)

	return err
}

type proxyBlock interface {
	toBlock() Block
}

type proxyBlockWithTransactions struct {
	Number           hexInt             `json:"number"`
	Hash             string             `json:"hash"`
	ParentHash       string             `json:"parentHash"`
	Nonce            string             `json:"nonce"`
	Sha3Uncles       string             `json:"sha3Uncles"`
	LogsBloom        string             `json:"logsBloom"`
	TransactionsRoot string             `json:"transactionsRoot"`
	StateRoot        string             `json:"stateRoot"`
	Miner            string             `json:"miner"`
	Difficulty       hexBig             `json:"difficulty"`
	TotalDifficulty  hexBig             `json:"totalDifficulty"`
	ExtraData        string             `json:"extraData"`
	Size             hexInt             `json:"size"`
	GasLimit         hexInt             `json:"gasLimit"`
	GasUsed          hexInt             `json:"gasUsed"`
	Timestamp        hexInt             `json:"timestamp"`
	Uncles           []string           `json:"uncles"`
	Transactions     []proxyTransaction `json:"transactions"`
}

func (proxy *proxyBlockWithTransactions) toBlock() Block {
	return *(*Block)(unsafe.Pointer(proxy))
}

type proxyBlockWithoutTransactions struct {
	Number           hexInt   `json:"number"`
	Hash             string   `json:"hash"`
	ParentHash       string   `json:"parentHash"`
	Nonce            string   `json:"nonce"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	LogsBloom        string   `json:"logsBloom"`
	TransactionsRoot string   `json:"transactionsRoot"`
	StateRoot        string   `json:"stateRoot"`
	Miner            string   `json:"miner"`
	Difficulty       hexBig   `json:"difficulty"`
	TotalDifficulty  hexBig   `json:"totalDifficulty"`
	ExtraData        string   `json:"extraData"`
	Size             hexInt   `json:"size"`
	GasLimit         hexInt   `json:"gasLimit"`
	GasUsed          hexInt   `json:"gasUsed"`
	Timestamp        hexInt   `json:"timestamp"`
	Uncles           []string `json:"uncles"`
	Transactions     []string `json:"transactions"`
}

func (proxy *proxyBlockWithoutTransactions) toBlock() Block {
	block := Block{
		Number:           int(proxy.Number),
		Hash:             proxy.Hash,
		ParentHash:       proxy.ParentHash,
		Nonce:            proxy.Nonce,
		Sha3Uncles:       proxy.Sha3Uncles,
		LogsBloom:        proxy.LogsBloom,
		TransactionsRoot: proxy.TransactionsRoot,
		StateRoot:        proxy.StateRoot,
		Miner:            proxy.Miner,
		Difficulty:       big.Int(proxy.Difficulty),
		TotalDifficulty:  big.Int(proxy.TotalDifficulty),
		ExtraData:        proxy.ExtraData,
		Size:             int(proxy.Size),
		GasLimit:         int(proxy.GasLimit),
		GasUsed:          int(proxy.GasUsed),
		Timestamp:        int(proxy.Timestamp),
		Uncles:           proxy.Uncles,
	}

	block.Transactions = make([]Transaction, len(proxy.Transactions))
	for i := range proxy.Transactions {
		block.Transactions[i] = Transaction{
			Hash: proxy.Transactions[i],
		}
	}

	return block
}

type RelayErrorResponse struct {
	Error string `json:"error"`
}

type FlashbotsUserStats struct {
	IsHighPriority       bool   `json:"is_high_priority"`
	AllTimeMinerPayments string `json:"all_time_miner_payments"`
	AllTimeGasSimulated  string `json:"all_time_gas_simulated"`
	Last7dMinerPayments  string `json:"last_7d_miner_payments"`
	Last7dGasSimulated   string `json:"last_7d_gas_simulated"`
	Last1dMinerPayments  string `json:"last_1d_miner_payments"`
	Last1dGasSimulated   string `json:"last_1d_gas_simulated"`
}

type FlashbotsCallBundleParam struct {
	Txs              []string `json:"txs"`                 // Array[String], A list of signed transactions to execute in an atomic bundle
	BlockNumber      string   `json:"blockNumber"`         // String, a hex encoded block number for which this bundle is valid on
	StateBlockNumber string   `json:"stateBlockNumber"`    // String, either a hex encoded number or a block tag for which state to base this simulation on. Can use "latest"
	Timestamp        int64    `json:"timestamp,omitempty"` // Number, the timestamp to use for this bundle simulation, in seconds since the unix epoch
	Timeout          int64    `json:"timeout,omitempty"`
	GasLimit         uint64   `json:"gasLimit,omitempty"`
	Difficulty       uint64   `json:"difficulty,omitempty"`
	BaseFee          uint64   `json:"baseFee,omitempty"`
}

type FlashbotsCallBundleResult struct {
	CoinbaseDiff      string `json:"coinbaseDiff"`      // "2717471092204423",
	EthSentToCoinbase string `json:"ethSentToCoinbase"` // "0",
	FromAddress       string `json:"fromAddress"`       // "0x37ff310ab11d1928BB70F37bC5E3cf62Df09a01c",
	GasFees           string `json:"gasFees"`           // "2717471092204423",
	GasPrice          string `json:"gasPrice"`          // "43000001459",
	GasUsed           int64  `json:"gasUsed"`           // 63197,
	ToAddress         string `json:"toAddress"`         // "0xdAC17F958D2ee523a2206206994597C13D831ec7",
	TxHash            string `json:"txHash"`            // "0xe2df005210bdc204a34ff03211606e5d8036740c686e9fe4e266ae91cf4d12df",
	Value             string `json:"value"`             // "0x"
	Error             string `json:"error"`
	Revert            string `json:"revert"`
}

type FlashbotsCallBundleResponse struct {
	BundleGasPrice    string                      `json:"bundleGasPrice"`    // "43000001459",
	BundleHash        string                      `json:"bundleHash"`        // "0x2ca9c4d2ba00d8144d8e396a4989374443cb20fb490d800f4f883ad4e1b32158",
	CoinbaseDiff      string                      `json:"coinbaseDiff"`      // "2717471092204423",
	EthSentToCoinbase string                      `json:"ethSentToCoinbase"` // "0",
	GasFees           string                      `json:"gasFees"`           // "2717471092204423",
	Results           []FlashbotsCallBundleResult `json:"results"`           // [],
	StateBlockNumber  int64                       `json:"stateBlockNumber"`  // 12960319,
	TotalGasUsed      int64                       `json:"totalGasUsed"`      // 63197
}

// sendBundle
type FlashbotsSendBundleRequest struct {
	Txs          []string  `json:"txs"`                         // Array[String], A list of signed transactions to execute in an atomic bundle
	BlockNumber  string    `json:"blockNumber"`                 // String, a hex encoded block number for which this bundle is valid on
	MinTimestamp *uint64   `json:"minTimestamp,omitempty"`      // (Optional) Number, the minimum timestamp for which this bundle is valid, in seconds since the unix epoch
	MaxTimestamp *uint64   `json:"maxTimestamp,omitempty"`      // (Optional) Number, the maximum timestamp for which this bundle is valid, in seconds since the unix epoch
	RevertingTxs *[]string `json:"revertingTxHashes,omitempty"` // (Optional) Array[String], A list of tx hashes that are allowed to revert
}

type FlashbotsGetBundleStatsParam struct {
	BlockNumber string `json:"blockNumber"` // String, a hex encoded block number for which this bundle is valid on
	BundleHash  string `json:"bundleHash"`  // String, returned by the flashbots api when calling eth_sendBundle
}

type FlashbotsGetBundleStatsResponse struct {
	IsSimulated            bool                          `json:"isSimulated"`
	IsSentToMiners         bool                          `json:"isSentToMiners"`
	IsHighPriority         bool                          `json:"isHighPriority"`
	SimulatedAt            time.Time                     `json:"simulatedAt"`
	SubmittedAt            time.Time                     `json:"submittedAt"`
	SentToMinersAt         time.Time                     `json:"sentToMinersAt"`
	ConsideredByBuildersAt []*BuilderPubkeyWithTimestamp `json:"consideredByBuildersAt"`
	SealedByBuildersAt     []*BuilderPubkeyWithTimestamp `json:"sealedByBuildersAt"`
}

type FlashbotsGetBundleStatsResponseV2 struct {
	IsSimulated            bool                          `json:"isSimulated"`
	IsHighPriority         bool                          `json:"isHighPriority"`
	SimulatedAt            time.Time                     `json:"simulatedAt"`
	ReceivedAt             time.Time                     `json:"receivedAt"`
	ConsideredByBuildersAt []*BuilderPubkeyWithTimestamp `json:"consideredByBuildersAt"`
	SealedByBuildersAt     []*BuilderPubkeyWithTimestamp `json:"sealedByBuildersAt"`
}

type BuilderPubkeyWithTimestamp struct {
	Pubkey    string    `json:"pubkey"`
	Timestamp time.Time `json:"timestamp"`
}

type FlashbotsSendBundleResponse struct {
	BundleHash string `json:"bundleHash"`
}

type BuilderBroadcastResponse struct {
	BundleResponse FlashbotsSendBundleResponse `json:"bundleResponse"`
	Err            error                       `json:"err"`
}

// sendPrivateTransaction
type FlashbotsSendPrivateTransactionRequest struct {
	Tx          string                         `json:"tx"`
	Preferences *FlashbotsPrivateTxPreferences `json:"preferences,omitempty"`
}

type FlashbotsPrivateTxPreferences struct {
	Fast bool `json:"fast"`
}

// cancelPrivateTransaction
type FlashbotsCancelPrivateTransactionRequest struct {
	TxHash string `json:"txHash"`
}
