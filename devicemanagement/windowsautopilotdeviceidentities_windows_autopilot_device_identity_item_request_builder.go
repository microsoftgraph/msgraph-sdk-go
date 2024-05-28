package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters read properties and relationships of the windowsAutopilotDeviceIdentity object.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignUserToDevice provides operations to call the assignUserToDevice method.
// returns a *WindowsautopilotdeviceidentitiesItemAssignusertodeviceAssignUserToDeviceRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) AssignUserToDevice()(*WindowsautopilotdeviceidentitiesItemAssignusertodeviceAssignUserToDeviceRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemAssignusertodeviceAssignUserToDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder instantiates a new WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a windowsAutopilotDeviceIdentity.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-enrollment-windowsautopilotdeviceidentity-delete?view=graph-rest-1.0
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the windowsAutopilotDeviceIdentity object.
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-enrollment-windowsautopilotdeviceidentity-get?view=graph-rest-1.0
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable), nil
}
// Patch update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable), nil
}
// ToDeleteRequestInformation deletes a windowsAutopilotDeviceIdentity.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the windowsAutopilotDeviceIdentity object.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnassignUserFromDevice provides operations to call the unassignUserFromDevice method.
// returns a *WindowsautopilotdeviceidentitiesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignUserFromDevice()(*WindowsautopilotdeviceidentitiesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateDeviceProperties provides operations to call the updateDeviceProperties method.
// returns a *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) UpdateDeviceProperties()(*WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
