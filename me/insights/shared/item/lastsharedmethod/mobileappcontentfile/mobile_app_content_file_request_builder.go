package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i36f08867e8b18ed8ad2be0c5f8be9e22d09f34c35789865ece6e75f98f20595e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/mobileappcontentfile/renewupload"
    idf03bbd3d482050770d3d19fb55e07a913ee41fb407e8412a5be77171de5db18 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/mobileappcontentfile/commit"
)

// MobileAppContentFileRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.mobileAppContentFile
type MobileAppContentFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*idf03bbd3d482050770d3d19fb55e07a913ee41fb407e8412a5be77171de5db18.CommitRequestBuilder) {
    return idf03bbd3d482050770d3d19fb55e07a913ee41fb407e8412a5be77171de5db18.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMobileAppContentFileRequestBuilderInternal instantiates a new MobileAppContentFileRequestBuilder and sets the default values.
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.mobileAppContentFile";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppContentFileRequestBuilder instantiates a new MobileAppContentFileRequestBuilder and sets the default values.
func NewMobileAppContentFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppContentFileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*i36f08867e8b18ed8ad2be0c5f8be9e22d09f34c35789865ece6e75f98f20595e.RenewUploadRequestBuilder) {
    return i36f08867e8b18ed8ad2be0c5f8be9e22d09f34c35789865ece6e75f98f20595e.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
