package api

// BuildResponse represents the paginated response for builds endpoint
type BuildResponse struct {
	Results []BuildItem `json:"results"`
	Total   int         `json:"total"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
}

// BuildItem represents a Maven build artifact
type BuildItem struct {
	Group     string `json:"group"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Filename  string `json:"filename"`
	CreatedAt string `json:"created_at"`
}
