package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetQueryParameters the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomExtensionStageSettingId provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) ByCustomExtensionStageSettingId(customExtensionStageSettingId string)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customExtensionStageSettingId != "" {
        urlTplParams["customExtensionStageSetting%2Did"] = customExtensionStageSettingId
    }
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionStageSettings{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) Count()(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// returns a CustomExtensionStageSettingCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomExtensionStageSettingCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingCollectionResponseable), nil
}
// Post create new navigation property to customExtensionStageSettings for identityGovernance
// returns a CustomExtensionStageSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomExtensionStageSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable), nil
}
// ToGetRequestInformation the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to customExtensionStageSettings for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomExtensionStageSettingable, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
