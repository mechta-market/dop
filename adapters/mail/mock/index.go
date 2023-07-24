package mock

import (
	"sync"

	"github.com/mechta-market/dop/adapters/logger"
	"github.com/mechta-market/dop/adapters/mail"
)

type St struct {
	lg      logger.Lite
	testing bool
	q       []*mail.SendReqSt
	mu      sync.Mutex
}

func New(
	lg logger.Lite,
	testing bool,
) *St {
	return &St{
		lg:      lg,
		testing: testing,
		q:       make([]*mail.SendReqSt, 0),
	}
}

func (m *St) Send(data *mail.SendReqSt) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.testing {
		m.lg.Infow("Mail", "data", data)
	}

	if len(m.q) > 100 {
		m.q = make([]*mail.SendReqSt, 0)
	}

	m.q = append(m.q, data)

	return true
}

func (m *St) PullAll() []*mail.SendReqSt {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := m.q

	m.q = make([]*mail.SendReqSt, 0)

	return q
}

func (m *St) Clean() {
	_ = m.PullAll()
}
