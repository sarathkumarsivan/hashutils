package hash

import (
	"errors"
)

var ErrUnsupportedAlgorithm = errors.New("hashutils: unsupported hashing algorithm")

// Algorithm represents the type of hashing algorithms supported by
// this package.
type Algorithm string

const (
	Md5Hash    Algorithm = "md5"
	Fnv32Hash            = "fnv32"
	Fnv32aHash           = "fnv32a"
	Fnv64Hash            = "fnv64"
	Fnv64aHash           = "fnv64a"
	Sha1Hash             = "sha1"
	Sha256Hash           = "sha256"
	Sha224Hash           = "sha224"
	Sha512Hash           = "sha512"
	Sha384Hash           = "sha384"
	Crc32Hash            = "crc32"
)

// An Encoding is a radix 64 encoding/decoding scheme, defined by a
// 64-character alphabet.
type Encoding string

const (
	Hex    Encoding = "hex"
	Base64          = "base64"
)

type ExtHash interface {
	HashText(text string) (string, error)
	HashFile(path string) (string, error)
	HashFiles(paths ...string) (map[string]string, error)
	// HashDir(path string) (string, error)
	HashPath(path string) (string, error)
}

type ExtHashBuilder interface {
	Algorithm(Algorithm) ExtHashBuilder
	Encoding(Encoding) ExtHashBuilder
	Build() ExtHash
}

type hashBuilder struct {
	algorithm Algorithm
	encoding  Encoding
}

func (h *hashBuilder) Algorithm(algorithm Algorithm) ExtHashBuilder {
	h.algorithm = algorithm
	return h
}

func (h *hashBuilder) Encoding(encoding Encoding) ExtHashBuilder {
	h.encoding = encoding
	return h
}

func (h *hashBuilder) Build() ExtHash {
	return &hashMaker{
		algorithm: h.algorithm,
		encoding:  h.encoding,
	}
}

func New() ExtHashBuilder {
	return &hashBuilder{}
}

type hashMaker struct {
	algorithm Algorithm
	encoding  Encoding
}

func (m *hashMaker) HashText(text string) (string, error) {
	if m.algorithm == Md5Hash && m.encoding == Hex {
		return Md5Hex(text)
	}
	if m.algorithm == Sha1Hash && m.encoding == Hex {
		return Sha1Hex(text)
	}
	if m.algorithm == Sha224Hash && m.encoding == Hex {
		return Sha224Hex(text)
	}
	if m.algorithm == Sha256Hash && m.encoding == Hex {
		return Sha256Hex(text)
	}
	if m.algorithm == Sha384Hash && m.encoding == Hex {
		return Sha384Hex(text)
	}
	if m.algorithm == Sha512Hash && m.encoding == Hex {
		return Sha512Hex(text)
	}

	if m.algorithm == Md5Hash && m.encoding == Base64 {
		return Md5Base64StdEnc(text)
	}
	if m.algorithm == Sha1Hash && m.encoding == Base64 {
		return Sha1Base64StdEnc(text)
	}
	if m.algorithm == Sha224Hash && m.encoding == Base64 {
		return Sha224Base64StdEnc(text)
	}
	if m.algorithm == Sha256Hash && m.encoding == Base64 {
		return Sha256Base64StdEnc(text)
	}
	if m.algorithm == Sha384Hash && m.encoding == Base64 {
		return Sha384Base64StdEnc(text)
	}
	if m.algorithm == Sha512Hash && m.encoding == Base64 {
		return Sha512Base64StdEnc(text)
	}

	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashFile(path string) (string, error) {
	if m.algorithm == Md5Hash && m.encoding == Hex {
		return Md5FileHex(path)
	}
	if m.algorithm == Sha1Hash && m.encoding == Hex {
		return Sha1FileHex(path)
	}
	if m.algorithm == Sha224Hash && m.encoding == Hex {
		return Sha224FileHex(path)
	}
	if m.algorithm == Sha256Hash && m.encoding == Hex {
		return Sha256FileHex(path)
	}
	if m.algorithm == Sha384Hash && m.encoding == Hex {
		return Sha384FileHex(path)
	}
	if m.algorithm == Sha512Hash && m.encoding == Hex {
		return Sha512FileHex(path)
	}

	if m.algorithm == Md5Hash && m.encoding == Base64 {
		return Md5FileBase64StdEnc(path)
	}
	if m.algorithm == Sha1Hash && m.encoding == Base64 {
		return Sha1FileBase64StdEnc(path)
	}
	if m.algorithm == Sha224Hash && m.encoding == Base64 {
		return Sha224FileBase64StdEnc(path)
	}
	if m.algorithm == Sha256Hash && m.encoding == Base64 {
		return Sha256FileBase64StdEnc(path)
	}
	if m.algorithm == Sha384Hash && m.encoding == Base64 {
		return Sha384FileBase64StdEnc(path)
	}
	if m.algorithm == Sha512Hash && m.encoding == Base64 {
		return Sha512FileBase64StdEnc(path)
	}

	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashFiles(paths ...string) (map[string]string, error) {
	pathHashes := make(map[string]string, len(paths))
	for _, path := range paths {
		if m.algorithm == Md5Hash && m.encoding == Hex {
			hash, err := Md5FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha1Hash && m.encoding == Hex {
			hash, err := Sha1FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha224Hash && m.encoding == Hex {
			hash, err := Sha224FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha256Hash && m.encoding == Hex {
			hash, err := Sha256FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha384Hash && m.encoding == Hex {
			hash, err := Sha384FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha512Hash && m.encoding == Hex {
			hash, err := Sha512FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}

		if m.algorithm == Md5Hash && m.encoding == Base64 {
			hash, err := Md5FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha1Hash && m.encoding == Base64 {
			hash, err := Sha1FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha224Hash && m.encoding == Base64 {
			hash, err := Sha224FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha256Hash && m.encoding == Base64 {
			hash, err := Sha256FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha384Hash && m.encoding == Base64 {
			hash, err := Sha384FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == Sha512Hash && m.encoding == Base64 {
			hash, err := Sha512FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
	}
	return pathHashes, nil
}

func (m *hashMaker) HashDir(path string) (string, error) {
	if m.algorithm == Md5Hash && m.encoding == Hex {
		return Md5DirHex(path)
	}
	if m.algorithm == Sha1Hash && m.encoding == Hex {
		return Sha1DirHex(path)
	}
	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashPath(path string) (string, error) {
	return "", ErrUnsupportedAlgorithm
}
