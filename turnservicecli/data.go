package turnservicecli

import (
	"time"
)

// CredentialsResponse defines a REST response containing TURN data.
type CredentialsResponse struct {
	Success bool             `json:"success"`
	Nonce   string           `json:"nonce"`
	Expires *time.Time       `json:"expires,omitempty"`
	Turn    *CredentialsData `json:"turn"`
	Session string           `json:"session,omitempty"`
}

// CredentialsData defines TURN credentials with servers.
type CredentialsData struct {
	TTL      int64         `json:"ttl"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Servers  []*URNsWithID `json:"servers,omitempty"`
	GeoURI   string        `json:"geo_uri,omitempty"`
}

// URNsWithID defines TURN servers groups with ID.
type URNsWithID struct {
	ID    string            `json:"id"`
	URNs  []string          `json:"urns"`
	Prio  int               `json:"prio"`
	Label string            `json:"label,omitempty"`
	I18N  map[string]string `json:"i18n,omitempty"`
}

// GeoResponse defines a REST response containing TURN geo.
type GeoResponse struct {
	Success bool     `json:"success"`
	Nonce   string   `json:"nonce"`
	Geo     *GeoData `json:"geo,omitempty"`
}

// GeoData defines ordered TURN IDs.
type GeoData struct {
	Prefer []string `json:"prefer"`
}
