package wildberies

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// const (
// 	apiURL = "https://common-api.wildberries.ru"
// )

type (
	// Doer is an interface that can do http request
	Doer interface {
		Do(req *http.Request) (*http.Response, error)
	}
	// A Client manages communication with the Yandex.Music API.
	Client struct {
		// HTTP client used to communicate with the API.
		client Doer
		// Access token to Wildberies API
		accessToken string
		// Debug sets should library print debug messages or not
		Debug bool
		// Services
		income   *IncomeService
		supplier *SupplierService
	}
)

var deblog = log.New(os.Stdout, "[DEBUG]\t", log.Ldate|log.Ltime|log.Lshortfile)

// NewClient returns a new API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(options ...func(*Client)) *Client {
	c := &Client{
		client: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	c.income = &IncomeService{client: c}
	c.supplier = &SupplierService{client: c}

	return c
}

// HTTPClient sets http client for Wildberies client
func HTTPClient(httpClient Doer) func(*Client) {
	return func(c *Client) {
		if httpClient != nil {
			c.client = httpClient
		}
	}
}

// AccessToken sets user_id and access token for Yandex.Music client
func AccessToken(accessToken string) func(*Client) {
	return func(c *Client) {
		if accessToken != "" {
			c.accessToken = accessToken
		}
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body, except when body is url.Values. If it is url.Values, it is
// encoded as application/x-www-form-urlencoded and included in request
// headers.
func (c *Client) NewRequest(
	method,
	urlStr string,
	body any,
) (*http.Request, error) {
	uri, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var reader io.Reader
	var isForm bool
	if body != nil {
		switch v := body.(type) {
		case url.Values:
			reader = strings.NewReader(v.Encode())
			isForm = true
		default:
			buf := new(bytes.Buffer)
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}

			reader = buf
		}
	}

	req, err := http.NewRequest(method, uri.String(), reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.accessToken)
	if isForm && method == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(
	ctx context.Context,
	req *http.Request,
	v any,
) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// if resp.StatusCode == http.StatusOK {
	// 	bodyBytes, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bodyString := string(bodyBytes)
	// 	fmt.Println(bodyString)
	// }

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			dat, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewReader(dat))
			err = json.Unmarshal(dat, v)
			if err == io.EOF {
				if c.Debug {
					deblog.Println("Got empty")
				}
				// Ignore EOF errors caused by empty response body.
				err = nil //nolint:ineffassign
			} else if err != nil {
				// Try parse XML if it's not JSON.
				err = xml.Unmarshal(dat, v) //nolint:ineffassign,staticcheck
			}
		}
	}

	return resp, err
}

// Income returns income service
func (c *Client) Income() *IncomeService {
	return c.income
}

// Income returns income service
func (c *Client) Supplier() *SupplierService {
	return c.supplier
}

// General types
type (
	// InvocationInfo is base info in all requests
	InvocationInfo struct {
		Hostname string `json:"hostname"`
		ReqID    string `json:"req-id"`
		// ExecDurationMillis sometimes int, sometimes string so saving interface{}
		ExecDurationMillis any `json:"exec-duration-millis"`
	}
	// Error is struct with error type and message.
	Error struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}
)
