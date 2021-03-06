package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// Sha384 returns SHA-384 checksum of a text as bytes.
func Sha384(text string) ([]byte, error) {
	hash := sha512.New384()
	return hashText(hash, text)
}

// Sha384Hex returns the SHA-384 checksum of a text in
// hexadecimal encoding format.
func Sha384Hex(text string) (string, error) {
	hash, err := Sha384(text)
	return hex.EncodeToString(hash), err
}

// Sha384Base64StdEnc returns the SHA-384 checksum of a text in
// standard base64 encoding, as defined in RFC 4648.
func Sha384Base64StdEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha384Base64URLEnc returns the SHA-384 checksum of a text in
// an alternate base64 encoding defined in RFC 4648.
func Sha384Base64URLEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Sha384Base64RawURLEnc returns the SHA-384 checksum of a text in
// a padded alternate base64 encoding defined in RFC 4648.
func Sha384Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Sha384Base64RawStdEnc returns the SHA-384 checksum of a text in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Sha384Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Sha384File returns SHA-384 checksum of a file as bytes.
func Sha384File(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashFile(hash, path)
}

// Sha384FileHex returns the SHA-384 checksum of a file in
// hexadecimal encoding format.
func Sha384FileHex(path string) (string, error) {
	hash, err := Sha384File(path)
	return hex.EncodeToString(hash), err
}

// Sha384FileBase64StdEnc returns the SHA-384 checksum of a file in
// standard base64 encoding, as defined in RFC 4648.
func Sha384FileBase64StdEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha384FileBase64URLEnc returns the SHA-384 checksum of a file in
// an alternate base64 encoding defined in RFC 4648.
func Sha384FileBase64URLEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Sha384FileBase64RawURLEnc returns the SHA-384 checksum of a file in
// a padded alternate base64 encoding defined in RFC 4648.
func Sha384FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Sha384FileBase64RawStdEnc returns the SHA-384 checksum of a file in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Sha384FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Sha384Dir returns SHA-384 checksum of a directory as bytes.
func Sha384Dir(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashDir(hash, path)
}

// Sha384DirHex returns the SHA-384 checksum of a directory in
// hexadecimal encoding format.
func Sha384DirHex(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return hex.EncodeToString(hash), err
}

// Sha384DirBase64StdEnc returns the SHA-384 checksum of a directory in
// standard base64 encoding, as defined in RFC 4648.
func Sha384DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha384DirBase64URLEnc returns the SHA-384 checksum of a directory in
// an alternate base64 encoding defined in RFC 4648.
func Sha384DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Sha384DirBase64RawURLEnc returns the SHA-384 checksum of a directory in
// a padded alternate base64 encoding defined in RFC 4648.
func Sha384DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// SHA384DirBase64RawStdEnc returns the SHA-384 checksum of a directory in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func SHA384DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Sha384Path returns SHA-384 checksum of a path as bytes.
func Sha384Path(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashPath(hash, path)
}

// Sha384PathHex returns the SHA-384 checksum of a path in
// hexadecimal encoding format.
func Sha384PathHex(path string) (string, error) {
	hash, err := Sha384Path(path)
	return hex.EncodeToString(hash), err
}

// Sha384PathBase64StdEnc returns the SHA-384 checksum of a path in
// standard base64 encoding, as defined in RFC 4648.
func Sha384PathBase64StdEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return hex.EncodeToString(hash), err
}

// Sha384PathBase64URLEnc returns the SHA-384 checksum of a path in
// an alternate base64 encoding defined in RFC 4648.
func Sha384PathBase64URLEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Sha384PathBase64RawURLEnc returns the SHA-384 checksum of a path in
// a padded alternate base64 encoding defined in RFC 4648.
func Sha384PathBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Sha384PathBase64RawStdEnc returns the SHA-384 checksum of a path in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Sha384PathBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
