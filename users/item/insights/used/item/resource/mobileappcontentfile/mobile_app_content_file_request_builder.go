package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3fa8197af17856c636438004f8d47a0566dc47090e1e94696897b67e9a4116cc "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/mobileappcontentfile/renewupload"
    i5483af733a41a8399a56974aa7ffd24f5399af0d397302b72631e2da5bcfada9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/mobileappcontentfile/commit"
)

type MobileAppContentFileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*i5483af733a41a8399a56974aa7ffd24f5399af0d397302b72631e2da5bcfada9.CommitRequestBuilder) {
    return i5483af733a41a8399a56974aa7ffd24f5399af0d397302b72631e2da5bcfada9.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.mobileAppContentFile";
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
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*i3fa8197af17856c636438004f8d47a0566dc47090e1e94696897b67e9a4116cc.RenewUploadRequestBuilder) {
    return i3fa8197af17856c636438004f8d47a0566dc47090e1e94696897b67e9a4116cc.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
