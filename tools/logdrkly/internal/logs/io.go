package logs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func loadCodes(path string) (codes LdLogCodesJson, err error) {
	f, err := os.Open(path)
	defer func(f *os.File) {
		err = f.Close()
	}(f)

	if err != nil {
		err = fmt.Errorf("could not open \"%s\"", path)
		return
	}

	readJson, err := io.ReadAll(f)

	if err != nil {
		err = fmt.Errorf("could not load \"%s\"", path)
		return
	}

	err = json.Unmarshal(readJson, &codes)
	if err != nil {
		err = fmt.Errorf("could not unmarshal \"%s\"", path)
		return
	}
	return codes, err
}

func writeLogCodes(path string, codes *LdLogCodesJson) (err error) {
	bytes, err := json.MarshalIndent(codes, "", "  ")

	if err != nil {
		return
	}

	of, err := os.Create(path)
	defer func(of *os.File) {
		err = of.Close()
	}(of)

	_, err = of.Write(bytes)

	return err
}

func UpdateCodes(path string, fn func(codes *LdLogCodesJson) error) error {
	codes, err := loadCodes(path)
	if err != nil {
		return err
	}
	err = fn(&codes)
	if err != nil {
		return err
	}
	err = writeLogCodes(path, &codes)
	return err
}

func WithCodes(path string, fn func(codes *LdLogCodesJson) error) error {
	codes, err := loadCodes(path)
	if err != nil {
		return err
	}
	return fn(&codes)
}
