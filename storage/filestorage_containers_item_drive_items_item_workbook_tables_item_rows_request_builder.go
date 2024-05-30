package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderGetQueryParameters represents a collection of all the rows in the table. Read-only.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Add provides operations to call the add method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsAddRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) Add()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsAddRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByWorkbookTableRowId provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsWorkbookTableRowItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) ByWorkbookTableRowId(workbookTableRowId string)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsWorkbookTableRowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workbookTableRowId != "" {
        urlTplParams["workbookTableRow%2Did"] = workbookTableRowId
    }
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsWorkbookTableRowItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/rows{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to call the count method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) Count()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents a collection of all the rows in the table. Read-only.
// returns a WorkbookTableRowCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableRowCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowCollectionResponseable), nil
}
// ItemAtWithIndex provides operations to call the itemAt method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) ItemAtWithIndex(index *int32)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsItematwithindexItemAtWithIndexRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsItematwithindexItemAtWithIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, index)
}
// Post create new navigation property to rows for storage
// returns a WorkbookTableRowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableRowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowable), nil
}
// ToGetRequestInformation represents a collection of all the rows in the table. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to rows for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
