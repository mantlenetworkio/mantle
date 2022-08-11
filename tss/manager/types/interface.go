package types

type SignService interface {
	SignStateBatch(request SignStateRequest) ([]byte, error)
	SignTxBatch() error
}
