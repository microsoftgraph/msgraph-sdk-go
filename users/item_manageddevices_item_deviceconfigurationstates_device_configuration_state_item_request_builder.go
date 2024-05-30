package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
type ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters device configuration states for this device.
type ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal instantiates a new ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    m := &ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/deviceConfigurationStates/{deviceConfigurationState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder instantiates a new ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceConfigurationStates for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, error) {
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
// Patch update the navigation property deviceConfigurationStates in users
// returns a DeviceConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, requestConfiguration *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, error) {
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
// ToDeleteRequestInformation delete navigation property deviceConfigurationStates for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceConfigurationStates in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationStateable, requestConfiguration *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder when successful
func (m *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    return NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
