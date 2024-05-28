package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder provides operations to manage the eligibilityScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderGetQueryParameters get a list of the privilegedAccessGroupEligibilityScheduleInstance objects and their properties.
type PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderGetQueryParameters struct {
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
// PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrivilegedAccessGroupEligibilityScheduleInstanceId provides operations to manage the eligibilityScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupEligibilityscheduleinstancesPrivilegedAccessGroupEligibilityScheduleInstanceItemRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) ByPrivilegedAccessGroupEligibilityScheduleInstanceId(privilegedAccessGroupEligibilityScheduleInstanceId string)(*PrivilegedaccessGroupEligibilityscheduleinstancesPrivilegedAccessGroupEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegedAccessGroupEligibilityScheduleInstanceId != "" {
        urlTplParams["privilegedAccessGroupEligibilityScheduleInstance%2Did"] = privilegedAccessGroupEligibilityScheduleInstanceId
    }
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesPrivilegedAccessGroupEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderInternal instantiates a new PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) {
    m := &PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder instantiates a new PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedaccessGroupEligibilityscheduleinstancesCountRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) Count()(*PrivilegedaccessGroupEligibilityscheduleinstancesCountRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PrivilegedaccessGroupEligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedaccessGroupEligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get a list of the privilegedAccessGroupEligibilityScheduleInstance objects and their properties.
// returns a PrivilegedAccessGroupEligibilityScheduleInstanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-list-eligibilityscheduleinstances?view=graph-rest-1.0
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleInstanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleInstanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleInstanceCollectionResponseable), nil
}
// Post create new navigation property to eligibilityScheduleInstances for identityGovernance
// returns a PrivilegedAccessGroupEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleInstanceable, requestConfiguration *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleInstanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleInstanceable), nil
}
// ToGetRequestInformation get a list of the privilegedAccessGroupEligibilityScheduleInstance objects and their properties.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to eligibilityScheduleInstances for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleInstanceable, requestConfiguration *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
