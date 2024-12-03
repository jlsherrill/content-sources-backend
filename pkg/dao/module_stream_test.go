package dao

import (
	"context"
	"testing"

	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/content-sources-backend/pkg/seeds"
	"github.com/content-services/yummy/pkg/yum"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ModuleStreamSuite struct {
	*DaoSuite
	repoConfig  *models.RepositoryConfiguration
	repo        *models.Repository
	repoPrivate *models.Repository
}

func (s *ModuleStreamSuite) SetupTest() {
	s.DaoSuite.SetupTest()

	repo := repoPublicTest.DeepCopy()
	if err := s.tx.Create(repo).Error; err != nil {
		s.FailNow("Preparing Repository record: %w", err)
	}
	s.repo = repo

	repoPrivate := repoPrivateTest.DeepCopy()
	if err := s.tx.Create(repoPrivate).Error; err != nil {
		s.FailNow("Preparing private Repository record: %w", err)
	}
	s.repoPrivate = repoPrivate

	repoConfig := repoConfigTest1.DeepCopy()
	repoConfig.RepositoryUUID = repo.Base.UUID
	if err := s.tx.Create(repoConfig).Error; err != nil {
		s.FailNow("Preparing RepositoryConfiguration record: %w", err)
	}
	s.repoConfig = repoConfig
}

func TestModuleStreamSuite(t *testing.T) {
	m := DaoSuite{}
	r := ModuleStreamSuite{DaoSuite: &m}
	suite.Run(t, &r)
}

func testYumModuleMD() yum.ModuleMD {
	return yum.ModuleMD{
		Document: "",
		Version:  0,
		Data: yum.Stream{
			Name:        "myModule",
			Stream:      "myStream",
			Version:     "Version",
			Context:     "lksdfoisdjf",
			Arch:        "x86_64",
			Summary:     "something short",
			Description: "something long",
			Artifacts: yum.Artifacts{Rpms: []string{"ruby-0:2.5.5-106.module+el8.3.0+7153+c6f6daa5.i686",
				"ruby-irb-0:2.5.5-106.module+el8.3.0+7153+c6f6daa5.noarch"}},
		},
	}
}

func (s *ModuleStreamSuite) TestInsertForRepository() {
	t := s.Suite.T()
	tx := s.tx

	mods := []yum.ModuleMD{testYumModuleMD()}

	dao := GetModuleStreamDao(tx)
	cnt, err := dao.InsertForRepository(context.Background(), s.repo.UUID, mods)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), cnt)

	created := models.ModuleStream{}
	res := tx.Where("context = ?", mods[0].Data.Context).Find(&created)
	assert.NoError(t, res.Error)
	assert.NotEmpty(t, created.UUID)
	assert.Equal(t, created.PackageNames[0], "ruby")
	assert.Equal(t, created.PackageNames[1], "ruby-irb")

	// re-run and expect only 1
	cnt, err = dao.InsertForRepository(context.Background(), s.repo.UUID, mods)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), cnt)
	var count int64
	res = tx.Model(&models.ModuleStream{}).Where("context = ?", mods[0].Data.Context).Count(&count)
	assert.NoError(t, res.Error)
	assert.Equal(t, int64(1), count)
}

func (s *ModuleStreamSuite) TestOrphanCleanup() {
	mod1 := models.ModuleStream{
		Name:         "mod1",
		Stream:       "mod1",
		Version:      "mod1",
		Context:      "mod1",
		Arch:         "mod1",
		Summary:      "mod1",
		Description:  "mod1",
		PackageNames: []string{"foo1"},
		HashValue:    random.String(10),
		Repositories: nil,
	}
	mod2 := models.ModuleStream{
		Name:         "mod2",
		Stream:       "mod2",
		Version:      "mod2",
		Context:      "mod2",
		Arch:         "mod2",
		Summary:      "mod2",
		Description:  "mod12",
		PackageNames: []string{"foo2"},
		HashValue:    random.String(10),
		Repositories: nil,
	}

	require.NoError(s.T(), s.tx.Create(&mod1).Error)
	require.NoError(s.T(), s.tx.Create(&mod2).Error)

	repos, err := seeds.SeedRepositoryConfigurations(s.tx, 1, seeds.SeedOptions{})
	require.NoError(s.T(), err)
	repo := repos[0]

	err = s.tx.Create(&models.RepositoryModuleStream{
		RepositoryUUID:   repo.RepositoryUUID,
		ModuleStreamUUID: mod1.UUID,
	}).Error
	require.NoError(s.T(), err)

	dao := GetModuleStreamDao(s.tx)
	err = dao.OrphanCleanup(context.Background())
	require.NoError(s.T(), err)

	// verify mod1 exists and mod2 doesn't
	mods := []models.ModuleStream{}
	err = s.tx.Where("uuid in (?)", []string{mod1.UUID, mod2.UUID}).Find(&mods).Error
	require.NoError(s.T(), err)

	assert.Len(s.T(), mods, 1)
	assert.Equal(s.T(), mod1.UUID, mods[0].UUID)
}
