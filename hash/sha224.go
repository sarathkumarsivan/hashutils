package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Sha1 returns SHA-224 checksum of a text as bytes.
func Sha224(text string) ([]byte, error) {
	hash := sha256.New224()
	return hashText(hash, text)
}

// Sha224Hex returns the SHA-224 checksum of a text in
// hexadecimal encoding format.
func Sha224Hex(text string) (string, error) {
	hash, err := Sha224(text)
	return hex.EncodeToString(hash), err
}

func Sha224Base64StdEnc(text string) (string, error) {
	hash, err := Sha224(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha224Base64URLEnc(text string) (string, error) {
	hash, err := Sha224(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha224Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha224(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha224Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha224(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha224File(path string) ([]byte, error) {
	hash := sha256.New224()
	return hashFile(hash, path)
}

func Sha224FileHex(path string) (string, error) {
	hash, err := Sha224File(path)
	return hex.EncodeToString(hash), err
}

func Sha224FileBase64StdEnc(path string) (string, error) {
	hash, err := Sha224File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha224FileBase64URLEnc(path string) (string, error) {
	hash, err := Sha224File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha224FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha224File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha224FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha224File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha224Dir(path string) ([]byte, error) {
	hash := sha256.New224()
	return hashDir(hash, path)
}

func Sha224DirHex(path string) (string, error) {
	hash, err := Sha224Dir(path)
	return hex.EncodeToString(hash), err
}

func Sha224DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha224Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha224DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha224Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha224DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha224Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha224DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha224Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha224Path(path string) ([]byte, error) {
	hash := sha256.New224()
	return hashPath(hash, path)
}

func Sha224PathHex(path string) (string, error) {
	hash, err := Sha224Path(path)
	return hex.EncodeToString(hash), err
}

func Sha224PathBase64StdEnc(path string) (string, error) {
	hash, err := Sha224Path(path)
	return hex.EncodeToString(hash), err
}

func Sha224PathBase64URLEnc(path string) (string, error) {
	hash, err := Sha224Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha224PathBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha224Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha224PathBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha224Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
