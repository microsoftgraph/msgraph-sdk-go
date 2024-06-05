package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
type AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters read the properties and relationships of an accessReviewInstanceDecisionItem object.
type AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters
}
// AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property decisions for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an accessReviewInstanceDecisionItem object.
// returns a AccessReviewInstanceDecisionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstancedecisionitem-get?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceDecisionItemable), nil
}
// Insights provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
// returns a *AccessreviewsDefinitionsItemInstancesItemDecisionsItemInsightsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Insights()(*AccessreviewsDefinitionsItemInstancesItemDecisionsItemInsightsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property decisions in identityGovernance
// returns a AccessReviewInstanceDecisionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceDecisionItemable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceDecisionItemable), nil
}
// ToDeleteRequestInformation delete navigation property decisions for identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an accessReviewInstanceDecisionItem object.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property decisions in identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceDecisionItemable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
