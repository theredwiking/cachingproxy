package origin

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"

	"github.com/theredwiking/cacheproxy/internal/pkg/slices"
)

// New Request created from origin url and provided path.
// Example is origin.Request("Get", "test")
func (origin *Origin) Request(method string, path string, headers map[string][]string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	url := fmt.Sprintf("%s%s", origin.Url, path)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		error := fmt.Sprintf("Failed to create request %s, error: %v", url, err)
		return nil, errors.New(error)
	}

	req.Header = headers

	resp, err := client.Do(req)

	if err != nil {
		error := fmt.Sprintf("Failed to complete request %s, error: %v", url, err)
		return nil, errors.New(error)
	}

	ok := slices.ContainsInt(validCodes[:], resp.StatusCode)

	if !ok {
		error := fmt.Sprintf("Invalid statuscode: %d, Only accepted status codes are: %v", resp.StatusCode, validCodes)
		return nil, errors.New(error)
	}

	return resp, nil
}
