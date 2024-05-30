package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder provides operations to manage the lastModifiedBy property of the microsoft.graph.identityGovernance.workflowBase entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderGetQueryParameters the unique identifier of the Microsoft Entra identity that last modified the workflow.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}/lastModifiedBy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the unique identifier of the Microsoft Entra identity that last modified the workflow.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) MailboxSettings()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) ServiceProvisioningErrors()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the unique identifier of the Microsoft Entra identity that last modified the workflow.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
