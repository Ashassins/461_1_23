// url/models.go
package url 

import {
	"time"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
}

// const {
// 	license types
// }

type URL struct {
	ID string // or PK in go, haven't found how to do this yet
	StoredURL string
	AltURL string // Alternative URL
	Name string
	License string // could change this to an enum
	Source string // GitHub/NPM/etc
	ResponseTime int // Store as hours, / 24 = days, etc
	NumContributors int
	// CorrectnessMetric int // not sure how we're measuring this yet
	LicenseCompatible bool
	IsValidURL bool
	// LastUpdated time.Time
}

func GenID() string {
	return uuid.New().String()
}

func (url URL) GetID() string {
	return url.ID
}

func (url URL) GetStoredURL() string {
	return url.StoredURL
}

func (url URL) GetAltURL() string {
	return url.AltURL
}

func (url URL) GetName() string {
	return url.Name
}

func (url URL) GetLicense() string {
	return url.License
}

func (url URL) GetSource() string {
	return url.Source
}

func (url URL) GetResponseTime() int {
	return url.ResponseTime
}

func (url URL) GetNumContributors() int {
	return url.NumContributors
}

func (url URL) GetLicenseCompatible() bool {
	return url.LicenseCompatible
}

func (url URL) GetIsValidURL() bool {
	return url.IsValid
}

// func (url URL) GetLastUpdated() time.Time{
// 	return url.LastUpdated
// }