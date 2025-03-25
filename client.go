package gofall

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// NewClient creates a new Client.
func NewClient(startingClient *http.Client) *Client {
	defaultMaxRetries := 5
	defaultMaxRequests := 5
	defaultWindow := time.Second
	defaultTimeoutSeconds := 30

	if startingClient == nil {
		startingClient = &http.Client{
			Timeout:       time.Second * time.Duration(defaultTimeoutSeconds),
			CheckRedirect: nil,
			Transport:     nil,
			Jar:           nil,
		}
	}

	transport := http.DefaultTransport
	if startingClient.Transport != nil {
		transport = startingClient.Transport
	}

	httpClient := &http.Client{
		Transport: &roundTripper{
			maxRetries: defaultMaxRetries,
			limiter:    newRateLimiter(defaultWindow, defaultMaxRequests),
			inner:      transport,
		},
		CheckRedirect: startingClient.CheckRedirect,
		Jar:           startingClient.Jar,
		Timeout:       startingClient.Timeout,
	}

	return &Client{
		Card:     &CardClient{client: httpClient},
		BulkData: &BulkDataClient{client: httpClient},
		Rulings:  &RulingClient{client: httpClient},
	}
}

// Client is the interface to Scryfall's API.
type Client struct {
	Card     *CardClient
	BulkData *BulkDataClient
	Rulings  *RulingClient
}

func doRequest(client *http.Client, req *http.Request, into interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		apiErr.Status = resp.StatusCode

		if err := decoder.Decode(&apiErr); err != nil {
			return fmt.Errorf("failed to decode respone: %w", err)
		}

		if err := matchAPIError(&apiErr); err != nil {
			return err
		}

		return &apiErr
	}

	if err := decoder.Decode(into); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
