package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder provides operations to call the getVirtualAppointmentJoinWebUrl method.
type OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderInternal instantiates a new GetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    m := &OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/getVirtualAppointmentJoinWebUrl()", pathParameters),
    }
    return m
}
// NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder instantiates a new GetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getVirtualAppointmentJoinWebUrl
func (m *OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseable), nil
}
// ToGetRequestInformation invoke function getVirtualAppointmentJoinWebUrl
func (m *OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
