package routes

type Route struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	File    string `json:"file,omitempty"`
	Handler string `json:"handler,omitempty"`
}
