package rest

import "fmt"

type Mime struct {
	Version string
}

func (m *Mime) Accept() string {
	if m.Version == "" {
		return "application/json"
	}

	return fmt.Sprintf("application/vnd.kyse+json; version=%s", m.Version)
}

func NewMime(version string) *Mime {
	return &Mime{Version: version}
}
