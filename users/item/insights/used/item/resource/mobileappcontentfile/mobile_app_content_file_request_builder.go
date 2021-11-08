package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3fa8197af17856c636438004f8d47a0566dc47090e1e94696897b67e9a4116cc "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/mobileappcontentfile/renewupload"
    i5483af733a41a8399a56974aa7ffd24f5399af0d397302b72631e2da5bcfada9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/mobileappcontentfile/commit"
)

// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.mobileAppContentFile
type MobileAppContentFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*i5483af733a41a8399a56974aa7ffd24f5399af0d397302b72631e2da5bcfada9.CommitRequestBuilder) {
    return i5483af733a41a8399a56974aa7ffd24f5399af0d397302b72631e2da5bcfada9.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new MobileAppContentFileRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.mobileAppContentFile";
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
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*i3fa8197af17856c636438004f8d47a0566dc47090e1e94696897b67e9a4116cc.RenewUploadRequestBuilder) {
    return i3fa8197af17856c636438004f8d47a0566dc47090e1e94696897b67e9a4116cc.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
