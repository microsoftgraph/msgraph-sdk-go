package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder provides operations to call the createDownloadUrl method.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal instantiates a new MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    m := &MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}/createDownloadUrl", pathParameters),
    }
    return m
}
// NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder instantiates a new MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// returns a AppLogCollectionDownloadDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-applogcollectionrequest-createdownloadurl?view=graph-rest-1.0
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionDownloadDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
