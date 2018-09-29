package iampigeon

import (
	"net/url"
)

// MQTT ...
type MQTT struct {
	Topic   string      `json:"mqtt_topic"`
	Payload interface{} `json:"mqtt_payload,omitempty"`
}

// SMS ...
type SMS struct {
	Phones []string `json:"phones"`
	Text   string   `json:"text"`
}

// HTTP ...
type HTTP struct {
	URL     *url.URL               `json:"url"`
	Body    map[string]interface{} `json:"body,omitempty"`
	Headers map[string]string      `json:"headers,omitempty"`
}

// Push ...
type Push struct {
	DeviceID string `json:"device_id"`
}
