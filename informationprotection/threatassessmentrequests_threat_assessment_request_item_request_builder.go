package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
type ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a specified threatAssessmentRequest object. A threat assessment request can be one of the following types:
type ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderGetQueryParameters
}
// ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderInternal instantiates a new ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder and sets the default values.
func NewThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) {
    m := &ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder instantiates a new ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder and sets the default values.
func NewThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property threatAssessmentRequests for informationProtection
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a specified threatAssessmentRequest object. A threat assessment request can be one of the following types:
// returns a ThreatAssessmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/threatassessmentrequest-get?view=graph-rest-1.0
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateThreatAssessmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable), nil
}
// Patch update the navigation property threatAssessmentRequests in informationProtection
// returns a ThreatAssessmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateThreatAssessmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable), nil
}
// Results provides operations to manage the results property of the microsoft.graph.threatAssessmentRequest entity.
// returns a *ThreatassessmentrequestsItemResultsRequestBuilder when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) Results()(*ThreatassessmentrequestsItemResultsRequestBuilder) {
    return NewThreatassessmentrequestsItemResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property threatAssessmentRequests for informationProtection
// returns a *RequestInformation when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a specified threatAssessmentRequest object. A threat assessment request can be one of the following types:
// returns a *RequestInformation when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property threatAssessmentRequests in informationProtection
// returns a *RequestInformation when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) WithUrl(rawUrl string)(*ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) {
    return NewThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
