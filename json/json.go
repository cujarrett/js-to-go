package json

// Config is what a user writes. The tags decide what the JSON and YAML look
// like, which is exactly how a Kubernetes CRD field gets its name.
type Config struct {
	// TODO: tag these so they serialise as sourceSecret, replicas, targets and
	// note. Replicas and note should disappear when empty; targets should not.
	SourceSecret string
	Replicas     int
	Targets      []string
	Note         string

	internal string // lowercase, so it is never serialised
}

// Encode returns c as JSON.
func Encode(c Config) ([]byte, error) {
	// TODO
	return nil, nil
}

// Decode parses JSON into a Config.
func Decode(data []byte) (Config, error) {
	// TODO
	return Config{}, nil
}
