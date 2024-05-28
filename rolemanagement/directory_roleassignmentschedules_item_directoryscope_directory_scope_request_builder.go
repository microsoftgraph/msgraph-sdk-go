package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
type DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderGetQueryParameters the directory object that is the scope of the role eligibility or assignment. Read-only.
type DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderGetQueryParameters
}
// NewDirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderInternal instantiates a new DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    m := &DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentSchedules/{unifiedRoleAssignmentSchedule%2Did}/directoryScope{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder instantiates a new DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the directory object that is the scope of the role eligibility or assignment. Read-only.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// ToGetRequestInformation the directory object that is the scope of the role eligibility or assignment. Read-only.
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
