package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder provides operations to call the revokeSignInSessions method.
type MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilderInternal instantiates a new RevokeSignInSessionsRequestBuilder and sets the default values.
func NewMicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder) {
    m := &MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.revokeSignInSessions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder instantiates a new RevokeSignInSessionsRequestBuilder and sets the default values.
func NewMicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action revokeSignInSessions
func (m *MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration)(MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsResponseable), nil
}
// ToPostRequestInformation invoke action revokeSignInSessions
func (m *MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphRevokeSignInSessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
