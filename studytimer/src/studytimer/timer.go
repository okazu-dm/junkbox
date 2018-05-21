package studytimer

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func ReadTime(filename string) (int64, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return -1, err
	}
	t, err := strconv.ParseInt(strings.TrimRight(string(content), "\n"), 10, 64)
	if err != nil {
		return -1, err
	}
	return t, nil
}

func WriteTime(filename string, t int64) error {
	content := strconv.FormatInt(t, 10)
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
