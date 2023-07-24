package mail

type SendReqSt struct {
	Receivers []string `json:"receivers"`
	Subject   string   `json:"subject"`
	Message   string   `json:"message"`
	Sync      bool     `json:"sync"`
}
