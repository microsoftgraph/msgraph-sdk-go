package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder provides operations to manage the sites property of the microsoft.graph.site entity.
type ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetQueryParameters get a collection of subsites defined for a [site][].
type ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetQueryParameters struct {
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
// ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetQueryParameters
}
// NewItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderInternal instantiates a new SitesRequestBuilder and sets the default values.
func NewItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) {
    m := &ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/sites{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder instantiates a new SitesRequestBuilder and sets the default values.
func NewItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a collection of subsites defined for a [site][].
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-list-subsites?view=graph-rest-1.0
func (m *ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SiteCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SiteCollectionResponseable), nil
}
// ToGetRequestInformation get a collection of subsites defined for a [site][].
func (m *ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) WithUrl(rawUrl string)(*ItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) {
    return NewItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
