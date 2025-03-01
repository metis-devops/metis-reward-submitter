package main

import (
	"flag"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metis-devops/metis-reward-submitter/internal/contracts/sequencerset"
)

type Epoch struct {
	Number uint64
	Start  uint64
	End    uint64
	Signer common.Address
}

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

	remote := strings.Split(L2geth, ",")

	if len(remote) == 0 {
		fmt.Println("No node provided")
		return
	}

	var blockCount uint64 = 0
	var blocks = make(map[common.Address]uint64)

	for i := Start; i <= End; i++ {
		epochs := make([]Epoch, len(remote))

		for idx, node := range remote {
			metisClient, err := ethclient.Dial(node)
			if err != nil {
				panic(err)
			}

			seqset, err := sequencerset.NewSequencerset(common.HexToAddress(SeqSet), metisClient)
			if err != nil {
				panic(err)
			}

			epoch, err := seqset.Epochs(nil, big.NewInt(i))
			if err != nil {
				panic(err)
			}

			epochs[idx] = Epoch{Number: epoch.Number.Uint64(), Start: epoch.StartBlock.Uint64(), End: epoch.EndBlock.Uint64(), Signer: epoch.Signer}
		}

		for idx, cur := range epochs {
			if idx == 0 {
				number := cur.End - cur.Start + 1
				blocks[cur.Signer] += number
				blockCount += number
				continue
			}
			if pre := epochs[idx-1]; cur != pre {
				fmt.Println("diff", "idx", idx, "cur", cur, "pre", pre)
			}
		}
	}

	fmt.Println()
	fmt.Println("===================")

	for addr, count := range blocks {
		fmt.Println(addr, count)
	}
}
