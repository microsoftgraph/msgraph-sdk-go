package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder provides operations to manage the names property of the microsoft.graph.workbook entity.
type FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderGetQueryParameters represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
type FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderGetQueryParameters struct {
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
// FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Add provides operations to call the add method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookNamesAddRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) Add()(*FilestorageContainersItemDriveItemsItemWorkbookNamesAddRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookNamesAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddFormulaLocal provides operations to call the addFormulaLocal method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookNamesAddformulalocalAddFormulaLocalRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) AddFormulaLocal()(*FilestorageContainersItemDriveItemsItemWorkbookNamesAddformulalocalAddFormulaLocalRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookNamesAddformulalocalAddFormulaLocalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByWorkbookNamedItemId provides operations to manage the names property of the microsoft.graph.workbook entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookNamesWorkbookNamedItemItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) ByWorkbookNamedItemId(workbookNamedItemId string)(*FilestorageContainersItemDriveItemsItemWorkbookNamesWorkbookNamedItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workbookNamedItemId != "" {
        urlTplParams["workbookNamedItem%2Did"] = workbookNamedItemId
    }
    return NewFilestorageContainersItemDriveItemsItemWorkbookNamesWorkbookNamedItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/names{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookNamesCountRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) Count()(*FilestorageContainersItemDriveItemsItemWorkbookNamesCountRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookNamesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
// returns a WorkbookNamedItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookNamedItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookNamedItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookNamedItemCollectionResponseable), nil
}
// Post create new navigation property to names for storage
// returns a WorkbookNamedItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookNamedItemable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookNamedItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookNamedItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookNamedItemable), nil
}
// ToGetRequestInformation represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to names for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookNamedItemable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookNamesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
