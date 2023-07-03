package test_pr

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
	"trustedStorage/grpsServer"
	"trustedStorage/mempool"
	"trustedStorage/p2pCommunication"
	"trustedStorage/settings"
	"trustedStorage/transaction"
)

//var Sh = shell.NewShell("localhost:5001")

type myInvoicerServer struct {
	grpsServer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) AddTxToBlockchain(ctx context.Context, req *grpsServer.CreateTx) (*grpsServer.CreateResponse, error) {
	var err error
	tx := transaction.Transaction{
		SenderAddress: req.SenderAddress,
		Data:          req.Data,
		PubKey:        req.PubKey,
		Signature:     req.Signature,
		Nonce:         req.Nonce,
		Cid:           req.Cid,
	}
	//documentBytes := req.DocumentBytes
	log.Printf("Transaction from %x received", req.SenderAddress)

	//if len(documentBytes) > 3000000 {
	//	err = errors.New("document is too big")
	//	log.Println(err)
	//	return &grpsServer.CreateResponse{
	//		Response: err.Error(),
	//	}, nil
	//}

	err = mempool.MemPoolIns.AddTxToMempool(tx)
	if err != nil {
		log.Println(err)
		return &grpsServer.CreateResponse{
			Response: err.Error(),
		}, nil
	}
	p2pCommunication.ProposeTransaction(tx)
	//fmt.Println(tx.ToString())

	return &grpsServer.CreateResponse{
		Response: "Tx added to mempool\nTx hash: " + hex.EncodeToString(tx.GetTxHash()),
	}, nil
}

func (s *myInvoicerServer) GetLastBlock(ctx context.Context, emp *emptypb.Empty) (*grpsServer.CreateResponse, error) {
	log.Println("Sending last block")
	return &grpsServer.CreateResponse{
		//Response: (*BlockChainIns)[len(*BlockChainIns)-1].ToString(),
		Response: blockchain.BlockChainIns.GetLastBlock().ToString(),
	}, nil
}

func (s *myInvoicerServer) GetBlockchain(ctx context.Context, emp *emptypb.Empty) (*grpsServer.CreateResponse, error) {
	log.Println("Sending blockchain history")
	return &grpsServer.CreateResponse{

		Response: blockchain.BlockChainIns.ToString(),
	}, nil
}

func (s *myInvoicerServer) GetUserTxHistory(ctx context.Context, req *grpsServer.User) (*grpsServer.CreateResponse, error) {
	log.Println("Sending user tx history")
	return &grpsServer.CreateResponse{
		Response: blockchain.GetUserTxHistory(blockchain.BlockChainIns, req.SenderAddress),
	}, nil
}

func (s *myInvoicerServer) GetTxByHash(ctx context.Context, req *grpsServer.Transaction) (*grpsServer.CreateResponse, error) {
	log.Printf("Finding transaction %x by hash\n", req.TxHash)

	var res string

	IsFind, blockNum, tx := blockchain.FindTransactionByHash(req.TxHash)
	if IsFind {
		res = fmt.Sprintf("Tx has been found in block on height %v\nTx Info:\n%v", strconv.Itoa(blockNum), tx.ToString())
	} else {
		res = "Can't find tx. It does not exist or has not been confirmed yet"
	}
	return &grpsServer.CreateResponse{
		Response: res,
	}, nil
}

func (s *myInvoicerServer) FindDocumentByHash(ctx context.Context, req *grpsServer.Document) (*grpsServer.CreateResponse, error) {
	log.Printf("Finding document %x by hash\n", req.DocumentHash)

	var res string

	IsFind, blockNum, tx := blockchain.FindDocumentByHash(req.DocumentHash)
	if IsFind {
		res = fmt.Sprintf("Document has been found in block on height %v\nTx Info:\n%v", strconv.Itoa(blockNum), tx.ToString())
	} else {
		res = "Can't find document. It does not exist or has not been confirmed yet"
	}
	return &grpsServer.CreateResponse{
		Response: res,
	}, nil
}

func StartGrpc() {

	lis, err := net.Listen("tcp", settings.GetRpcNodeAddress())
	if err != nil {
		log.Fatalf("Cannot create listener : %s", err)
		//lis, err = net.Listen("tcp", ":0")
	}

	ServerRegistrar := grpc.NewServer()

	service := &myInvoicerServer{}
	grpsServer.RegisterInvoicerServer(ServerRegistrar, service)

	log.Println("RPC listening on ", lis.Addr())

	err = ServerRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}

}
