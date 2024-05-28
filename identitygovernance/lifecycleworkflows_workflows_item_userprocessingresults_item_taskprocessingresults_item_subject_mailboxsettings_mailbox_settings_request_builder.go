package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder builds and executes requests for operations under \identityGovernance\lifecycleWorkflows\workflows\{workflow-id}\userProcessingResults\{userProcessingResult-id}\taskProcessingResults\{taskProcessingResult-id}\subject\mailboxSettings
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetQueryParameters settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetQueryParameters
}
// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/subject/mailboxSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
// returns a MailboxSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMailboxSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable), nil
}
// Patch update property mailboxSettings value.
// returns a MailboxSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMailboxSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable), nil
}
// ToGetRequestInformation settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update property mailboxSettings value.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
