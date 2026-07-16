package llm

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type SSEEvent struct {
	Event string
	Data  string
	ID    string
	Retry int
}

type SSEDecoder struct {
	reader *bufio.Reader
}

func NewSSEDecoder(r io.Reader) *SSEDecoder {
	return &SSEDecoder{
		reader: bufio.NewReader(r),
	}
}

func (d *SSEDecoder) Decode() (*SSEEvent, error) {
	var event SSEEvent

	for {
		line, err := d.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimRight(line, "\r\n")

		if line == "" {
			if event.Data != "" || event.Event != "" {
				return &event, nil
			}
			continue
		}

		if strings.HasPrefix(line, "event:") {
			event.Event = strings.TrimSpace(line[6:])
		} else if strings.HasPrefix(line, "data:") {
			data := strings.TrimSpace(line[5:])
			if event.Data != "" {
				event.Data += "\n" + data
			} else {
				event.Data = data
			}
		} else if strings.HasPrefix(line, "id:") {
			event.ID = strings.TrimSpace(line[3:])
		} else if strings.HasPrefix(line, "retry:") {
			fmt.Sscanf(line[6:], "%d", &event.Retry)
		}
	}
}
