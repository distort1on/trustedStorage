package test_pr

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"trustedStorage/blockchain"
	"trustedStorage/database"
	"trustedStorage/grpsServer"
	"trustedStorage/serialization"
	"trustedStorage/transaction"
)

type myInvoicerServer struct {
	grpsServer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) AddTxToBlockchain(ctx context.Context, req *grpsServer.CreateTx) (*grpsServer.CreateResponse, error) {

	tx := transaction.Transaction{
		SenderAddress: req.SenderAddress,
		Data:          req.Data,
		PubKey:        req.PubKey,
		Signature:     req.Signature,
	}
	documentBytes := req.DocumentBytes

	err := memPool.AddTxToMempool(tx, documentBytes, sh)
	if err != nil {
		log.Println(err)
		return &grpsServer.CreateResponse{
			Response: "tx invalid",
		}, nil
	}

	//todo add goroutine
	if len(memPool) > numOfTransactionsInBlock {
		block := blockchain.CreateBlock(1, (*blockChain)[len(*blockChain)-1].GetBlockHash(), memPool.FormTransactionsList(numOfTransactionsInBlock))
		err := blockChain.AcceptingBlock(&block)
		if err != nil {
			fmt.Println(err)
		}
	}

	return &grpsServer.CreateResponse{
		Response: "tx added",
	}, nil
}

func (s *myInvoicerServer) GetLastBlock(ctx context.Context, emp *emptypb.Empty) (*grpsServer.CreateResponse, error) {
	return &grpsServer.CreateResponse{
		Response: (*blockChain)[len(*blockChain)-1].ToString(),
	}, nil
}

func (s *myInvoicerServer) GetBlockchain(ctx context.Context, emp *emptypb.Empty) (*grpsServer.CreateResponse, error) {
	return &grpsServer.CreateResponse{
		Response: blockChain.ToString(),
	}, nil
}

func (s *myInvoicerServer) GetUserTxHistory(ctx context.Context, req *grpsServer.User) (*grpsServer.CreateResponse, error) {
	return &grpsServer.CreateResponse{
		Response: blockchain.GetUserTxHistory(blockChain, req.SenderAddress),
	}, nil
}

func StartGrpc() {
	bcBytes := database.GetFromDB("blockchain")
	err := serialization.DeSerialize(&blockChain, bcBytes)
	if err != nil {
		panic(err)
	}

	mpBytes := database.GetFromDB("mempool")
	err = serialization.DeSerialize(&memPool, mpBytes)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener : %s", err)

	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	grpsServer.RegisterInvoicerServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}

	/*for {
		fmt.Println("a + a")
		time.Sleep(1)
	}*/

}
