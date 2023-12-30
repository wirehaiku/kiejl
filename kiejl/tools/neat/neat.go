// Package neat implements string sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
)

// Body returns a clean file body string.
func Body(body string) string {
	body = strings.TrimSpace(body)
	return body + "\n"

}

// Extn returns a clean file extension string.
func Extn(extn string) string {
	extn = strings.TrimSpace(extn)
	extn = strings.ToLower(extn)
	return "." + strings.TrimPrefix(extn, ".")
}

// Name returns a clean file name string.
func Name(name string) string {
	name = strings.TrimSpace(name)
	return strings.ToLower(name)
}

// Path returns a clean file path string.
func Path(dire string) string {
	dire = strings.TrimSpace(dire)
	dire, _ = filepath.Abs(dire)
	return dire
}
