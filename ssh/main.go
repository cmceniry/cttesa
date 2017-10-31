package main

import (
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

func getKey() ssh.Signer {
	keyFile := "/home/got/.ssh/id_rsa"
	pemKey, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(pemKey)
	if err != nil {
		panic(err)
	}
	return signer
}

func connect(host string, signer ssh.Signer) *ssh.Client {
	config := &ssh.ClientConfig{
		User:            "got",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	signer := getKey()
	client := connect("127.0.0.1:22", signer)

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	sin, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	err = session.Start(...)
	if err != nil {
		panic(err)
	}
	...
	sin.Close()
	session.Wait()
}
