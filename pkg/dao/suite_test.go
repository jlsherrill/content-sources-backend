package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/lib/pq"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDSNWithOptions(user string, password string, dbname string, host string, port int) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		user,
		password,
		dbname,
		host,
		port,
	)
}

func getDSNWithConfig(c *config.Configuration) string {
	if c == nil {
		return ""
	}
	return getDSNWithOptions(
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.Host,
		c.Database.Port,
	)
}

func getDSNDefault() string {
	config := config.Get()
	return getDSNWithConfig(config)
}

func getDbConnection() *gorm.DB {
	dsn := getDSNDefault()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

type DaoSuite struct {
	suite.Suite
	db                        *gorm.DB
	tx                        *gorm.DB
	skipDefaultTransactionOld bool
}

const orgIdTest = "acme"
const accountIdTest = "817342"

var repoConfigTest1 = models.RepositoryConfiguration{
	Base: models.Base{
		UUID:      "67eb30d9-9264-4726-9d90-8959e0945a55",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	Name:      "Demo Repository Config",
	URL:       "https://www.redhat.com",
	Arch:      "x86_64",
	Versions:  pq.StringArray{"6", "7", "8", "9"},
	AccountID: accountIdTest,
	OrgID:     orgIdTest,
}

var repoTest1 = models.Repository{
	Base: models.Base{
		UUID:      "55bc5f6b-b5e6-45cb-9953-425b6d4102a0",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	URL:             "https://www.redhat.com",
	LastReadTime:    nil,
	LastReadError:   nil,
	ReferRepoConfig: pointy.String(repoConfigTest1.Base.UUID),
}

var repoRpmTest1 = models.RepositoryRpm{
	Name:        "test-package",
	Arch:        "x86_64",
	Version:     "1.0.0",
	Release:     "123",
	Epoch:       pointy.Int32(1),
	Summary:     "Test package summary",
	Description: "Test package summary",
}

var repoRpmTest2 = models.RepositoryRpm{
	Name:        "demo-package",
	Arch:        "noarch",
	Version:     "2.0.0",
	Release:     "321",
	Epoch:       pointy.Int32(2),
	Summary:     "Demo package summary",
	Description: "Demo package summary",
}

//
//
//

func (suite *DaoSuite) SetupTest() {
	// suite.savedDB = db.DB
	suite.db = getDbConnection()
	suite.skipDefaultTransactionOld = suite.db.SkipDefaultTransaction
	suite.db.SkipDefaultTransaction = true
	suite.tx = suite.db.Begin()

	// Remove the content for the 3 involved tables
	suite.db.Where("1=1").Delete(models.RepositoryRpm{})
	suite.db.Where("1=1").Delete(models.Repository{})
	suite.db.Where("1=1").Delete(models.RepositoryConfiguration{})

	// // Add RepositoryConfig record
	// suite.db.Create(&repoConfigTest1)

	// // Add Repository
	// suite.db.Create(&repoTest1)

}

func (s *DaoSuite) TearDownTest() {
	//Rollback and reset db.DB
	s.tx.Rollback()
	s.db.SkipDefaultTransaction = s.skipDefaultTransactionOld
}

func TestDaoSuite(t *testing.T) {
	suite.Run(t, new(DaoSuite))
}