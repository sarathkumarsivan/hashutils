package hash

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMD5Hash(t *testing.T) {
	hash, err := Md5Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", hash)

	hash, err = Md5Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "rL0Y20zC+Fzt72VPzMSk2A==", hash)

	hash, err = Md5Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "rL0Y20zC+Fzt72VPzMSk2A", hash)

	hash, err = Md5Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "rL0Y20zC-Fzt72VPzMSk2A", hash)

	hash, err = Md5Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "rL0Y20zC-Fzt72VPzMSk2A==", hash)
}

func TestMD5HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := Md5FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", hash)

	hash, err = Md5FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg==", hash)

	hash, err = Md5FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg==", hash)

	hash, err = Md5FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg", hash)

	hash, err = Md5FileBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg", hash)
}

func TestMD5HashDir(t *testing.T) {
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

	hash, err := Md5DirHex(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5DirBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5DirBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5DirBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5DirBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", Md5Hash)
	assert.NotEmpty(t, hash)
}

func TestMD5HashPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	hash, err := Md5PathHex(dir)
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)

	hash, err = Md5PathBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Md5Hash)
	assert.NotEmpty(t, hash)
}
