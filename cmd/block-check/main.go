package main

import (
	"context"
	"flag"
	"fmt"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

type Tx struct {
	Hash hexutil.Bytes  `json:"hash"`
	From common.Address `json:"from"`
	SeqR hexutil.Big    `json:"seqR"`
	SeqS hexutil.Big    `json:"seqS"`
	SeqV hexutil.Big    `json:"seqV"`
}

type Block struct {
	Txs  []*Tx       `json:"transactions"`
	Time hexutil.Big `json:"timestamp"`
}

func main() {
	height := flag.Uint64("block", 0, "height")
	remote := flag.String("rpc", "https://andromeda.metis.io", "rpc")
	flag.Parse()

	addrs := map[common.Address]string{
		common.HexToAddress("0xEcA7Ae7dE0d1978DF299a547Ee66c4503fBa474D"): "genesis1",
		common.HexToAddress("0xa233Cc81fC6C12e3318eA71EC5D7bBA78C706b04"): "genesis2",
		common.HexToAddress("0xAfF606251d8540f97Ca2Db12774C0147A170aB9e"): "genesis3",
		common.HexToAddress("0x31e623DCb8B43aD4d05aAA6209574cf336980590"): "artemis",
		common.HexToAddress("0x36f10B20781bb1E78278e84f3e1E97Acf92FA302"): "enki",
	}

	if *height == 0 {
		fmt.Println("no height")
		return
	}

	ctx := context.Background()
	jsonrpc, err := rpc.Dial(*remote)
	if err != nil {
		panic(err)
	}

	var block Block
	err = jsonrpc.CallContext(ctx, &block,
		"eth_getBlockByNumber", hexutil.Uint64(*height), true)
	if err != nil {
		panic(err)
	}

	fmt.Println("timestamp", time.Unix(block.Time.ToInt().Int64(), 0))

	for i, tx := range block.Txs {
		if i != 0 {
			return
		}

		pubkey, err := crypto.SigToPub(
			tx.Hash,
			slices.Concat(
				tx.SeqR.ToInt().FillBytes(make([]byte, 32)),
				tx.SeqS.ToInt().FillBytes(make([]byte, 32)),
				[]byte{byte(tx.SeqV.ToInt().Uint64())},
			))
		if err != nil {
			panic(err)
		}

		addr := crypto.PubkeyToAddress(*pubkey)
		fmt.Println("signer", addrs[addr], addr)
	}
}
