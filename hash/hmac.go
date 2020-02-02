package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// HmacMd5 computes a Hash-based Message Authentication Code (HMAC) by using
// the MD5 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacMd5(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(md5.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacMd5Hex(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacMd5Base64StdEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacMd5Base64RawStdEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacMd5Base64URLEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacMd5Base64RawURLEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// HmacSha1 computes a Hash-based Message Authentication Code (HMAC) by using
// the SHA-1 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacSha1(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacSha1Hex(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacSha1Base64StdEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacSha1Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacSha1Base64URLEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacSha1Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HmacSha256(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacSha256Hex(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacSha256Base64StdEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacSha256Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacSha256Base64URLEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacSha256Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HmacSha224(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New224, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacSha224Hex(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacSha224Base64StdEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacSha224Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacSha224Base64URLEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacSha224Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HmacSha512(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacSha512Hex(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacSha512Base64StdEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacSha512Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacSha512Base64URLEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacSha512Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HmacSha384(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New384, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacSha384Hex(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacSha384Base64StdEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacSha384Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacSha384Base64URLEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacSha384Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}
