package nft

import (
	"github.com/irismod/nft/keeper"
	"github.com/irismod/nft/types"
)

const (
	QuerySupply     = types.QuerySupply
	QueryOwner      = types.QueryOwner
	QueryCollection = types.QueryCollection
	QueryDenoms     = types.QueryDenoms
	QueryNFT        = types.QueryNFT
	ModuleName      = types.ModuleName
	StoreKey        = types.StoreKey
	QuerierRoute    = types.QuerierRoute
	RouterKey       = types.RouterKey
)

var (
	// functions aliases
	RegisterInvariants = keeper.RegisterInvariants
	AllInvariants      = keeper.AllInvariants
	SupplyInvariant    = keeper.SupplyInvariant
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier

	RegisterCodec            = types.RegisterCodec
	NewCollection            = types.NewCollection
	NewCollections           = types.NewCollections
	NewGenesisState          = types.NewGenesisState
	GetOwnerKey              = types.KeyOwner
	NewMsgIssueDenom         = types.NewMsgIssueDenom
	NewMsgTransferNFT        = types.NewMsgTransferNFT
	NewMsgEditNFTMetadata    = types.NewMsgEditNFT
	NewMsgMintNFT            = types.NewMsgMintNFT
	NewMsgBurnNFT            = types.NewMsgBurnNFT
	NewBaseNFT               = types.NewBaseNFT
	NewNFTs                  = types.NewNFTs
	NewIDCollection          = types.NewIDCollection
	NewOwner                 = types.NewOwner
	NewOwners                = types.NewOwners
	NewQuerySupplyParams     = types.NewQuerySupplyParams
	NewQueryCollectionParams = types.NewQueryCollectionParams
	NewQueryNFTParams        = types.NewQueryNFTParams

	//validate field function
	ValidateDenom    = types.ValidateDenom
	ValidateTokenID  = types.ValidateTokenID
	ValidateTokenURI = types.ValidateTokenURI

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	EventTypeTransfer        = types.EventTypeTransfer
	EventTypeEditNFTMetadata = types.EventTypeEditNFT
	EventTypeMintNFT         = types.EventTypeMintNFT
	EventTypeBurnNFT         = types.EventTypeBurnNFT
	AttributeValueCategory   = types.AttributeValueCategory
	AttributeKeySender       = types.AttributeKeySender
	AttributeKeyRecipient    = types.AttributeKeyRecipient
	AttributeKeyOwner        = types.AttributeKeyOwner
	AttributeKeyNFTID        = types.AttributeKeyTokenID
	AttributeKeyNFTTokenURI  = types.AttributeKeyTokenURI
	AttributeKeyDenom        = types.AttributeKeyDenom

	// error
	ErrInvalidCollection = types.ErrInvalidCollection
	ErrUnknownCollection = types.ErrUnknownCollection
	ErrInvalidNFT        = types.ErrInvalidNFT
	ErrNFTAlreadyExists  = types.ErrNFTAlreadyExists
	ErrUnknownNFT        = types.ErrUnknownNFT
	ErrEmptyMetadata     = types.ErrEmptyMetadata
)

type (
	Keeper                = keeper.Keeper
	Collection            = types.Collection
	Collections           = types.Collections
	GenesisState          = types.GenesisState
	MsgTransferNFT        = types.MsgTransferNFT
	MsgEditNFT            = types.MsgEditNFT
	MsgMintNFT            = types.MsgMintNFT
	MsgBurnNFT            = types.MsgBurnNFT
	BaseNFT               = types.BaseNFT
	NFTs                  = types.NFTs
	Denom                 = types.Denom
	IDCollection          = types.IDCollection
	IDCollections         = types.IDCollections
	Owner                 = types.Owner
	QuerySupplyParams     = types.QuerySupplyParams
	QueryCollectionParams = types.QueryCollectionParams
	QueryNFTParams        = types.QueryNFTParams
)
