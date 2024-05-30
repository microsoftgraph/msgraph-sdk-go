package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
type ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters device configuration states for this device.
type ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters
}
// ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal instantiates a new ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder and sets the default values.
func NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    m := &ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceConfigurationStates/{deviceConfigurationState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder instantiates a new ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder and sets the default values.
func NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceConfigurationStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get device configuration states for this device.
// returns a DeviceConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable), nil
}
// Patch update the navigation property deviceConfigurationStates in deviceManagement
// returns a DeviceConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, requestConfiguration *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable), nil
}
// ToDeleteRequestInformation delete navigation property deviceConfigurationStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation device configuration states for this device.
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceConfigurationStates in deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, requestConfiguration *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder when successful
func (m *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    return NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
