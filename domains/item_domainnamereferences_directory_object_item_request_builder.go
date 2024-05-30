package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemDomainnamereferencesDirectoryObjectItemRequestBuilder provides operations to manage the domainNameReferences property of the microsoft.graph.domain entity.
type ItemDomainnamereferencesDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDomainnamereferencesDirectoryObjectItemRequestBuilderGetQueryParameters the objects such as users and groups that reference the domain ID. Read-only, Nullable. Supports $expand and $filter by the OData type of objects returned. For example, /domains/{domainId}/domainNameReferences/microsoft.graph.user and /domains/{domainId}/domainNameReferences/microsoft.graph.group.
type ItemDomainnamereferencesDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDomainnamereferencesDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDomainnamereferencesDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDomainnamereferencesDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemDomainnamereferencesDirectoryObjectItemRequestBuilderInternal instantiates a new ItemDomainnamereferencesDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDomainnamereferencesDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDomainnamereferencesDirectoryObjectItemRequestBuilder) {
    m := &ItemDomainnamereferencesDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/domains/{domain%2Did}/domainNameReferences/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemDomainnamereferencesDirectoryObjectItemRequestBuilder instantiates a new ItemDomainnamereferencesDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDomainnamereferencesDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDomainnamereferencesDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDomainnamereferencesDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the objects such as users and groups that reference the domain ID. Read-only, Nullable. Supports $expand and $filter by the OData type of objects returned. For example, /domains/{domainId}/domainNameReferences/microsoft.graph.user and /domains/{domainId}/domainNameReferences/microsoft.graph.group.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDomainnamereferencesDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDomainnamereferencesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// ToGetRequestInformation the objects such as users and groups that reference the domain ID. Read-only, Nullable. Supports $expand and $filter by the OData type of objects returned. For example, /domains/{domainId}/domainNameReferences/microsoft.graph.user and /domains/{domainId}/domainNameReferences/microsoft.graph.group.
// returns a *RequestInformation when successful
func (m *ItemDomainnamereferencesDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDomainnamereferencesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDomainnamereferencesDirectoryObjectItemRequestBuilder when successful
func (m *ItemDomainnamereferencesDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*ItemDomainnamereferencesDirectoryObjectItemRequestBuilder) {
    return NewItemDomainnamereferencesDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
