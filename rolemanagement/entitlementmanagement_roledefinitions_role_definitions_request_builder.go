package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderGetQueryParameters get a list of unifiedRoleDefinition objects for the provider. The following RBAC providers are currently supported:- directory (Microsoft Entra ID)- entitlement management (Microsoft Entra Entitlement Management)
type EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleDefinitionId provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoledefinitionsUnifiedRoleDefinitionItemRequestBuilder when successful
func (m *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) ByUnifiedRoleDefinitionId(unifiedRoleDefinitionId string)(*EntitlementmanagementRoledefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleDefinitionId != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = unifiedRoleDefinitionId
    }
    return NewEntitlementmanagementRoledefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderInternal instantiates a new EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder and sets the default values.
func NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) {
    m := &EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder instantiates a new EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder and sets the default values.
func NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementRoledefinitionsCountRequestBuilder when successful
func (m *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) Count()(*EntitlementmanagementRoledefinitionsCountRequestBuilder) {
    return NewEntitlementmanagementRoledefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of unifiedRoleDefinition objects for the provider. The following RBAC providers are currently supported:- directory (Microsoft Entra ID)- entitlement management (Microsoft Entra Entitlement Management)
// returns a UnifiedRoleDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/rbacapplication-list-roledefinitions?view=graph-rest-1.0
func (m *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionCollectionResponseable, error) {
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
// Post create new navigation property to roleDefinitions for roleManagement
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, error) {
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
func (m *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleDefinitions for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder when successful
func (m *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) {
    return NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
