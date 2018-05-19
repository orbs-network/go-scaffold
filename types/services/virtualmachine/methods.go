package virtualmachine

type Methods interface {
	ProcessTransaction(*ProcessTransactionInput) (*ProcessTransactionOutput, error)
	CallContract(*CallContractInput) (*CallContractOutput, error)
}
