package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4b0b164de4b4a42b3808f8e834302218608f2836d053ca703b5b4daf32925b00 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/mobileappcontentfile/commit"
    ibe150cff9f90b2a81dc029759fd133643d8e920fb8fffb9af034606973e5829d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/mobileappcontentfile/renewupload"
)

// Builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.mobileAppContentFile
type MobileAppContentFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*i4b0b164de4b4a42b3808f8e834302218608f2836d053ca703b5b4daf32925b00.CommitRequestBuilder) {
    return i4b0b164de4b4a42b3808f8e834302218608f2836d053ca703b5b4daf32925b00.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new MobileAppContentFileRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/trending/{trending_id}/resource/microsoft.graph.mobileAppContentFile";
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
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*ibe150cff9f90b2a81dc029759fd133643d8e920fb8fffb9af034606973e5829d.RenewUploadRequestBuilder) {
    return ibe150cff9f90b2a81dc029759fd133643d8e920fb8fffb9af034606973e5829d.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
