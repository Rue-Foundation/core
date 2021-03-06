package commands

import (
	"os"

	"github.com/sonm-io/core/accounts"
	"github.com/sonm-io/core/util"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Open or generate Etherum keys",
	Run: func(cmd *cobra.Command, _ []string) {
		ko, err := accounts.DefaultKeyOpener(cmd, cfg.KeyStore(), cfg.PassPhrase())
		if err != nil {
			showError(cmd, "Cannot init KeyOpener", err)
			os.Exit(1)
		}

		created, err := ko.OpenKeystore()
		if err != nil {
			showError(cmd, "Cannot open KeyStore", err)
			os.Exit(1)
		}

		if created {
			cmd.Printf("Keystore successfully created\r\n")
		} else {
			cmd.Printf("Keystore successfully opened\r\n")
		}

		key, err := ko.GetKey()
		if err != nil {
			showError(cmd, "Cannot get key from the keystore", err)
			os.Exit(1)
		}

		cmd.Printf("Eth address: %s\r\n", util.PubKeyToAddr(key.PublicKey).Hex())
	},
}
