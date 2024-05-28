package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentScheduleInstance entity.
type DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderGetQueryParameters if the request is from an eligible administrator to activate a role, this parameter shows the related eligible assignment for that activation. Otherwise, it's null. Supports $expand and $select nested in $expand.
type DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderGetQueryParameters
}
// NewDirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderInternal instantiates a new DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) {
    m := &DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance%2Did}/activatedUsing{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder instantiates a new DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get if the request is from an eligible administrator to activate a role, this parameter shows the related eligible assignment for that activation. Otherwise, it's null. Supports $expand and $select nested in $expand.
// returns a UnifiedRoleEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable), nil
}
// ToGetRequestInformation if the request is from an eligible administrator to activate a role, this parameter shows the related eligible assignment for that activation. Otherwise, it's null. Supports $expand and $select nested in $expand.
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) {
    return NewDirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
