package cryptography

type Encrypter interface {
	Encrypt(payload map[string]interface{}) string
}
