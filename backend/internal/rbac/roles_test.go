package rbac

import (
	"testing"
)

func TestHasPermission(t *testing.T) {
	// Viewer
	if !HasPermission(Viewer, ViewQueue) {
		t.Error("Viewer should have ViewQueue permission")
	}
	if HasPermission(Viewer, PublishApp) {
		t.Error("Viewer should not have PublishApp permission")
	}

	// Contributor
	if !HasPermission(Contributor, CreateDraft) {
		t.Error("Contributor should have CreateDraft permission")
	}
	if HasPermission(Contributor, EditAnyDraft) {
		t.Error("Contributor should not have EditAnyDraft permission")
	}

	// Editor
	if !HasPermission(Editor, PublishApp) {
		t.Error("Editor should have PublishApp permission")
	}
	if HasPermission(Editor, ManageUsers) {
		t.Error("Editor should not have ManageUsers permission")
	}

	// Administrator
	if !HasPermission(Administrator, ManageSystem) {
		t.Error("Administrator should have ManageSystem permission")
	}
	if !HasPermission(Administrator, PublishApp) {
		t.Error("Administrator should have PublishApp permission")
	}
}
