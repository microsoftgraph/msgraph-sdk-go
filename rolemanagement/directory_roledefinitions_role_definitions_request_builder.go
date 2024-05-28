package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoledefinitionsRoleDefinitionsRequestBuilder provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
type DirectoryRoledefinitionsRoleDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoledefinitionsRoleDefinitionsRequestBuilderGetQueryParameters get a list of unifiedRoleDefinition objects for the provider. The following RBAC providers are currently supported:- directory (Microsoft Entra ID)- entitlement management (Microsoft Entra Entitlement Management)
type DirectoryRoledefinitionsRoleDefinitionsRequestBuilderGetQueryParameters struct {
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
// DirectoryRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoledefinitionsRoleDefinitionsRequestBuilderGetQueryParameters
}
// DirectoryRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleDefinitionId provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryRoledefinitionsUnifiedRoleDefinitionItemRequestBuilder when successful
func (m *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) ByUnifiedRoleDefinitionId(unifiedRoleDefinitionId string)(*DirectoryRoledefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleDefinitionId != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = unifiedRoleDefinitionId
    }
    return NewDirectoryRoledefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoledefinitionsRoleDefinitionsRequestBuilderInternal instantiates a new DirectoryRoledefinitionsRoleDefinitionsRequestBuilder and sets the default values.
func NewDirectoryRoledefinitionsRoleDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) {
    m := &DirectoryRoledefinitionsRoleDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDirectoryRoledefinitionsRoleDefinitionsRequestBuilder instantiates a new DirectoryRoledefinitionsRoleDefinitionsRequestBuilder and sets the default values.
func NewDirectoryRoledefinitionsRoleDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoledefinitionsRoleDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryRoledefinitionsCountRequestBuilder when successful
func (m *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) Count()(*DirectoryRoledefinitionsCountRequestBuilder) {
    return NewDirectoryRoledefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of unifiedRoleDefinition objects for the provider. The following RBAC providers are currently supported:- directory (Microsoft Entra ID)- entitlement management (Microsoft Entra Entitlement Management)
// returns a UnifiedRoleDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/rbacapplication-list-roledefinitions?view=graph-rest-1.0
func (m *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionCollectionResponseable), nil
}
// Post create a new custom unifiedRoleDefinition object. This feature requires a Microsoft Entra ID P1 or P2 license.
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/rbacapplication-post-roledefinitions?view=graph-rest-1.0
func (m *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *DirectoryRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable), nil
}
// ToGetRequestInformation get a list of unifiedRoleDefinition objects for the provider. The following RBAC providers are currently supported:- directory (Microsoft Entra ID)- entitlement management (Microsoft Entra Entitlement Management)
// returns a *RequestInformation when successful
func (m *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new custom unifiedRoleDefinition object. This feature requires a Microsoft Entra ID P1 or P2 license.
// returns a *RequestInformation when successful
func (m *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *DirectoryRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder when successful
func (m *DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoledefinitionsRoleDefinitionsRequestBuilder) {
    return NewDirectoryRoledefinitionsRoleDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
