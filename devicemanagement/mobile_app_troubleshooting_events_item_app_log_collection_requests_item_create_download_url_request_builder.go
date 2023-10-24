package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder provides operations to call the createDownloadUrl method.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal instantiates a new CreateDownloadUrlRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    m := &MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}/createDownloadUrl", pathParameters),
    }
    return m
}
// NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder instantiates a new CreateDownloadUrlRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-applogcollectionrequest-createdownloadurl?view=graph-rest-1.0
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionDownloadDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppLogCollectionDownloadDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionDownloadDetailsable), nil
}
// ToPostRequestInformation not yet documented
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    return NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
