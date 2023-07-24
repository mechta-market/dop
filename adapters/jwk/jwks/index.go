package jwks

import (
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mechta-market/dop/adapters/logger"
)

type St struct {
	lg logger.WarnAndError

	jwks *keyfunc.JWKS
}

func NewByUrl(lg logger.WarnAndError, url string, refreshInterval time.Duration) (*St, error) {
	jwks, err := keyfunc.Get(url, keyfunc.Options{
		RefreshInterval: refreshInterval,
		RefreshTimeout:  10 * time.Second,
		RefreshErrorHandler: func(err error) {
			lg.Errorw("Jwks refresh error", err)
		},
	})
	if err != nil {
		return nil, err
	}

	return &St{
		lg:   lg,
		jwks: jwks,
	}, nil
}

func (p *St) Validate(token string) (bool, error) {
	jwtToken, err := jwt.Parse(token, p.jwks.Keyfunc)
	if err != nil {
		return false, err
	}

	return jwtToken.Valid, nil
}
