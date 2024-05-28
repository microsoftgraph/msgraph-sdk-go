package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
type DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
type DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters
}
// DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderInternal instantiates a new DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewDirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) {
    m := &DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleDefinitions/{unifiedRoleDefinition%2Did}/inheritsPermissionsFrom/{unifiedRoleDefinition%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder instantiates a new DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewDirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property inheritsPermissionsFrom for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property inheritsPermissionsFrom in roleManagement
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property inheritsPermissionsFrom for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
// returns a *RequestInformation when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property inheritsPermissionsFrom in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, requestConfiguration *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder when successful
func (m *DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder) {
    return NewDirectoryRoledefinitionsItemInheritspermissionsfromUnifiedRoleDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
