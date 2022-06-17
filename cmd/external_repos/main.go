package main

import (
	"os"
	"sort"

	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/content-services/content-sources-backend/pkg/db"
	"github.com/content-services/content-sources-backend/pkg/external_repos"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	args := os.Args
	config.Load()
	err := db.Connect()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to connect to database")
	}

	if len(args) < 2 {
		log.Fatal().Msg("Requires arguments: import.")
	}
	if args[1] == "download" {
		if len(args) < 3 {
			log.Fatal().Msg("Usage:  ./external_repos import /path/to/jsons/")
		}
		importRepos(args[2])
	} else if args[1] == "import" {
		err = external_repos.SaveToDB()
		if err != nil {
			log.Panic().Err(err).Msg("Failed to save repositories")
		}
		log.Debug().Msg("Successfully loaded external repositories.")
	} else if args[1] == "introspect" {
		if len(args) < 3 {
			log.Fatal().Msg("Usage:  ./external_repos introspect URL")
		}
		count, err := external_repos.IntrospectUrl(args[2])
		if err != nil {
			log.Panic().Err(err).Msg("Failed to introspect repositories")
		}
		log.Debug().Msgf("Successfully Inserted %d packages", count)
	} else if args[1] == "introspect-all" {
		count, errors := external_repos.IntrospectAll()

		for i := 0; i < len(errors); i++ {
			log.Panic().Err(errors[i]).Msg("Failed to introspect repositories")
		}

		log.Debug().Msgf("Successfully Inserted %d packages", count)
	}
}

func importRepos(path string) {
	urls, err := external_repos.IBUrlsFromDir(path)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to import repositories")
	}
	sort.Strings(urls)
	err = external_repos.SaveToFile(urls)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to import repositories")
	}
	log.Info().Msg("Saved External Repositories")
}
