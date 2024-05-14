package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metisprotocol/metis-reward-submitter/internal/contracts/sequencerset"
)

func main() {
	var (
		L2geth string
		SeqSet string
		Start  int64
		End    int64
	)

	flag.StringVar(&L2geth, "l2geth", "https://andromeda.metis.io", "the l2geth jsonrpc url")
	flag.StringVar(&SeqSet, "seqset", "0x0fe382b74C3894B65c10E5C12ae60Bbd8FAf5b48", "the seqset contract")
	flag.Int64Var(&Start, "start", 0, "start")
	flag.Int64Var(&End, "end", 0, "end")
	flag.Parse()

	metisClient, err := ethclient.Dial(L2geth)
	if err != nil {
		panic(err)
	}

	seqset, err := sequencerset.NewSequencerset(common.HexToAddress(SeqSet), metisClient)
	if err != nil {
		panic(err)
	}

	var blocks = make(map[common.Address]uint64)

	var blockCount uint64 = 0

	for i := Start; i <= End; i++ {
		epoch, err := seqset.Epochs(nil, big.NewInt(i))
		if err != nil {
			panic(err)
		}

		number := epoch.EndBlock.Uint64() - epoch.StartBlock.Uint64() + 1
		blocks[epoch.Signer] += number
		blockCount += number
		fmt.Println("epoch", i, "signer", epoch.Signer, "start", epoch.StartBlock, "end", epoch.EndBlock)
	}

	fmt.Println()
	fmt.Println("===================")

	for addr, count := range blocks {
		fmt.Println(addr, count)
	}
}
