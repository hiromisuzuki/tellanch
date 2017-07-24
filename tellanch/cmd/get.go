// Copyright © 2017 Suzuki Hiromi
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/hiromisuzuki/tellanch/config"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get branch name from git",
	Long:  `Get branch name from git`,
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Manager
		c.Load()
		for _, v := range c {
			connect(v)
		}
	},
}

func connect(v *config.Host) {
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
		Timeout: time.Second * 30,
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

	session, err := conn.NewSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	for _, p := range v.Path {
		fmt.Println(branch(session, p))
	}
}

func branch(session *ssh.Session, path string) string {
	var b bytes.Buffer
	session.Stdout = &b
	c := "pwd " + path
	if err := session.Run(c); err != nil {
		panic("Failed to run [" + c + "]: " + err.Error())
	}
	return b.String()
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
