// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	serverURL string
	http      HTTPClient
}

func NewClient(serverURL string) *Client {
	return &Client{
		serverURL: serverURL,
		http: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c *Client) Caching(ctx context.Context, params CachingParams) (*WorldObjects, error) {
	path := c.serverURL
	path += "/cached-worlds"

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := r.URL.Query()
	{
		s := conv.Int64ToString(params.Count)
		q.Set("count", s)
	}
	r.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeCachingResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) DB(ctx context.Context) (*WorldObject, error) {
	path := c.serverURL
	path += "/db"

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeDBResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) JSON(ctx context.Context) (*HelloWorld, error) {
	path := c.serverURL
	path += "/json"

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeJSONResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) Queries(ctx context.Context, params QueriesParams) (*WorldObjects, error) {
	path := c.serverURL
	path += "/queries"

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := r.URL.Query()
	{
		s := conv.Int64ToString(params.Queries)
		q.Set("queries", s)
	}
	r.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeQueriesResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) Updates(ctx context.Context, params UpdatesParams) (*WorldObjects, error) {
	path := c.serverURL
	path += "/updates"

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := r.URL.Query()
	{
		s := conv.Int64ToString(params.Queries)
		q.Set("queries", s)
	}
	r.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeUpdatesResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}