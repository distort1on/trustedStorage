package grpsServer

import (
	"context"
	"encoding/hex"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"strconv"
	"trustedStorage/blockchain"
	"trustedStorage/mempool"
	"trustedStorage/p2pCommunication"
	"trustedStorage/settings"
	"trustedStorage/transaction"
)

//var Sh = shell.NewShell("localhost:5001")

type myInvoicerServer struct {
	UnimplementedInvoicerServer
}

func (s *myInvoicerServer) AddTxToBlockchain(ctx context.Context, req *CreateTx) (*CreateResponse, error) {
	var err error
	tx := transaction.Transaction{
		SenderAddress: req.SenderAddress,
		Data:          req.Data,
		PubKey:        req.PubKey,
		Signature:     req.Signature,
		Nonce:         req.Nonce,
		Cid:           req.Cid,
	}
	log.Printf("Transaction from %x received", req.SenderAddress)
	err = mempool.MemPoolIns.AddTxToMempool(tx)
	if err != nil {
		log.Println(err)
		return &CreateResponse{
			Response: err.Error(),
		}, nil
	}
	p2pCommunication.ProposeTransaction(tx)

	return &CreateResponse{
		Response: "Tx added to mempool\nTx hash: " + hex.EncodeToString(tx.GetTxHash()),
	}, nil
}

func (s *myInvoicerServer) GetLastBlock(ctx context.Context, emp *emptypb.Empty) (*CreateResponse, error) {
	log.Println("Sending last block")
	return &CreateResponse{
		//Response: (*BlockChainIns)[len(*BlockChainIns)-1].ToString(),
		Response: blockchain.BlockChainIns.GetLastBlock().ToString(),
	}, nil
}

func (s *myInvoicerServer) GetBlockchain(ctx context.Context, emp *emptypb.Empty) (*CreateResponse, error) {
	log.Println("Sending blockchain history")
	return &CreateResponse{

		Response: blockchain.BlockChainIns.ToString(),
	}, nil
}

func (s *myInvoicerServer) GetUserTxHistory(ctx context.Context, req *User) (*CreateResponse, error) {
	log.Println("Sending user tx history")
	return &CreateResponse{
		Response: blockchain.GetUserTxHistory(blockchain.BlockChainIns, req.SenderAddress),
	}, nil
}

func (s *myInvoicerServer) GetTxByHash(ctx context.Context, req *Transaction) (*CreateResponse, error) {
	log.Printf("Finding transaction %x by hash\n", req.TxHash)

	var res string

	IsFind, blockNum, tx := blockchain.FindTransactionByHash(req.TxHash)
	if IsFind {
		res = fmt.Sprintf("Tx has been found in block on height %v\nTx Info:\n%v", strconv.Itoa(blockNum), tx.ToString())
	} else {
		res = "Can't find tx. It does not exist or has not been confirmed yet"
	}
	return &CreateResponse{
		Response: res,
	}, nil
}

func (s *myInvoicerServer) FindDocumentByHash(ctx context.Context, req *Document) (*CreateResponse, error) {
	log.Printf("Finding document %x by hash\n", req.DocumentHash)

	var res string

	IsFind, blockNum, tx := blockchain.FindDocumentByHash(req.DocumentHash)
	if IsFind {
		res = fmt.Sprintf("Document has been found in block on height %v\nTx Info:\n%v", strconv.Itoa(blockNum), tx.ToString())
	} else {
		res = "Can't find document. It does not exist or has not been confirmed yet"
	}
	return &CreateResponse{
		Response: res,
	}, nil
}

func StartGrpc() {

	lis, err := net.Listen("tcp", settings.GetRpcNodeAddress())
	if err != nil {
		log.Fatalf("Cannot create listener : %s", err)
	}

	ServerRegistrar := grpc.NewServer()

	service := &myInvoicerServer{}
	RegisterInvoicerServer(ServerRegistrar, service)

	log.Println("RPC listening on ", lis.Addr())

	err = ServerRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}

}
