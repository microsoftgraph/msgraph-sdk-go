package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.appManagementPolicy entity.
type AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderGetQueryParameters collection of applications and service principals to which the policy is applied.
type AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewAppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal instantiates a new AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder and sets the default values.
func NewAppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) {
    m := &AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/appManagementPolicies/{appManagementPolicy%2Did}/appliesTo/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder instantiates a new AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder and sets the default values.
func NewAppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get collection of applications and service principals to which the policy is applied.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
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
// ToGetRequestInformation collection of applications and service principals to which the policy is applied.
// returns a *RequestInformation when successful
func (m *AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder when successful
func (m *AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*AppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) {
    return NewAppmanagementpoliciesItemAppliestoDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
