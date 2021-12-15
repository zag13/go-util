package ujson

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func JsonLoad(fname string, data interface{}) (err error) {
	fp, err := os.Open(fname)
	if err != nil {
		return
	}
	defer fp.Close()

	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return
	}

	return
}

func JsonDump(fname string, data interface{}) (n int, err error) {
	fp, err := os.Create(fname)
	if err != nil {
		return
	}
	defer fp.Close()

	v, err := json.Marshal(data)
	if err != nil {
		return
	}

	n, err = fp.Write(v)
	if err != nil {
		return
	}

	return
}
