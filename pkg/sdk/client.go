package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseUrl string
}

type Forum struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	TopicKeyword string  `json:"topickeyword"`
	Users        []uint8 `json:"users"`
}

func (f Forum) String() string {
	return f.Name
}

func (c *Client) ListForums(ctx context.Context) ([]Forum, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseUrl+"/forums", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	var res []Forum
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) CreateUser(ctx context.Context, name string) error {
	var p Forum
	p.Name = name
	bb := new(bytes.Buffer)
	err := json.NewEncoder(bb).Encode(&p)
	if err != nil {
		panic(err)
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseUrl+"/forums", bb)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	return nil
}
