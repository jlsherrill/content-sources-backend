package dao

import (
	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/content-sources-backend/pkg/seeds"
	"github.com/stretchr/testify/assert"
)

func (suite *RepositorySuite) TestCreate() {
	name := "Updated"
	url := "http://someUrl.com"
	orgId := "111"
	accountId := "222"

	t := suite.T()

	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	err := GetRepositoryDao(suite.tx).Create(api.RepositoryRequest{
		Name:      &name,
		URL:       &url,
		OrgID:     &orgId,
		AccountID: &accountId,
	})
	assert.Nil(t, err)

	suite.tx.First(&found)
	assert.Equal(t, name, found.Name)
	assert.Equal(t, url, found.URL)
	assert.Equal(t, orgId, found.OrgID)
}

func (suite *RepositorySuite) TestRepositoryCreateAlreadyExists() {
	t := suite.T()
	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)

	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	err = GetRepositoryDao(suite.tx).Create(api.RepositoryRequest{
		Name:      &found.Name,
		URL:       &found.URL,
		OrgID:     &found.OrgID,
		AccountID: &found.AccountID,
	})

	assert.NotNil(t, err)
	daoError, ok := err.(*Error)
	assert.True(t, ok)
	assert.True(t, daoError.BadValidation)
}

func (suite *RepositorySuite) TestRepositoryCreateBlankTest() {
	t := suite.T()

	blank := ""
	name := "name"
	url := "http://foobar.com"
	OrgID := "34"
	AccountID := "34"

	blankItems := []api.RepositoryRequest{
		api.RepositoryRequest{
			Name:      &blank,
			URL:       &url,
			OrgID:     &OrgID,
			AccountID: &AccountID,
		},
		api.RepositoryRequest{
			Name:      &name,
			URL:       &blank,
			OrgID:     &OrgID,
			AccountID: &AccountID,
		},
		api.RepositoryRequest{
			Name:      &name,
			URL:       &url,
			OrgID:     &blank,
			AccountID: &AccountID,
		},
		api.RepositoryRequest{
			Name:      &name,
			URL:       &url,
			OrgID:     &OrgID,
			AccountID: &blank,
		},
	}
	for i := 0; i < len(blankItems); i++ {
		err := GetRepositoryDao(suite.db).Create(blankItems[i])
		assert.NotNil(t, err)
		daoError, ok := err.(*Error)
		assert.True(t, ok)
		// assert.True(t, daoError.BadValidation)
		assert.Contains(t, daoError.Message, "ERROR: null value in column")
	}
}

func (suite *RepositorySuite) TestUpdate() {
	name := "Updated"
	url := "http://someUrl.com"
	t := suite.T()
	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)
	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	err = GetRepositoryDao(suite.tx).Update(found.OrgID, found.UUID,
		api.RepositoryRequest{
			Name: &name,
			URL:  &url,
		})
	assert.Nil(t, err)

	suite.tx.First(&found)
	assert.Equal(t, "Updated", found.Name)
	assert.Equal(t, "http://someUrl.com", found.URL)
}

func (suite *RepositorySuite) TestUpdateEmpty() {
	name := "Updated"
	arch := ""
	t := suite.T()
	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)
	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	assert.NotEmpty(t, found.Arch)
	err = GetRepositoryDao(suite.tx).Update(found.OrgID, found.UUID,
		api.RepositoryRequest{
			Name:             &name,
			DistributionArch: &arch,
		})
	assert.Nil(t, err)

	suite.tx.First(&found)
	assert.Equal(t, name, found.Name)
	assert.Empty(t, found.Arch)
}

func (suite *RepositorySuite) TestDuplicateUpdate() {
	name := "unique"
	t := suite.T()
	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)
	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	err = GetRepositoryDao(suite.tx).Create(api.RepositoryRequest{OrgID: &found.OrgID, AccountID: &found.AccountID, Name: &name, URL: &name})
	assert.Nil(t, err)

	err = GetRepositoryDao(suite.tx).Update(found.OrgID, found.UUID,
		api.RepositoryRequest{
			Name: &name,
			URL:  &name,
		})

	assert.NotNil(t, err)
	daoError, ok := err.(*Error)
	assert.True(t, ok)
	assert.True(t, daoError.BadValidation)
}

func (suite *RepositorySuite) TestUpdateNotFound() {
	name := "unique"
	t := suite.T()
	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)
	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	err = GetRepositoryDao(suite.tx).Update("Wrong OrgID!! zomg hacker", found.UUID,
		api.RepositoryRequest{
			Name: &name,
			URL:  &name,
		})

	assert.NotNil(t, err)
	daoError, ok := err.(*Error)
	assert.True(t, ok)
	assert.True(t, daoError.NotFound)
}

func (suite *RepositorySuite) TestFetch() {
	t := suite.T()
	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)
	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	fetched, err := GetRepositoryDao(suite.tx).Fetch(found.OrgID, found.UUID)
	assert.Nil(t, err)
	assert.Equal(t, found.UUID, fetched.UUID)
	assert.Equal(t, found.Name, fetched.Name)
}

func (suite *RepositorySuite) TestFetchNotFound() {
	t := suite.T()
	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)
	found := models.RepositoryConfiguration{}
	suite.tx.First(&found)

	_, err = GetRepositoryDao(suite.tx).Fetch("bad org id", found.UUID)
	assert.NotNil(t, err)
	daoError, ok := err.(*Error)
	assert.True(t, ok)
	assert.True(t, daoError.NotFound)
}

func (suite *RepositorySuite) TestList() {
	t := suite.T()
	repoConfig := models.RepositoryConfiguration{}
	orgID := "1028"
	var total int64

	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{OrgID: orgID})
	assert.Nil(t, err)

	result := suite.tx.Where("org_id = ?", orgID).Find(&repoConfig).Count(&total)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(1), total)

	response, total, err := GetRepositoryDao(suite.tx).List(orgID, 100, 0)
	assert.Nil(t, err)
	assert.Equal(t, repoConfig.Name, response.Data[0].Name)
	assert.Equal(t, repoConfig.URL, response.Data[0].URL)
	assert.Equal(t, int64(1), total)
}

func (suite *RepositorySuite) TestListNoRepositories() {
	t := suite.T()
	repoConfigs := make([]models.RepositoryConfiguration, 0)
	orgID := "1028"
	limit := 100
	offset := 0
	var total int64

	result := suite.tx.Where("org_id = ?", orgID).Find(&repoConfigs).Count(&total)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(0), total)

	response, total, err := GetRepositoryDao(suite.tx).List(orgID, limit, offset)
	assert.Nil(t, err)
	assert.Empty(t, response.Data)
	assert.Equal(t, int64(0), total)
}

func (suite *RepositorySuite) TestListPageLimit() {
	t := suite.T()
	repoConfigs := make([]models.RepositoryConfiguration, 0)
	orgID := "1028"
	limit := 10
	offset := 0
	var total int64

	err := seeds.SeedRepositoryConfigurations(suite.tx, 20, seeds.SeedOptions{OrgID: orgID})
	assert.Nil(t, err)

	result := suite.tx.Where("org_id = ?", orgID).Find(&repoConfigs).Count(&total)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(20), total)

	response, total, err := GetRepositoryDao(suite.tx).List(orgID, limit, offset)
	assert.Nil(t, err)
	assert.Equal(t, len(response.Data), limit)
	assert.Equal(t, int64(20), total)
}

func (suite *RepositorySuite) TestDelete() {
	t := suite.T()

	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)

	repoConfig := models.RepositoryConfiguration{}
	err = suite.tx.First(&repoConfig).Error
	assert.Nil(t, err)

	err = GetRepositoryDao(suite.tx).Delete(repoConfig.OrgID, repoConfig.UUID)
	assert.Nil(t, err)

	repoConfig2 := models.RepositoryConfiguration{}
	err = suite.tx.Where("org_id = ? AND uuid = ?", repoConfig.OrgID, repoConfig.UUID).
		First(&repoConfig2).Error
	assert.NotNil(t, err)
}

func (suite *RepositorySuite) TestDeleteNotFound() {
	t := suite.T()

	err := seeds.SeedRepositoryConfigurations(suite.tx, 1, seeds.SeedOptions{})
	assert.Nil(t, err)

	found := models.RepositoryConfiguration{}
	result := suite.tx.First(&found)
	assert.Nil(t, result.Error)

	err = GetRepositoryDao(suite.tx).Delete("bad org id", found.UUID)
	assert.NotNil(t, err)
	daoError, ok := err.(*Error)
	assert.True(t, ok)
	assert.True(t, daoError.NotFound)

	result = suite.tx.First(&found)
	assert.Nil(t, result.Error)
}
