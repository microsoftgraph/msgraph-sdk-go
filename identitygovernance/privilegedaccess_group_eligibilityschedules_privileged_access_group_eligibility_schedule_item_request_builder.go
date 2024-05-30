package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder provides operations to manage the eligibilitySchedules property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
type PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal instantiates a new PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    m := &PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilitySchedules/{privilegedAccessGroupEligibilitySchedule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder instantiates a new PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property eligibilitySchedules for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupeligibilityschedule-get?view=graph-rest-1.0
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupEligibilitySchedule entity.
// returns a *PrivilegedaccessGroupEligibilityschedulesItemGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Group()(*PrivilegedaccessGroupEligibilityschedulesItemGroupRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulesItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property eligibilitySchedules in identityGovernance
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupEligibilitySchedule entity.
// returns a *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Principal()(*PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property eligibilitySchedules for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property eligibilitySchedules in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
