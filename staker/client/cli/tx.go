package staker

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authctx "github.com/cosmos/cosmos-sdk/x/auth/client/context"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/maticnetwork/heimdall/staker"
)

func GetCmdCreateMaticValidator(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "godmode",
		Args:  cobra.ExactArgs(1),
		Short: "Create a matic validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			txCtx := authctx.NewTxContextFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithLogger(os.Stdout).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

			fromAddress, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}
			pkStr := args[0]
			if len(pkStr) == 0 {
				return fmt.Errorf("must use --pubkey flag")
			}

			var x secp256k1.PubKeySecp256k1
			k, _ := hex.DecodeString(pkStr)
			copy(x[:], k[:])

			fmt.Printf("The address is %v", x.Address())

			//todo send message
			msg := staker.NewCreateMaticValidator(fromAddress, x.Address(), x, int64(3))

			//// build and sign the transaction, then broadcast to Tendermint
			return utils.SendTx(txCtx, cliCtx, []sdk.Msg{msg})
		},
	}

	return cmd
}
