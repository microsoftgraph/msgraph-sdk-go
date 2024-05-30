package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
type WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetQueryParameters read properties and relationships of the windowsInformationProtectionAppLearningSummary object.
type WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal instantiates a new WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    m := &WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsInformationProtectionAppLearningSummaries/{windowsInformationProtectionAppLearningSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder instantiates a new WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a windowsInformationProtectionAppLearningSummary.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-wip-windowsinformationprotectionapplearningsummary-delete?view=graph-rest-1.0
func (m *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the windowsInformationProtectionAppLearningSummary object.
// returns a WindowsInformationProtectionAppLearningSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-wip-windowsinformationprotectionapplearningsummary-get?view=graph-rest-1.0
func (m *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable), nil
}
// Patch update the properties of a windowsInformationProtectionAppLearningSummary object.
// returns a WindowsInformationProtectionAppLearningSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-wip-windowsinformationprotectionapplearningsummary-update?view=graph-rest-1.0
func (m *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, requestConfiguration *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable), nil
}
// ToDeleteRequestInformation deletes a windowsInformationProtectionAppLearningSummary.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the windowsInformationProtectionAppLearningSummary object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a windowsInformationProtectionAppLearningSummary object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, requestConfiguration *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder when successful
func (m *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    return NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
