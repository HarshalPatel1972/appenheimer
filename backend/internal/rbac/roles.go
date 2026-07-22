package rbac

type Role string

const (
	Viewer        Role = "viewer"
	Contributor   Role = "contributor"
	Editor        Role = "editor"
	Administrator Role = "administrator"
)

var rolePermissions = map[Role][]Permission{
	Viewer: {
		ViewQueue,
	},
	Contributor: {
		ViewQueue,
		CreateDraft,
		EditOwnDraft,
	},
	Editor: {
		ViewQueue,
		CreateDraft,
		EditOwnDraft,
		EditAnyDraft,
		RequestChanges,
		PublishApp,
		ArchiveApp,
	},
	Administrator: {
		ViewQueue,
		CreateDraft,
		EditOwnDraft,
		EditAnyDraft,
		RequestChanges,
		PublishApp,
		ArchiveApp,
		ManageUsers,
		ManageSystem,
	},
}

func HasPermission(role Role, perm Permission) bool {
	perms, exists := rolePermissions[role]
	if !exists {
		return false
	}
	for _, p := range perms {
		if p == perm {
			return true
		}
	}
	return false
}
