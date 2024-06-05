package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters read properties and relationships of the deviceEnrollmentConfiguration object.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *DeviceenrollmentconfigurationsItemAssignRequestBuilder when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Assign()(*DeviceenrollmentconfigurationsItemAssignRequestBuilder) {
    return NewDeviceenrollmentconfigurationsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceEnrollmentConfiguration entity.
// returns a *DeviceenrollmentconfigurationsItemAssignmentsRequestBuilder when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Assignments()(*DeviceenrollmentconfigurationsItemAssignmentsRequestBuilder) {
    return NewDeviceenrollmentconfigurationsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal instantiates a new DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    m := &DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder instantiates a new DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a deviceEnrollmentLimitConfiguration.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-deviceenrollmentlimitconfiguration-delete?view=graph-rest-1.0
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the deviceEnrollmentConfiguration object.
// returns a DeviceEnrollmentConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-deviceenrollmentconfiguration-get?view=graph-rest-1.0
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable), nil
}
// Patch update the properties of a deviceEnrollmentPlatformRestrictionsConfiguration object.
// returns a DeviceEnrollmentConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-deviceenrollmentplatformrestrictionsconfiguration-update?view=graph-rest-1.0
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable), nil
}
// SetPriority provides operations to call the setPriority method.
// returns a *DeviceenrollmentconfigurationsItemSetprioritySetPriorityRequestBuilder when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) SetPriority()(*DeviceenrollmentconfigurationsItemSetprioritySetPriorityRequestBuilder) {
    return NewDeviceenrollmentconfigurationsItemSetprioritySetPriorityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a deviceEnrollmentLimitConfiguration.
// returns a *RequestInformation when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the deviceEnrollmentConfiguration object.
// returns a *RequestInformation when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a deviceEnrollmentPlatformRestrictionsConfiguration object.
// returns a *RequestInformation when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    return NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
