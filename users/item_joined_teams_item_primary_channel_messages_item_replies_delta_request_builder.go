package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder provides operations to call the delta method.
type ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters invoke function delta
type ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters struct {
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
// ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters
}
// NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal instantiates a new ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) {
    m := &ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder instantiates a new ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaResponseable), nil
}
// GetAsDeltaGetResponse invoke function delta
// returns a ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder when successful
func (m *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) {
    return NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
