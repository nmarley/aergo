/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */

package cmd

import (
	"context"
	"errors"

	"github.com/aergoio/aergo/cmd/aergocli/util"
	"github.com/aergoio/aergo/types"
	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
)

var sendtxCmd = &cobra.Command{
	Use:   "sendtx",
	Short: "Send transaction",
	Args:  cobra.MinimumNArgs(0),
	RunE:  execSendTX,
}
var chainIdHash string

func init() {
	rootCmd.AddCommand(sendtxCmd)
	sendtxCmd.Flags().StringVar(&from, "from", "", "Sender account address")
	sendtxCmd.MarkFlagRequired("from")
	sendtxCmd.Flags().StringVar(&to, "to", "", "Recipient account address")
	sendtxCmd.MarkFlagRequired("to")
	sendtxCmd.Flags().StringVar(&amount, "amount", "0", "How much in AER")
	sendtxCmd.MarkFlagRequired("amount")
	sendtxCmd.Flags().Uint64Var(&nonce, "nonce", 0, "setting nonce manually")
	sendtxCmd.Flags().StringVar(&chainIdHash, "chainidhash", "", "hash value of chain id in the block")
	sendtxCmd.Flags().Uint64VarP(&gas, "gaslimit", "g", 0, "Gas limit")
}

func execSendTX(cmd *cobra.Command, args []string) error {
	account, err := types.DecodeAddress(from)
	if err != nil {
		return errors.New("Wrong address in --from flag\n" + err.Error())
	}
	recipient, err := types.DecodeAddress(to)
	if err != nil {
		return errors.New("Wrong address in --to flag\n" + err.Error())
	}
	amountBigInt, err := util.ParseUnit(amount)
	if err != nil {
		return errors.New("Wrong value in --amount flag\n" + err.Error())
	}
	tx := &types.Tx{Body: &types.TxBody{
		Type:      types.TxType_TRANSFER,
		Account:   account,
		Recipient: recipient,
		Amount:    amountBigInt.Bytes(),
		Nonce:     nonce,
		GasLimit:  gas,
	}}
	if chainIdHash != "" {
		cid, err := base58.Decode(chainIdHash)
		if err != nil {
			return errors.New("Wrong value in --chainidhash flag\n" + err.Error())
		}
		tx.GetBody().ChainIdHash = cid
	}
	msg, err := client.SendTX(context.Background(), tx)
	if err != nil {
		cmd.Println(err.Error())
		return nil
	}
	cmd.Println(util.JSON(msg))
	return nil
}
