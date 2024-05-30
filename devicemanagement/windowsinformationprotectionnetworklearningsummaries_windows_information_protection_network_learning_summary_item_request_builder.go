package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetQueryParameters read properties and relationships of the windowsInformationProtectionNetworkLearningSummary object.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal instantiates a new WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    m := &WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/{windowsInformationProtectionNetworkLearningSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder instantiates a new WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a windowsInformationProtectionNetworkLearningSummary.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-wip-windowsinformationprotectionnetworklearningsummary-delete?view=graph-rest-1.0
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the windowsInformationProtectionNetworkLearningSummary object.
// returns a WindowsInformationProtectionNetworkLearningSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-wip-windowsinformationprotectionnetworklearningsummary-get?view=graph-rest-1.0
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionNetworkLearningSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionNetworkLearningSummaryable), nil
}
// Patch update the properties of a windowsInformationProtectionNetworkLearningSummary object.
// returns a WindowsInformationProtectionNetworkLearningSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-wip-windowsinformationprotectionnetworklearningsummary-update?view=graph-rest-1.0
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionNetworkLearningSummaryable, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionNetworkLearningSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionNetworkLearningSummaryable), nil
}
// ToDeleteRequestInformation deletes a windowsInformationProtectionNetworkLearningSummary.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the windowsInformationProtectionNetworkLearningSummary object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a windowsInformationProtectionNetworkLearningSummary object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionNetworkLearningSummaryable, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    return NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
