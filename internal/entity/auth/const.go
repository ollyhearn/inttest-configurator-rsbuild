package auth

// fixme: мне это абсолютно не нравится.
// При малейшем обновлении в миграции придется менять
// этот список енумов. Но в то же время в БЛ вообще
// никак не определить какие пермишны есть у пользователя

type EPermission string

const (
	PermissionListUser   EPermission = "perm_list_user"
	PermissionCreateUser EPermission = "perm_create_user"
	PermissionEditUser   EPermission = "perm_edit_user"
	PermissionDeleteUser EPermission = "perm_delete_user"

	PermissionCreateProject EPermission = "perm_create_project"
	PermissionEditProject   EPermission = "perm_edit_project"
	PermissionDeleteProject EPermission = "perm_delete_project"
)
