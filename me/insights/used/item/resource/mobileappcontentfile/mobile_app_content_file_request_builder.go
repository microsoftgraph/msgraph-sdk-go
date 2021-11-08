package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i8b206fdc1146a88c217f365a8d0576caf4988d44fb7d4aa4de0eeea515861506 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/mobileappcontentfile/commit"
    ia34ec4f6adbc00cb601c79cf4bbbb91f9f477453881a1bfcaa29805b26c86063 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/mobileappcontentfile/renewupload"
)

// Builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.mobileAppContentFile
type MobileAppContentFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*i8b206fdc1146a88c217f365a8d0576caf4988d44fb7d4aa4de0eeea515861506.CommitRequestBuilder) {
    return i8b206fdc1146a88c217f365a8d0576caf4988d44fb7d4aa4de0eeea515861506.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new MobileAppContentFileRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/used/{usedInsight_id}/resource/microsoft.graph.mobileAppContentFile";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new MobileAppContentFileRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMobileAppContentFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppContentFileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*ia34ec4f6adbc00cb601c79cf4bbbb91f9f477453881a1bfcaa29805b26c86063.RenewUploadRequestBuilder) {
    return ia34ec4f6adbc00cb601c79cf4bbbb91f9f477453881a1bfcaa29805b26c86063.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
