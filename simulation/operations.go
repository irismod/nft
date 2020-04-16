package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irismod/nft/keeper"
	"github.com/irismod/nft/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgMintNFT     = "op_weight_msg_mint_nft"
	OpWeightMsgEditNFT     = "op_weight_msg_edit_nft_metadata"
	OpWeightMsgTransferNFT = "op_weight_msg_transfer_nft"
	OpWeightMsgBurnNFT     = "op_weight_msg_transfer_burn_nft"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simulation.AppParams,
	cdc *codec.Codec,
	k keeper.Keeper, ak auth.AccountKeeper) simulation.WeightedOperations {

	var weightMint, weightEdit, weightBurn, weightTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgMintNFT, &weightMint, nil,
		func(_ *rand.Rand) {
			weightMint = 100
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgEditNFT, &weightEdit, nil,
		func(_ *rand.Rand) {
			weightEdit = 50
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgTransferNFT, &weightTransfer, nil,
		func(_ *rand.Rand) {
			weightTransfer = 50
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBurnNFT, &weightBurn, nil,
		func(_ *rand.Rand) {
			weightBurn = 10
		},
	)

	return simulation.WeightedOperations{
		//simulation.NewWeightedOperation(
		//	weightMint,
		//	SimulateMsgMintNFT(k, ak),
		//),
		simulation.NewWeightedOperation(
			weightEdit,
			SimulateMsgEditNFTMetadata(k, ak),
		),
		//simulation.NewWeightedOperation(
		//	weightTransfer,
		//	SimulateMsgTransferNFT(k, ak),
		//),
		//simulation.NewWeightedOperation(
		//	weightBurn,
		//	SimulateMsgBurnNFT(k, ak),
		//),
	}
}

// SimulateMsgTransferNFT simulates the transfer of an NFT
func SimulateMsgTransferNFT(k keeper.Keeper, ak auth.AccountKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, chainID string) (opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		ownerAddr, denom, nftID := getRandomNFTFromOwner(ctx, k, r)
		if ownerAddr.Empty() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		simAccount, _ := simulation.RandomAcc(r, accs)

		msg := types.NewMsgTransferNFT(
			ownerAddr,          // sender
			simAccount.Address, // recipient
			denom,
			nftID,
			"",
		)
		account := ak.GetAccount(ctx, msg.Sender)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err = app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgEditNFTMetadata simulates an edit metadata transaction
func SimulateMsgEditNFTMetadata(k keeper.Keeper, ak auth.AccountKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, chainID string) (opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		ownerAddr, denom, nftID := getRandomNFTFromOwner(ctx, k, r)
		if ownerAddr.Empty() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgEditNFT(
			ownerAddr,
			nftID,
			denom,
			simulation.RandStringOfLength(r, 45), // tokenURI
		)

		account := ak.GetAccount(ctx, msg.Sender)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		simAccount, found := simulation.FindAccount(accs, ownerAddr)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account %s not found", ownerAddr)
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err = app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgMintNFT simulates a mint of an NFT
func SimulateMsgMintNFT(k keeper.Keeper, ak auth.AccountKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, chainID string) (opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {

		randomSender, _ := simulation.RandomAcc(r, accs)
		randomRecipient, _ := simulation.RandomAcc(r, accs)

		msg := types.NewMsgMintNFT(
			randomSender.Address,                 // sender
			randomRecipient.Address,              // recipient
			simulation.RandStringOfLength(r, 10), // nft ID
			simulation.RandStringOfLength(r, 10), // denom
			simulation.RandStringOfLength(r, 45), // tokenURI
		)

		account := ak.GetAccount(ctx, msg.Sender)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		simAccount, found := simulation.FindAccount(accs, msg.Sender)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account %s not found", msg.Sender)
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err = app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgBurnNFT simulates a burn of an existing NFT
func SimulateMsgBurnNFT(k keeper.Keeper, ak auth.AccountKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, chainID string) (opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		ownerAddr, denom, nftID := getRandomNFTFromOwner(ctx, k, r)
		if ownerAddr.Empty() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgBurnNFT(ownerAddr, nftID, denom)

		account := ak.GetAccount(ctx, msg.Sender)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		simAccount, found := simulation.FindAccount(accs, msg.Sender)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account %s not found", msg.Sender)
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err = app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func getRandomNFTFromOwner(ctx sdk.Context, k keeper.Keeper, r *rand.Rand) (address sdk.AccAddress, denom, nftID string) {
	owners := k.GetOwners(ctx)

	ownersLen := len(owners)
	if ownersLen == 0 {
		return nil, "", ""
	}

	// get random owner
	i := r.Intn(ownersLen)
	owner := owners[i]

	idCollectionsLen := len(owner.IDCollections)
	if idCollectionsLen == 0 {
		return nil, "", ""
	}

	// get random collection from owner's balance
	i = r.Intn(idCollectionsLen)
	idsCollection := owner.IDCollections[i] // nfts IDs
	denom = idsCollection.Denom

	idsLen := len(idsCollection.IDs)
	if idsLen == 0 {
		return nil, "", ""
	}

	// get random nft from collection
	i = r.Intn(idsLen)
	nftID = idsCollection.IDs[i]

	return owner.Address, denom, nftID
}
