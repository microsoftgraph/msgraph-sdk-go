package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ia5932d07c37b043589675035830f8d97ceab69020a15751c8659d274b2865077 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/schedulechangerequest/decline"
    iee277c96a0ecc6f8f85e4b78b4af2a54bb5de24503bd0a912443d697a175af6b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/schedulechangerequest/approve"
)

// scheduleChangeRequestRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\resource\microsoft.graph.scheduleChangeRequest
type ScheduleChangeRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*iee277c96a0ecc6f8f85e4b78b4af2a54bb5de24503bd0a912443d697a175af6b.ApproveRequestBuilder) {
    return iee277c96a0ecc6f8f85e4b78b4af2a54bb5de24503bd0a912443d697a175af6b.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewScheduleChangeRequestRequestBuilderInternal instantiates a new ScheduleChangeRequestRequestBuilder and sets the default values.
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.scheduleChangeRequest";
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
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*ia5932d07c37b043589675035830f8d97ceab69020a15751c8659d274b2865077.DeclineRequestBuilder) {
    return ia5932d07c37b043589675035830f8d97ceab69020a15751c8659d274b2865077.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
