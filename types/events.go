package types

// NFT module event types
var (
	EventTypeTransfer = "transfer_nft"
	EventTypeEditNFT  = "edit_nft_"
	EventTypeMintNFT  = "mint_nft"
	EventTypeBurnNFT  = "burn_nft"

	AttributeValueCategory = ModuleName

	AttributeKeySender      = "sender"
	AttributeKeyRecipient   = "recipient"
	AttributeKeyOwner       = "owner"
	AttributeKeyNFTTokenID  = "token-id"
	AttributeKeyNFTTokenURI = "token-uri"
	AttributeKeyDenom       = "denom"
)
