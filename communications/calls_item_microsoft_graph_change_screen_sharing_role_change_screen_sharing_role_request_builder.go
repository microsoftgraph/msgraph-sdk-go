package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder provides operations to call the changeScreenSharingRole method.
type CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilderInternal instantiates a new ChangeScreenSharingRoleRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder) {
    m := &CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}/microsoft.graph.changeScreenSharingRole";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder instantiates a new ChangeScreenSharingRoleRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post allow applications to share screen content with the participants of a group call.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/call-changescreensharingrole?view=graph-rest-1.0
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder) Post(ctx context.Context, body CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation allow applications to share screen content with the participants of a group call.
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRoleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
