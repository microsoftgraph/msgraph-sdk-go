package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4811442522f2179f1b114b07d01826d821788985c71b038c016ba6d5508488c0 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/mobileappcontentfile/renewupload"
    ib1e234b4ff408390909249e7bb9f4912b7ceb64c4587784c93396fe7dc824c48 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/mobileappcontentfile/commit"
)

type MobileAppContentFileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*ib1e234b4ff408390909249e7bb9f4912b7ceb64c4587784c93396fe7dc824c48.CommitRequestBuilder) {
    return ib1e234b4ff408390909249e7bb9f4912b7ceb64c4587784c93396fe7dc824c48.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/trending/{trending_id}/resource/microsoft.graph.mobileAppContentFile";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewMobileAppContentFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppContentFileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*i4811442522f2179f1b114b07d01826d821788985c71b038c016ba6d5508488c0.RenewUploadRequestBuilder) {
    return i4811442522f2179f1b114b07d01826d821788985c71b038c016ba6d5508488c0.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
