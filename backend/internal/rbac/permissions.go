package rbac

type Permission int

const (
	ViewQueue Permission = iota
	CreateDraft
	EditOwnDraft
	EditAnyDraft
	RequestChanges
	PublishApp
	ArchiveApp
	ManageUsers
	ManageSystem
)
