package scaffolding

// User stores the basic github user information
type User struct {
	Email     string
	Name      string
	SigninKey string
}

// LoadLocal reads the local .gitconfig file to get basic github user information.
func (user *User) LoadLocal() {

}
