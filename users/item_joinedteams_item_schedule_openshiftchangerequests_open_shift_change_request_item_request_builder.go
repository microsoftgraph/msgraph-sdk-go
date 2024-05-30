package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
type ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderGetQueryParameters the open shift requests in the schedule.
type ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderGetQueryParameters
}
// ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderInternal instantiates a new ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) {
    m := &ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/schedule/openShiftChangeRequests/{openShiftChangeRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder instantiates a new ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property openShiftChangeRequests for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the open shift requests in the schedule.
// returns a OpenShiftChangeRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OpenShiftChangeRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOpenShiftChangeRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OpenShiftChangeRequestable), nil
}
// Patch update the navigation property openShiftChangeRequests in users
// returns a OpenShiftChangeRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OpenShiftChangeRequestable, requestConfiguration *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OpenShiftChangeRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOpenShiftChangeRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OpenShiftChangeRequestable), nil
}
// ToDeleteRequestInformation delete navigation property openShiftChangeRequests for users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the open shift requests in the schedule.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property openShiftChangeRequests in users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OpenShiftChangeRequestable, requestConfiguration *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder) {
    return NewItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
