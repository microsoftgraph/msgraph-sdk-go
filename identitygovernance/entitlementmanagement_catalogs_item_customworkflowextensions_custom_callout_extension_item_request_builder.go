package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder provides operations to manage the customWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
type EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetQueryParameters read the properties and relationships of an accessPackageAssignmentRequestWorkflowExtension object.
type EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderInternal instantiates a new EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) {
    m := &EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/customWorkflowExtensions/{customCalloutExtension%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder instantiates a new EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an accessPackageAssignmentRequestWorkflowExtension object. The custom workflow extension must first be removed from any associated policies before it can be deleted. Follow these steps to remove the custom workflow extension from any associated policies:
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequestworkflowextension-delete?view=graph-rest-1.0
func (m *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an accessPackageAssignmentRequestWorkflowExtension object.
// returns a CustomCalloutExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequestworkflowextension-get?view=graph-rest-1.0
func (m *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, error) {
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
// Patch update the properties of an accessPackageAssignmentWorkflowExtension object.
// returns a CustomCalloutExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentworkflowextension-update?view=graph-rest-1.0
func (m *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, requestConfiguration *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, error) {
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
// ToDeleteRequestInformation delete an accessPackageAssignmentRequestWorkflowExtension object. The custom workflow extension must first be removed from any associated policies before it can be deleted. Follow these steps to remove the custom workflow extension from any associated policies:
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an accessPackageAssignmentRequestWorkflowExtension object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an accessPackageAssignmentWorkflowExtension object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomCalloutExtensionable, requestConfiguration *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder when successful
func (m *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemCustomworkflowextensionsCustomCalloutExtensionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
