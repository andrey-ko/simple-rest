package build

import (
	"bytes"
	"io"
	"os"
	"strings"
)

func ParseStdin() (map[string]string, error) {

	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, os.Stdin); err != nil {
		return nil, err
	}
	res := map[string]string{}
	kvs := strings.Split(buf.String(), ";")
	for _, kv := range kvs {
		if i := strings.Index(kv, "="); i < 0 {
			res[strings.TrimSpace(kv)] = ""
		} else {
			res[strings.TrimSpace(kv[:i])] = strings.TrimSpace(kv[i+1:])
		}
	}
	return res, nil
}
