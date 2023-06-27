package util

import (
	"bufio"
	"io"
	"os"
)

func ReadFromFile(path string) (*string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			logger.Panic("file `%v` closing failed", path)
		}
	}()

	rd := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	var text string

	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		text += string(buf[:n])
	}

	return &text, nil
}

// WriteToFile Overwrite text into path
func WriteToFile(path, text string) error {
	fo, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err := fo.Close()
		if err != nil {
			logger.Panic("file `%v` closing failed", path)
		}
	}()

	if _, err := fo.Write([]byte(text)); err != nil {
		return err
	}

	return nil
}
