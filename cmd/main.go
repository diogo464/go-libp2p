package main

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name: "example",
	Commands: []*cli.Command{
		{
			Name:   "initiator",
			Action: actionInitiator,
		},
		{
			Name:   "receiver",
			Action: actionReceiver,
		},
	},
}

func main() {
	app.RunAndExitOnError()
}

func actionInitiator(c *cli.Context) error {
	h, err := libp2p.New(libp2p.NoListenAddrs)
	if err != nil {
		return err
	}

	receiver, err := peer.AddrInfoFromString(c.Args().Get(0))
	if err != nil {
		return err
	}

	fmt.Println("Connecting to receiver at ", receiver.String())
	if err := h.Connect(c.Context, *receiver); err != nil {
		return err
	}

	protocols, err := h.Network().Peerstore().GetProtocols(receiver.ID)
	if err != nil {
		return err
	}

	fmt.Println("Protocols:", protocols)
	time.Sleep(time.Second)

	return nil
}

func actionReceiver(c *cli.Context) error {
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	if err != nil {
		return err
	}

	for _, addr := range h.Addrs() {
		println("Listening at:", addr.String()+"/p2p/"+h.ID().Pretty())
	}
	fmt.Println("Sleeping for 60 seconds...")
	time.Sleep(60 * time.Second)

	return nil
}
