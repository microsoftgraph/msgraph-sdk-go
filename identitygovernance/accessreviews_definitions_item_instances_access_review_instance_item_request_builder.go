package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReviewScheduleDefinition entity.
type AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters read the properties and relationships of an accessReviewInstance object.
type AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters
}
// AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
// returns a *AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyDecisions provides operations to call the applyDecisions method.
// returns a *AccessreviewsDefinitionsItemInstancesItemApplydecisionsApplyDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*AccessreviewsDefinitionsItemInstancesItemApplydecisionsApplyDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemApplydecisionsApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
// returns a *AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
// returns a *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
// returns a *AccessreviewsDefinitionsItemInstancesItemDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Decisions()(*AccessreviewsDefinitionsItemInstancesItemDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property instances for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an accessReviewInstance object.
// returns a AccessReviewInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-get?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable), nil
}
// Patch update the properties of an accessReviewInstance object. Only the reviewers and fallbackReviewers properties can be updated but the scope property is also required in the request body. You can only add reviewers to the fallbackReviewers property but can't remove existing fallbackReviewers. To update an accessReviewInstance, it's status must be InProgress.
// returns a AccessReviewInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-update?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, requestConfiguration *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable), nil
}
// ResetDecisions provides operations to call the resetDecisions method.
// returns a *AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ResetDecisions()(*AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendReminder provides operations to call the sendReminder method.
// returns a *AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) SendReminder()(*AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
// returns a *AccessreviewsDefinitionsItemInstancesItemStagesRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Stages()(*AccessreviewsDefinitionsItemInstancesItemStagesRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemStagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *AccessreviewsDefinitionsItemInstancesItemStopRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Stop()(*AccessreviewsDefinitionsItemInstancesItemStopRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property instances for identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an accessReviewInstance object.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an accessReviewInstance object. Only the reviewers and fallbackReviewers properties can be updated but the scope property is also required in the request body. You can only add reviewers to the fallbackReviewers property but can't remove existing fallbackReviewers. To update an accessReviewInstance, it's status must be InProgress.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, requestConfiguration *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
