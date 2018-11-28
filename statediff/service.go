package statediff

import (
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethdb"
	"log"
	"fmt"
)

type StateDiffService struct {
	builder *StateDiffBuilder
	blockchain *core.BlockChain
}

func NewStateDiffService(db ethdb.Database, blockChain *core.BlockChain) (*StateDiffService, error) {
	builder, err := NewStateDiffBuilder(db)
	if err != nil {
		log.Fatal("Error creating state diff builder")
	}

	return &StateDiffService{
		builder:    builder,
		blockchain: blockChain,
	}, nil
}

func (StateDiffService) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (StateDiffService) APIs() []rpc.API {
	return []rpc.API{}
}

func (sds *StateDiffService) Start(server *p2p.Server) error {
	fmt.Println("starting the state diff service")
	blockChannel := make(chan core.ChainHeadEvent)
	sds.blockchain.SubscribeChainHeadEvent(blockChannel)
	fmt.Println("block channel", blockChannel)
	//for {
	//	select {
	//	case <-blockChannel:
	//		headOfChainEvent := <-blockChannel
	//		currentBlock := headOfChainEvent.Block
	//		blockNumber := currentBlock.Number()
	//		stateDiff, _ := sds.StateDiffBuilder.CreateStateDiff(currentBlock.Root(), currentBlock.Root(), *blockNumber, currentBlock.Hash())
	//		fmt.Println(stateDiff)
	//	}
	//}
	return nil
}

func (StateDiffService) Stop() error {
	fmt.Println("stopping the state diff service")
	return nil
}


