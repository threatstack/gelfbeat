package beater

import (
	"bytes"
	"encoding/json"
	"strings"
)

// GELFMessage is the struct we extract GELF chunks into
type GELFMessage struct {
	id     string
	data   string
	number int
	count  int
}

// GELFDecodedMessage allows us to write a custom unmarshaler
type GELFDecodedMessage map[string]interface{}

// UnmarshalJSON lets us remap _ fields to not having a _
func (rawGELFmsg *GELFDecodedMessage) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, (*map[string]interface{})(rawGELFmsg)); err != nil {
		return err
	}

	rawGELFAccessor := *rawGELFmsg
	keys := Keys(*rawGELFmsg)

	for i := range keys {
		if strings.HasPrefix(keys[i], "_") {
			field := keys[i][1:]
			rawGELFAccessor[field] = rawGELFAccessor[keys[i]]
			delete(rawGELFAccessor, keys[i])
		}
	}

	return nil
}

// Keys returns a slice of key names
func Keys(m map[string]interface{}) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func msgBuildWorker(c chan []byte, ac chan GELFMessage, nn chan []byte, numer int) {
	var chunkByte = []byte{0xef}
	msg := <-nn
	result := msgCheck(msg)
	if bytes.Equal(result, chunkByte) {
		ac <- extractGELFChunk(msg)
	} else {
		c <- result
	}
}

func buildMessage(c chan []byte, ac chan GELFMessage) {
	buffer := map[string][]GELFMessage{}
	for {
		chunk := <-ac
		buffer[chunk.id] = append(buffer[chunk.id], chunk)

		if chunk.count == len(buffer[chunk.id]) {
			var msg string
			for _, item := range buffer[chunk.id] {
				msg += item.data
			}
			delete(buffer, chunk.id)
			c <- msgCheck([]byte(msg))
		} else {
			c <- []byte{}
		}
	}
}

func printer(c chan []byte, nn chan []byte) {
	for {
		msg := <-c
		nn <- msg
	}
}

func extractGELFChunk(chunked []byte) GELFMessage {
	var chunk = GELFMessage{}
	chunk.id = string(chunked[2:10])
	chunk.number = int(chunked[10:11][0])
	chunk.count = int(chunked[11:12][0])
	chunk.data = string(chunked[12:])
	return chunk
}
