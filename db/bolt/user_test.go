package bolt

import (
	"testing"
	"time"

	"github.com/semaphoreui/semaphore/db"
)

func TestBoltDb_UpdateProjectUser(t *testing.T) {
	store := CreateTestStore()

	usr, err := store.CreateUser(db.UserWithPwd{
		Pwd: "123456",
		User: db.User{
			Email:    "denguk@example.com",
			Name:     "Denis Gukov",
			Username: "fiftin",
		},
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	proj1, err := store.CreateProject(db.Project{
		Created: time.Now(),
		Name:    "Test1",
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	projUser, err := store.CreateProjectUser(db.ProjectUser{
		ProjectID: proj1.ID,
		UserID:    usr.ID,
		Role:      db.ProjectOwner,
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	projUser.Role = db.ProjectOwner
	err = store.UpdateProjectUser(projUser)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetUsers(t *testing.T) {
	store := CreateTestStore()

	_, err := store.CreateUser(db.UserWithPwd{
		Pwd: "123456",
		User: db.User{
			Email:    "denguk@example.com",
			Name:     "Denis Gukov",
			Username: "fiftin",
		},
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	found, err := store.GetUsers(db.RetrieveQueryParams{})

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(found) != 1 {
		t.Fatal(err.Error())
	}

}

func TestGetUser(t *testing.T) {
	store := CreateTestStore()

	usr, err := store.CreateUser(db.UserWithPwd{
		Pwd: "123456",
		User: db.User{
			Email:    "denguk@example.com",
			Name:     "Denis Gukov",
			Username: "fiftin",
		},
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	found, err := store.GetUser(usr.ID)

	if err != nil {
		t.Fatal(err.Error())
	}

	if found.Username != "fiftin" {
		t.Fatal(err.Error())
	}

	err = store.DeleteUser(usr.ID)

	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestGetUserCount(t *testing.T) {
	store := CreateTestStore()

	// Create first user
	_, err := store.CreateUser(db.UserWithPwd{
		Pwd: "123456",
		User: db.User{
			Email:    "user1@example.com",
			Name:     "User One",
			Username: "userone",
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	// Create second user
	_, err = store.CreateUser(db.UserWithPwd{
		Pwd: "123456",
		User: db.User{
			Email:    "user2@example.com",
			Name:     "User Two",
			Username: "usertwo",
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	// Get user count
	count, err := store.GetUserCount()
	if err != nil {
		t.Fatal(err.Error())
	}

	// Verify the count
	if count != 2 {
		t.Fatalf("expected 2 users, got %d", count)
	}
}
func TestBoltDb_DeleteUser(t *testing.T) {
	store := CreateTestStore()

	// Create a user
	usr, err := store.CreateUser(db.UserWithPwd{
		Pwd: "123456",
		User: db.User{
			Email:    "deleteuser@example.com",
			Name:     "Delete User",
			Username: "deleteuser",
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	// Create a project
	proj, err := store.CreateProject(db.Project{
		Created: time.Now(),
		Name:    "DeleteUserProject",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	// Associate the user with the project
	_, err = store.CreateProjectUser(db.ProjectUser{
		ProjectID: proj.ID,
		UserID:    usr.ID,
		Role:      db.ProjectOwner,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	// Delete the user
	err = store.DeleteUser(usr.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Verify the user is deleted
	_, err = store.GetUser(usr.ID)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Verify the project-user association is deleted
	_, err = store.GetProjectUser(proj.ID, usr.ID)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
