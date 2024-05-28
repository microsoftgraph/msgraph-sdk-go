package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderGetQueryParameters read the properties and relationships of a cloudPcUserSetting object.
type VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderGetQueryParameters
}
// VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *VirtualendpointUsersettingsItemAssignRequestBuilder when successful
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) Assign()(*VirtualendpointUsersettingsItemAssignRequestBuilder) {
    return NewVirtualendpointUsersettingsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.cloudPcUserSetting entity.
// returns a *VirtualendpointUsersettingsItemAssignmentsRequestBuilder when successful
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) Assignments()(*VirtualendpointUsersettingsItemAssignmentsRequestBuilder) {
    return NewVirtualendpointUsersettingsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderInternal instantiates a new VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder and sets the default values.
func NewVirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) {
    m := &VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/userSettings/{cloudPcUserSetting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder instantiates a new VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder and sets the default values.
func NewVirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a cloudPcUserSetting object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcusersetting-delete?view=graph-rest-1.0
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a cloudPcUserSetting object.
// returns a CloudPcUserSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcusersetting-get?view=graph-rest-1.0
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcUserSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingable), nil
}
// Patch update the properties of a cloudPcUserSetting object.
// returns a CloudPcUserSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcusersetting-update?view=graph-rest-1.0
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingable, requestConfiguration *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcUserSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingable), nil
}
// ToDeleteRequestInformation delete a cloudPcUserSetting object.
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a cloudPcUserSetting object.
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a cloudPcUserSetting object.
// returns a *RequestInformation when successful
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcUserSettingable, requestConfiguration *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder when successful
func (m *VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder) {
    return NewVirtualendpointUsersettingsCloudPcUserSettingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
