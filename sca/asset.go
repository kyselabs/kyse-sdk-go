package sca

import "github.com/kyselabs/kyse-sdk-go/entities"

type Resource struct {
	Vendor string  `json:"vendor" example:"PyPI"`
	Assets []Asset `json:"assets"`
}

type VerdictedResource struct {
	Vendor string           `json:"vendor" example:"PyPI"`
	Assets []VerdictedAsset `json:"assets"`
}

type Asset struct {
	Package string `json:"package" example:"requests"`
	Version string `json:"version" example:"1.2.3"`
}

type VerdictedAsset struct {
	Asset
	Verdict Verdict `json:"verdict"`
}

type Verdict struct {
	DesiredVersion  string          `json:"desired_version"`
	Vulnerabilities []entities.Vuln `json:"vulnerabilities"`
}
