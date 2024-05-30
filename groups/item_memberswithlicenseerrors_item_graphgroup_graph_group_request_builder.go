package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder casts the previous resource to group.
type ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.group
type ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderGetQueryParameters
}
// NewItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderInternal instantiates a new ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder and sets the default values.
func NewItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder) {
    m := &ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}/graph.group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder instantiates a new ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder and sets the default values.
func NewItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.group
// returns a *RequestInformation when successful
func (m *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder) WithUrl(rawUrl string)(*ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder) {
    return NewItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
