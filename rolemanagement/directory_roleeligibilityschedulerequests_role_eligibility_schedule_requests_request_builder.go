package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters in PIM, retrieve the requests for role eligibilities for principals made through the unifiedRoleEligibilityScheduleRequest object.
type DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters struct {
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
// DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters
}
// DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleEligibilityScheduleRequestId provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) ByUnifiedRoleEligibilityScheduleRequestId(unifiedRoleEligibilityScheduleRequestId string)(*DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleEligibilityScheduleRequestId != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = unifiedRoleEligibilityScheduleRequestId
    }
    return NewDirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal instantiates a new DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    m := &DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder instantiates a new DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryRoleeligibilityschedulerequestsCountRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) Count()(*DirectoryRoleeligibilityschedulerequestsCountRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *DirectoryRoleeligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*DirectoryRoleeligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get in PIM, retrieve the requests for role eligibilities for principals made through the unifiedRoleEligibilityScheduleRequest object.
// returns a UnifiedRoleEligibilityScheduleRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/rbacapplication-list-roleeligibilityschedulerequests?view=graph-rest-1.0
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestCollectionResponseable), nil
}
// Post in PIM, request for a role eligibility for a principal through the unifiedRoleEligibilityScheduleRequest object. This operation allows both admins and eligible users to add, revoke, or extend eligible assignments.
// returns a UnifiedRoleEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/rbacapplication-post-roleeligibilityschedulerequests?view=graph-rest-1.0
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable), nil
}
// ToGetRequestInformation in PIM, retrieve the requests for role eligibilities for principals made through the unifiedRoleEligibilityScheduleRequest object.
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation in PIM, request for a role eligibility for a principal through the unifiedRoleEligibilityScheduleRequest object. This operation allows both admins and eligible users to add, revoke, or extend eligible assignments.
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
