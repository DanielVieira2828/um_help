package validation

import (
	"encoding/json"
	"errors"
	"io"
)

func VerifyNewUserRequest(rc io.ReadCloser) (r *req.NewUser, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("invalid read closer")
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, errors.New("invalid json payload")
	}

	if len(r.FirstName) <= 2 {
		return nil, errors.New("first name too short, please type at least 3 characters")
	}

	if len(r.FirstName) > 40 {
		return nil, errors.New("first name too long, please keep it at under 40 characters")
	}

	if len(r.LastName) <= 2 {
		return nil, errors.New("last name too short, please type at least 3 characters")
	}

	if len(r.LastName) > 40 {
		return nil, errors.New("last name too long, please keep it at under 40 characters")
	}

	if len(r.DocumentNumber) != 11 {
		return nil, errors.New("document number invalid, please type 11 characters")
	}

	return r, nil
}
