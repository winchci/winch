package transom

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/switch-bit/winch/config"
	"io/ioutil"
	"net/http"
	"os"
)

type LoginRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	ClientCode string `json:"clientCode"`
}

type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
}

type PublishRequest struct {
	Org      string `json:"org"`
	App      string `json:"app"`
	Version  string `json:"version"`
	Contents []byte `json:"asset"`
}

type Version struct {
	Org      string `json:"org"`
	App      string `json:"app"`
	Version  string `json:"version"`
	Checksum string `json:"checksum"`
}

type PublishResponse struct {
	Version Version `json:"version"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type Transom struct {
	shipyardUrl string
	url         string
	token       string
}

func NewTransom(cfg *config.TransomConfig, name string) (*Transom, error) {
	if len(cfg.Shipyard) == 0 {
		cfg.Shipyard = os.Getenv("TRANSOM_SHIPYARD")
	}

	if len(cfg.Shipyard) == 0 {
		cfg.Shipyard = "shipyard.b10s.io"
	}

	if len(cfg.Server) == 0 {
		cfg.Server = os.Getenv("TRANSOM_SERVER")
	}

	if len(cfg.Server) == 0 {
		cfg.Server = "transom.b10s.io"
	}

	if len(cfg.Organization) == 0 {
		cfg.Organization = os.Getenv("TRANSOM_ORGANIZATION")
	}

	if len(cfg.Organization) == 0 {
		cfg.Organization = name
	}

	if len(cfg.Organization) == 0 {
		return nil, fmt.Errorf("the Transom organization is required")
	}

	if len(cfg.Application) == 0 {
		cfg.Application = os.Getenv("TRANSOM_APPLICATION")
	}

	if len(cfg.Application) == 0 {
		cfg.Application = name
	}

	if len(cfg.Application) == 0 {
		return nil, fmt.Errorf("the Transom application is required")
	}

	if len(cfg.Token) == 0 {
		cfg.Token = os.Getenv("TRANSOM_TOKEN")
	}

	if len(cfg.Token) == 0 {
		if len(cfg.Username) == 0 {
			cfg.Username = os.Getenv("TRANSOM_USERNAME")
		}

		if len(cfg.Username) == 0 {
			return nil, fmt.Errorf("the Transom username or token is required")
		}

		if len(cfg.Password) == 0 {
			cfg.Password = os.Getenv("TRANSOM_PASSWORD")
		}

		if len(cfg.Password) == 0 {
			return nil, fmt.Errorf("the Transom password or token is required")
		}
	}

	return &Transom{
		shipyardUrl: fmt.Sprintf("https://%s/shipyard/", cfg.Shipyard),
		url:         fmt.Sprintf("https://%s/transom/", cfg.Server),
	}, nil
}

func (t *Transom) SetToken(token string) {
	t.token = token
}

func (t Transom) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	var resp LoginResponse
	err := t.do(ctx, t.shipyardUrl, "tenants/transom/login", in, &resp, nil)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ComputeChecksum computes a SHA-256 checksum of the given byte slice.
func ComputeChecksum(b []byte) string {
	sha := sha256.New()
	sha.Write(b)
	src := sha.Sum(nil)
	dest := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dest, src)
	return string(dest)
}

func (t Transom) Publish(ctx context.Context, in *PublishRequest) (*PublishResponse, error) {
	var resp PublishResponse

	checksum := ComputeChecksum(in.Contents)

	err := t.do(ctx, t.url, fmt.Sprintf("orgs/%s/apps/%s/versions/%s?force=true&checksum=%s", in.Org,
		in.App, in.Version, checksum), in.Contents, &resp, nil)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (t *Transom) do(ctx context.Context, baseUrl string, url string, in interface{}, out interface{}, headers http.Header) error {
	b := new(bytes.Buffer)
	var contentType string

	if _, ok := in.([]byte); ok {
		b.Write(in.([]byte))
		contentType = "application/octet-stream"
	} else {
		err := json.NewEncoder(b).Encode(in)
		if err != nil {
			return err
		}

		contentType = "application/json"
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, baseUrl+url, b)
	if err != nil {
		return err
	}

	if headers != nil {
		r.Header = headers
	}

	if t.token != "" {
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))
	}

	r.Header.Add("Content-Type", contentType)
	r.Header.Add("Accept", "application/json")
	r.Body = ioutil.NopCloser(b)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var errResp errorResponse
	err = json.Unmarshal(respBytes, &errResp)
	if err != nil {
		return err
	}

	if errResp.Error != "" {
		return errors.New(errResp.Error)
	}

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(respBytes, out)
		if err != nil {
			return err
		}
	}

	return nil
}
