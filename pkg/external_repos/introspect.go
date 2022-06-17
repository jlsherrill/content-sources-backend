package external_repos

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/content-services/content-sources-backend/pkg/db"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/yummy/pkg/yum"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func IntrospectUrl(url string) (int64, error) {
	//TODO Use DAO
	repo := models.Repository{}
	result := db.DB.Where("URL = ?", url).First(&repo)
	if result.Error != nil {
		return 0, result.Error
	}
	return Introspect(repo)
}

func Introspect(repo models.Repository) (int64, error) {
	client, err := httpClient(strings.Contains(repo.URL, "cdn.redhat.com"))
	if err != nil {
		return 0, err
	}
	pkgs, err := yum.ExtractPackageData(client, repo.URL)
	if err != nil {
		return 0, err
	}

	return PagedInsert(repo, pkgs)
}

func IntrospectAll() (int64, []error) {
	//TODO Move to dao
	var repos []models.Repository
	var errors []error
	var total int64
	result := db.DB.Find(&repos)
	if result.Error != nil {
		return 0, []error{result.Error}
	}
	for i := 0; i < len(repos); i++ {
		fmt.Println(repos[i].URL)
		count, err := Introspect(repos[i])
		total += count
		if err != nil {
			errors = append(errors, err)
		}
	}
	return total, errors
}

func httpClient(useCert bool) (http.Client, error) {
	if useCert {
		filename := "/home/jlsherri/cdncert/cert.pem"
		caFile := "/home/jlsherri/cdncert/ca.pem"
		cert, err := tls.LoadX509KeyPair(filename, filename)
		if err != nil {
			return http.Client{}, err
		}

		caCert, err := ioutil.ReadFile(caFile)
		if err != nil {
			return http.Client{}, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		}

		//tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		return http.Client{Transport: transport, Timeout: 600}, nil
	} else {
		return http.Client{}, nil
	}
}

func PagedInsert(repo models.Repository, pkgs []yum.Package) (int64, error) {
	var count int64
	chunk := 5000
	for i := 0; i < len(pkgs); i += chunk {
		end := i + chunk
		if i+chunk > len(pkgs) {
			end = len(pkgs)
		}
		added, err := Insert(repo, pkgs[i:end])
		if err != nil {
			return count, err
		}
		count += added
	}
	return count, nil
}

func Insert(repo models.Repository, pkgs []yum.Package) (int64, error) {
	dbPkgs := Convert(repo, pkgs)
	result := db.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "refer_repo"},
			{Name: "name"},
			{Name: "version"},
			{Name: "release"},
			{Name: "arch"},
			{Name: "epoch"}},
		DoNothing: true}).Create(&dbPkgs)

	return result.RowsAffected, result.Error
}

func Convert(repo models.Repository, yumPkgs []yum.Package) []models.RepositoryRpm {
	var dbPkgs []models.RepositoryRpm
	for i := 1; i < len(yumPkgs); i++ {
		yumPkg := yumPkgs[i]
		epoch := Epoch(yumPkg.Version.Epoch)
		dbPkgs = append(dbPkgs, models.RepositoryRpm{
			Base: models.Base{
				UUID: uuid.NewString(),
			},
			Name:      yumPkg.Name,
			Arch:      yumPkg.Arch,
			Version:   yumPkg.Version.Version,
			Release:   yumPkg.Version.Release,
			Epoch:     &epoch,
			ReferRepo: repo.UUID,
		})
	}
	return dbPkgs
}

func Epoch(epoch string) int {
	intEpoch, err := strconv.Atoi(epoch)
	if err != nil {
		return 0
	}
	return intEpoch
}
