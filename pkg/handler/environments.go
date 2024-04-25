package handler

import (
	"net/http"

	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/dao"
	ce "github.com/content-services/content-sources-backend/pkg/errors"
	"github.com/content-services/content-sources-backend/pkg/rbac"
	"github.com/labstack/echo/v4"
	"github.com/openlyinc/pointy"
)

type RepositoryEnvironmentHandler struct {
	Dao dao.DaoRegistry
}

func RegisterEnvironmentRoutes(engine *echo.Group, rDao *dao.DaoRegistry) {
	rh := RepositoryEnvironmentHandler{
		Dao: *rDao,
	}

	addRoute(engine, http.MethodGet, "/repositories/:uuid/environments", rh.listRepositoriesEnvironments, rbac.RbacVerbRead)
	addRoute(engine, http.MethodPost, "/environments/names", rh.searchEnvironmentByName, rbac.RbacVerbRead)
	addRoute(engine, http.MethodPost, "/snapshots/environments/names", rh.searchSnapshotEnvironments, rbac.RbacVerbRead)
}

// searchEnvironmentByName godoc
// @Summary      Search environments
// @ID           searchEnvironments
// @Description  This enables users to search for environments in a given list of repositories.
// @Tags         repositories,environments
// @Accept       json
// @Produce      json
// @Param        body  body   api.ContentUnitSearchRequest  true  "request body"
// @Success      200 {object} []api.SearchEnvironmentResponse
// @Failure      400 {object} ce.ErrorResponse
// @Failure      401 {object} ce.ErrorResponse
// @Failure      404 {object} ce.ErrorResponse
// @Failure      415 {object} ce.ErrorResponse
// @Failure      500 {object} ce.ErrorResponse
// @Router       /environments/names [post]
func (rh *RepositoryEnvironmentHandler) searchEnvironmentByName(c echo.Context) error {
	_, orgId := getAccountIdOrgId(c)
	dataInput := api.ContentUnitSearchRequest{}
	if err := c.Bind(&dataInput); err != nil {
		return ce.NewErrorResponse(c, http.StatusBadRequest, "Error binding parameters", err.Error())
	}
	preprocessInput(&dataInput)

	apiResponse, err := rh.Dao.Environment.Search(orgId, dataInput)
	if err != nil {
		return ce.NewErrorResponse(c, ce.HttpCodeForDaoError(err), "Error searching environments", err.Error())
	}

	return c.JSON(200, apiResponse)
}

// listRepositoriesEnvironments godoc
// @Summary      List Repositories Environments
// @ID           listRepositoriesEnvironments
// @Description  List environments in a repository.
// @Tags         repositories,environments
// @Accept       json
// @Produce      json
// @Param		 uuid	path string true "Repository ID."
// @Param		 limit query int false "Number of items to include in response. Use it to control the number of items, particularly when dealing with large datasets. Default value: `100`."
// @Param		 offset query int false "Starting point for retrieving a subset of results. Determines how many items to skip from the beginning of the result set. Default value:`0`."
// @Param		 search query string false "Term to filter and retrieve items that match the specified search criteria. Search term can include name."
// @Param		 sort_by query string false "Sort the response based on specific repository parameters. Sort criteria can include `id`, `name`, and `description`."
// @Success      200 {object} api.RepositoryEnvironmentCollectionResponse
// @Failure      400 {object} ce.ErrorResponse
// @Failure      401 {object} ce.ErrorResponse
// @Failure      404 {object} ce.ErrorResponse
// @Failure      500 {object} ce.ErrorResponse
// @Router       /repositories/{uuid}/environments [get]
func (rh *RepositoryEnvironmentHandler) listRepositoriesEnvironments(c echo.Context) error {
	// Read input information
	environmentInput := api.ContentUnitListRequest{}
	if err := c.Bind(&environmentInput); err != nil {
		return ce.NewErrorResponse(c, http.StatusInternalServerError, "Error binding parameters", err.Error())
	}

	_, orgId := getAccountIdOrgId(c)
	page := ParsePagination(c)

	// Request record from database
	apiResponse, total, err := rh.Dao.Environment.List(orgId, environmentInput.UUID, page.Limit, page.Offset, environmentInput.Search, environmentInput.SortBy)

	if err != nil {
		return ce.NewErrorResponse(c, ce.HttpCodeForDaoError(err), "Error listing environments", err.Error())
	}

	return c.JSON(200, setCollectionResponseMetadata(&apiResponse, c, total))
}

// searchSnapshotEnvironments godoc
// @Summary      Search environments within snapshots
// @ID           searchSnapshotEnvironments
// @Description  This enables users to search for environments in a given list of snapshots.
// @Tags         snapshots,environments
// @Accept       json
// @Produce      json
// @Param        body  body   api.SnapshotSearchRpmRequest  true  "request body"
// @Success      200 {object} []api.SearchEnvironmentResponse
// @Failure      400 {object} ce.ErrorResponse
// @Failure      401 {object} ce.ErrorResponse
// @Failure      404 {object} ce.ErrorResponse
// @Failure      415 {object} ce.ErrorResponse
// @Failure      500 {object} ce.ErrorResponse
// @Router       /snapshots/environments/names [post]
func (rh *RepositoryEnvironmentHandler) searchSnapshotEnvironments(c echo.Context) error {
	_, orgId := getAccountIdOrgId(c)
	dataInput := api.SnapshotSearchRpmRequest{}

	var err error
	err = CheckSnapshotAccessible(c)
	if err != nil {
		return err
	}

	if err = c.Bind(&dataInput); err != nil {
		return ce.NewErrorResponse(c, http.StatusBadRequest, "Error binding parameters", err.Error())
	}
	if dataInput.Limit == nil || *dataInput.Limit > api.SearchRpmRequestLimitDefault {
		dataInput.Limit = pointy.Pointer(api.SearchRpmRequestLimitDefault)
	}

	resp, err := rh.Dao.Environment.SearchSnapshotEnvironments(c.Request().Context(), orgId, dataInput)
	if err != nil {
		return ce.NewErrorResponse(c, ce.HttpCodeForDaoError(err), "Error searching environments", err.Error())
	}
	return c.JSON(200, resp)
}
