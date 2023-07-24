package jwk

type Jwk interface {
	Validate(token string) (bool, error)
}
