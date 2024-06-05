package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceconfigurationsDeviceConfigurationItemRequestBuilder provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceconfigurationsDeviceConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceconfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters read properties and relationships of the iosCertificateProfile object.
type DeviceconfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceconfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceconfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *DeviceconfigurationsItemAssignRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) Assign()(*DeviceconfigurationsItemAssignRequestBuilder) {
    return NewDeviceconfigurationsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceConfiguration entity.
// returns a *DeviceconfigurationsItemAssignmentsRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) Assignments()(*DeviceconfigurationsItemAssignmentsRequestBuilder) {
    return NewDeviceconfigurationsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceconfigurationsDeviceConfigurationItemRequestBuilderInternal instantiates a new DeviceconfigurationsDeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceconfigurationsDeviceConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsDeviceConfigurationItemRequestBuilder) {
    m := &DeviceconfigurationsDeviceConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationsDeviceConfigurationItemRequestBuilder instantiates a new DeviceconfigurationsDeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceconfigurationsDeviceConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsDeviceConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsDeviceConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a windows10GeneralConfiguration.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-windows10generalconfiguration-delete?view=graph-rest-1.0
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceconfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceConfiguration entity.
// returns a *DeviceconfigurationsItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) DeviceSettingStateSummaries()(*DeviceconfigurationsItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilder) {
    return NewDeviceconfigurationsItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.deviceConfiguration entity.
// returns a *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatuses()(*DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewDeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatusOverview provides operations to manage the deviceStatusOverview property of the microsoft.graph.deviceConfiguration entity.
// returns a *DeviceconfigurationsItemDevicestatusoverviewDeviceStatusOverviewRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatusOverview()(*DeviceconfigurationsItemDevicestatusoverviewDeviceStatusOverviewRequestBuilder) {
    return NewDeviceconfigurationsItemDevicestatusoverviewDeviceStatusOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the iosCertificateProfile object.
// returns a DeviceConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-ioscertificateprofile-get?view=graph-rest-1.0
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationable), nil
}
// GetOmaSettingPlainTextValueWithSecretReferenceValueId provides operations to call the getOmaSettingPlainTextValue method.
// returns a *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) GetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return NewDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, secretReferenceValueId)
}
// Patch update the properties of a windowsUpdateForBusinessConfiguration object.
// returns a DeviceConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-windowsupdateforbusinessconfiguration-update?view=graph-rest-1.0
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationable, requestConfiguration *DeviceconfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationable), nil
}
// ToDeleteRequestInformation deletes a windows10GeneralConfiguration.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the iosCertificateProfile object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a windowsUpdateForBusinessConfiguration object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationable, requestConfiguration *DeviceconfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.deviceConfiguration entity.
// returns a *DeviceconfigurationsItemUserstatusesUserStatusesRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) UserStatuses()(*DeviceconfigurationsItemUserstatusesUserStatusesRequestBuilder) {
    return NewDeviceconfigurationsItemUserstatusesUserStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserStatusOverview provides operations to manage the userStatusOverview property of the microsoft.graph.deviceConfiguration entity.
// returns a *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) UserStatusOverview()(*DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) {
    return NewDeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceconfigurationsDeviceConfigurationItemRequestBuilder when successful
func (m *DeviceconfigurationsDeviceConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsDeviceConfigurationItemRequestBuilder) {
    return NewDeviceconfigurationsDeviceConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
