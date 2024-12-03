package dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/yummy/pkg/yum"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type moduleStreamDaoImpl struct {
	db *gorm.DB
}

func GetModuleStreamDao(db *gorm.DB) ModuleStreamDao {
	// Return DAO instance
	return moduleStreamDaoImpl{
		db: db,
	}
}

func (r moduleStreamDaoImpl) fetchRepo(ctx context.Context, uuid string) (models.Repository, error) {
	found := models.Repository{}
	if err := r.db.WithContext(ctx).
		Where("UUID = ?", uuid).
		First(&found).
		Error; err != nil {
		return found, err
	}
	return found, nil
}

// Converts an rpm NVREA into just the name
func extractRpmName(nvrea string) string {
	// rubygem-bson-debugsource-0:4.3.0-2.module+el8.1.0+3656+f80bfa1d.x86_64
	split := strings.Split(nvrea, "-")
	if len(split) < 3 {
		return nvrea
	}
	split = split[0 : len(split)-2]
	return strings.Join(split, "-")
}

func ModuleMdToModuleStreams(moduleMds []yum.ModuleMD) (moduleStreams []models.ModuleStream) {
	for _, m := range moduleMds {
		mStream := models.ModuleStream{
			Name:         m.Data.Name,
			Stream:       m.Data.Stream,
			Version:      m.Data.Version,
			Context:      m.Data.Context,
			Arch:         m.Data.Arch,
			Summary:      m.Data.Summary,
			Description:  m.Data.Description,
			PackageNames: []string{},
		}
		for _, p := range m.Data.Artifacts.Rpms {
			mStream.PackageNames = append(mStream.PackageNames, extractRpmName(p))
		}
		slices.Sort(mStream.PackageNames) // Sort the package names so the hash is consistent
		mStream.HashValue = generateHash(mStream.ToHashString())
		moduleStreams = append(moduleStreams, mStream)
	}
	return moduleStreams
}

// InsertForRepository inserts a set of yum module streams for a given repository
// and removes any that are not in the list.  This will involve inserting the package groups
// if not present, and adding or removing any associations to the Repository
// Returns a count of new package groups added to the system (not the repo), as well as any error
func (r moduleStreamDaoImpl) InsertForRepository(ctx context.Context, repoUuid string, modules []yum.ModuleMD) (int64, error) {
	var (
		err  error
		repo models.Repository
	)
	ctxDb := r.db.WithContext(ctx)

	// Retrieve Repository record
	if repo, err = r.fetchRepo(ctx, repoUuid); err != nil {
		return 0, fmt.Errorf("failed to fetchRepo: %w", err)
	}

	moduleStreams := ModuleMdToModuleStreams(modules)

	err = ctxDb.Model(&models.ModuleStream{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "hash_value"}},
		DoNothing: true}).
		Create(moduleStreams).Error
	if err != nil {
		return 0, fmt.Errorf("failed to insert module streams: %w", err)
	}

	hashes := make([]string, len(moduleStreams))
	for _, m := range moduleStreams {
		hashes = append(hashes, m.HashValue)
	}
	uuids := make([]string, len(moduleStreams))

	// insert any modules streams, ignoring any hash conflicts
	if err = r.db.WithContext(ctx).
		Where("hash_value in (?)", hashes).
		Model(&models.ModuleStream{}).
		Pluck("uuid", &uuids).Error; err != nil {
		return 0, fmt.Errorf("failed retrieving existing ids in module_streams: %w", err)
	}

	// Delete repository module stream entries not needed
	err = r.deleteUnneeded(ctx, repo, uuids)
	if err != nil {
		return 0, fmt.Errorf("failed to delete unneeded module streams: %w", err)
	}

	// Add any needed repo module stream entries
	repoModStreams := make([]models.RepositoryModuleStream, len(moduleStreams))
	for i, uuid := range uuids {
		repoModStreams[i] = models.RepositoryModuleStream{
			RepositoryUUID:   repo.UUID,
			ModuleStreamUUID: uuid,
		}
	}
	err = ctxDb.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "repository_uuid"}, {Name: "module_stream_uuid"}},
		DoNothing: true}).
		Create(repoModStreams).Error
	if err != nil {
		return 0, fmt.Errorf("failed to insert repo module streams: %w", err)
	}
	return int64(len(repoModStreams)), nil
}

// deleteUnneeded removes any RepositoryPackageGroup entries that are not in the list of package_group_uuids
func (r moduleStreamDaoImpl) deleteUnneeded(ctx context.Context, repo models.Repository, moduleStreamUUIDs []string) error {
	if err := r.db.WithContext(ctx).Model(&models.RepositoryModuleStream{}).
		Where("repository_uuid = ?", repo.UUID).
		Where("module_stream_uuid NOT IN (?)", moduleStreamUUIDs).
		Error; err != nil {
		return err
	}
	return nil
}

func (r moduleStreamDaoImpl) OrphanCleanup(ctx context.Context) error {
	if err := r.db.WithContext(ctx).
		Model(&models.ModuleStream{}).
		Where("NOT EXISTS (select from repositories_module_streams where module_streams.uuid = repositories_module_streams.module_stream_uuid )").
		Delete(&models.ModuleStream{}).Error; err != nil {
		return err
	}
	return nil
}
