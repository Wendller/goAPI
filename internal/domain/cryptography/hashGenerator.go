package cryptography

type HashGenerator interface {
	hash(plainText string) ([]byte, error)
}
