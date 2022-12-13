package keygen

type Request struct {
	Keys      []string `json:"keys"`
	ThresHold int      `json:"thres_hold"`
}

func NewRequest(keys []string, threshold int) Request {
	return Request{
		Keys:      keys,
		ThresHold: threshold,
	}
}
