package hash

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA384Hash(t *testing.T) {
	hash, err := Sha384Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.Equal(t, "98c11ffdfdd540676b1a137cb1a22b2a70350c9a44171d6b1180c6be5cbb2ee3f79d532c8a1dd9ef2e8e08e752a3babb", hash)

	hash, err = Sha384Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Sha512Hash)
	assert.Equal(t, "mMEf/f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)

	hash, err = Sha384Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Sha512Hash)
	assert.Equal(t, "mMEf/f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)

	hash, err = Sha384Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Sha512Hash)
	assert.Equal(t, "mMEf_f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)

	hash, err = Sha384Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Sha512Hash)
	assert.Equal(t, "mMEf_f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)
}

func TestSHA384HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := Sha384FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.Equal(t, "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", hash)

	hash, err = Sha384FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.Equal(t, "OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLlb", hash)

	hash, err = Sha384FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.Equal(t, "OLBgp1GsljhM2TJ-sbHjaiH9txEUvgdDTAzHv2P24donTt6_529l-9Ua0vFImLlb", hash)

	hash, err = Sha384FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.Equal(t, "OLBgp1GsljhM2TJ-sbHjaiH9txEUvgdDTAzHv2P24donTt6_529l-9Ua0vFImLlb", hash)

	hash, err = Sha384FileBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.Equal(t, "OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLlb", hash)
}

func TestSHA384HashDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	bar, err := ioutil.TempFile(dir, "bar.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = bar.WriteString("bar")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(bar.Name())

	hash, err := Sha384DirHex(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384DirBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384DirBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384DirBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA384DirBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)
}

func TestSHA384HashPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	hash, err := Sha384PathHex(dir)
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha384PathBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Sha384Hash)
	assert.NotEmpty(t, hash)
}
