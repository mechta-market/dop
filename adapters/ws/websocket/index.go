package websocket

import (
	"time"

	"github.com/mechta-market/dop/adapters/client/httpc"
	"github.com/mechta-market/dop/adapters/ws"
)

type St struct {
	httpc httpc.HttpC
}

func New(httpc httpc.HttpC) *St {
	return &St{
		httpc: httpc,
	}
}

func (p *St) Send2User(usrId int64, data any) error {
	return p.Send2Users([]int64{usrId}, data)
}

func (p *St) Send2Users(usrIds []int64, data any) error {
	if len(usrIds) == 0 {
		return nil
	}

	reqObj := ws.SendReqSt{
		UsrIds:  usrIds,
		Message: data,
	}

	_, _, err := p.httpc.SendJson(reqObj, httpc.OptionsSt{
		Method:        "POST",
		Path:          "send",
		LogPrefix:     "Send: ",
		RetryCount:    1,
		RetryInterval: 3 * time.Second,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *St) GetConnectionCount() (int64, error) {
	repObj := ws.ConnectionCountRepSt{}

	_, _, err := p.httpc.SendRecvJson(nil, &repObj, nil, httpc.OptionsSt{
		Method:        "POST",
		Path:          "connection_count",
		LogPrefix:     "ConnectionCount: ",
		RetryCount:    1,
		RetryInterval: time.Second,
	})
	if err != nil {
		return 0, err
	}

	return repObj.Value, nil
}
