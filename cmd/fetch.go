/*
Copyright © 2021 Graham Plata <graham.plata@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"eta-multitool/pkg/components"
	"eta-multitool/pkg/config"
	"eta-multitool/pkg/text"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fetchCmd represents the base verb command for fetch keyword
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch details about 'GRIM' artifacts",
	Long:  "Fetch details about 'GRIM' artifacts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eta-multitool fetch")
	},
}

// tokenCmd token action
var tokenCmd = &cobra.Command{
	Use:   "tokens",
	Short: "Returns all tokens from the given authority",
	Long: `Returns all tokens from the given authority

This includes tokens from the Original Grim collection, 
the Daemons of the 'Lurkers of the Abyss' and finally 
the Ingredients of Lordrym's Workshop. 

This operation is very slow unless a private node is used.
https://docs.solana.com/cluster/rpc-endpoints#rate-limits-2`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		config.SetOutput(output)
		components.GetAllMetaplexTokenByAuthority()
	},
}

// metadataCmd metadata action
var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Fetches token Metadata",
	Long: `Fetches detailed NFT token Metadata

Given a token address returns a formatted json blob from arweave.
`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		config.SetOutput(output)
		if len(args) < 1 {
			logrus.Error(fmt.Sprintf("%s: fetch metadata requires an argument", text.GrimError))
		}
		components.GetMetadataByAddress(args[0])
	},
}

// walletCmd metadata action
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "fetch the 'GRIM' community wallet",
	Long: `fetch the 'GRIM' community wallet

'GRIM'
community wallet address:    Es1YghGkHZNJ8A9r6oFEHbWsRHbqs4rz6gfkRJ9V4bYf
cold storage wallet address: RTp26f9wY2fXxeWRE7FkS9iVrsuxgdUJfDYH8GgoBH9`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		config.SetOutput(output)
		// add cold storage wallet
		components.GetWallet()
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.AddCommand(tokenCmd, metadataCmd, walletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")
	rootCmd.PersistentFlags().String("all", "", "output type of the cli (default is ...)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
