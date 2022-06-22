package dao

import (
	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type repositoryDaoImpl struct {
	db *gorm.DB
}

func GetRepositoryDao(db *gorm.DB) RepositoryDao {
	return repositoryDaoImpl{
		db: db,
	}
}

func DBErrorToApi(e error) error {
	pgError, ok := e.(*pgconn.PgError)
	if ok {
		if pgError.Code == "23505" {
			return &Error{BadValidation: true, Message: "Repository with this URL already belongs to organization"}
		}
	}
	dbError, ok := e.(models.Error)
	if ok {
		return &Error{BadValidation: dbError.Validation, Message: dbError.Message}
	}
	return &Error{Message: e.Error()}
}

func (r repositoryDaoImpl) Create(newRepo api.RepositoryRequest) error {
	newRepoConfig := models.RepositoryConfiguration{
		AccountID: *newRepo.AccountID,
		OrgID:     *newRepo.OrgID,
	}
	ApiFieldsToModel(&newRepo, &newRepoConfig)

	if err := r.db.Create(&newRepoConfig).Error; err != nil {
		return DBErrorToApi(err)
	}
	return nil
}

func (r repositoryDaoImpl) List(OrgID string, limit int, offset int) (api.RepositoryCollectionResponse, int64, error) {
	var totalRepos int64
	repoConfigs := make([]models.RepositoryConfiguration, 0)

	result := r.db.Where("org_id = ?", OrgID).Find(&repoConfigs).Count(&totalRepos)
	if result.Error != nil {
		return api.RepositoryCollectionResponse{}, totalRepos, result.Error
	}

	result = r.db.Where("org_id = ?", OrgID).Limit(limit).Offset(offset).Find(&repoConfigs)
	if result.Error != nil {
		return api.RepositoryCollectionResponse{}, totalRepos, result.Error
	}

	repos := convertToResponses(repoConfigs)
	return api.RepositoryCollectionResponse{Data: repos}, totalRepos, nil
}

func (r repositoryDaoImpl) Fetch(orgID string, uuid string) (api.RepositoryResponse, error) {
	repo := api.RepositoryResponse{}
	repoConfig, err := r.fetchRepoConfig(orgID, uuid)
	if err != nil {
		return repo, err
	}
	ModelToApiFields(repoConfig, &repo)
	return repo, err
}

func (r repositoryDaoImpl) fetchRepoConfig(orgID string, uuid string) (models.RepositoryConfiguration, error) {
	found := models.RepositoryConfiguration{}
	result := r.db.Where("UUID = ? AND ORG_ID = ?", uuid, orgID).First(&found)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return found, &Error{NotFound: true, Message: "Could not find repository with UUID " + uuid}
		} else {
			return found, result.Error
		}
	}
	return found, nil
}

func (r repositoryDaoImpl) Update(orgID string, uuid string, repoParams api.RepositoryRequest) error {
	repoConfig, err := r.fetchRepoConfig(orgID, uuid)
	if err != nil {
		return err
	}
	ApiFieldsToModel(&repoParams, &repoConfig)
	result := r.db.Model(&repoConfig).Updates(repoConfig.MapForUpdate())
	if result.Error != nil {
		return DBErrorToApi(result.Error)
	}
	return nil
}

func (r repositoryDaoImpl) Delete(orgID string, uuid string) error {
	repoConfig := models.RepositoryConfiguration{}
	if err := r.db.Debug().Where("UUID = ? AND ORG_ID = ?", uuid, orgID).First(&repoConfig).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &Error{NotFound: true, Message: "Could not find repository with UUID " + uuid}
		} else {
			return err
		}
	}
	if err := r.db.Delete(&repoConfig).Error; err != nil {
		return err
	}
	return nil
}

func ApiFieldsToModel(apiRepo *api.RepositoryRequest, repoConfig *models.RepositoryConfiguration) {
	if apiRepo.Name != nil {
		repoConfig.Name = *apiRepo.Name
	}
	if apiRepo.URL != nil {
		repoConfig.URL = *apiRepo.URL
	}
	if apiRepo.DistributionArch != nil {
		repoConfig.Arch = *apiRepo.DistributionArch
	}
	if apiRepo.DistributionVersions != nil {
		repoConfig.Versions = *apiRepo.DistributionVersions
	}
}

func ModelToApiFields(repoConfig models.RepositoryConfiguration, apiRepo *api.RepositoryResponse) {
	apiRepo.UUID = repoConfig.UUID
	apiRepo.Name = repoConfig.Name
	apiRepo.URL = repoConfig.URL
	apiRepo.DistributionVersions = repoConfig.Versions
	apiRepo.DistributionArch = repoConfig.Arch
	apiRepo.AccountID = repoConfig.AccountID
	apiRepo.OrgID = repoConfig.OrgID
}

// Converts the database models to our response objects
func convertToResponses(repoConfigs []models.RepositoryConfiguration) []api.RepositoryResponse {
	repos := make([]api.RepositoryResponse, len(repoConfigs))
	for i := 0; i < len(repoConfigs); i++ {
		ModelToApiFields(repoConfigs[i], &repos[i])
	}
	return repos
}
