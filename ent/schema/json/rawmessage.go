package rawmessage

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalRawMessage(t json.RawMessage) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		s, _ := t.MarshalJSON()
		_, _ = io.WriteString(w, string(s))
	})
}

func UnmarshalRawMessage(v interface{}) (json.RawMessage, error) {
	switch v := v.(type) {
	case json.RawMessage:
		return v, nil
	case string:
		return json.RawMessage(v), nil
	default:
		return nil, fmt.Errorf("unsupported type: %+v", v)
	}
}
