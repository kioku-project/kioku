package handler_test

import (
	"testing"

	pbcollab "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/services/user/handler"
	"github.com/kioku-project/kioku/store"
)

func TestUser(t *testing.T) {
	var mockUserStore store.UserStore = nil
	var mockCollaborationService pbcollab.CollaborationService = nil

	userService := handler.New(mockUserStore, mockCollaborationService)

	if userService == nil {
		t.Errorf("Received invalid User Service\n")
	}
}