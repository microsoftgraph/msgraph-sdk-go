package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.cloudPcUserSetting entity.
type VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderGetQueryParameters represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand. For an example, see Get cloudPcUserSetting.
type VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderGetQueryParameters
}
// VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderInternal instantiates a new VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder and sets the default values.
func NewVirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) {
    m := &VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/userSettings/{cloudPcUserSetting%2Did}/assignments/{cloudPcUserSettingAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder instantiates a new VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder and sets the default values.
func NewVirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand. For an example, see Get cloudPcUserSetting.
// returns a CloudPcUserSettingAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcUserSettingAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable), nil
}
// Patch update the navigation property assignments in deviceManagement
// returns a CloudPcUserSettingAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcUserSettingAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand. For an example, see Get cloudPcUserSetting.
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder when successful
func (m *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) {
    return NewVirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
