package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/ryodanqqe/flight-bookings/pkg/service"
	mock_service "github.com/ryodanqqe/flight-bookings/pkg/service/mocks"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                string
		headerName          string
		headerValue         string
		token               string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			headerName:  "Cookie",
			headerValue: "token=valid_token",
			token:       "valid_token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return("user_id", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "user_id",
		},
		{
			name:                "No_Token",
			headerName:          "Cookie",
			headerValue:         "",
			token:               "",
			mockBehaviour:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"token cookie not found"}`,
		},
		{
			name:        "Invalid_Token",
			headerName:  "Cookie",
			headerValue: "token=invalid_token",
			token:       "invalid_token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return("", errors.New("invalid token"))
			},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"invalid token"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			test.mockBehaviour(auth, test.token)

			services := &service.Service{Authorization: auth}
			handler := Handler{services: services}

			// Test Server
			r := gin.New()
			r.GET("/identity", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get("user_id")
				c.String(200, "%s", id)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			// Perform Requests
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())
		})
	}
}
