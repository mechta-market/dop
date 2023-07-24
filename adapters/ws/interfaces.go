package ws

type Ws interface {
	Send2User(usrId int64, data any) error
	Send2Users(usrIds []int64, data any) error
	GetConnectionCount() (int64, error)
}
