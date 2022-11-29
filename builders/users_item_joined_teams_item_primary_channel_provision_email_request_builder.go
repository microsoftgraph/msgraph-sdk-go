package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder provides operations to call the provisionEmail method.
type UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderInternal instantiates a new ProvisionEmailRequestBuilder and sets the default values.
func NewUsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder) {
    m := &UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/microsoft.graph.provisionEmail";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder instantiates a new ProvisionEmailRequestBuilder and sets the default values.
func NewUsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation provision an email address for a channel. Microsoft Teams doesn't automatically provision an email address for a **channel** by default. To have Teams provision an email address, you can call **provisionEmail**, or through the Teams user interface, select **Get email address**, which triggers Teams to generate an email address if it has not already provisioned one. To remove the email address of a **channel**, use the removeEmail method.
func (m *UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post provision an email address for a channel. Microsoft Teams doesn't automatically provision an email address for a **channel** by default. To have Teams provision an email address, you can call **provisionEmail**, or through the Teams user interface, select **Get email address**, which triggers Teams to generate an email address if it has not already provisioned one. To remove the email address of a **channel**, use the removeEmail method.
func (m *UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProvisionChannelEmailResultable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateProvisionChannelEmailResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProvisionChannelEmailResultable), nil
}
