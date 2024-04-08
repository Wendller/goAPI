package cryptography

type HashComparer interface {
	compare(hash, plainText string) bool
}
