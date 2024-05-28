package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
type FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
type FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Application provides operations to manage the application property of the microsoft.graph.workbook entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookApplicationRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Application()(*FilestorageContainersItemDriveItemsItemWorkbookApplicationRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloseSession provides operations to call the closeSession method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookClosesessionCloseSessionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) CloseSession()(*FilestorageContainersItemDriveItemsItemWorkbookClosesessionCloseSessionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookClosesessionCloseSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments provides operations to manage the comments property of the microsoft.graph.workbook entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookCommentsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Comments()(*FilestorageContainersItemDriveItemsItemWorkbookCommentsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateSession provides operations to call the createSession method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookCreatesessionCreateSessionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) CreateSession()(*FilestorageContainersItemDriveItemsItemWorkbookCreatesessionCreateSessionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookCreatesessionCreateSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property workbook for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Functions()(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
// returns a Workbookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Names()(*FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.workbook entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookOperationsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Operations()(*FilestorageContainersItemDriveItemsItemWorkbookOperationsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property workbook in storage
// returns a Workbookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookRefreshsessionRefreshSessionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) RefreshSession()(*FilestorageContainersItemDriveItemsItemWorkbookRefreshsessionRefreshSessionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookRefreshsessionRefreshSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SessionInfoResourceWithKey provides operations to call the sessionInfoResource method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) SessionInfoResourceWithKey(key *string)(*FilestorageContainersItemDriveItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, key)
}
// TableRowOperationResultWithKey provides operations to call the tableRowOperationResult method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablerowoperationresultwithkeyTableRowOperationResultWithKeyRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) TableRowOperationResultWithKey(key *string)(*FilestorageContainersItemDriveItemsItemWorkbookTablerowoperationresultwithkeyTableRowOperationResultWithKeyRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablerowoperationresultwithkeyTableRowOperationResultWithKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, key)
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbook entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Tables()(*FilestorageContainersItemDriveItemsItemWorkbookTablesRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property workbook for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheets provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookRequestBuilder) Worksheets()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
