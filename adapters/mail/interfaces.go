package mail

type Mail interface {
	Send(data *SendReqSt) bool
}
