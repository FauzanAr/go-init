package helper

import "github.com/alexedwards/argon2id"

func Hash(text string) (string, error) {
	hash, err := argon2id.CreateHash(text, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func Compare(text string, hash string) (bool, error) {
	res, err := argon2id.ComparePasswordAndHash(text, hash)
	if err != nil {
		return false, err
	}

	return res, nil
}
