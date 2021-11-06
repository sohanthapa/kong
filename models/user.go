package models

//User contains the info about a User
type User struct {
	ID          string // unique value
	Name        string
	Permissions []string
}

// HasPermission ranges over the User Permissions and checks for the perm
func (user *User) HasPermission(perm string) bool {
	for _, userPerm := range user.Permissions {
		if userPerm == perm {
			return true
		}
	}
	return false
}
