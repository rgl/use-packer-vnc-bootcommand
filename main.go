package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/hashicorp/packer-plugin-sdk/bootcommand"
	"github.com/mitchellh/go-vnc"
)

func main() {
	vncAddress := flag.String("address", "127.0.0.1:5901", "VNC server address")
	vncPassword := flag.String("password", "password", "VNC server password")
	bootCommand := flag.String("boot-command", "<leftAltOn><f2><leftAltOff>", "Boot command to send tru VNC")

	flag.Parse()

	bootCommandSequence, err := bootcommand.GenerateExpressionSequence(*bootCommand)
	if err != nil {
		log.Fatalf("ERROR failed to parse the boot command: %v", err)
	}

	vncAddr, err := net.ResolveTCPAddr("tcp", *vncAddress)
	if err != nil {
		log.Fatalf("ERROR failed to resolve to %s: %v", *vncAddress, err)
	}

	vncConn, err := net.DialTCP("tcp", nil, vncAddr)
	if err != nil {
		log.Fatalf("ERROR failed to connect to %s: %v", *vncAddress, err)
	}

	vncConfig := &vnc.ClientConfig{
		//Exclusive: true,
	}
	if *vncPassword != "" {
		vncConfig.Auth = []vnc.ClientAuth{
			&vnc.PasswordAuth{Password: *vncPassword},
		}
	}
	vncClient, err := vnc.Client(vncConn, vncConfig)
	if err != nil {
		log.Fatalf("ERROR failed to connect to VNC server at %s: %v", *vncAddress, err)
	}
	defer vncClient.Close()

	log.Printf(
		"Connected to VNC server %s screen %s (%dx%d)",
		*vncAddress,
		vncClient.DesktopName,
		vncClient.FrameBufferWidth,
		vncClient.FrameBufferHeight)

	vncDriver := bootcommand.NewVNCDriver(vncClient, bootcommand.PackerKeyDefault)

	log.Println("Sending the boot command...")
	err = bootCommandSequence.Do(context.Background(), vncDriver)
	if err != nil {
		log.Fatalf("ERROR failed to send the boot command: %v", err)
	}
}
