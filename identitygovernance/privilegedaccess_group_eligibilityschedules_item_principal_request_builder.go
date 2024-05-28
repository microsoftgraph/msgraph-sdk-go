package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupEligibilitySchedule entity.
type PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderGetQueryParameters references the principal that's in the scope of this membership or ownership eligibility request to the group that's governed by PIM. Supports $expand.
type PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderGetQueryParameters
}
// NewPrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderInternal instantiates a new PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder) {
    m := &PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilitySchedules/{privilegedAccessGroupEligibilitySchedule%2Did}/principal{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder instantiates a new PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get references the principal that's in the scope of this membership or ownership eligibility request to the group that's governed by PIM. Supports $expand.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
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
// ToGetRequestInformation references the principal that's in the scope of this membership or ownership eligibility request to the group that's governed by PIM. Supports $expand.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulesItemPrincipalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
