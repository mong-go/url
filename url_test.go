package url

import (
	"github.com/nowk/assert"
	"testing"
)

func TestParseBasicMongoURL(t *testing.T) {
	u, err := Parse("mongodb://user:pass@127.0.0.1:30001/databasename")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "mongodb", u.Scheme)
	assert.Equal(t, "127.0.0.1", u.Host())
	assert.Equal(t, "30001", u.Port())
	assert.Equal(t, "databasename", u.Database())

	ui := u.User
	assert.Equal(t, "user", ui.Username())

	p, set := ui.Password()
	assert.Equal(t, true, set)
	assert.Equal(t, "pass", p)
}

func TestDefaultPort(t *testing.T) {
	u, err := Parse("mongodb://127.0.0.1/databasename")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "27017", u.Port())
}

func TestShortStringCompilesURLWithoutDatabase(t *testing.T) {
	{
		u, err := Parse("mongodb://user:pass@127.0.0.1:30001/databasename")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "mongodb://user:pass@127.0.0.1:30001/databasename", u.String())
		assert.Equal(t, "mongodb://user:pass@127.0.0.1:30001", u.ShortString())
	}

	{
		u, err := Parse("mongodb://127.0.0.1:30001/databasename")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "mongodb://127.0.0.1:30001/databasename", u.String())
		assert.Equal(t, "mongodb://127.0.0.1:30001", u.ShortString())
	}
}
