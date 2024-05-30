package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryResourcenamespacesResourceNamespacesRequestBuilder provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
type DirectoryResourcenamespacesResourceNamespacesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryResourcenamespacesResourceNamespacesRequestBuilderGetQueryParameters get resourceNamespaces from roleManagement
type DirectoryResourcenamespacesResourceNamespacesRequestBuilderGetQueryParameters struct {
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
// DirectoryResourcenamespacesResourceNamespacesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryResourcenamespacesResourceNamespacesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryResourcenamespacesResourceNamespacesRequestBuilderGetQueryParameters
}
// DirectoryResourcenamespacesResourceNamespacesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryResourcenamespacesResourceNamespacesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRbacResourceNamespaceId provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder when successful
func (m *DirectoryResourcenamespacesResourceNamespacesRequestBuilder) ByUnifiedRbacResourceNamespaceId(unifiedRbacResourceNamespaceId string)(*DirectoryResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRbacResourceNamespaceId != "" {
        urlTplParams["unifiedRbacResourceNamespace%2Did"] = unifiedRbacResourceNamespaceId
    }
    return NewDirectoryResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryResourcenamespacesResourceNamespacesRequestBuilderInternal instantiates a new DirectoryResourcenamespacesResourceNamespacesRequestBuilder and sets the default values.
func NewDirectoryResourcenamespacesResourceNamespacesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryResourcenamespacesResourceNamespacesRequestBuilder) {
    m := &DirectoryResourcenamespacesResourceNamespacesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/resourceNamespaces{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDirectoryResourcenamespacesResourceNamespacesRequestBuilder instantiates a new DirectoryResourcenamespacesResourceNamespacesRequestBuilder and sets the default values.
func NewDirectoryResourcenamespacesResourceNamespacesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryResourcenamespacesResourceNamespacesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryResourcenamespacesResourceNamespacesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryResourcenamespacesCountRequestBuilder when successful
func (m *DirectoryResourcenamespacesResourceNamespacesRequestBuilder) Count()(*DirectoryResourcenamespacesCountRequestBuilder) {
    return NewDirectoryResourcenamespacesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get resourceNamespaces from roleManagement
// returns a UnifiedRbacResourceNamespaceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryResourcenamespacesResourceNamespacesRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryResourcenamespacesResourceNamespacesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRbacResourceNamespaceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRbacResourceNamespaceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRbacResourceNamespaceCollectionResponseable), nil
}
// Post create new navigation property to resourceNamespaces for roleManagement
// returns a UnifiedRbacResourceNamespaceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryResourcenamespacesResourceNamespacesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRbacResourceNamespaceable, requestConfiguration *DirectoryResourcenamespacesResourceNamespacesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRbacResourceNamespaceable), nil
}
// ToGetRequestInformation get resourceNamespaces from roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryResourcenamespacesResourceNamespacesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryResourcenamespacesResourceNamespacesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to resourceNamespaces for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryResourcenamespacesResourceNamespacesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRbacResourceNamespaceable, requestConfiguration *DirectoryResourcenamespacesResourceNamespacesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryResourcenamespacesResourceNamespacesRequestBuilder when successful
func (m *DirectoryResourcenamespacesResourceNamespacesRequestBuilder) WithUrl(rawUrl string)(*DirectoryResourcenamespacesResourceNamespacesRequestBuilder) {
    return NewDirectoryResourcenamespacesResourceNamespacesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
