package keysign

type Request struct {
	PoolPubKey    string   `json:"pool_pub_key"`
	Message       []byte   `json:"message"`
	SignerPubKeys []string `json:"signer_pub_keys"`
}

func NewRequest(pk string, msg []byte, signers []string) Request {
	return Request{
		PoolPubKey:    pk,
		Message:       msg,
		SignerPubKeys: signers,
	}
}
