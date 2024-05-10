package transport

import (
	"Tasks_Users_Vk_test/internal/model"
	"Tasks_Users_Vk_test/internal/repository"
	mock_repository "Tasks_Users_Vk_test/internal/repository/mocks"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockUser, user model.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test", "balance": 1000}`,
			inputUser: model.User{
				Name:    "Test",
				Balance: 1000,
			},
			mockBehavior: func(s *mock_repository.MockUser, user model.User) {
				s.EXPECT().CreateUser(user).Return(nil)
			},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: "",
		},
		{
			name:                "Empty fields",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_repository.MockUser, user model.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"fields are required"}` + "\n",
		},
		{
			name:                "Incorrect fields",
			inputBody:           `{"n": }`,
			mockBehavior:        func(s *mock_repository.MockUser, user model.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid request body"}` + "\n",
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

			require.Equal(t, testCase.expectedStatusCode, rr.Code)
			require.Equal(t, testCase.expectedRequestBody, rr.Body.String())

		})
	}
}
