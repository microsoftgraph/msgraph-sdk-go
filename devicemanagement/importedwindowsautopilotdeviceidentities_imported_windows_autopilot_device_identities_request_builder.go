package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters list properties and relationships of the importedWindowsAutopilotDeviceIdentity objects.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters struct {
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
// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters
}
// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByImportedWindowsAutopilotDeviceIdentityId provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ByImportedWindowsAutopilotDeviceIdentityId(importedWindowsAutopilotDeviceIdentityId string)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if importedWindowsAutopilotDeviceIdentityId != "" {
        urlTplParams["importedWindowsAutopilotDeviceIdentity%2Did"] = importedWindowsAutopilotDeviceIdentityId
    }
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal instantiates a new ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    m := &ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder instantiates a new ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ImportedwindowsautopilotdeviceidentitiesCountRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Count()(*ImportedwindowsautopilotdeviceidentitiesCountRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the importedWindowsAutopilotDeviceIdentity objects.
// returns a ImportedWindowsAutopilotDeviceIdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-enrollment-importedwindowsautopilotdeviceidentity-list?view=graph-rest-1.0
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateImportedWindowsAutopilotDeviceIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable), nil
}
// ImportEscaped provides operations to call the import method.
// returns a *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ImportEscaped()(*ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new importedWindowsAutopilotDeviceIdentity object.
// returns a ImportedWindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-enrollment-importedwindowsautopilotdeviceidentity-create?view=graph-rest-1.0
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable), nil
}
// ToGetRequestInformation list properties and relationships of the importedWindowsAutopilotDeviceIdentity objects.
// returns a *RequestInformation when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new importedWindowsAutopilotDeviceIdentity object.
// returns a *RequestInformation when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) WithUrl(rawUrl string)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
