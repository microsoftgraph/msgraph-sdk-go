package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderGetQueryParameters get a list of the privilegedAccessGroupAssignmentScheduleInstance objects and their properties.
type PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderGetQueryParameters struct {
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
// PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrivilegedAccessGroupAssignmentScheduleInstanceId provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) ByPrivilegedAccessGroupAssignmentScheduleInstanceId(privilegedAccessGroupAssignmentScheduleInstanceId string)(*PrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegedAccessGroupAssignmentScheduleInstanceId != "" {
        urlTplParams["privilegedAccessGroupAssignmentScheduleInstance%2Did"] = privilegedAccessGroupAssignmentScheduleInstanceId
    }
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleInstances{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesCountRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) Count()(*PrivilegedaccessGroupAssignmentscheduleinstancesCountRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedaccessGroupAssignmentscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get a list of the privilegedAccessGroupAssignmentScheduleInstance objects and their properties.
// returns a PrivilegedAccessGroupAssignmentScheduleInstanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-list-assignmentscheduleinstances?view=graph-rest-1.0
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleInstanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceCollectionResponseable), nil
}
// Post create new navigation property to assignmentScheduleInstances for identityGovernance
// returns a PrivilegedAccessGroupAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable), nil
}
// ToGetRequestInformation get a list of the privilegedAccessGroupAssignmentScheduleInstance objects and their properties.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to assignmentScheduleInstances for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleInstanceable, requestConfiguration *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
