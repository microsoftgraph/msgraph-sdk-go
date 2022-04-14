package downloadapplepushnotificationcertificatesigningrequest

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DownloadApplePushNotificationCertificateSigningRequestRequestBuilder provides operations to call the downloadApplePushNotificationCertificateSigningRequest method.
type DownloadApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetOptions options for Get
type DownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new DownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    m := &DownloadApplePushNotificationCertificateSigningRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/applePushNotificationCertificate/microsoft.graph.downloadApplePushNotificationCertificateSigningRequest()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilder instantiates a new DownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation download Apple push notification certificate signing request
func (m *DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) CreateGetRequestInformation(options *DownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get download Apple push notification certificate signing request
func (m *DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) Get(options *DownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetOptions)(DownloadApplePushNotificationCertificateSigningRequestResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateDownloadApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(DownloadApplePushNotificationCertificateSigningRequestResponseable), nil
}
