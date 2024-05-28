package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
type ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetQueryParameters the device protection status. This property is read-only.
type ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal instantiates a new ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder and sets the default values.
func NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    m := &ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/windowsProtectionState{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder instantiates a new ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder and sets the default values.
func NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsProtectionState for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemManageddevicesItemWindowsprotectionstateDetectedmalwarestateDetectedMalwareStateRequestBuilder when successful
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) DetectedMalwareState()(*ItemManageddevicesItemWindowsprotectionstateDetectedmalwarestateDetectedMalwareStateRequestBuilder) {
    return NewItemManageddevicesItemWindowsprotectionstateDetectedmalwarestateDetectedMalwareStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the device protection status. This property is read-only.
// returns a WindowsProtectionStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, error) {
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
// Patch update the navigation property windowsProtectionState in users
// returns a WindowsProtectionStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, requestConfiguration *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, error) {
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
// ToDeleteRequestInformation delete navigation property windowsProtectionState for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsProtectionState in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsProtectionStateable, requestConfiguration *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder when successful
func (m *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    return NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
