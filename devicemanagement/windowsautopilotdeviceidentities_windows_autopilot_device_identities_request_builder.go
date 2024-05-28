package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters list properties and relationships of the windowsAutopilotDeviceIdentity objects.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsAutopilotDeviceIdentityId provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) ByWindowsAutopilotDeviceIdentityId(windowsAutopilotDeviceIdentityId string)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsAutopilotDeviceIdentityId != "" {
        urlTplParams["windowsAutopilotDeviceIdentity%2Did"] = windowsAutopilotDeviceIdentityId
    }
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderInternal instantiates a new WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    m := &WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder instantiates a new WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsautopilotdeviceidentitiesCountRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) Count()(*WindowsautopilotdeviceidentitiesCountRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the windowsAutopilotDeviceIdentity objects.
// returns a WindowsAutopilotDeviceIdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-enrollment-windowsautopilotdeviceidentity-list?view=graph-rest-1.0
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsAutopilotDeviceIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityCollectionResponseable), nil
}
// Post create a new windowsAutopilotDeviceIdentity object.
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-enrollment-windowsautopilotdeviceidentity-create?view=graph-rest-1.0
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation list properties and relationships of the windowsAutopilotDeviceIdentity objects.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new windowsAutopilotDeviceIdentity object.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
