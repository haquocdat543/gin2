package user

const (

	// Success messages
	MsgUserCreated       = "User created successfully"
	MsgUserFetched       = "Users fetched successfully"
	MsgUserInfoFetched   = "User info fetched successfully"
	MsgLoginSuccess      = "login successfully"
	MsgDeleteSuccess     = "delete user successfully"
	MsgUpdateUserSuccess = "update user successfully"
	MsgPatchDeleteUserSuccess  = "patch delete user successfully"
)

const (

	// Error messages
	ErrEmailAlreadyExists = "Email already exists"
	ErrInvalidRequest     = "Invalid request payload"
	ErrInternalServer     = "Internal server error"

	ErrInvalidPassword = "Invalid password"
)
