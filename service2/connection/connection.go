package connection

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

// Connection .
type Connection struct {
	ConnectionID string `json:"connection_id"`
}

// GetConnection .
func GetConnection(ctx context.Context) (Connection, error) {

	///////////////////////////////// Trace /////////////////////////////////////////////
	client := &http.Client{
		Transport: &ochttp.Transport{},
	}
	////////////////////////////////////////////////////////////////////////////////////

	conn := Connection{}

	req, err := http.NewRequest("GET", "http://localhost:8087/connection", nil)
	if err != nil {
		return conn, err
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return conn, err
	}
	defer resp.Body.Close() //nolint: errcheck

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return conn, err
	}

	if resp.StatusCode != http.StatusOK {
		return conn, fmt.Errorf("checker error: %s", string(body))
	}

	err = json.Unmarshal(body, &conn)
	if err != nil {
		return conn, err
	}

	return conn, nil
}
