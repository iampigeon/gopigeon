package iampigeon

// ResponseMessages represents the information of every message
// by channel and id
type ResponseMessages struct {
	ID      string `json:"id"`
	Channel string `json:"channel"`
	Err     string `json:"error,omitempty"`
}

// TODO: find a smart way to do this.
// declare some kind of interface and a smart way to decode it

// Data represents the entire options that a responde of pigeon can contain
//type Data struct {
//Messages []*MessageResponse `json:"messages,omitempty"`
//}

//// Meta ...
//type Meta struct{}

// PigeonResponse ...
type PigeonResponse struct {
	Data map[string]interface{} `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
