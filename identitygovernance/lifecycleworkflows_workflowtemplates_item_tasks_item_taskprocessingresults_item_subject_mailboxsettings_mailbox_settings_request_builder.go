package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder builds and executes requests for operations under \identityGovernance\lifecycleWorkflows\workflowTemplates\{workflowTemplate-id}\tasks\{task-id}\taskProcessingResults\{taskProcessingResult-id}\subject\mailboxSettings
type LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetQueryParameters settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
type LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetQueryParameters
}
// LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) {
    m := &LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflowTemplates/{workflowTemplate%2Did}/tasks/{task%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/subject/mailboxSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder instantiates a new LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
// returns a MailboxSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, error) {
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
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, error) {
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
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailboxSettingsable, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsItemSubjectMailboxsettingsMailboxSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
