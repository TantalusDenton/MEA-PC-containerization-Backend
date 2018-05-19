package main

import (
	"fmt"
	"io/ioutil"

	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func checkFileForError(e error) {
	if e != nil {
		panic(e)
	}
}

func connectToLXDserver() error {

	// Connection parameters - LXD API needs to know client vert, key and server cert
	ClientCertFile, errcert := ioutil.ReadFile("/home/tantalus/lxd-api-access-cert-key-files/lxd-type3access.crt")
	checkFileForError(errcert)
	ClientCertString := string(ClientCertFile)

	ClientKeyFile, errkey := ioutil.ReadFile("/home/tantalus/lxd-api-access-cert-key-files/lxd-type3access.key")
	checkFileForError(errkey)
	ClientKeyString := string(ClientKeyFile)

	ServerCertFile, errservercert := ioutil.ReadFile("/home/tantalus/lxd-api-access-cert-key-files/lxd-type3server.crt")
	checkFileForError(errservercert)
	ServerCertString := string(ServerCertFile)

	argumentsToPass := &lxd.ConnectionArgs{
		TLSClientCert: ClientCertString,
		TLSClientKey:  ClientKeyString,
		TLSServerCert: ServerCertString,
		/*InsecureSkipVerify: true*/}

	// Connect to LXD over http
	c, err := lxd.ConnectLXD("https://127.0.0.1:8443", argumentsToPass)
	if err != nil {
		fmt.Print("Could not connect because of error: ", err)
		fmt.Print("server cert is: ", ServerCertString)
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
		fmt.Print("Could not create container because of error: ", err)
		return err
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		fmt.Print("Could not wait for operation because of error: ", err)
		return err
	}

	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	op, err = c.UpdateContainerState("madewithapi", reqState, "")
	if err != nil {
		fmt.Print("Could not update container status because of error: ", err)
		return err
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		fmt.Print("Could not wait because of error: ", err)
		return err
	}
	return err
}

func main() {
	connectToLXDserver()
}
