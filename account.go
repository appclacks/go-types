package types

type ChangeAccountPassword struct {
	NewPassword string `json:"new-password" description:"User new password" validate:"required,min=10,max=255"`
}
