package downloadapplepushnotificationcertificatesigningrequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// DownloadApplePushNotificationCertificateSigningRequestRequestBuilder provides operations to call the downloadApplePushNotificationCertificateSigningRequest method.
type DownloadApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetOptions options for Get
type DownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new DownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
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
func NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation download Apple push notification certificate signing request
func (m *DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) CreateGetRequestInformation(options *DownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
