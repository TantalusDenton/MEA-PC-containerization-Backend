package main

import (
	"fmt"

	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func connectToLXDserver() error {
	// Connect to LXD over the Unix socket
	argumentsToPass := &lxd.ConnectionArgs{TLSClientCert: "/home/lxd-api-access-cert-key-files/lxd-type3access.crt", InsecureSkipVerify: true}

	c, err := lxd.ConnectLXD("127.0.0.1", argumentsToPass)
	if err != nil {
		fmt.Print("Could not connect")
		return err
	}

	// Container creation request
	req := api.ContainersPost{
		Name: "madewithapi",
		Source: api.ContainerSource{
			Type:  "image",
			Alias: "image4go",
		},
	}

	// Get LXD to create the container (background operation)
	op, err := c.CreateContainer(req)
	if err != nil {
		return err
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return err
	}

	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	op, err = c.UpdateContainerState("madewithapi", reqState, "")
	if err != nil {
		return err
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return err
	}
	return err
}

func main() {
	connectToLXDserver()
}
