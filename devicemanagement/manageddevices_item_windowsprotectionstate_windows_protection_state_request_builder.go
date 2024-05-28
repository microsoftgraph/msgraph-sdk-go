package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
type ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetQueryParameters the device protection status. This property is read-only.
type ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetQueryParameters
}
// ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal instantiates a new ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder and sets the default values.
func NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    m := &ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/windowsProtectionState{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder instantiates a new ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder and sets the default values.
func NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsProtectionState for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration)(error) {
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
// DetectedMalwareState provides operations to manage the detectedMalwareState property of the microsoft.graph.windowsProtectionState entity.
// returns a *ManageddevicesItemWindowsprotectionstateDetectedmalwarestateDetectedMalwareStateRequestBuilder when successful
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) DetectedMalwareState()(*ManageddevicesItemWindowsprotectionstateDetectedmalwarestateDetectedMalwareStateRequestBuilder) {
    return NewManageddevicesItemWindowsprotectionstateDetectedmalwarestateDetectedMalwareStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the device protection status. This property is read-only.
// returns a WindowsProtectionStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsProtectionStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable), nil
}
// Patch update the navigation property windowsProtectionState in deviceManagement
// returns a WindowsProtectionStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, requestConfiguration *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsProtectionStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable), nil
}
// ToDeleteRequestInformation delete navigation property windowsProtectionState for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the device protection status. This property is read-only.
// returns a *RequestInformation when successful
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsProtectionState in deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, requestConfiguration *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder when successful
func (m *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    return NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
