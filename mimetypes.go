package mimetypes

import (
	"encoding/json"
	"path"
	"regexp"
	"strings"
)

var (
	types         = make(map[string]MimeValue)
	extensions    = make(map[string]string)
	rxExtractType = regexp.MustCompile(`^\s*([^;\s]*)(?:;|\s|$)`)
)

func init() {
	if err := json.Unmarshal([]byte(mimetypemap), &types); err != nil {
		panic(err)
	}
	for k, v := range types {
		for _, ext := range v.Extensions {
			extensions[ext] = k
		}
	}
}

// MimeValue ...
type MimeValue struct {
	Source       string   `json:"source,omitempty"`
	Compressible bool     `json:"compressible,omitempty"`
	Extensions   []string `json:"extensions"`
	Charset      string   `json:"charset,omitempty"`
}

// Lookup lookup the MIME type for a file path/extension.
func Lookup(name string) string {
	ext := strings.Replace(path.Ext(name), ".", "", 1)
	return extensions[strings.ToLower(ext)]
}

// Extension lookup the default extension for a MIME type.
func Extension(mimetype string) string {
	match := rxExtractType.FindAllStringSubmatch(mimetype, -1)
	m := types[strings.Trim(strings.ToLower(match[0][1]), " ")]
	if len(m.Extensions) == 0 {
		return ""
	}
	return m.Extensions[0]
}
