package gobase64

import "encoding/base64"

//goland:noinspection GoUnusedExportedFunction
func Decode(s string) string {
	last := s[len(s)-1:]

	var r []byte
	if "=" == last {
		r, _ = base64.StdEncoding.DecodeString(s)
	} else {
		r, _ = base64.RawURLEncoding.DecodeString(s)
	}

	return string(r)
}
