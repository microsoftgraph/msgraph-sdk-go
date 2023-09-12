package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingBusinessesItemGetStaffAvailabilityRequestBuilder provides operations to call the getStaffAvailability method.
type BookingBusinessesItemGetStaffAvailabilityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingBusinessesItemGetStaffAvailabilityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemGetStaffAvailabilityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingBusinessesItemGetStaffAvailabilityRequestBuilderInternal instantiates a new GetStaffAvailabilityRequestBuilder and sets the default values.
func NewBookingBusinessesItemGetStaffAvailabilityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemGetStaffAvailabilityRequestBuilder) {
    m := &BookingBusinessesItemGetStaffAvailabilityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/getStaffAvailability", pathParameters),
    }
    return m
}
// NewBookingBusinessesItemGetStaffAvailabilityRequestBuilder instantiates a new GetStaffAvailabilityRequestBuilder and sets the default values.
func NewBookingBusinessesItemGetStaffAvailabilityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemGetStaffAvailabilityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesItemGetStaffAvailabilityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the availability information of staff members of a Microsoft Bookings calendar.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-getstaffavailability?view=graph-rest-1.0
func (m *BookingBusinessesItemGetStaffAvailabilityRequestBuilder) Post(ctx context.Context, body BookingBusinessesItemGetStaffAvailabilityPostRequestBodyable, requestConfiguration *BookingBusinessesItemGetStaffAvailabilityRequestBuilderPostRequestConfiguration)(BookingBusinessesItemGetStaffAvailabilityResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBookingBusinessesItemGetStaffAvailabilityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(BookingBusinessesItemGetStaffAvailabilityResponseable), nil
}
// ToPostRequestInformation get the availability information of staff members of a Microsoft Bookings calendar.
func (m *BookingBusinessesItemGetStaffAvailabilityRequestBuilder) ToPostRequestInformation(ctx context.Context, body BookingBusinessesItemGetStaffAvailabilityPostRequestBodyable, requestConfiguration *BookingBusinessesItemGetStaffAvailabilityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *BookingBusinessesItemGetStaffAvailabilityRequestBuilder) WithUrl(rawUrl string)(*BookingBusinessesItemGetStaffAvailabilityRequestBuilder) {
    return NewBookingBusinessesItemGetStaffAvailabilityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
