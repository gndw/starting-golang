package models

import "errors"

type TestRequest struct {
	UserID int64 `json:"user_id"`
}

func (m *TestRequest) Validate() error {
	if m.UserID <= 0 {
		return errors.New("[TestRequest] user_id cannot be <= 0")
	}
	return nil
}

type TestResponse struct {
	UserID   int64  `json:"user_id"`
	FullName string `json:"full_name"`
}
