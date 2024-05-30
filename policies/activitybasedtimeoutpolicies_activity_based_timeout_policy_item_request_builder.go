package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
type ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetQueryParameters get the properties of an activityBasedTimeoutPolicy object.
type ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetQueryParameters
}
// ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppliesTo provides operations to manage the appliesTo property of the microsoft.graph.stsPolicy entity.
// returns a *ActivitybasedtimeoutpoliciesItemAppliestoAppliesToRequestBuilder when successful
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) AppliesTo()(*ActivitybasedtimeoutpoliciesItemAppliestoAppliesToRequestBuilder) {
    return NewActivitybasedtimeoutpoliciesItemAppliestoAppliesToRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal instantiates a new ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder and sets the default values.
func NewActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    m := &ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/activityBasedTimeoutPolicies/{activityBasedTimeoutPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder instantiates a new ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder and sets the default values.
func NewActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an activityBasedTimeoutPolicy object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/activitybasedtimeoutpolicy-delete?view=graph-rest-1.0
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties of an activityBasedTimeoutPolicy object.
// returns a ActivityBasedTimeoutPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/activitybasedtimeoutpolicy-get?view=graph-rest-1.0
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityBasedTimeoutPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateActivityBasedTimeoutPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityBasedTimeoutPolicyable), nil
}
// Patch update the properties of an activityBasedTimeoutPolicy object.
// returns a ActivityBasedTimeoutPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/activitybasedtimeoutpolicy-update?view=graph-rest-1.0
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityBasedTimeoutPolicyable, requestConfiguration *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityBasedTimeoutPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateActivityBasedTimeoutPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityBasedTimeoutPolicyable), nil
}
// ToDeleteRequestInformation delete an activityBasedTimeoutPolicy object.
// returns a *RequestInformation when successful
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties of an activityBasedTimeoutPolicy object.
// returns a *RequestInformation when successful
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an activityBasedTimeoutPolicy object.
// returns a *RequestInformation when successful
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityBasedTimeoutPolicyable, requestConfiguration *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder when successful
func (m *ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    return NewActivitybasedtimeoutpoliciesActivityBasedTimeoutPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
