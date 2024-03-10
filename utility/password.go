package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(plaintextPassword, hashedPassword string) bool {
	hashedPlaintextPassword, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), hashedPlaintextPassword)
	return err == nil
}
