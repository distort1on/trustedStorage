package p2pCommunication

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"io"
	"log"
	"trustedStorage/settings"
)

//func KademiaDHT() {
//	ctx := context.Background()
//	kademliaDHT, err := dht.New(ctx, Node)
//	if err != nil {
//		panic(err)
//	}
//
//	if err = kademliaDHT.Bootstrap(ctx); err != nil {
//		panic(err)
//	}
//
//	var wg sync.WaitGroup
//	for _, peerAddr := range dht.DefaultBootstrapPeers {
//		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			if err := Node.Connect(ctx, *peerinfo); err != nil {
//				log.Println(err)
//			} else {
//				log.Println("Connection established with bootstrap node:", *peerinfo)
//			}
//		}()
//	}
//	wg.Wait()
//
//	routingDiscovery := drouting.NewRoutingDiscovery(kademliaDHT)
//	dutil.Advertise(ctx, routingDiscovery, "test")
//	peerChan, err := routingDiscovery.FindPeers(ctx, "test")
//	if err != nil {
//		panic(err)
//	}
//
//	for p := range peerChan {
//
//		log.Println(p.ID.String())
//	}
//}

var Node host.Host

type myNotifee struct {
	h host.Host
}

type PeersResponse struct {
}

func (mn myNotifee) HandlePeerFound(info peer.AddrInfo) {

	fmt.Println("found peer", info.Addrs, info.ID)

	err := mn.h.Connect(context.Background(), info)
	//mn.h.peer().AddAddrs(info.ID, info.Addrs, pstore.PermanentAddrTTL)

	if err != nil {
		log.Println(err)
	}

	//fmt.Println(mn.h.Network().Peers())
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
		log.Println(err)
	}

}

func PeerStreamHandler(s network.Stream) {
	//mu.Lock()
	//defer mu.Unlock()

	log.Println("Received message"+" from", s.Conn().RemotePeer())

	data, err := io.ReadAll(s)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(data[:3]))
	NodeActionDecision(data, s.Conn().RemotePeer(), false)

	//ReadData(s)
	//if err != nil {
	//	s.Reset()
	//} else {
	//	s.Close()
	//}
}

func LaunchP2PPeer() {
	var err error

	keyString, err := hex.DecodeString(settings.GetNodePrivateKey())
	if err != nil {
		log.Println(err)
	}

	nodePrivKey, err := crypto.UnmarshalECDSAPrivateKey(keyString)
	if err != nil {
		log.Println(err)
	}

	Node, err = libp2p.New(
		libp2p.Identity(nodePrivKey),
		libp2p.ListenAddrStrings(settings.GetP2PNodeAddress()),
	)

	if err != nil {
		panic(err)
	}

	peerInfo := peer.AddrInfo{
		ID:    Node.ID(),
		Addrs: Node.Addrs(),
	}

	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("libp2p node address:", addrs[0])

	Node.SetStreamHandler("/send/1.0.0", PeerStreamHandler)

	StartPeerDiscovery(&Node)

	//test_pr.StartGrpc()

}
