package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
type FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
type FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Application provides operations to manage the application property of the microsoft.graph.workbook entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookApplicationRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Application()(*FileStorageContainersItemDriveItemsItemWorkbookApplicationRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloseSession provides operations to call the closeSession method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookCloseSessionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) CloseSession()(*FileStorageContainersItemDriveItemsItemWorkbookCloseSessionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookCloseSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments provides operations to manage the comments property of the microsoft.graph.workbook entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookCommentsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Comments()(*FileStorageContainersItemDriveItemsItemWorkbookCommentsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateSession provides operations to call the createSession method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookCreateSessionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) CreateSession()(*FileStorageContainersItemDriveItemsItemWorkbookCreateSessionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookCreateSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property workbook for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Functions provides operations to manage the functions property of the microsoft.graph.workbook entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Functions()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
// returns a Workbookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable), nil
}
// Names provides operations to manage the names property of the microsoft.graph.workbook entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Names()(*FileStorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.workbook entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookOperationsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Operations()(*FileStorageContainersItemDriveItemsItemWorkbookOperationsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property workbook in storage
// returns a Workbookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable), nil
}
// RefreshSession provides operations to call the refreshSession method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookRefreshSessionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) RefreshSession()(*FileStorageContainersItemDriveItemsItemWorkbookRefreshSessionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookRefreshSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SessionInfoResourceWithKey provides operations to call the sessionInfoResource method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookSessionInfoResourceWithKeyRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) SessionInfoResourceWithKey(key *string)(*FileStorageContainersItemDriveItemsItemWorkbookSessionInfoResourceWithKeyRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookSessionInfoResourceWithKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, key)
}
// TableRowOperationResultWithKey provides operations to call the tableRowOperationResult method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) TableRowOperationResultWithKey(key *string)(*FileStorageContainersItemDriveItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, key)
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbook entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Tables()(*FileStorageContainersItemDriveItemsItemWorkbookTablesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property workbook for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property workbook in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheets provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookRequestBuilder) Worksheets()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
