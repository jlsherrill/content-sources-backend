package seeds

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/openlyinc/pointy"
	"gorm.io/gorm"
)

type SeedOptions struct {
	OrgID string
}

const (
	batchSize = 500
)

func randomRepositoryConfigurationName() string {
	return fmt.Sprintf("%s - %s - %s", RandStringBytes(2), "TestRepo", RandStringBytes(10))
}

func randomURL() string {
	return fmt.Sprintf("https://%s.com/%s", RandStringBytes(20), RandStringBytes(5))
}

func randomAccountId() string {
	return fmt.Sprintf("%d", rand.Intn(9999))
}

func randomRepositoryRpmName() string {
	return fmt.Sprintf("%s", RandStringBytes(12))
}

var (
	archs []string = []string{
		"x86_64",
		"noarch",
	}
)

func randomRepositoryRpmArch() string {
	return archs[rand.Int()%len(archs)]
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
	result := db.Create(&repos)
	if result.Error != nil {
		return errors.New("could not save seed")
	}
	return nil
}

func SeedRepository(db *gorm.DB, size int) error {
	var repoConfigs []models.RepositoryConfiguration
	var repos []models.Repository

	// Retrieve all the repos
	if r := db.Find(&repoConfigs); r != nil && r.Error != nil {
		return r.Error
	}

	// For each repo add 'size' rpm random packages
	countRecords := 0
	for _, repoConfig := range repoConfigs {
		referRepoConfig := pointy.String(repoConfig.UUID)
		for i := 0; i < size; i++ {
			repo := models.Repository{
				URL:             randomURL(),
				ReferRepoConfig: referRepoConfig,
			}
			repos = append(repos, repo)
			if len(repos) >= batchSize {
				if r := db.Create(repos); r != nil && r.Error != nil {
					return r.Error
				}
				countRecords += len(repos)
				repos = []models.Repository{}
				fmt.Printf("repoConfig: %d        \r", countRecords)
			}
		}
	}

	// Add remaining records
	if len(repos) > 0 {
		if r := db.Create(repos); r != nil && r.Error != nil {
			return r.Error
		}
		countRecords += len(repos)
		repos = []models.Repository{}
		fmt.Printf("repoConfig: %d        \r", countRecords)
	}

	return nil
}

// SeedRepositoryRpms Populate database with random package information
// db The database descriptor.
// size The number of rpm packages per repository to be generated.
func SeedRepositoryRpms(db *gorm.DB, size int) error {
	var repos []models.Repository
	var rpms []models.RepositoryRpm

	// Retrieve all the repos
	if r := db.Find(&repos); r != nil && r.Error != nil {
		return r.Error
	}

	// For each repo add 'size' rpm random packages
	countRecords := 0
	for _, repo := range repos {
		for i := 0; i < size; i++ {
			rpm := models.RepositoryRpm{
				Name:      randomRepositoryRpmName(),
				Arch:      randomRepositoryRpmArch(),
				Version:   fmt.Sprintf("%d.%d.%d", rand.Int()%6, rand.Int()%16, rand.Int()%64),
				Release:   fmt.Sprintf("%d", rand.Int()%128),
				Epoch:     nil,
				ReferRepo: repo.Base.UUID,
				// Repo:      repo,
			}
			rpms = append(rpms, rpm)
			if len(rpms) >= batchSize {
				if r := db.Create(rpms); r != nil && r.Error != nil {
					return r.Error
				}
				countRecords += len(rpms)
				rpms = []models.RepositoryRpm{}
				fmt.Printf("RepositoryRpm: %d        \r", countRecords)
			}
		}
	}

	// Add remaining records
	if len(rpms) > 0 {
		if r := db.Create(rpms); r != nil && r.Error != nil {
			return r.Error
		}
		countRecords += len(rpms)
		rpms = []models.RepositoryRpm{}
		fmt.Printf("RepositoryRpm: %d        \r", countRecords)
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
