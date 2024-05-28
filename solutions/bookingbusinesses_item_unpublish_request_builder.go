package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingbusinessesItemUnpublishRequestBuilder provides operations to call the unpublish method.
type BookingbusinessesItemUnpublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingbusinessesItemUnpublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemUnpublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingbusinessesItemUnpublishRequestBuilderInternal instantiates a new BookingbusinessesItemUnpublishRequestBuilder and sets the default values.
func NewBookingbusinessesItemUnpublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemUnpublishRequestBuilder) {
    m := &BookingbusinessesItemUnpublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/unpublish", pathParameters),
    }
    return m
}
// NewBookingbusinessesItemUnpublishRequestBuilder instantiates a new BookingbusinessesItemUnpublishRequestBuilder and sets the default values.
func NewBookingbusinessesItemUnpublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemUnpublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingbusinessesItemUnpublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post make the scheduling page of this business not available to external customers. Set the isPublished property to false, and the publicUrl property to null.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-unpublish?view=graph-rest-1.0
func (m *BookingbusinessesItemUnpublishRequestBuilder) Post(ctx context.Context, requestConfiguration *BookingbusinessesItemUnpublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation make the scheduling page of this business not available to external customers. Set the isPublished property to false, and the publicUrl property to null.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemUnpublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *BookingbusinessesItemUnpublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BookingbusinessesItemUnpublishRequestBuilder when successful
func (m *BookingbusinessesItemUnpublishRequestBuilder) WithUrl(rawUrl string)(*BookingbusinessesItemUnpublishRequestBuilder) {
    return NewBookingbusinessesItemUnpublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
