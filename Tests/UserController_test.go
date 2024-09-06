package Tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"goSample/Types/Requests"
	"net/http/httptest"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	loginDto := Requests.LoginDto{
		Username: "tuanvan",
		Password: "password",
	}
	jsonData, _ := json.Marshal(loginDto)

	req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(jsonData))
	req.Ca
	assert.Equal(t, 200, resp.StatusCode)
}
