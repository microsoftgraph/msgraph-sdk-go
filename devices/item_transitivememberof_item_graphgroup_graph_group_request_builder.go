package devices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder casts the previous resource to group.
type ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderGetQueryParameters get the groups and administrative units that the device is a member of. This API request is transitive, and will also return all groups and administrative units the device is a nested member of.
type ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderGetQueryParameters
}
// NewItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderInternal instantiates a new ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder and sets the default values.
func NewItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder) {
    m := &ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/devices/{device%2Did}/transitiveMemberOf/{directoryObject%2Did}/graph.group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder instantiates a new ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder and sets the default values.
func NewItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the groups and administrative units that the device is a member of. This API request is transitive, and will also return all groups and administrative units the device is a nested member of.
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/device-list-transitivememberof?view=graph-rest-1.0
func (m *ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
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
// ToGetRequestInformation get the groups and administrative units that the device is a member of. This API request is transitive, and will also return all groups and administrative units the device is a nested member of.
// returns a *RequestInformation when successful
func (m *ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder) WithUrl(rawUrl string)(*ItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder) {
    return NewItemTransitivememberofItemGraphgroupGraphGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
