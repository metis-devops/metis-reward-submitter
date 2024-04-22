package main

import (
	"context"
	"flag"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metisprotocol/metis-reward-submitter/internal/contracts/sequencerset"
	"github.com/metisprotocol/metis-reward-submitter/internal/themis"
)

func main() {
	var (
		RestHost string
		L2geth   string
		SeqSet   string
	)

	flag.StringVar(&RestHost, "themis", "", "the themis rest url")
	flag.StringVar(&L2geth, "l2geth", "", "the l2geth jsonrpc url")
	flag.StringVar(&SeqSet, "seqset", "", "the seqset contract")
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
			slog.Info("Signer not match", "epoch", i, "l2geth", epoch.Signer, "themis", themisSigner)
		}

		if epoch.EndBlock.Uint64() != restEpoch.EndBlock {
			slog.Info("EndBlock not match", "epoch", i, "l2geth", epoch.EndBlock, "themis", restEpoch.EndBlock)
		}

		if epoch.StartBlock.Uint64() != restEpoch.StartBlock {
			slog.Info("StartBlock not match", "epoch", i, "l2geth", epoch.StartBlock, "themis", restEpoch.StartBlock)
		}
	}
}
