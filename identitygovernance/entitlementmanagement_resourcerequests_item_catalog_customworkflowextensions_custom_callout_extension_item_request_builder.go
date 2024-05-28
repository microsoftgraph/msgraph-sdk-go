package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder provides operations to manage the customWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
type EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetQueryParameters get customWorkflowExtensions from identityGovernance
type EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderInternal instantiates a new EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) {
    m := &EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resourceRequests/{accessPackageResourceRequest%2Did}/catalog/customWorkflowExtensions/{customCalloutExtension%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder instantiates a new EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property customWorkflowExtensions for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get customWorkflowExtensions from identityGovernance
// returns a CustomCalloutExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomCalloutExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable), nil
}
// Patch update the navigation property customWorkflowExtensions in identityGovernance
// returns a CustomCalloutExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomCalloutExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable), nil
}
// ToDeleteRequestInformation delete navigation property customWorkflowExtensions for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get customWorkflowExtensions from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property customWorkflowExtensions in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder when successful
func (m *EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) {
    return NewEntitlementmanagementResourcerequestsItemCatalogCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
