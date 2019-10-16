package connection

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

// Connection .
type Connection struct {
	ConnectionID string `json:"connection_id"`
}

// GetConnection .
func GetConnection(ctx context.Context) (Connection, error) {
	span := trace.FromContext(ctx)
	defer span.End()

	conn := Connection{}
	client := &http.Client{
		Transport: &ochttp.Transport{},
	}

	jsonData, err := json.Marshal(map[string]string{"test": "test"})
	if err != nil {
		return conn, fmt.Errorf("marshal error: %s", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8087/connection", bytes.NewBuffer(jsonData))
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
