package event

import (
	"encoding/json"
	"time"
)

// Topc Type
type Topic string

// Supported Topic Names
var (
	VolumeIncreaseTopic  = Topic("volume:increase")
	VolumeDecreaseTopic  = Topic("volume:decrease")
	VolumeIncreasedTopic = Topic("volume:increased")
	VolumeDecreasedTopic = Topic("volume:decreased")
)

// SFM 2.0 Event JSON structure
type Event struct {
	Topic   Topic           `json:"topic"`             // Name of the Event
	Created time.Time       `json:"created"`           // Date the event was created
	Payload json.RawMessage `json:"payload,omitempty"` // Any other data
}

// All volume events will use this common payload
type VolumeLevelPayload struct {
	Level int `json:"level"`
}
