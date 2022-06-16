package models

type RepositoryRpm struct {
	Base
	// The rpm package name
	Name string `json:"name" gorm:"not null"`
	// The architecture that this package belong to
	Arch string `json:"arch" gorm:"not null"`
	// The version for this package
	Version string `json:"version" gorm:"not null"`
	// The release for this package
	Release string `json:"release" gorm:"null"`
	// Epoch is a way to define weighted dependencies based
	// on version numbers. It's default value is 0 and this
	// is assumed if an Epoch directive is not listed in the RPM SPEC file.
	// https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/packaging_and_distributing_software/advanced-topics#packaging-epoch_epoch-scriplets-and-triggers
	Epoch       *int32     `json:"epoch" gorm:"default:0;not null"`
	Summary     string     `json:"summary" gorm:"not null"`
	Description string     `json:"description" gorm:"not null"`
	ReferRepo   string     `gorm:"not null"`
	Repo        Repository `gorm:"foreignKey:UUID;references:ReferRepo"`
}