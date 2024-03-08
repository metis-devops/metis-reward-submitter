package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metisprotocol/metis-seq-reward-submitter/internal/contracts/sequencerset"
	"github.com/metisprotocol/metis-seq-reward-submitter/internal/themis"
)

func main() {
	var (
		RestHost string
		L2geth   string
		SeqSet   string
	)

	flag.StringVar(&RestHost, "rest", "http://10.128.5.71:1317", "the themis rest url")
	flag.StringVar(&L2geth, "l2geth", "http://10.128.5.71:8545", "the l2geth jsonrpc url")
	flag.StringVar(&SeqSet, "seqset", "0xdE8d56212118906a0CeCD331e842429714b4c47B", "the seqset contract")
	flag.Parse()

	rest, err := themis.NewClient(RestHost)
	if err != nil {
		panic(err)
	}

	metisClient, err := ethclient.Dial(L2geth)
	if err != nil {
		panic(err)
	}

	seqset, err := sequencerset.NewSequencerset(common.HexToAddress(SeqSet), metisClient)
	if err != nil {
		panic(err)
	}

	number, err := seqset.CurrentEpochNumber(nil)
	if err != nil {
		panic(err)
	}

	slog.Info("seqset", "epoch", number)

	for i := int64(0); i <= number.Int64(); i++ {
		restEpoch, err := rest.GetEpochByID(context.Background(), i)
		if err != nil {
			panic(err)
		}

		epoch, err := seqset.Epochs(nil, big.NewInt(i))
		if err != nil {
			panic(err)
		}

		if len(restEpoch.Producers) != 1 {
			panic("rest signer length != 1")
		}

		if themisSigner := restEpoch.Producers[0].Signer; themisSigner.Cmp(epoch.Signer) != 0 {
			fmt.Println("epoch", i, "signer not match", "l2", epoch.Signer, "themis", themisSigner)
		}

		if epoch.EndBlock.Uint64() != restEpoch.EndBlock {
			fmt.Println("epoch", i, "EndBlock not match")
		}

		if epoch.StartBlock.Uint64() != restEpoch.StartBlock {
			fmt.Println("epoch", i, "StartBlock not match")
		}
	}
}
