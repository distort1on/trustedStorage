package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	peerstore "github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/multiformats/go-multiaddr"
	"io"
	"log"
)

type myNotifee struct {
	h host.Host
}

func (mn myNotifee) HandlePeerFound(info peerstore.AddrInfo) {

	fmt.Println("found peer", info.Addrs, info.ID)

	err := mn.h.Connect(context.Background(), info)
	mn.h.Peerstore().AddAddrs(info.ID, info.Addrs, pstore.PermanentAddrTTL)

	if err != nil {
		panic(err)
	}

	fmt.Println(mn.h.Network().Peers())

}

func StartPeerDiscovery(node *host.Host) {
	n := myNotifee{h: *node}

	discoveryService := mdns.NewMdnsService(
		*node,
		"example",
		&n,
	)

	err := discoveryService.Start()
	if err != nil {
		return
	}

}

func ReadData(s network.Stream) {
	test, _ := io.ReadAll(s)
	fmt.Println(test)
	peerID := s.Conn().RemotePeer()
	fmt.Println(peerID)

	fmt.Println("Received data")
	buf := bufio.NewReader(s)
	header, err := buf.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}

	payload := make([]byte, header)
	n, err := io.ReadFull(buf, payload)
	log.Printf("payload has %d bytes", n)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("read: %s", payload)

	//log.Printf("Message from '%s': %s", connection.RemotePeer().String(), message)
}

func WriteData(stream *network.Stream, data []byte) {
	payload := data
	header := []byte{byte(len(payload))}
	_, err := (*stream).Write(header)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = (*stream).Write(payload)
	if err != nil {
		log.Println(err)
		return
	}
}

func SendDataToPeer(node *host.Host, targetPeerAddr string, data []byte) {
	log.Println("sending message to ", targetPeerAddr)
	addr, err := multiaddr.NewMultiaddr(targetPeerAddr)
	if err != nil {
		panic(err)
	}
	peer, err := peerstore.AddrInfoFromP2pAddr(addr)
	if err != nil {
		panic(err)
	}

	stream, err := (*node).NewStream(context.Background(), peer.ID, "/send/1.0.0")
	if err != nil {
		panic(err)
	}
	WriteData(&stream, data)
	stream.Close()
}

func main() {
	//addr, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/50030/p2p/12D3KooWJ7NWAWXzXyfpVYPUJDtoy8XnJRFgoJDBeg7i9cudKEsr")
	//if err != nil {
	//	panic(err)
	//}

	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
	)

	//fmt.Println(node.Peerstore().PrivKey(node.ID()))

	if err != nil {
		panic(err)
	}

	peerInfo := peerstore.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}

	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("libp2p node address:", addrs[0])

	node.SetStreamHandler("/send/1.0.0", func(s network.Stream) {
		//log.Printf("/send/1.0.0 stream created")
		ReadData(s)
		//if err != nil {
		//	s.Reset()
		//} else {
		//	s.Close()
		//}
	})

	StartPeerDiscovery(&node)
	var a string
	for {
		fmt.Scan(&a)
		SendDataToPeer(&node, a, []byte("Ping!"))
	}

	//test_pr.StartGrpc()

}
