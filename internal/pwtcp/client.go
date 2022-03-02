package pwtcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"jaytaylor.com/html2text"
	"net"
	"strings"
)

const serpStackApiKey = "1329e634cda359ccab9611c2f8929f20"

type Client interface {
	Get(string) (string, error)
	SearchGoogle(string) (string, error)
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) Get(url string) (string, error) {
	body, err := c.getBody(url, "")
	if err != nil {
		return "", err
	}

	cleanBody, err := html2text.FromString(body, html2text.Options{TextOnly: true, PrettyTables: true})
	if err != nil {
		return "", err
	}

	return cleanBody, nil
}

func (c *client) SearchGoogle(query string) (string, error) {
	body, err := c.getBody("api.serpstack.com", fmt.Sprintf("search?access_key=%s&query=%s", serpStackApiKey, query))
	if err != nil {
		return "", err
	}

	var response serpStackStruct

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", err
	}

	if !response.Request.Success {
		return "", errors.New("request unsuccessful")
	}

	var results string

	for _, entry := range response.OrganicResults {
		results += fmt.Sprintf("%d: %s\nURL: %s\n-------------------------------------------------------------\n\n", entry.Position, entry.Title, entry.URL)
	}

	return results, nil
}

func (c *client) getBody(domain, url string) (string, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:80", domain))
	if err != nil {
		return "", err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	request := []byte(
		fmt.Sprintf(
			"GET /%s HTTP/1.0\n"+
				"Host: %s\n"+
				"\n", url, domain))

	_, err = conn.Write(request)
	if err != nil {
		return "", err
	}

	reader, err := io.ReadAll(conn)
	if err != nil {
		return "", err
	}

	content := string(reader)

	bodyStart := strings.Index(content, "\r\n\r\n")

	if bodyStart == -1 {
		return "", errors.New("missing body")
	}

	return content[bodyStart+4:], nil
}
