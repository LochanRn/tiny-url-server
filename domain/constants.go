package domain

import "github.com/go-openapi/strfmt"

const (
	SiteClient = "SITE_CLIENT"
	Mock       = "MOCK"
)

type Time strfmt.DateTime

type DBType string

const (
	LocalStore DBType = "LOCAL_STORE"
)
