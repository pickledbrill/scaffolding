package scaffolding

// RepositoryPost is the post body for create repository request.
type RepositoryPost struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}
