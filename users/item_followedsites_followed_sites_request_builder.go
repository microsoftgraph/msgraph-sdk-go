package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemFollowedsitesFollowedSitesRequestBuilder provides operations to manage the followedSites property of the microsoft.graph.user entity.
type ItemFollowedsitesFollowedSitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFollowedsitesFollowedSitesRequestBuilderGetQueryParameters get followedSites from users
type ItemFollowedsitesFollowedSitesRequestBuilderGetQueryParameters struct {
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
// ItemFollowedsitesFollowedSitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFollowedsitesFollowedSitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFollowedsitesFollowedSitesRequestBuilderGetQueryParameters
}
// Add provides operations to call the add method.
// returns a *ItemFollowedsitesAddRequestBuilder when successful
func (m *ItemFollowedsitesFollowedSitesRequestBuilder) Add()(*ItemFollowedsitesAddRequestBuilder) {
    return NewItemFollowedsitesAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BySiteId provides operations to manage the followedSites property of the microsoft.graph.user entity.
// returns a *ItemFollowedsitesSiteItemRequestBuilder when successful
func (m *ItemFollowedsitesFollowedSitesRequestBuilder) BySiteId(siteId string)(*ItemFollowedsitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if siteId != "" {
        urlTplParams["site%2Did"] = siteId
    }
    return NewItemFollowedsitesSiteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemFollowedsitesFollowedSitesRequestBuilderInternal instantiates a new ItemFollowedsitesFollowedSitesRequestBuilder and sets the default values.
func NewItemFollowedsitesFollowedSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowedsitesFollowedSitesRequestBuilder) {
    m := &ItemFollowedsitesFollowedSitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/followedSites{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemFollowedsitesFollowedSitesRequestBuilder instantiates a new ItemFollowedsitesFollowedSitesRequestBuilder and sets the default values.
func NewItemFollowedsitesFollowedSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowedsitesFollowedSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFollowedsitesFollowedSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemFollowedsitesCountRequestBuilder when successful
func (m *ItemFollowedsitesFollowedSitesRequestBuilder) Count()(*ItemFollowedsitesCountRequestBuilder) {
    return NewItemFollowedsitesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get followedSites from users
// returns a SiteCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFollowedsitesFollowedSitesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFollowedsitesFollowedSitesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SiteCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Remove provides operations to call the remove method.
// returns a *ItemFollowedsitesRemoveRequestBuilder when successful
func (m *ItemFollowedsitesFollowedSitesRequestBuilder) Remove()(*ItemFollowedsitesRemoveRequestBuilder) {
    return NewItemFollowedsitesRemoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get followedSites from users
// returns a *RequestInformation when successful
func (m *ItemFollowedsitesFollowedSitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFollowedsitesFollowedSitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFollowedsitesFollowedSitesRequestBuilder when successful
func (m *ItemFollowedsitesFollowedSitesRequestBuilder) WithUrl(rawUrl string)(*ItemFollowedsitesFollowedSitesRequestBuilder) {
    return NewItemFollowedsitesFollowedSitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
