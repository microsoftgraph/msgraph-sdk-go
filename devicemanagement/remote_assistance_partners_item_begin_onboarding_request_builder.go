package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RemoteAssistancePartnersItemBeginOnboardingRequestBuilder provides operations to call the beginOnboarding method.
type RemoteAssistancePartnersItemBeginOnboardingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilderInternal instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) {
    m := &RemoteAssistancePartnersItemBeginOnboardingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner%2Did}/beginOnboarding", pathParameters),
    }
    return m
}
// NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilder instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) Post(ctx context.Context, requestConfiguration *RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
