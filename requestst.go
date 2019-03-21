package iampigeon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	subjectsEndpoint = "/api/v1/subjects"
)

func makeRequest(endpoint, apiKey, pigeonAPI string) (*http.Request, error) {
	u, err := url.Parse(fmt.Sprintf("%s%s", pigeonAPI, endpoint))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", apiKey)

	return req, nil
}

func getSubjectsList(apiKey, host string) ([]*Subject, error) {
	subjects := make([]*Subject, 0)

	req, err := makeRequest(subjectsEndpoint, apiKey, host)
	if err != nil {
		return subjects, err
	}

	cl := http.DefaultClient

	res, err := cl.Do(req)
	if err != nil {
		return subjects, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return subjects, err
	}
	defer res.Body.Close()

	pr := new(PigeonResponse)

	err = json.Unmarshal(body, pr)
	if err != nil {
		return subjects, err
	}

	rs := new(struct{ Subjects []*Subject })

	err = Decode(pr.Data, rs)
	if err != nil {
		return subjects, err
	}

	return rs.Subjects, nil
}
