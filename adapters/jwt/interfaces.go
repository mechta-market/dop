package jwt

type Jwt interface {
	Create(sub string, expSeconds int64, payload map[string]any) (string, error)
}
