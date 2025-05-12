package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Tx struct {
	Txid common.Hash  `json:"hash"`
	SeqR *hexutil.Big `json:"seqR"`
	SeqS *hexutil.Big `json:"seqS"`
	SeqV *hexutil.Big `json:"seqV"`
}

type NameService struct {
	Name    string         `json:"name"`
	Address common.Address `json:"seq_addr"`
}

func GetNames(chainId uint64) map[common.Address]string {
	target := fmt.Sprintf("https://metisprotocol.github.io/metis-sequencer-resources/%d/all.json", chainId)
	file, err := http.Get(target)
	if err != nil {
		panic(err)
	}
	defer file.Body.Close()

	var ns []NameService
	if err := json.NewDecoder(file.Body).Decode(&ns); err != nil {
		panic(err)
	}

	mapping := make(map[common.Address]string)
	for _, v := range ns {
		mapping[v.Address] = v.Name
	}
	return mapping
}

func main() {
	var (
		TxHash string
		Height uint64
		Rpc    string
	)

	flag.StringVar(&TxHash, "tx", "", "tx hash")
	flag.Uint64Var(&Height, "height", 0, "tx hash")
	flag.StringVar(&Rpc, "rpc", "https://andromeda.metis.io", "rpc")
	flag.Parse()

	client, err := rpc.Dial(Rpc)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	chainId, err := ethclient.NewClient(client).ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	if id := chainId.Uint64(); id != 1088 && id != 59902 {
		fmt.Println("Not supported chain", id)
		return
	}

	ns := GetNames(chainId.Uint64())

	var getTxSigner = func(tx *Tx) {
		if tx.SeqR == nil || tx.SeqS == nil || tx.SeqV == nil {
			fmt.Println("Unknown")
			return
		}

		sig := slices.Concat(
			tx.SeqR.ToInt().FillBytes(make([]byte, 32)),
			tx.SeqS.ToInt().FillBytes(make([]byte, 32)),
			tx.SeqV.ToInt().FillBytes(make([]byte, 1)),
		)

		seq, err := crypto.SigToPub(tx.Txid[:], sig)
		if err != nil {
			panic(err)
		}

		addr := crypto.PubkeyToAddress(*seq)
		fmt.Println(ns[addr], addr)
	}

	if TxHash != "" {
		var res Tx
		if err := client.Call(&res, "eth_getTransactionByHash", TxHash); err != nil {
			panic(err)
		}

		getTxSigner(&res)
		return
	}

	if Height != 0 {
		type Block struct {
			Tx []*Tx `json:"transactions"`
		}

		var res Block
		if err := client.Call(&res, "eth_getBlockByNumber", hexutil.Uint64(Height), true); err != nil {
			panic(err)
		}
		for _, tx := range res.Tx {
			getTxSigner(tx)
			return
		}
	}
}
