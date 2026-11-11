package json

import "encoding/json"

// Config is what a user writes. The tags decide what the JSON and YAML look
// like, which is exactly how a Kubernetes CRD field gets its name.
type Config struct {
	SourceSecret string   `json:"sourceSecret"`
	Replicas     int      `json:"replicas,omitempty"`
	Targets      []string `json:"targets"`
	Note         string   `json:"note,omitempty"`

	internal string // lowercase, so it is never serialised
}

// Encode returns c as JSON.
func Encode(c Config) ([]byte, error) {
	return json.Marshal(c)
}

// Decode parses JSON into a Config.
func Decode(data []byte) (Config, error) {
	var c Config
	err := json.Unmarshal(data, &c)
	return c, err
}
