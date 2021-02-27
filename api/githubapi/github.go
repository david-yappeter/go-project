package githubapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//GithubClient Github Client
type GithubClient struct {
	Host     string `json:"string"`
	Endpoint string `json:"endpoint"`
}

//InitClient Init Github Client
func InitClient(endpoint string) *GithubClient {
	return &GithubClient{
		Host:     os.Getenv("GITHUB_HOST_API"),
		Endpoint: endpoint,
	}
}

//Get Get Request With Github Client
func (c *GithubClient) Get(ctx context.Context, parameter map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodGet, c.Host+c.Endpoint, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query := request.URL.Query()

	for key, val := range parameter {
		query.Add(key, val)
	}

	request.URL.RawQuery = query.Encode()

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

//Post Post Request With Github Client
func (c *GithubClient) Post(ctx context.Context, parameter map[string]interface{}) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	parameterJSON, err := json.Marshal(parameter)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.Host+c.Endpoint, bytes.NewBuffer(parameterJSON))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}
