package csv

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

type Csv [][]string

func (csv Csv) Find(row, col int) (string, error) {
	if len(csv) <= row {
		return "", errors.New("error row")
	}
	if len(csv[row]) <= col {
		return "", errors.New("error col")
	}
	return strings.TrimSpace(csv[row][col]), nil
}

func (csv Csv) String(row, col int) (string, error) {
	return csv.Find(row, col)
}

func (csv Csv) Float(row, col int) (float64, error) {
	str, err := csv.Find(row, col)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(str, 64)
}

func (csv Csv) Int(row, col int) (int, error) {
	str, err := csv.Find(row, col)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str)
}

func New(filename, charset string) (Csv, error) {
	csv := make(Csv, 0)
	fs, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	br := bufio.NewReader(fs)
	for {

		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		b, err := toUTF8(charset, string(a))
		if err != nil {
			panic(err)
		}
		line := strings.Split(b, ",")
		csv = append(csv, line)
	}
	return csv, err
}

func toUTF8(srcCharset string, src string) (dst string, err error) {
	if srcCharset == "UTF-8" {
		return src, nil
	}
	if e := getEncoding(srcCharset); e != nil {
		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
		)
		if err != nil {
			return "", fmt.Errorf("%s to utf8 failed. %v", srcCharset, err)
		}
		return string(tmp), nil
	}
	return dst, fmt.Errorf("unsupport srcCharset: %s", srcCharset)

}

func getEncoding(charset string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(charset); err == nil && e != nil {
		return e
	}
	return nil
}
