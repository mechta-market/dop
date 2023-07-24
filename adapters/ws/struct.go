package ws

type SendReqSt struct {
	UsrIds  []int64 `json:"usr_ids"`
	Message any     `json:"message"`
}

type ConnectionCountRepSt struct {
	Value int64 `json:"value"`
}
