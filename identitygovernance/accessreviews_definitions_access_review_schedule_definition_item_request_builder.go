package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder provides operations to manage the definitions property of the microsoft.graph.accessReviewSet entity.
type AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters read the properties and relationships of an accessReviewScheduleDefinition object. To retrieve the instances of the access review series, use the list accessReviewInstance API.
type AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters
}
// AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderInternal instantiates a new AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) {
    m := &AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder instantiates a new AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes an accessReviewScheduleDefinition object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewscheduledefinition-delete?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an accessReviewScheduleDefinition object. To retrieve the instances of the access review series, use the list accessReviewInstance API.
// returns a AccessReviewScheduleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewscheduledefinition-get?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewScheduleDefinitionable), nil
}
// Instances provides operations to manage the instances property of the microsoft.graph.accessReviewScheduleDefinition entity.
// returns a *AccessreviewsDefinitionsItemInstancesRequestBuilder when successful
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Instances()(*AccessreviewsDefinitionsItemInstancesRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put update an existing accessReviewScheduleDefinition object to change one or more of its properties.
// returns a AccessReviewScheduleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewscheduledefinition-update?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewScheduleDefinitionable, requestConfiguration *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewScheduleDefinitionable), nil
}
// Stop provides operations to call the stop method.
// returns a *AccessreviewsDefinitionsItemStopRequestBuilder when successful
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Stop()(*AccessreviewsDefinitionsItemStopRequestBuilder) {
    return NewAccessreviewsDefinitionsItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes an accessReviewScheduleDefinition object.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an accessReviewScheduleDefinition object. To retrieve the instances of the access review series, use the list accessReviewInstance API.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update an existing accessReviewScheduleDefinition object to change one or more of its properties.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewScheduleDefinitionable, requestConfiguration *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder when successful
func (m *AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) {
    return NewAccessreviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
