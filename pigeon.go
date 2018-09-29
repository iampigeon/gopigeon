package iampigeon

import (
	"net/http"
)

// Pigeon ...
type Pigeon struct {
	APIKey   string
	client   *http.Client
	subjects []*subject
}

const host = "localhost:5050"

// NewPigeon returns a new pigeon instance. If cannot get the subject list
// will return a nil pointer and an error
func NewPigeon(apiKey string) (*Pigeon, error) {
	subjects, err := getSubjectsList(apiKey)
	if err != nil {
		return nil, err
	}

	return &Pigeon{
		APIKey: apiKey,
		client: http.DefaultClient,
	}
}

// TODO: get this types from pigeon

type subject struct {
	Name     string   `json:"name"`
	Channels []string `json:"channels"`
}

// Message ...
type Message struct {
	Subject string `json:"subject"`
	MQTT    *MQTT  `json:"mqtt,omitempty"`
	SMS     *SMS   `json:"sms,omitempty"`
	HTTP    *HTTP  `json:"http,omitempty"`
	PUSH    *Push  `json:"push,omitempty"`
}

// PigeonResponse ...
type PigeonResponse struct {
	Data map[string]interface{} `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

//Deliver ...
// the delay will be handled internally
func (p *Pigeon) Deliver(m Message) {
}

//get message by id
//cancel message by id
