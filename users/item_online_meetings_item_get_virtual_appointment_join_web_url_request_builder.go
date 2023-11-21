package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder provides operations to call the getVirtualAppointmentJoinWebUrl method.
type ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderInternal instantiates a new GetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    m := &ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/getVirtualAppointmentJoinWebUrl()", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder instantiates a new GetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getVirtualAppointmentJoinWebUrl
// Deprecated: This method is obsolete. Use GetAsGetVirtualAppointmentJoinWebUrlGetResponse instead.
func (m *ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseable), nil
}
// GetAsGetVirtualAppointmentJoinWebUrlGetResponse invoke function getVirtualAppointmentJoinWebUrl
func (m *ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) GetAsGetVirtualAppointmentJoinWebUrlGetResponse(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlGetResponseable), nil
}
// ToGetRequestInformation invoke function getVirtualAppointmentJoinWebUrl
func (m *ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewItemOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
