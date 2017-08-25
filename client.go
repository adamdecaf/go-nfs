package nfs

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

const (
	authHeaderKey = "X-NFSN-Authentication"
	nfsApiHost    = "api.nearlyfreespeech.net"
	nfsApiScheme  = "https://"
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
		apiKey:    apiKey,
		login:     login,
	}
}

// Errors represented by the NFS api
type Error struct {
	Debug string `json:"debug"`
	Error string `json:"error"`
}

func (e Error) Err() error {
	return fmt.Errorf("error - %s, %s", e.Error, e.Debug)
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
	return c.makeRequest("GET", p, strings.NewReader(""))
}

func (c Client) post(p string, params map[string]string) (*http.Response, error) {
	data := url.Values{}
	for k, v := range params {
		data.Add(k, v)
	}
	return c.makeRequest("POST", p, strings.NewReader(data.Encode()))
}

func (c Client) put(p, body string) (*http.Response, error) {
	return c.makeRequest("PUT", p, strings.NewReader(body))
}

func (c Client) makeRequest(method, p string, body io.Reader) (*http.Response, error) {
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	authHeaderValue, err := c.getAuthHeaderValue(p, hashfn(bs))
	if err != nil {
		return nil, err
	}

	u := nfsApiScheme + path.Join(nfsApiHost, p)
	req, err := http.NewRequest(method, u, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	// Set the auth header
	req.Header.Set(authHeaderKey, authHeaderValue)

	// Content-type
	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	// Make the http req
	resp, err := c.h.Do(req)

	// Only close response body on errors, otherwise let clients close
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}

	return resp, nil
}

// readResponse is designed to be wrapped around a `makeRequest` or similar call
// it returns a successful response as a string, but will propegate errors
// from the http call or in reading the response
func (c Client) readResponse(resp *http.Response, err error) (string, error) {
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return "", err
	}
	defer resp.Body.Close()

	err = c.checkErrors(resp, err)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil

}

// checkErrors is designed to be wrapped around a `makeRequest` or similar call
// it returns no error if no passed-in error was given and if the response json
// contains no debug/error message
// checkErrors does not modify the response body (except on failure)
func (c Client) checkErrors(resp *http.Response, err error) error {
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return err
	}

	if resp.StatusCode != 200 {
		dec := json.NewDecoder(resp.Body)

		e := Error{}
		err = dec.Decode(&e)
		if err != nil {
			return err
		}
		return e.Err()
	}

	return nil
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
