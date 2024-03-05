package mongodb

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client - for mongodb
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	// Define your authentication response structure here if needed
}

// NewClient -
func NewClient(host, port, database, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default MongoDB connection URL
		HostURL: "mongodb://localhost:27017",
	}

	if host != nil && port != nil {
		c.HostURL = fmt.Sprintf("mongodb://%s:%s", *host, *port)
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	// Perform any authentication or connection setup logic here

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	// Add any MongoDB-specific request logic here
	// This method is just a placeholder and may need modification based on MongoDB API requirements

	token := ""
	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
