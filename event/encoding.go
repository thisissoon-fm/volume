package event

import (
	"encoding/json"
	"io"
)

// Interfaces encoders and decoders must implement
type (
	Encoder interface {
		Encode(v interface{}) error
	}
	Decoder interface {
		Decode(v interface{}) error
	}
)

// Func types for consutructing new Encoders and Decoders
type (
	NewDecoderFn func(r io.Reader) Decoder
	NewEncoderFn func(w io.Writer) Encoder
)

// Default Decoder is JSON
var defaultNewDecoder = func(r io.Reader) Decoder {
	return json.NewDecoder(r)
}

// Set the global NewDecoder to use default new decoder
var NewDecoder NewDecoderFn = defaultNewDecoder

// Default Encoder is JSON
var defaultNewEncoder = func(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

// Set the global NewEncoder to use default new encoder
var NewEncoder NewEncoderFn = defaultNewEncoder
