package invitations

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder provides operations to manage the invitedUserSponsors property of the microsoft.graph.invitation entity.
type ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderGetQueryParameters the users or groups who are sponsors of the invited user. Sponsors are users and groups that are responsible for guest users' privileges in the tenant and for keeping the guest users' information and access up to date.
type ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the invitedUserSponsors property of the microsoft.graph.invitation entity.
// returns a *ItemInvitedusersponsorsDirectoryObjectItemRequestBuilder when successful
func (m *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemInvitedusersponsorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemInvitedusersponsorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderInternal instantiates a new ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder and sets the default values.
func NewItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) {
    m := &ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/invitations/{invitation%2Did}/invitedUserSponsors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder instantiates a new ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder and sets the default values.
func NewItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemInvitedusersponsorsCountRequestBuilder when successful
func (m *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) Count()(*ItemInvitedusersponsorsCountRequestBuilder) {
    return NewItemInvitedusersponsorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the users or groups who are sponsors of the invited user. Sponsors are users and groups that are responsible for guest users' privileges in the tenant and for keeping the guest users' information and access up to date.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// ToGetRequestInformation the users or groups who are sponsors of the invited user. Sponsors are users and groups that are responsible for guest users' privileges in the tenant and for keeping the guest users' information and access up to date.
// returns a *RequestInformation when successful
func (m *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder when successful
func (m *ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) WithUrl(rawUrl string)(*ItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder) {
    return NewItemInvitedusersponsorsInvitedUserSponsorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
