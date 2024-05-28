package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointUsersettingsItemAssignmentsRequestBuilder provides operations to manage the assignments property of the microsoft.graph.cloudPcUserSetting entity.
type VirtualendpointUsersettingsItemAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointUsersettingsItemAssignmentsRequestBuilderGetQueryParameters represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand. For an example, see Get cloudPcUserSetting.
type VirtualendpointUsersettingsItemAssignmentsRequestBuilderGetQueryParameters struct {
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
// VirtualendpointUsersettingsItemAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsItemAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointUsersettingsItemAssignmentsRequestBuilderGetQueryParameters
}
// VirtualendpointUsersettingsItemAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsItemAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudPcUserSettingAssignmentId provides operations to manage the assignments property of the microsoft.graph.cloudPcUserSetting entity.
// returns a *VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder when successful
func (m *VirtualendpointUsersettingsItemAssignmentsRequestBuilder) ByCloudPcUserSettingAssignmentId(cloudPcUserSettingAssignmentId string)(*VirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudPcUserSettingAssignmentId != "" {
        urlTplParams["cloudPcUserSettingAssignment%2Did"] = cloudPcUserSettingAssignmentId
    }
    return NewVirtualendpointUsersettingsItemAssignmentsCloudPcUserSettingAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointUsersettingsItemAssignmentsRequestBuilderInternal instantiates a new VirtualendpointUsersettingsItemAssignmentsRequestBuilder and sets the default values.
func NewVirtualendpointUsersettingsItemAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointUsersettingsItemAssignmentsRequestBuilder) {
    m := &VirtualendpointUsersettingsItemAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/userSettings/{cloudPcUserSetting%2Did}/assignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointUsersettingsItemAssignmentsRequestBuilder instantiates a new VirtualendpointUsersettingsItemAssignmentsRequestBuilder and sets the default values.
func NewVirtualendpointUsersettingsItemAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointUsersettingsItemAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointUsersettingsItemAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualendpointUsersettingsItemAssignmentsCountRequestBuilder when successful
func (m *VirtualendpointUsersettingsItemAssignmentsRequestBuilder) Count()(*VirtualendpointUsersettingsItemAssignmentsCountRequestBuilder) {
    return NewVirtualendpointUsersettingsItemAssignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand. For an example, see Get cloudPcUserSetting.
// returns a CloudPcUserSettingAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointUsersettingsItemAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcUserSettingAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentCollectionResponseable), nil
}
// Post create new navigation property to assignments for deviceManagement
// returns a CloudPcUserSettingAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointUsersettingsItemAssignmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand. For an example, see Get cloudPcUserSetting.
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsItemAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsItemAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingAssignmentable, requestConfiguration *VirtualendpointUsersettingsItemAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointUsersettingsItemAssignmentsRequestBuilder when successful
func (m *VirtualendpointUsersettingsItemAssignmentsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointUsersettingsItemAssignmentsRequestBuilder) {
    return NewVirtualendpointUsersettingsItemAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
