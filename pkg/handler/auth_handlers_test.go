package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/pkg/service"
	mock_service "github.com/ryodanqqe/flight-bookings/pkg/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, user models.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"admin", "password":"qwerty"}`,
			inputUser: models.User{
				Email:    "admin",
				Password: "qwerty",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return("some_id", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"userID":"some_id"}`,
		},
		{
			name:      "Bad Request",
			inputBody: `{"email":"user"}`,
			inputUser: models.User{
				Email: "user",
			},
			mockBehaviour:       func(s *mock_service.MockAuthorization, user models.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Internal Service Error",
			inputBody: `{"email":"user", "password":"qwerty"}`,
			inputUser: models.User{
				Email:    "user",
				Password: "qwerty",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			test.mockBehaviour(auth, test.inputUser)

			services := &service.Service{Authorization: auth}
			handler := Handler{services: services}

			// Test Server
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(test.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Body.String(), test.expectedRequestBody)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
