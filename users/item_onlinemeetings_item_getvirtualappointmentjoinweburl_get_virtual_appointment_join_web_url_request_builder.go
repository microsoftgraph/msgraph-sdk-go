package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder provides operations to call the getVirtualAppointmentJoinWebUrl method.
type ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal instantiates a new ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    m := &ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/getVirtualAppointmentJoinWebUrl()", pathParameters),
    }
    return m
}
// NewItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder instantiates a new ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a join web URL for a Microsoft Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// Deprecated: This method is obsolete. Use GetAsGetVirtualAppointmentJoinWebUrlGetResponse instead.
// returns a ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-getvirtualappointmentjoinweburl?view=graph-rest-1.0
func (m *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseable), nil
}
// GetAsGetVirtualAppointmentJoinWebUrlGetResponse get a join web URL for a Microsoft Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// returns a ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-getvirtualappointmentjoinweburl?view=graph-rest-1.0
func (m *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) GetAsGetVirtualAppointmentJoinWebUrlGetResponse(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseable), nil
}
// ToGetRequestInformation get a join web URL for a Microsoft Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
