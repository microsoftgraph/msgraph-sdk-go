package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i14d7f0005fc36613f7f1de40544c198c0b452c7d09e615920e261be8e18d6c87 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/schedulechangerequest/approve"
    ie2b87b8c4436d15c09e0b35812e1ab9fa4f36c57d99720ace13f51872f5dc0c2 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/schedulechangerequest/decline"
)

// Builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.scheduleChangeRequest
type ScheduleChangeRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*i14d7f0005fc36613f7f1de40544c198c0b452c7d09e615920e261be8e18d6c87.ApproveRequestBuilder) {
    return i14d7f0005fc36613f7f1de40544c198c0b452c7d09e615920e261be8e18d6c87.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ScheduleChangeRequestRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/used/{usedInsight_id}/resource/microsoft.graph.scheduleChangeRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ScheduleChangeRequestRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewScheduleChangeRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleChangeRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*ie2b87b8c4436d15c09e0b35812e1ab9fa4f36c57d99720ace13f51872f5dc0c2.DeclineRequestBuilder) {
    return ie2b87b8c4436d15c09e0b35812e1ab9fa4f36c57d99720ace13f51872f5dc0c2.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
