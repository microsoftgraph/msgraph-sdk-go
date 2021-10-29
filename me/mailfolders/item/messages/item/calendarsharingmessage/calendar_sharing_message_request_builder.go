package calendarsharingmessage

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i48bb4c4ad9b00e257ad888a69f5d2ebd9afa2982a69d66e19dc4cce9db02f8d7 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/calendarsharingmessage/accept"
)

// Builds and executes requests for operations under \me\mailFolders\{mailFolder-id}\messages\{message-id}\microsoft.graph.calendarSharingMessage
type CalendarSharingMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *CalendarSharingMessageRequestBuilder) Accept()(*i48bb4c4ad9b00e257ad888a69f5d2ebd9afa2982a69d66e19dc4cce9db02f8d7.AcceptRequestBuilder) {
    return i48bb4c4ad9b00e257ad888a69f5d2ebd9afa2982a69d66e19dc4cce9db02f8d7.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarSharingMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    m := &CalendarSharingMessageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/mailFolders/{mailFolder_id}/messages/{message_id}/microsoft.graph.calendarSharingMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarSharingMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarSharingMessageRequestBuilderInternal(urlParams, requestAdapter)
}
