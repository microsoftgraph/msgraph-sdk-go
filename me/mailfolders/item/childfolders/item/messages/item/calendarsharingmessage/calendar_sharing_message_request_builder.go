package calendarsharingmessage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i334c09b864ea8301045edf0ecb98ff6bba0ce2ee121be456e2b96ce656884b55 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/messages/item/calendarsharingmessage/accept"
)

// CalendarSharingMessageRequestBuilder builds and executes requests for operations under \me\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\microsoft.graph.calendarSharingMessage
type CalendarSharingMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accept the accept property
func (m *CalendarSharingMessageRequestBuilder) Accept()(*i334c09b864ea8301045edf0ecb98ff6bba0ce2ee121be456e2b96ce656884b55.AcceptRequestBuilder) {
    return i334c09b864ea8301045edf0ecb98ff6bba0ce2ee121be456e2b96ce656884b55.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCalendarSharingMessageRequestBuilderInternal instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
func NewCalendarSharingMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    m := &CalendarSharingMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/microsoft.graph.calendarSharingMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarSharingMessageRequestBuilder instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
func NewCalendarSharingMessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarSharingMessageRequestBuilderInternal(urlParams, requestAdapter)
}
