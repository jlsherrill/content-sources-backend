package seeds

import (
	"fmt"
	"math/rand"

	"github.com/content-services/content-sources-backend/pkg/models"
	"gorm.io/gorm"
)

type SeedOptions struct {
	OrgID string
}

func SeedRepositoryConfigurations(db *gorm.DB, size int, options SeedOptions) error {
	var repos []models.RepositoryConfiguration

	if options.OrgID == "" {
		options.OrgID = fmt.Sprintf("%d", rand.Intn(9999))
	}

	for i := 0; i < size; i++ {
		repoConfig := models.RepositoryConfiguration{
			Name:      fmt.Sprintf("%s - %s - %s", RandStringBytes(2), "TestRepo", RandStringBytes(10)),
			URL:       fmt.Sprintf("https://%s.com/%s", RandStringBytes(20), RandStringBytes(5)),
			Versions:  []string{"9"},
			Arch:      "x86_64",
			AccountID: fmt.Sprintf("%d", rand.Intn(9999)),
			OrgID:     options.OrgID,
		}
		repos = append(repos, repoConfig)
	}
	if result := db.Create(&repos); result.Error != nil {
		return result.Error
		// return errors.New("could not save seed")
	}
	return nil
}

func SeedRepository(db *gorm.DB, size int) error {
	var repoConfigs []models.RepositoryConfiguration
	var repos []models.Repository

	archs := []string{
		"amd64",
		"i386",
		"aarch64",
		"noarch",
	}

	// Retrieve all the repos
	if r := db.Find(&repoConfigs); r != nil && r.Error != nil {
		return r.Error
	}

	// For each repo add 'size' rpm random packages
	for repoIdx, repoConfig := range repoConfigs {
		fmt.Printf("repoConfig: %d        \r", repoIdx)
		for i := 0; i < size; i++ {
			arch := archs[rand.Int()%4]
			repo := models.Repository{
				URL:             fmt.Sprintf("https://%s.com/%s", RandStringBytes(12), arch),
				ReferRepoConfig: repoConfig.UUID,
			}
			repos = append(repos, repo)
			if len(repos) >= 10 {
				if r := db.Create(repos); r != nil && r.Error != nil {
					return r.Error
				}
				repos = []models.Repository{}
			}
		}
	}

	// Add remaining records
	if len(repos) > 0 {
		if r := db.Create(repos); r != nil && r.Error != nil {
			return r.Error
		}
		repos = []models.Repository{}
	}

	return nil
}

// SeedRepositoryRpms Populate database with random package information
// db The database descriptor.
// size The number of rpm packages per repository to be generated.
func SeedRepositoryRpms(db *gorm.DB, size int) error {
	var repos []models.Repository
	var rpms []models.RepositoryRpm

	archs := []string{
		"amd64",
		"i386",
		"aarch64",
		"noarch",
	}

	// Retrieve all the repos
	if r := db.Find(&repos); r != nil && r.Error != nil {
		return r.Error
	}

	// For each repo add 'size' rpm random packages
	for repoIdx, repo := range repos {
		fmt.Printf("RepositoryRpm: %d        \r", repoIdx)
		for i := 0; i < size; i++ {
			rpm := models.RepositoryRpm{
				Name:      fmt.Sprintf("%s", RandStringBytes(12)),
				Arch:      archs[rand.Int()%4],
				Version:   fmt.Sprintf("%d.%d.%d", rand.Int()%6, rand.Int()%16, rand.Int()%64),
				Release:   fmt.Sprintf("%d", rand.Int()%128),
				Epoch:     nil,
				ReferRepo: repo.Base.UUID,
				// Repo:      repo,
			}
			rpms = append(rpms, rpm)
			if len(rpms) >= 10 {
				if r := db.Create(rpms); r != nil && r.Error != nil {
					return r.Error
				}
				rpms = []models.RepositoryRpm{}
			}
		}
	}

	// Add remaining records
	if len(rpms) > 0 {
		if r := db.Create(rpms); r != nil && r.Error != nil {
			return r.Error
		}
		rpms = []models.RepositoryRpm{}
	}

	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
