package statestorage

type Methods interface {
	WriteKey(*WriteKeyInput) (*WriteKeyOutput, error)
	ReadKey(*ReadKeyInput) (*ReadKeyOutput, error)
}
