package session

import (
	"io/ioutil"
	"log"
	"net"
	"strconv"

	"github.com/hiromisuzuki/tellanch/config"
	"golang.org/x/crypto/ssh"
)

type SessionProvider struct {
	Host    *config.Host
	client  *ssh.Client
	session *ssh.Session
}

func (v *SessionProvider) NewSession() (*SessionProvider, error) {
	f, err := ioutil.ReadFile(v.Key)
	if err != nil {
		panic(err)
	}
	key, err := ssh.ParsePrivateKey(f)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: v.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error {
			return nil
		},
	}
	port := v.Port
	if port == 0 {
		port = 22
	}

	conn, err := ssh.Dial("tcp", v.Address+":"+strconv.Itoa(port), config)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	return v, nil
}

func (v *SessionProvider) Close() {
	v.session.Close()
	v.client.Close()
}
