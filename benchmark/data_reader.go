package benchmark

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)

func ReadData(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(gz)
	if err != nil {
		return nil, err
	}
	return data, nil
}
