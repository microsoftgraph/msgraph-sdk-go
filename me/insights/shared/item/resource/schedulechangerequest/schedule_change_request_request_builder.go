package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ia5932d07c37b043589675035830f8d97ceab69020a15751c8659d274b2865077 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/schedulechangerequest/decline"
    iee277c96a0ecc6f8f85e4b78b4af2a54bb5de24503bd0a912443d697a175af6b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/schedulechangerequest/approve"
)

type ScheduleChangeRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*iee277c96a0ecc6f8f85e4b78b4af2a54bb5de24503bd0a912443d697a175af6b.ApproveRequestBuilder) {
    return iee277c96a0ecc6f8f85e4b78b4af2a54bb5de24503bd0a912443d697a175af6b.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.scheduleChangeRequest";
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
func NewScheduleChangeRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleChangeRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*ia5932d07c37b043589675035830f8d97ceab69020a15751c8659d274b2865077.DeclineRequestBuilder) {
    return ia5932d07c37b043589675035830f8d97ceab69020a15751c8659d274b2865077.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
