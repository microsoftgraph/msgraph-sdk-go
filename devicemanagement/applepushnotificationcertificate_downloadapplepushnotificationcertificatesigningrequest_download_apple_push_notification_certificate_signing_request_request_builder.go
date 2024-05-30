package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder provides operations to call the downloadApplePushNotificationCertificateSigningRequest method.
type ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    m := &ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/applePushNotificationCertificate/downloadApplePushNotificationCertificateSigningRequest()", pathParameters),
    }
    return m
}
// NewApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder instantiates a new ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Get download Apple push notification certificate signing request
// Deprecated: This method is obsolete. Use GetAsDownloadApplePushNotificationCertificateSigningRequestGetResponse instead.
// returns a ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-applepushnotificationcertificate-downloadapplepushnotificationcertificatesigningrequest?view=graph-rest-1.0
func (m *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestResponseable), nil
}
// GetAsDownloadApplePushNotificationCertificateSigningRequestGetResponse download Apple push notification certificate signing request
// returns a ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-applepushnotificationcertificate-downloadapplepushnotificationcertificatesigningrequest?view=graph-rest-1.0
func (m *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) GetAsDownloadApplePushNotificationCertificateSigningRequestGetResponse(ctx context.Context, requestConfiguration *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestGetResponseable), nil
}
// ToGetRequestInformation download Apple push notification certificate signing request
// returns a *RequestInformation when successful
func (m *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder when successful
func (m *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) WithUrl(rawUrl string)(*ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    return NewApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
