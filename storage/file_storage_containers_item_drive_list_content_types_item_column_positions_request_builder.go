package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
type FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderGetQueryParameters column order information in a content type.
type FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderGetQueryParameters struct {
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
// FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderGetQueryParameters
}
// ByColumnDefinitionId provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
// returns a *FileStorageContainersItemDriveListContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) ByColumnDefinitionId(columnDefinitionId string)(*FileStorageContainersItemDriveListContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if columnDefinitionId != "" {
        urlTplParams["columnDefinition%2Did"] = columnDefinitionId
    }
    return NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderInternal instantiates a new FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) {
    m := &FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}/columnPositions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder instantiates a new FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FileStorageContainersItemDriveListContentTypesItemColumnPositionsCountRequestBuilder when successful
func (m *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) Count()(*FileStorageContainersItemDriveListContentTypesItemColumnPositionsCountRequestBuilder) {
    return NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get column order information in a content type.
// returns a ColumnDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateColumnDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionCollectionResponseable), nil
}
// ToGetRequestInformation column order information in a content type.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder when successful
func (m *FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) {
    return NewFileStorageContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
