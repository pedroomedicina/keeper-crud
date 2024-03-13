package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"keeper-crud/data/request"
	"keeper-crud/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUsersService struct {
	mock.Mock
}

func (m *MockUsersService) SignUp(req request.UserSignUpRequest) {
	m.Called(req)
}

func (m *MockUsersService) AuthenticateUser(email, password string) (*model.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestUsersController_Signup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUsersService)
	controller := NewUsersController(mockService)

	router := gin.Default()
	router.POST("/signup", controller.Signup)

	userSignUpRequest := request.UserSignUpRequest{
		Email:    "test@example.com",
		Name:     "Test Name",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(userSignUpRequest)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	mockService.On("SignUp", mock.AnythingOfType("request.UserSignUpRequest")).Return()

	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, req)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	mockService.AssertExpectations(t) // Verify that SignUp was called
}

func TestUsersController_Signin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUsersService)
	controller := NewUsersController(mockService)

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/signin", controller.Signin)

	loginDetails := request.UserSignInRequest{
		Email:    "user@example.com",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(loginDetails)
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	expectedUser := &model.User{Email: "user@example.com", Name: "Test User"}
	mockService.On("AuthenticateUser", "user@example.com", "password123").Return(expectedUser, nil)

	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, req)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	mockService.AssertExpectations(t)

	var responseBody map[string]string
	err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody)
	assert.NoError(t, err, "should decode response body without error")
	assert.Equal(t, "User signed in successfully", responseBody["message"], "response message should match")
}
