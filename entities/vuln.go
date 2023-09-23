package entities

type Vuln struct {
	Entity

	RepoName   string `json:"-"`
	ExternalID string `json:"external_id"`

	Summary  string `json:"summary,omitempty"`
	Details  string `json:"details,omitempty"`
	Severity string `json:"severity,omitempty"`
}
