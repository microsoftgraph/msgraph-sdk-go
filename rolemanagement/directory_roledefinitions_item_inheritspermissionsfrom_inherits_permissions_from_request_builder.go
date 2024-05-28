package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
type DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderGetQueryParameters read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
type DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderGetQueryParameters struct {
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
// DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderGetQueryParameters
}
// DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleDefinitionId1 provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
// returns a *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) ByUnifiedRoleDefinitionId1(unifiedRoleDefinitionId1 string)(*DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleDefinitionId1 != "" {
        urlTplParams["unifiedRoleDefinition%2Did1"] = unifiedRoleDefinitionId1
    }
    return NewDirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderInternal instantiates a new DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder and sets the default values.
func NewDirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) {
    m := &DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleDefinitions/{unifiedRoleDefinition%2Did}/inheritsPermissionsFrom{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder instantiates a new DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder and sets the default values.
func NewDirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryRoledefinitionsItemInheritspermissionsfromCountRequestBuilder when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) Count()(*DirectoryRoledefinitionsItemInheritspermissionsfromCountRequestBuilder) {
    return NewDirectoryRoledefinitionsItemInheritspermissionsfromCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
// returns a UnifiedRoleDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionCollectionResponseable, error) {
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
// Post create new navigation property to inheritsPermissionsFrom for roleManagement
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, error) {
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
// ToGetRequestInformation read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
// returns a *RequestInformation when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to inheritsPermissionsFrom for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder) {
    return NewDirectoryRoledefinitionsItemInheritspermissionsfromInheritsPermissionsFromRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
