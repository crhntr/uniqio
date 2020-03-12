package uniqio

import (
	"bufio"
	"io"
)

func Lines(w io.Writer, r io.Reader) error {
	var rd *bufio.Reader

	if bRd, ok := r.(*bufio.Reader); !ok {
		rd = bufio.NewReaderSize(r, 512)
	} else {
		rd = bRd
	}

	lines := make(map[string]struct{})

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}

		_, ok := lines[line]
		if ok {
			continue
		}

		if _, err := w.Write([]byte(line)); err != nil {
			return err
		}

		lines[line] = struct{}{}
	}
}
