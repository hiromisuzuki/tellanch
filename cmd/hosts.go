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
	"fmt"
	"strconv"
	"strings"

	"github.com/hiromisuzuki/tellanch/config"
	"github.com/spf13/cobra"
)

// hostsCmd represents the hosts command
var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Show host settings in .tellanch.yaml",
	Long:  `Show host settings in .tellanch.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		var c config.ConfigManager
		c.Load()
		for i, v := range c {
			fmt.Println("Host " + strconv.Itoa(i+1))
			fmt.Println("  User: " + v.User)
			fmt.Println("  Address: " + v.GetAddress())
			fmt.Println("  Key: " + v.Key)
			fmt.Println("  Path: " + strings.Join(v.Path, ", "))
		}
	},
}

func init() {
	RootCmd.AddCommand(hostsCmd)
}
