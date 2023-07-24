package smss

type SendReqSt struct {
	To   string `json:"to"`
	Text string `json:"text"`
	Sync bool   `json:"sync"`
}

type SendRepSt struct {
	ID string `json:"id"`
}

type ErrorRepSt struct {
	ErrorCode string `json:"error_code"`
}
