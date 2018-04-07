package util

import (
	"encoding/json"
	"strings"

	"github.com/jmoiron/jsonq"
)

// Jsonsable represents an string utility to jsonQ
type Jsonsable interface {
	// ToJsonQ get JsonQuery from string
	ToJSONQ(body string) *jsonq.JsonQuery
}

// ToJSONQ return JsonQuery from String formatted
func ToJSONQ(body string) *jsonq.JsonQuery {
	data := map[string]interface{}{}
	bodyFormatted := strings.Replace(body, "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(bodyFormatted))
	dec.Decode(&data)
	return jsonq.NewQuery(data)
}
