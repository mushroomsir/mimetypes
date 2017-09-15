package mimetypes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMimetype(t *testing.T) {
	assert := assert.New(t)
	dat := make(map[string]MimeValue)
	if err := json.Unmarshal([]byte(mimetypemap), &dat); err != nil {
		panic(err)
	}
	assert.Equal(838, len(dat))
	for k, v := range dat {
		if len(v.Extensions) == 0 {
			delete(dat, k)
		}
	}
	assert.Equal(838, len(dat))
	json.Marshal(dat)
	assert.Equal(1065, len(extensions))
}

func TestLookup(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		result string
		val    string
	}{
		{"text/html", ".html"},
		{"application/javascript", ".js"},
		{"application/json", ".json"},
	}
	for _, c := range cases {
		if !assert.Equal(c.result, Lookup(c.val)) {
			fmt.Println(c.val, c.result)
		}
	}
}

func TestExtension(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		result string
		val    string
	}{
		{"html", "text/html"},
		{"html", "text/html "},
		{"html", " text/html"},
		{"html", "text/html;charset=UTF-8"},
		{"html", "text/HTML; charset=UTF-8"},
		{"html", "text/html; charset=UTF-8"},
		{"html", "text/html; charset=UTF-8 "},
		{"html", "text/html ; charset=UTF-8"},
		{"", "application/x-bogus"},
		{"", ""},
		{"", "bogus"},
		{"", "{}"},
	}
	for _, c := range cases {
		if !assert.Equal(c.result, Extension(c.val)) {
			fmt.Println(c.val, c.result)
		}
	}
}
