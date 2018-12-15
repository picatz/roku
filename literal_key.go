package roku

// LiteralKey formats a given (plaintext/string) key for Keypress
// that helps simulate the on-screen keyboard for roku input.
func LiteralKey(key string) string {
	return "Lit_" + key
}
