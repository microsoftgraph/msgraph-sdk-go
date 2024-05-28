package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder provides operations to manage the softwareUpdateStatusSummary property of the microsoft.graph.deviceManagement entity.
type SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters read properties and relationships of the softwareUpdateStatusSummary object.
type SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters
}
// NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderInternal instantiates a new SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder and sets the default values.
func NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) {
    m := &SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/softwareUpdateStatusSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder instantiates a new SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder and sets the default values.
func NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read properties and relationships of the softwareUpdateStatusSummary object.
// returns a SoftwareUpdateStatusSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-softwareupdatestatussummary-get?view=graph-rest-1.0
func (m *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SoftwareUpdateStatusSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSoftwareUpdateStatusSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SoftwareUpdateStatusSummaryable), nil
}
// ToGetRequestInformation read properties and relationships of the softwareUpdateStatusSummary object.
// returns a *RequestInformation when successful
func (m *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder when successful
func (m *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) WithUrl(rawUrl string)(*SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) {
    return NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
