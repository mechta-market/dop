package mock

import (
	"encoding/base64"
	"encoding/json"

	"github.com/mechta-market/dop/adapters/logger"
)

type St struct {
	lg      logger.WarnAndError
	testing bool
}

func New(lg logger.WarnAndError, testing bool) *St {
	return &St{
		lg:      lg,
		testing: testing,
	}
}

func (p *St) Create(sub string, expSeconds int64, payload map[string]any) (string, error) {
	pld := make(map[string]any, len(payload)+1)

	for k, v := range payload {
		pld[k] = v
	}

	pld["sub"] = sub

	payloadRaw, err := json.Marshal(pld)
	if err != nil {
		p.lg.Errorw("Fail to marshal data", err)
		return "", err
	}

	payloadB64 := base64.RawURLEncoding.EncodeToString(payloadRaw)

	return "XXX." + payloadB64 + ".YYY", nil
}
