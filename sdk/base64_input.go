package sdk

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// Base64FilePath distinguishes a path from a string that is already base64 encoded.
type Base64FilePath string

// EncodeBase64FileInput accepts a path, a byte slice, a reader, or a pre-encoded
// string. Seekable readers are restored to their original position.
func EncodeBase64FileInput(value any) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case Base64FilePath:
		raw, err := os.ReadFile(string(v))
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(raw), nil
	case []byte:
		return base64.StdEncoding.EncodeToString(v), nil
	case io.Reader:
		var position int64
		seeker, seekable := v.(io.Seeker)
		if seekable {
			var err error
			position, err = seeker.Seek(0, io.SeekCurrent)
			if err != nil {
				seekable = false
			}
		}
		raw, err := io.ReadAll(v)
		if seekable {
			_, restoreErr := seeker.Seek(position, io.SeekStart)
			if restoreErr != nil {
				return "", restoreErr
			}
		}
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(raw), nil
	default:
		return "", fmt.Errorf("unsupported base64 input %T", value)
	}
}
