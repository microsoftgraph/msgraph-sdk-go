package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder provides operations to call the beginOnboarding method.
type RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilderInternal instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder) {
    m := &RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner%2Did}/microsoft.graph.beginOnboarding";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder) Post(ctx context.Context, requestConfiguration *RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingBeginOnboardingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
