package validation

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/DanielVieirass/um_help/presenter/req"
)

func VerifyLoginRequest(rc io.ReadCloser) (r *req.LoginRequest, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("invalid read closer")
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, errors.New("invalid json payload")
	}

	re := regexp.MustCompile(`^[0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}$`)
	if !re.MatchString(r.DocumentNumber) {
		return nil, errors.New("document number invalid, please type 14 characters in the format 123.456.789-00")
	}

	if len(r.Password) < 8 {
		return nil, errors.New("password too short, please type at least 8 characters")
	}

	return r, nil
}