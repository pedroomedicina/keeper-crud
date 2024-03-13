package repository

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"keeper-crud/model"
	"testing"
)

func TestUsersRepositoryImplementation_SignUpAndFindByEmail(t *testing.T) {
	// Setup in-memory DB
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatalf("Failed to migrate schema: %v", err)
	}

	repo := NewUsersRepositoryImplementation(db)

	// Create a new user
	testUser := model.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password",
	}
	repo.SignUp(testUser)

	// Retrieve the user by email
	retrievedUser, err := repo.FindByEmail("johndoe@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, retrievedUser)
	assert.Equal(t, "John Doe", retrievedUser.Name)
}
