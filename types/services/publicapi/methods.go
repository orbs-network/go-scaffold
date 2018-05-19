package publicapi

type Methods interface {
	Transfer(*TransferInput) (*TransferOutput, error)
	GetBalance(*GetBalanceInput) (*GetBalanceOutput, error)
}
