package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder provides operations to count the resources in the collection.
type FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderGetQueryParameters get the number of the resource
type FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder) {
    m := &FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}/baseTypes/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder instantiates a new FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemBasetypesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
