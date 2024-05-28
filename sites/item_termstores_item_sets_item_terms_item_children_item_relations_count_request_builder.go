package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder provides operations to count the resources in the collection.
type ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderGetQueryParameters
}
// NewItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderInternal instantiates a new ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder and sets the default values.
func NewItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder) {
    m := &ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/terms/{term%2Did}/children/{term%2Did1}/relations/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder instantiates a new ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder and sets the default values.
func NewItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder when successful
func (m *ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder) WithUrl(rawUrl string)(*ItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder) {
    return NewItemTermstoresItemSetsItemTermsItemChildrenItemRelationsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
