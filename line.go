package uniqio

import (
	"bufio"
	"io"
)

// Lines reads from r and only prints the line to w if
// it has not seen that line before. It uses '\n' as the delimiter
// so it may not work where a different line delimiter is expected.
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
			if err == io.EOF {
				return nil
			}
			return err
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
