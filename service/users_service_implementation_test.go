package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"keeper-crud/data/request"
	"keeper-crud/model"
	"testing"
)

// MockUsersRepository is a mock type for the UsersRepository
type MockUsersRepository struct {
	mock.Mock
}

func (m *MockUsersRepository) SignUp(user model.User) {
	m.Called(user)
}

func (m *MockUsersRepository) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestUsersServiceImplementation_SignUp(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	validate := validator.New()
	service := NewUsersServiceImplementation(mockRepo, validate)

	userSignUpRequest := request.UserSignUpRequest{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "securepassword",
	}

	// Expectation: SignUp is called with a User model.
	mockRepo.On("SignUp", mock.AnythingOfType("model.User")).Return()

	service.SignUp(userSignUpRequest)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestUsersServiceImplementation_AuthenticateUser(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	validate := validator.New()
	service := NewUsersServiceImplementation(mockRepo, validate)

	expectedUser := &model.User{
		Email:    "user@example.com",
		Password: "$2a$14$2mxdLNoK10VyONRnK93DweKvQDm/yEFO16MIDsYltrLLBdV62zkcW",
	}

	mockRepo.On("FindByEmail", "user@example.com").Return(expectedUser, nil)

	// Simulate successful authentication
	user, err := service.AuthenticateUser("user@example.com", "password")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "user@example.com", user.Email)

	// Simulate failed authentication (e.g., wrong password)
	_, err = service.AuthenticateUser("user@example.com", "wrongpassword")
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
