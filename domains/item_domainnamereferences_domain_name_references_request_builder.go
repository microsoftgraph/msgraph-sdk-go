package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemDomainnamereferencesDomainNameReferencesRequestBuilder provides operations to manage the domainNameReferences property of the microsoft.graph.domain entity.
type ItemDomainnamereferencesDomainNameReferencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDomainnamereferencesDomainNameReferencesRequestBuilderGetQueryParameters retrieve a list of directoryObject with a reference to the domain. The returned list will contain all directory objects that have a dependency on the domain.
type ItemDomainnamereferencesDomainNameReferencesRequestBuilderGetQueryParameters struct {
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
// ItemDomainnamereferencesDomainNameReferencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDomainnamereferencesDomainNameReferencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDomainnamereferencesDomainNameReferencesRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the domainNameReferences property of the microsoft.graph.domain entity.
// returns a *ItemDomainnamereferencesDirectoryObjectItemRequestBuilder when successful
func (m *ItemDomainnamereferencesDomainNameReferencesRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemDomainnamereferencesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemDomainnamereferencesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDomainnamereferencesDomainNameReferencesRequestBuilderInternal instantiates a new ItemDomainnamereferencesDomainNameReferencesRequestBuilder and sets the default values.
func NewItemDomainnamereferencesDomainNameReferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDomainnamereferencesDomainNameReferencesRequestBuilder) {
    m := &ItemDomainnamereferencesDomainNameReferencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/domains/{domain%2Did}/domainNameReferences{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemDomainnamereferencesDomainNameReferencesRequestBuilder instantiates a new ItemDomainnamereferencesDomainNameReferencesRequestBuilder and sets the default values.
func NewItemDomainnamereferencesDomainNameReferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDomainnamereferencesDomainNameReferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDomainnamereferencesDomainNameReferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemDomainnamereferencesCountRequestBuilder when successful
func (m *ItemDomainnamereferencesDomainNameReferencesRequestBuilder) Count()(*ItemDomainnamereferencesCountRequestBuilder) {
    return NewItemDomainnamereferencesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of directoryObject with a reference to the domain. The returned list will contain all directory objects that have a dependency on the domain.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/domain-list-domainnamereferences?view=graph-rest-1.0
func (m *ItemDomainnamereferencesDomainNameReferencesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDomainnamereferencesDomainNameReferencesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// ToGetRequestInformation retrieve a list of directoryObject with a reference to the domain. The returned list will contain all directory objects that have a dependency on the domain.
// returns a *RequestInformation when successful
func (m *ItemDomainnamereferencesDomainNameReferencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDomainnamereferencesDomainNameReferencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDomainnamereferencesDomainNameReferencesRequestBuilder when successful
func (m *ItemDomainnamereferencesDomainNameReferencesRequestBuilder) WithUrl(rawUrl string)(*ItemDomainnamereferencesDomainNameReferencesRequestBuilder) {
    return NewItemDomainnamereferencesDomainNameReferencesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
