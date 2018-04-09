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

// Format Json to readable string
func Format(body string) string {
	r := strings.NewReplacer("&#34;", "\"")
	return r.Replace(body)
}

// ToJSONQ return JsonQuery from String formatted
func ToJSONQ(body string) *jsonq.JsonQuery {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(Format(body)))
	dec.Decode(&data)
	return jsonq.NewQuery(data)
}
