package iampigeon

import (
	"net/url"
)

// MQTT ...
type MQTT struct {
	Payload map[string]interface{} `json:"mqtt_payload,omitempty"`
}

// SMS ...
type SMS struct {
	Phone string `json:"phone"`
	Text  string `json:"text"`
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

// Channels represents the channels key
type Channels struct {
	MQTT *MQTT `json:"mqtt,omitempty"`
	SMS  *SMS  `json:"sms,omitempty"`
	HTTP *HTTP `json:"http,omitempty"`
	Push *Push `json:"push,omitempty"`
}
