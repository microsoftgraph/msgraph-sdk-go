package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i74c17e639f4e03d11a4a2f9c47d62b4e4b043f1f220fdf80ca96614f57bae748 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/mobileappcontentfile/commit"
    iff4ca3615452eb7b8b3d04e6b2463f326803d21ec15a2d29eb77d868f2d45c6f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/mobileappcontentfile/renewupload"
)

// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.mobileAppContentFile
type MobileAppContentFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*i74c17e639f4e03d11a4a2f9c47d62b4e4b043f1f220fdf80ca96614f57bae748.CommitRequestBuilder) {
    return i74c17e639f4e03d11a4a2f9c47d62b4e4b043f1f220fdf80ca96614f57bae748.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new MobileAppContentFileRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.mobileAppContentFile";
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
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*iff4ca3615452eb7b8b3d04e6b2463f326803d21ec15a2d29eb77d868f2d45c6f.RenewUploadRequestBuilder) {
    return iff4ca3615452eb7b8b3d04e6b2463f326803d21ec15a2d29eb77d868f2d45c6f.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
