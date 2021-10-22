package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i338a88d74f10a1ee4f139da369171285a9a24ed637b166a2d685d7ba1f8154d4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/schedulechangerequest/approve"
    i3f10ce10afb39c3e7a5d0d00a34a04b716e8707216ce9094b8d2b1c4d8ed50b2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/schedulechangerequest/decline"
)

type ScheduleChangeRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*i338a88d74f10a1ee4f139da369171285a9a24ed637b166a2d685d7ba1f8154d4.ApproveRequestBuilder) {
    return i338a88d74f10a1ee4f139da369171285a9a24ed637b166a2d685d7ba1f8154d4.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.scheduleChangeRequest";
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
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*i3f10ce10afb39c3e7a5d0d00a34a04b716e8707216ce9094b8d2b1c4d8ed50b2.DeclineRequestBuilder) {
    return i3f10ce10afb39c3e7a5d0d00a34a04b716e8707216ce9094b8d2b1c4d8ed50b2.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
