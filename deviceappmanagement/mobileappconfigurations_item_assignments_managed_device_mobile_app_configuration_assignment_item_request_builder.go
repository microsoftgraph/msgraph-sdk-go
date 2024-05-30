package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
type MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetQueryParameters read properties and relationships of the managedDeviceMobileAppConfigurationAssignment object.
type MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetQueryParameters
}
// MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal instantiates a new MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    m := &MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}/assignments/{managedDeviceMobileAppConfigurationAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder instantiates a new MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a managedDeviceMobileAppConfigurationAssignment.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-manageddevicemobileappconfigurationassignment-delete?view=graph-rest-1.0
func (m *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the managedDeviceMobileAppConfigurationAssignment object.
// returns a ManagedDeviceMobileAppConfigurationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-manageddevicemobileappconfigurationassignment-get?view=graph-rest-1.0
func (m *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationAssignmentable), nil
}
// Patch update the properties of a managedDeviceMobileAppConfigurationAssignment object.
// returns a ManagedDeviceMobileAppConfigurationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-manageddevicemobileappconfigurationassignment-update?view=graph-rest-1.0
func (m *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationAssignmentable, requestConfiguration *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationAssignmentable), nil
}
// ToDeleteRequestInformation deletes a managedDeviceMobileAppConfigurationAssignment.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the managedDeviceMobileAppConfigurationAssignment object.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a managedDeviceMobileAppConfigurationAssignment object.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationAssignmentable, requestConfiguration *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder when successful
func (m *MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*MobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    return NewMobileappconfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
