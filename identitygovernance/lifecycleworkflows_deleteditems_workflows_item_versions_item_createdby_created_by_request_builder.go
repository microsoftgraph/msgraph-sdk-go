package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderGetQueryParameters the user who created the workflow.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}/createdBy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the user who created the workflow.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) MailboxSettings()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) ServiceProvisioningErrors()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the user who created the workflow.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
