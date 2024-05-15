package validation

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"

	"github.com/DanielVieirass/um_help/presenter/req"
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

	re := regexp.MustCompile(`^[0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}$`)
	if !re.MatchString(r.DocumentNumber) {
		return nil, errors.New("document number invalid, please type 14 characters in the format 123.456.789-00")
	}

	return r, nil
}
