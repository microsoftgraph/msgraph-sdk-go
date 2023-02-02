package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder provides operations to call the publish method.
type BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilderInternal instantiates a new PublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder) {
    m := &BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/microsoft.graph.publish";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewBookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder instantiates a new PublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post make the scheduling page of a business available to external customers. Set the **isPublished** property to `true`, and the **publicUrl** property to the URL of the scheduling page.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingbusiness-publish?view=graph-rest-1.0
func (m *BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder) Post(ctx context.Context, requestConfiguration *BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation make the scheduling page of a business available to external customers. Set the **isPublished** property to `true`, and the **publicUrl** property to the URL of the scheduling page.
func (m *BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesItemMicrosoftGraphPublishPublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
