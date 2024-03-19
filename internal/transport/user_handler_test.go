package transport

import (
	"Tasks_Users_Vk_test/internal/domain"
	"Tasks_Users_Vk_test/internal/repository"
	mock_repository "Tasks_Users_Vk_test/internal/repository/mocks"
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockUser, user domain.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           domain.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test", "balance": 1000}`,
			inputUser: domain.User{
				Name:    "Test",
				Balance: 1000,
			},
			mockBehavior: func(s *mock_repository.MockUser, user domain.User) {
				s.EXPECT().CreateUser(user).Return(nil)
			},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: "",
		},
		{
			name:                "Empty fields",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_repository.MockUser, user domain.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"fields are required"}`,
		},
		{
			name:                "Incorrect fields",
			inputBody:           `{"n": "test"}`,
			mockBehavior:        func(s *mock_repository.MockUser, user domain.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid request body"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_repository.NewMockUser(c)
			testCase.mockBehavior(user, testCase.inputUser)

			repository := &repository.Repositories{User: user}
			handler := NewUserHandler(repository)

			rr := httptest.NewRecorder()

			// Создание фейкового http.Request с телом запроса
			reqBody := bytes.NewBufferString(testCase.inputBody)
			req := httptest.NewRequest(http.MethodPost, "/users", reqBody)

			// Вызов обработчика с фейковыми объектами http.ResponseWriter и http.Request
			handler.CreateUser(rr, req)

			fmt.Println("Actual response body:", rr.Body.String())
			fmt.Println("Expected response body:", testCase.expectedRequestBody)
			assert.Equal(t, testCase.expectedStatusCode, rr.Code)
			assert.Equal(t, testCase.expectedRequestBody, rr.Body.String())

		})
	}
}
