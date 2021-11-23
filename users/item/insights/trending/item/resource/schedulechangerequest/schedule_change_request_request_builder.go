package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i23c29a9732230b31120768fb044badc4f6ea1d861eec7d031f3996e041e846d8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/schedulechangerequest/approve"
    i38789783a77336f314b8520d585101276577c65d51151af78b10f3c8c7ba1e41 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/schedulechangerequest/decline"
)

// ScheduleChangeRequestRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.scheduleChangeRequest
type ScheduleChangeRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*i23c29a9732230b31120768fb044badc4f6ea1d861eec7d031f3996e041e846d8.ApproveRequestBuilder) {
    return i23c29a9732230b31120768fb044badc4f6ea1d861eec7d031f3996e041e846d8.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewScheduleChangeRequestRequestBuilderInternal instantiates a new ScheduleChangeRequestRequestBuilder and sets the default values.
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/trending/{trending_id}/resource/microsoft.graph.scheduleChangeRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleChangeRequestRequestBuilder instantiates a new ScheduleChangeRequestRequestBuilder and sets the default values.
func NewScheduleChangeRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleChangeRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*i38789783a77336f314b8520d585101276577c65d51151af78b10f3c8c7ba1e41.DeclineRequestBuilder) {
    return i38789783a77336f314b8520d585101276577c65d51151af78b10f3c8c7ba1e41.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
