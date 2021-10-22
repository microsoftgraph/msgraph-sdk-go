package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i14d7f0005fc36613f7f1de40544c198c0b452c7d09e615920e261be8e18d6c87 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/schedulechangerequest/approve"
    ie2b87b8c4436d15c09e0b35812e1ab9fa4f36c57d99720ace13f51872f5dc0c2 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/schedulechangerequest/decline"
)

type ScheduleChangeRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*i14d7f0005fc36613f7f1de40544c198c0b452c7d09e615920e261be8e18d6c87.ApproveRequestBuilder) {
    return i14d7f0005fc36613f7f1de40544c198c0b452c7d09e615920e261be8e18d6c87.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/used/{usedInsight_id}/resource/microsoft.graph.scheduleChangeRequest";
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
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*ie2b87b8c4436d15c09e0b35812e1ab9fa4f36c57d99720ace13f51872f5dc0c2.DeclineRequestBuilder) {
    return ie2b87b8c4436d15c09e0b35812e1ab9fa4f36c57d99720ace13f51872f5dc0c2.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
