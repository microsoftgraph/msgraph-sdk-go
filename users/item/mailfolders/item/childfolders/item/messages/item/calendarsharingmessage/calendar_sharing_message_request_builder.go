package calendarsharingmessage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibaa6a71ab032a685fe85ce7cefea89706af5749ad58bb2bb38038e54cd94633b "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/calendarsharingmessage/accept"
)

// CalendarSharingMessageRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\microsoft.graph.calendarSharingMessage
type CalendarSharingMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accept the accept property
func (m *CalendarSharingMessageRequestBuilder) Accept()(*ibaa6a71ab032a685fe85ce7cefea89706af5749ad58bb2bb38038e54cd94633b.AcceptRequestBuilder) {
    return ibaa6a71ab032a685fe85ce7cefea89706af5749ad58bb2bb38038e54cd94633b.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCalendarSharingMessageRequestBuilderInternal instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
func NewCalendarSharingMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    m := &CalendarSharingMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/microsoft.graph.calendarSharingMessage";
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
