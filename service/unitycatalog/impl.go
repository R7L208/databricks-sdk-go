// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Catalogs API methods
type catalogsImpl struct {
	client *client.DatabricksClient
}

func (a *catalogsImpl) Create(ctx context.Context, request CreateCatalog) (*CreateCatalogResponse, error) {
	var createCatalogResponse CreateCatalogResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createCatalogResponse)
	return &createCatalogResponse, err
}

func (a *catalogsImpl) Delete(ctx context.Context, request DeleteCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *catalogsImpl) Get(ctx context.Context, request GetCatalogRequest) (*GetCatalogResponse, error) {
	var getCatalogResponse GetCatalogResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getCatalogResponse)
	return &getCatalogResponse, err
}

func (a *catalogsImpl) List(ctx context.Context) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *catalogsImpl) Update(ctx context.Context, request UpdateCatalog) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just ExternalLocations API methods
type externalLocationsImpl struct {
	client *client.DatabricksClient
}

func (a *externalLocationsImpl) Create(ctx context.Context, request CreateExternalLocation) (*CreateExternalLocationResponse, error) {
	var createExternalLocationResponse CreateExternalLocationResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createExternalLocationResponse)
	return &createExternalLocationResponse, err
}

func (a *externalLocationsImpl) Delete(ctx context.Context, request DeleteExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *externalLocationsImpl) Get(ctx context.Context, request GetExternalLocationRequest) (*GetExternalLocationResponse, error) {
	var getExternalLocationResponse GetExternalLocationResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getExternalLocationResponse)
	return &getExternalLocationResponse, err
}

func (a *externalLocationsImpl) List(ctx context.Context) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *externalLocationsImpl) Update(ctx context.Context, request UpdateExternalLocation) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Grants API methods
type grantsImpl struct {
	client *client.DatabricksClient
}

func (a *grantsImpl) Get(ctx context.Context, request GetGrantRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}

func (a *grantsImpl) Update(ctx context.Context, request UpdatePermissions) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Metastores API methods
type metastoresImpl struct {
	client *client.DatabricksClient
}

func (a *metastoresImpl) Assign(ctx context.Context, request CreateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *metastoresImpl) Create(ctx context.Context, request CreateMetastore) (*CreateMetastoreResponse, error) {
	var createMetastoreResponse CreateMetastoreResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createMetastoreResponse)
	return &createMetastoreResponse, err
}

func (a *metastoresImpl) Delete(ctx context.Context, request DeleteMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *metastoresImpl) Get(ctx context.Context, request GetMetastoreRequest) (*GetMetastoreResponse, error) {
	var getMetastoreResponse GetMetastoreResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getMetastoreResponse)
	return &getMetastoreResponse, err
}

func (a *metastoresImpl) List(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *metastoresImpl) Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

func (a *metastoresImpl) Unassign(ctx context.Context, request UnassignRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *metastoresImpl) Update(ctx context.Context, request UpdateMetastore) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *metastoresImpl) UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Providers API methods
type providersImpl struct {
	client *client.DatabricksClient
}

func (a *providersImpl) Create(ctx context.Context, request CreateProvider) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createProviderResponse)
	return &createProviderResponse, err
}

func (a *providersImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *providersImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getProviderResponse)
	return &getProviderResponse, err
}

func (a *providersImpl) List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *providersImpl) ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
	var listProviderSharesResponse ListProviderSharesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &listProviderSharesResponse)
	return &listProviderSharesResponse, err
}

func (a *providersImpl) Update(ctx context.Context, request UpdateProvider) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just RecipientActivation API methods
type recipientActivationImpl struct {
	client *client.DatabricksClient
}

func (a *recipientActivationImpl) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	err := a.client.Do(ctx, http.MethodGet, path, request, nil)
	return err
}

func (a *recipientActivationImpl) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	var retrieveTokenResponse RetrieveTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation/%v", request.ActivationUrl)
	err := a.client.Do(ctx, http.MethodGet, path, request, &retrieveTokenResponse)
	return &retrieveTokenResponse, err
}

// unexported type that holds implementations of just Recipients API methods
type recipientsImpl struct {
	client *client.DatabricksClient
}

func (a *recipientsImpl) Create(ctx context.Context, request CreateRecipient) (*CreateRecipientResponse, error) {
	var createRecipientResponse CreateRecipientResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createRecipientResponse)
	return &createRecipientResponse, err
}

func (a *recipientsImpl) Delete(ctx context.Context, request DeleteRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *recipientsImpl) Get(ctx context.Context, request GetRecipientRequest) (*GetRecipientResponse, error) {
	var getRecipientResponse GetRecipientResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getRecipientResponse)
	return &getRecipientResponse, err
}

func (a *recipientsImpl) List(ctx context.Context, request ListRecipientsRequest) (*ListRecipientsResponse, error) {
	var listRecipientsResponse ListRecipientsResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listRecipientsResponse)
	return &listRecipientsResponse, err
}

func (a *recipientsImpl) RotateToken(ctx context.Context, request RotateRecipientToken) (*RotateRecipientTokenResponse, error) {
	var rotateRecipientTokenResponse RotateRecipientTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
	err := a.client.Do(ctx, http.MethodPost, path, request, &rotateRecipientTokenResponse)
	return &rotateRecipientTokenResponse, err
}

func (a *recipientsImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	var getRecipientSharePermissionsResponse GetRecipientSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getRecipientSharePermissionsResponse)
	return &getRecipientSharePermissionsResponse, err
}

func (a *recipientsImpl) Update(ctx context.Context, request UpdateRecipient) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Schemas API methods
type schemasImpl struct {
	client *client.DatabricksClient
}

func (a *schemasImpl) Create(ctx context.Context, request CreateSchema) (*CreateSchemaResponse, error) {
	var createSchemaResponse CreateSchemaResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createSchemaResponse)
	return &createSchemaResponse, err
}

func (a *schemasImpl) Delete(ctx context.Context, request DeleteSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *schemasImpl) Get(ctx context.Context, request GetSchemaRequest) (*GetSchemaResponse, error) {
	var getSchemaResponse GetSchemaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getSchemaResponse)
	return &getSchemaResponse, err
}

func (a *schemasImpl) List(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

func (a *schemasImpl) Update(ctx context.Context, request UpdateSchema) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Shares API methods
type sharesImpl struct {
	client *client.DatabricksClient
}

func (a *sharesImpl) Create(ctx context.Context, request CreateShare) (*CreateShareResponse, error) {
	var createShareResponse CreateShareResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createShareResponse)
	return &createShareResponse, err
}

func (a *sharesImpl) Delete(ctx context.Context, request DeleteShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *sharesImpl) Get(ctx context.Context, request GetShareRequest) (*GetShareResponse, error) {
	var getShareResponse GetShareResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getShareResponse)
	return &getShareResponse, err
}

func (a *sharesImpl) List(ctx context.Context) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listSharesResponse)
	return &listSharesResponse, err
}

func (a *sharesImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetSharePermissionsResponse, error) {
	var getSharePermissionsResponse GetSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getSharePermissionsResponse)
	return &getSharePermissionsResponse, err
}

func (a *sharesImpl) Update(ctx context.Context, request UpdateShare) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *sharesImpl) UpdatePermissions(ctx context.Context, request UpdateSharePermissions) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just StorageCredentials API methods
type storageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *storageCredentialsImpl) Create(ctx context.Context, request CreateStorageCredential) (*CreateStorageCredentialResponse, error) {
	var createStorageCredentialResponse CreateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createStorageCredentialResponse)
	return &createStorageCredentialResponse, err
}

func (a *storageCredentialsImpl) Delete(ctx context.Context, request DeleteStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *storageCredentialsImpl) Get(ctx context.Context, request GetStorageCredentialRequest) (*GetStorageCredentialResponse, error) {
	var getStorageCredentialResponse GetStorageCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getStorageCredentialResponse)
	return &getStorageCredentialResponse, err
}

func (a *storageCredentialsImpl) List(ctx context.Context) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}

func (a *storageCredentialsImpl) Update(ctx context.Context, request UpdateStorageCredential) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Tables API methods
type tablesImpl struct {
	client *client.DatabricksClient
}

func (a *tablesImpl) Create(ctx context.Context, request CreateTable) (*CreateTableResponse, error) {
	var createTableResponse CreateTableResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createTableResponse)
	return &createTableResponse, err
}

func (a *tablesImpl) CreateStagingTable(ctx context.Context, request CreateStagingTable) (*CreateStagingTableResponse, error) {
	var createStagingTableResponse CreateStagingTableResponse
	path := "/api/2.1/unity-catalog/staging-tables"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createStagingTableResponse)
	return &createStagingTableResponse, err
}

func (a *tablesImpl) Delete(ctx context.Context, request DeleteTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *tablesImpl) Get(ctx context.Context, request GetTableRequest) (*GetTableResponse, error) {
	var getTableResponse GetTableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getTableResponse)
	return &getTableResponse, err
}

func (a *tablesImpl) List(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listTablesResponse)
	return &listTablesResponse, err
}

func (a *tablesImpl) TableSummaries(ctx context.Context, request TableSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}

func (a *tablesImpl) Update(ctx context.Context, request UpdateTable) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
