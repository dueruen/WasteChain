package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

func Compress(data []byte) (base64Output string, err error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		return "", err
	}
	if err := gz.Flush(); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	return str, nil
}

func Decompress(base64Input string) ([]byte, error) {
	data, _ := base64.StdEncoding.DecodeString(base64Input)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return s, nil
}
