package common

// CredentialIface is the common interface for all credential types.
// Implementations: Credential (HMAC signing), StsCredential (HMAC signing with a
// security token), TokenCredential (Bearer token).
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

// SecurityTokenCredential is implemented by credentials that carry a security
// token. The client sends the token in the x-zc-security-token header and
// includes that header in the signature.
type SecurityTokenCredential interface {
	CredentialIface
	GetSecurityToken() string
}

// StsCredential authenticates via HMAC-SHA256 request signing with a temporary
// access key issued by AssumeRole, together with its security token.
// Never log this structure.
type StsCredential struct {
	SecretKeyId       string
	SecretKeyPassword string
	SecurityToken     string
}

func NewStsCredential(secretKeyId, secretKeyPassword, securityToken string) *StsCredential {
	return &StsCredential{
		SecretKeyId:       secretKeyId,
		SecretKeyPassword: secretKeyPassword,
		SecurityToken:     securityToken,
	}
}

func (c *StsCredential) GetSecretKeyId() string       { return c.SecretKeyId }
func (c *StsCredential) GetSecretKeyPassword() string { return c.SecretKeyPassword }
func (c *StsCredential) GetToken() string             { return "" }
func (c *StsCredential) GetSecurityToken() string     { return c.SecurityToken }

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
