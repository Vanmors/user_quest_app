package transport

import (
	"Tasks_Users_Vk_test/internal/domain"
	repository2 "Tasks_Users_Vk_test/internal/repository"
	"Tasks_Users_Vk_test/internal/service"
	mock_service "Tasks_Users_Vk_test/internal/service/mocks"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCompletedQuestsHandler_CompleteTask(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCompletedQuests, record domain.RecordCompleted)

	testTable := []struct {
		name                string
		inputBody           string
		inputRecord         domain.RecordCompleted
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"userID": 1, "questID": 1}`,
			inputRecord: domain.RecordCompleted{
				UserID:  1,
				QuestID: 1,
			},
			mockBehavior: func(s *mock_service.MockCompletedQuests, record domain.RecordCompleted) {
				s.EXPECT().CompleteTask(domain.RecordCompleted{
					UserID:  1,
					QuestID: 1,
				}).Return(nil)
			},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: "",
		},
		{
			name:                "Empty fields",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_service.MockCompletedQuests, record domain.RecordCompleted) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"fields are required"}` + "\n",
		},
		{
			name:                "Incorrect fields",
			inputBody:           `{"n": }`,
			mockBehavior:        func(s *mock_service.MockCompletedQuests, record domain.RecordCompleted) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid request body"}` + "\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			record := mock_service.NewMockCompletedQuests(c)
			testCase.mockBehavior(record, testCase.inputRecord)

			repository := &repository2.Repositories{}
			service := &service.Services{CompletedQuests: record}

			handler := NewCompletedQuestsHandler(repository, service)

			rr := httptest.NewRecorder()

			// Создание фейкового http.Request с телом запроса
			reqBody := bytes.NewBufferString(testCase.inputBody)
			req := httptest.NewRequest(http.MethodPost, "/complete", reqBody)

			// Вызов обработчика с фейковыми объектами http.ResponseWriter и http.Request
			handler.CompleteTask(rr, req)

			require.Equal(t, testCase.expectedStatusCode, rr.Code)
			require.Equal(t, testCase.expectedRequestBody, rr.Body.String())

		})
	}
}
