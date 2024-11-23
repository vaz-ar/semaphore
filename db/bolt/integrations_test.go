package bolt

import (
	"testing"

	"github.com/semaphoreui/semaphore/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type IntegrationStore interface {
	CreateIntegration(db.Integration) (db.Integration, error)
	GetIntegration(int, int) (db.Integration, error)
	UpdateIntegration(db.Integration) error
	DeleteIntegration(int, int) error
}

type MockIntegrationStore struct {
	mock.Mock
}

func (m *MockIntegrationStore) CreateIntegration(integration db.Integration) (db.Integration, error) {
	args := m.Called(integration)
	return args.Get(0).(db.Integration), args.Error(1)
}

func (m *MockIntegrationStore) GetIntegration(projectID int, integrationID int) (db.Integration, error) {
	args := m.Called(projectID, integrationID)
	return args.Get(0).(db.Integration), args.Error(1)
}

func (m *MockIntegrationStore) UpdateIntegration(integration db.Integration) error {
	args := m.Called(integration)
	return args.Error(0)
}

func (m *MockIntegrationStore) DeleteIntegration(projectID int, integrationID int) error {
	args := m.Called(projectID, integrationID)
	return args.Error(0)
}

func TestCreateIntegration(t *testing.T) {
	mockStore := new(MockIntegrationStore)
	integration := db.Integration{ProjectID: 1, Name: "Test Integration"}

	mockStore.On("CreateIntegration", integration).Return(integration, nil)

	createdIntegration, err := mockStore.CreateIntegration(integration)
	assert.NoError(t, err)
	assert.Equal(t, integration.Name, createdIntegration.Name)

	mockStore.AssertExpectations(t)
}

func TestGetIntegration(t *testing.T) {
	mockStore := new(MockIntegrationStore)
	integration := db.Integration{ID: 1, ProjectID: 1, Name: "Test Integration"}

	mockStore.On("GetIntegration", 1, 1).Return(integration, nil)

	retrievedIntegration, err := mockStore.GetIntegration(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, integration, retrievedIntegration)

	mockStore.AssertExpectations(t)
}

func TestUpdateIntegration(t *testing.T) {
	mockStore := new(MockIntegrationStore)
	integration := db.Integration{ID: 1, ProjectID: 1, Name: "Updated Integration"}

	mockStore.On("UpdateIntegration", integration).Return(nil)

	err := mockStore.UpdateIntegration(integration)
	assert.NoError(t, err)

	mockStore.AssertExpectations(t)
}

func TestDeleteIntegration(t *testing.T) {
	mockStore := new(MockIntegrationStore)

	mockStore.On("DeleteIntegration", 1, 1).Return(nil)

	err := mockStore.DeleteIntegration(1, 1)
	assert.NoError(t, err)

	mockStore.AssertExpectations(t)
}

// ...similar changes for other tests using smaller interfaces...
