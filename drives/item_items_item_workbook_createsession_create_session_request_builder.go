package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder provides operations to call the createSession method.
type ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilderInternal instantiates a new ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder) {
    m := &ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/createSession", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder instantiates a new ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new workbook session.  Excel APIs can be called in one of two modes:  To represent the session in the API, use the workbook-session-id: {session-id} header.  In some cases, creating a new session requires an indeterminate time to complete. Microsoft Graph also provides a long running operations pattern. This pattern provides a way to poll for creation status updates, without waiting for the creation to complete. The following are the steps:
// returns a WorkbookSessionInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/workbook-createsession?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookCreatesessionCreateSessionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSessionInfoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookSessionInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSessionInfoable), nil
}
// ToPostRequestInformation create a new workbook session.  Excel APIs can be called in one of two modes:  To represent the session in the API, use the workbook-session-id: {session-id} header.  In some cases, creating a new session requires an indeterminate time to complete. Microsoft Graph also provides a long running operations pattern. This pattern provides a way to poll for creation status updates, without waiting for the creation to complete. The following are the steps:
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookCreatesessionCreateSessionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder when successful
func (m *ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder) {
    return NewItemItemsItemWorkbookCreatesessionCreateSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
