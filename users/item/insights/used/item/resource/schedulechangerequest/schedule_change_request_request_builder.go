package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i5886815ee7c2ee60a2390f20c5a807bb28b7d2ceff9ffb6b29da6a4dc57b5809 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/schedulechangerequest/approve"
    ieb3ce81e48c891f489003912749bba0993082fced76ce6be9c1ef7b366d0251e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/schedulechangerequest/decline"
)

type ScheduleChangeRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*i5886815ee7c2ee60a2390f20c5a807bb28b7d2ceff9ffb6b29da6a4dc57b5809.ApproveRequestBuilder) {
    return i5886815ee7c2ee60a2390f20c5a807bb28b7d2ceff9ffb6b29da6a4dc57b5809.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.scheduleChangeRequest";
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
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*ieb3ce81e48c891f489003912749bba0993082fced76ce6be9c1ef7b366d0251e.DeclineRequestBuilder) {
    return ieb3ce81e48c891f489003912749bba0993082fced76ce6be9c1ef7b366d0251e.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
