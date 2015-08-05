// Copyright 2015 go-dockerclient authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package docker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Network represents a network.
type Network struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Type      string `json:"type"`
	Endpoints []struct {
		Name    string `json:"name"`
		ID      string `json:"id"`
		Network string `json:"network"`
	} `json:"endpoints"`
}

// ListNetworks returns all networks.
func (c *Client) ListNetworks() ([]Network, error) {
	path := "/networks"
	body, _, err := c.do("GET", path, doOptions{})
	if err != nil {
		return nil, err
	}
	var networks []Network
	if err := json.Unmarshal(body, &networks); err != nil {
		return nil, err
	}
	return networks, nil
}

// NetworkInfo returns information about a network by its ID.
func (c *Client) NetworkInfo(id string) (*Network, error) {
	path := "/networks/" + id
	body, status, err := c.do("GET", path, doOptions{})
	if status == http.StatusNotFound {
		return nil, &NoSuchNetwork{ID: path}
	}
	if err != nil {
		return nil, err
	}
	var network Network
	if err := json.Unmarshal(body, &network); err != nil {
		return nil, err
	}
	return &network, nil
}

// NoSuchNetwork is the error returned when a given network does not exist.
type NoSuchNetwork struct {
	ID string
}

func (err *NoSuchNetwork) Error() string {
	return fmt.Sprintf("No such network: %s", err.ID)
}
