package transport

import (
	"Tasks_Users_Vk_test/internal/domain"
	"Tasks_Users_Vk_test/internal/repository"
	mock_repository "Tasks_Users_Vk_test/internal/repository/mocks"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuestHandler_CreateQuest(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockQuest, quest domain.Quest)

	testTable := []struct {
		name                string
		inputBody           string
		inputQuest          domain.Quest
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "TestQuest", "cost": 1000, "stages": 2}`,
			inputQuest: domain.Quest{
				Name:   "TestQuest",
				Cost:   1000,
				Stages: 2,
			},
			mockBehavior: func(s *mock_repository.MockQuest, quest domain.Quest) {
				s.EXPECT().CreateQuest(quest).Return(nil)
			},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: "",
		},
		{
			name:                "Empty fields",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_repository.MockQuest, quest domain.Quest) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error": "fields are required"}`,
		},
		{
			name:                "Incorrect fields",
			inputBody:           `{"nameeee": "test"}`,
			mockBehavior:        func(s *mock_repository.MockQuest, quest domain.Quest) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error": "invalid request body"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			quest := mock_repository.NewMockQuest(c)
			testCase.mockBehavior(quest, testCase.inputQuest)

			repository := &repository.Repositories{Quest: quest}
			handler := NewUserHandler(repository)

			rr := httptest.NewRecorder()

			// Создание фейкового http.Request с телом запроса
			reqBody := bytes.NewBufferString(testCase.inputBody)
			req := httptest.NewRequest(http.MethodPost, "/quest", reqBody)

			// Вызов обработчика с фейковыми объектами http.ResponseWriter и http.Request
			handler.CreateUser(rr, req)

			assert.Equal(t, testCase.expectedStatusCode, rr.Code)

		})
	}
}
