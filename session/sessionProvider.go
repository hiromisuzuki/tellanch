package session

import (
	"bytes"
	"io/ioutil"
	"net"
	"strings"

	"github.com/hiromisuzuki/tellanch/config"
	"golang.org/x/crypto/ssh"
)

//SessionProvider SessionProvider
type SessionProvider struct {
	Host    *config.Host
	client  *ssh.Client
	session *ssh.Session
}

func (v *SessionProvider) NewSession() (*SessionProvider, error) {
	config, err := v.getConfig()
	if err != nil {
		return v, err
	}
	v.client, err = ssh.Dial("tcp", v.Host.GetAddress(), config)
	if err != nil {
		return v, err
	}
	v.session, err = v.client.NewSession()
	if err != nil {
		return v, err
	}
	return v, nil
}

func (v *SessionProvider) Run() string {
	var b bytes.Buffer
	v.session.Stdout = &b
	c := v.getOneLineCommand()
	if err := v.session.Run(c); err != nil {
		panic("Failed to run [" + c + "]: " + err.Error())
	}
	return b.String()
}

func (v *SessionProvider) Close() {
	v.session.Close()
	v.client.Close()
}

func (v *SessionProvider) getOneLineCommand() string {
	var cmd []string
	for _, path := range v.Host.Path {
		c := "pushd " + path + "; cat .git/HEAD; popd"
		cmd = append(cmd, c)
	}
	return strings.Join(cmd, ";")
}

func (v *SessionProvider) getConfig() (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{
		User: v.Host.User,
		HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error {
			return nil
		},
	}
	if key := v.Host.GetKey(); key != nil {
		f, err := ioutil.ReadFile(*key)
		if err != nil {
			return nil, err
		}
		pem, err := ssh.ParsePrivateKey(f)
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(pem),
		}
	}
	return config, nil
}
