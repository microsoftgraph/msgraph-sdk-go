package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.termsAndConditions entity.
type TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderGetQueryParameters read properties and relationships of the termsAndConditionsAssignment object.
type TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderGetQueryParameters
}
// TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderInternal instantiates a new TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder and sets the default values.
func NewTermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) {
    m := &TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}/assignments/{termsAndConditionsAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder instantiates a new TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder and sets the default values.
func NewTermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a termsAndConditionsAssignment.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-companyterms-termsandconditionsassignment-delete?view=graph-rest-1.0
func (m *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the termsAndConditionsAssignment object.
// returns a TermsAndConditionsAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-companyterms-termsandconditionsassignment-get?view=graph-rest-1.0
func (m *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTermsAndConditionsAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsAssignmentable), nil
}
// Patch update the properties of a termsAndConditionsAssignment object.
// returns a TermsAndConditionsAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-companyterms-termsandconditionsassignment-update?view=graph-rest-1.0
func (m *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsAssignmentable, requestConfiguration *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTermsAndConditionsAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsAssignmentable), nil
}
// ToDeleteRequestInformation deletes a termsAndConditionsAssignment.
// returns a *RequestInformation when successful
func (m *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the termsAndConditionsAssignment object.
// returns a *RequestInformation when successful
func (m *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a termsAndConditionsAssignment object.
// returns a *RequestInformation when successful
func (m *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsAssignmentable, requestConfiguration *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder when successful
func (m *TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*TermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) {
    return NewTermsandconditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
