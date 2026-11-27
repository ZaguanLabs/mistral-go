package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2/google"
)

// GCSWorkflowBlobStore uses the JSON storage API and supports a custom HTTPS API root.
// A nil client selects Google Application Default Credentials.
type GCSWorkflowBlobStore struct {
	client               *http.Client
	root, bucket, prefix string
}

func NewGCSWorkflowBlobStore(ctx context.Context, bucket, prefix, apiRoot string, client *http.Client) (*GCSWorkflowBlobStore, error) {
	if bucket == "" {
		return nil, fmt.Errorf("GCS bucket is required")
	}
	if apiRoot == "" {
		apiRoot = "https://storage.googleapis.com"
	}
	base, err := url.Parse(apiRoot)
	if err != nil || base.Scheme != "https" || base.Host == "" || base.User != nil || base.RawQuery != "" || base.Fragment != "" {
		return nil, fmt.Errorf("GCS API root must be an HTTPS URL")
	}
	if client == nil {
		client, err = google.DefaultClient(ctx, "https://www.googleapis.com/auth/devstorage.read_write")
		if err != nil {
			return nil, err
		}
	}
	return &GCSWorkflowBlobStore{client: client, root: strings.TrimRight(apiRoot, "/"), bucket: url.PathEscape(bucket), prefix: strings.Trim(prefix, "/")}, nil
}
func (s *GCSWorkflowBlobStore) name(key string) string {
	if s.prefix != "" && !strings.HasPrefix(key, s.prefix+"/") {
		return s.prefix + "/" + key
	}
	return key
}
func (s *GCSWorkflowBlobStore) PutIfAbsent(ctx context.Context, key string, data []byte) error {
	query := url.Values{"uploadType": {"media"}, "name": {s.name(key)}, "ifGenerationMatch": {"0"}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.root+"/upload/storage/v1/b/"+s.bucket+"/o?"+query.Encode(), bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	response, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusPreconditionFailed {
		return nil
	}
	if response.StatusCode/100 != 2 {
		return fmt.Errorf("upload GCS workflow payload: %s", response.Status)
	}
	_, err = io.Copy(io.Discard, response.Body)
	return err
}
func (s *GCSWorkflowBlobStore) Get(ctx context.Context, key string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.root+"/storage/v1/b/"+s.bucket+"/o/"+url.PathEscape(s.name(key))+"?alt=media", nil)
	if err != nil {
		return nil, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode/100 != 2 {
		return nil, fmt.Errorf("read GCS workflow payload: %s", response.Status)
	}
	return io.ReadAll(response.Body)
}
