package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i82a69798c7c95d4d17a218f432021f7a9ebf393c2c757b4834bf76ac217e37f1 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/schedulechangerequest/approve"
    i98757369a3b78c4985128c975f68f48a002e40e2eacd3b1dc1fd292a07f4164c "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/schedulechangerequest/decline"
)

type ScheduleChangeRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*i82a69798c7c95d4d17a218f432021f7a9ebf393c2c757b4834bf76ac217e37f1.ApproveRequestBuilder) {
    return i82a69798c7c95d4d17a218f432021f7a9ebf393c2c757b4834bf76ac217e37f1.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.scheduleChangeRequest";
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
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*i98757369a3b78c4985128c975f68f48a002e40e2eacd3b1dc1fd292a07f4164c.DeclineRequestBuilder) {
    return i98757369a3b78c4985128c975f68f48a002e40e2eacd3b1dc1fd292a07f4164c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
