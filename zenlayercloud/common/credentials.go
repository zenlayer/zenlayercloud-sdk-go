package common

// CredentialIface is the common interface for all credential types.
// Implementations: Credential (HMAC signing), TokenCredential (Bearer token).
type CredentialIface interface {
	GetSecretKeyId() string
	GetSecretKeyPassword() string
	GetToken() string
}

// Credential authenticates via HMAC-SHA256 request signing.
type Credential struct {
	SecretKeyId       string
	SecretKeyPassword string
}

func NewCredential(secretKeyId, secretKeyPassword string) *Credential {
	return &Credential{
		SecretKeyId:       secretKeyId,
		SecretKeyPassword: secretKeyPassword,
	}
}

func (c *Credential) GetSecretKeyId() string       { return c.SecretKeyId }
func (c *Credential) GetSecretKeyPassword() string { return c.SecretKeyPassword }
func (c *Credential) GetToken() string             { return "" }

// TokenCredential authenticates via a Bearer token (Authorization: Bearer <token>).
type TokenCredential struct {
	Token string
}

func NewTokenCredential(token string) *TokenCredential {
	return &TokenCredential{Token: token}
}

func (c *TokenCredential) GetSecretKeyId() string       { return "" }
func (c *TokenCredential) GetSecretKeyPassword() string { return "" }
func (c *TokenCredential) GetToken() string             { return c.Token }
