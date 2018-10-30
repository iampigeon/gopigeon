package iampigeon

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Pigeon ...
type Pigeon struct {
	client   *Client
	Subjects []*Subject
}

// NewPigeon returns a new pigeon instance. If cannot get the Subject list
// will return a nil pointer and an error
func NewPigeon(apiKey, host string) (*Pigeon, error) {
	subjects, err := getSubjectsList(apiKey, host)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Client: http.DefaultClient,
		Host:   host,
		APIKey: apiKey,
	}

	return &Pigeon{
		client:   client,
		Subjects: subjects,
	}, nil
}

// TODO: get this types from pigeon

// Subject ...
type Subject struct {
	Name     string   `json:"name"`
	Channels []string `json:"channels"`
}

// Message ...
type Message struct {
	SubjectName string    `json:"subject_name"`
	Channels    *Channels `json:"channels"`
}

// MessageResponse ...
type MessageResponse struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	SubjectID string `json:"subject_id"`
}

//Deliver ...
// the delay will be handled internally
func (p *Pigeon) Deliver(m *Message) (*PigeonResponse, error) {
	endpoint := "/api/v1/messages"
	res, err := p.client.Request(http.MethodPost, endpoint, m)
	if err != nil {
		return nil, err
	}

	pr := new(PigeonResponse)

	err = UnmarshalJSON(res.Body, pr)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

// FetchStatusByID ...
func (p *Pigeon) FetchStatusByID(id string) (*PigeonResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/messages/%s/status", id)
	res, err := p.client.Request(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	pr := new(PigeonResponse)

	err = UnmarshalJSON(res.Body, pr)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

// CancelMessage ...
func (p *Pigeon) CancelMessage(id string) (*PigeonResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/messages/%s/cancel", id)
	res, err := p.client.Request(http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}

	pr := new(PigeonResponse)

	err = UnmarshalJSON(res.Body, pr)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

// UnmarshalJSON returns an error if cannot unmarshal the json
// allocated on r
func UnmarshalJSON(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

//get message by id
//cancel message by id
