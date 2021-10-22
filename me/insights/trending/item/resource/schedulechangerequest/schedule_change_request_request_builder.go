package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ibe97c6a9e6d9d1164c43365a405dd9c4400c558fbda38c4afa7dab1cc9d8b1ef "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/schedulechangerequest/approve"
    if9ac75f014e4b2dde39fa8838a004daa8c792d8023bead39b08408a615c0bc53 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/schedulechangerequest/decline"
)

type ScheduleChangeRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*ibe97c6a9e6d9d1164c43365a405dd9c4400c558fbda38c4afa7dab1cc9d8b1ef.ApproveRequestBuilder) {
    return ibe97c6a9e6d9d1164c43365a405dd9c4400c558fbda38c4afa7dab1cc9d8b1ef.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/trending/{trending_id}/resource/microsoft.graph.scheduleChangeRequest";
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
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*if9ac75f014e4b2dde39fa8838a004daa8c792d8023bead39b08408a615c0bc53.DeclineRequestBuilder) {
    return if9ac75f014e4b2dde39fa8838a004daa8c792d8023bead39b08408a615c0bc53.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
