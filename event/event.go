package event

import (
	"encoding/json"
	"time"
)

// Topc Type
type Topic string

// Supported Topic Names
var (
	VolumeUpdateTopic  = Topic("volume:update")
	VolumeUpdatedTopic = Topic("volume:updated")
	VolumeMuteTopic    = Topic("volume:mute")
	VolumeMutedTopic   = Topic("volume:muted")
	VolumeUnmuteTopic  = Topic("volume:unmute")
	VolumeUnmutedTopic = Topic("volume:unmuted")
)

// SFM 2.0 Event JSON structure
type Event struct {
	Topic   Topic           `json:"topic"`             // Name of the Event
	Created time.Time       `json:"created"`           // Date the event was created
	Payload json.RawMessage `json:"payload,omitempty"` // Any other data
}

// All volume events will use this common payload
type VolumeLevelPayload struct {
	Level uint8 `json:"level"`
}
