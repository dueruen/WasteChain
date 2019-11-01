package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

func Compress(data string) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(data)); err != nil {
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

func Decompress(inputData string) ([]byte, error) {
	data, _ := base64.StdEncoding.DecodeString(inputData)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return s, nil
}
