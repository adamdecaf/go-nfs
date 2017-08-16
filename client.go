package nfs

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"
)

const (
	authHeaderKey = "X-NFSN-Authentication"
	nfsApiHost = "api.nearlyfreespeech.net"
	nfsApiScheme = "https://"
)

type Client struct {
	h *http.Client

	// auth information
	accountId, apiKey, login string
}

func NewClient(apiKey, login string) *Client {
	return NewClientForAccount("", apiKey, login)
}

func NewClientForAccount(accountId, apiKey, login string) *Client {
	return &Client{
		h: http.DefaultClient,
		// auth info
		accountId: accountId,
		apiKey: apiKey,
		login: login,
	}
}

// getAuthHeader returns the value half for a NFS auth header.
// The format of this header is found in their docs:
// https://members.nearlyfreespeech.net/wiki/API/Introduction
func (c Client) getAuthHeaderValue(requestUri, bodyHash string) (string, error) {
	timestamp := time.Now().Unix()
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	// hash: login;timestamp;salt;api-key;request-uri;body-hash
	hash := fmt.Sprintf("%s;%d;%s;%s;%s;%s", c.login, timestamp, salt, c.apiKey, requestUri, bodyHash)
	hash = hashfn([]byte(hash))

	// X-NFSN-Authentication: login;timestamp;salt;hash
	value := fmt.Sprintf("%s;%d;%s;%s", c.login, timestamp, salt, hash)
	return value, nil
}

func (c Client) get(p string) (*http.Response, error) {
	return c.makeRequest("GET", p, "")
}

func (c Client) put(p, body string) (*http.Response, error) {
	return c.makeRequest("PUT", p, body)
}

func (c Client) makeRequest(method, p, body string) (*http.Response, error) {
	authHeaderValue, err := c.getAuthHeaderValue(p, hashfn([]byte(body)))
	if err != nil {
		return nil, err
	}

	u := nfsApiScheme + path.Join(nfsApiHost, p)
	req, err := http.NewRequest(method, u, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Set the auth header
	req.Header.Set(authHeaderKey, authHeaderValue)

	// Make the http req
	resp, err := c.h.Do(req)

	// Only close response body on errors, otherwise let clients close
	if err != nil {
		if resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}

	return resp, nil
}

// readResponse is designed to be wrapped around a `makeRequest`, `get` or `put` call
// it returns a successful response as a string, but will propegate errors
// from the http call or in reading the response
func (c Client) readResponse(resp *http.Response, err error) (string, error) {
	if err != nil {
		if resp.Body != nil {
			resp.Body.Close()
		}
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil

}

func generateSalt() (string, error) {
	b := make([]byte, 1)
	n, err := rand.Read(b)
	if err != nil || n != len(b) {
		return "", errors.New("error: unable to generate salt")
	}
	return hashfn(b), nil
}

// hashfn returns the hex encoded result of the sha1 hash function over
// the given byte array
func hashfn(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
