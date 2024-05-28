package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemContactfoldersItemChildfoldersDeltaRequestBuilder provides operations to call the delta method.
type ItemContactfoldersItemChildfoldersDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetQueryParameters get a set of contact folders that have been added, deleted, or removed from the user's mailbox. A delta function call for contact folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the contact folders. This allows you to maintain and synchronize a local store of a user's contact folders without having to fetch all the contact folders of that mailbox from the server every time.
type ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetQueryParameters struct {
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
// ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetQueryParameters
}
// NewItemContactfoldersItemChildfoldersDeltaRequestBuilderInternal instantiates a new ItemContactfoldersItemChildfoldersDeltaRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersDeltaRequestBuilder) {
    m := &ItemContactfoldersItemChildfoldersDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/childFolders/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemContactfoldersItemChildfoldersDeltaRequestBuilder instantiates a new ItemContactfoldersItemChildfoldersDeltaRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContactfoldersItemChildfoldersDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of contact folders that have been added, deleted, or removed from the user's mailbox. A delta function call for contact folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the contact folders. This allows you to maintain and synchronize a local store of a user's contact folders without having to fetch all the contact folders of that mailbox from the server every time.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemContactfoldersItemChildfoldersDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contactfolder-delta?view=graph-rest-1.0
func (m *ItemContactfoldersItemChildfoldersDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetRequestConfiguration)(ItemContactfoldersItemChildfoldersDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContactfoldersItemChildfoldersDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContactfoldersItemChildfoldersDeltaResponseable), nil
}
// GetAsDeltaGetResponse get a set of contact folders that have been added, deleted, or removed from the user's mailbox. A delta function call for contact folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the contact folders. This allows you to maintain and synchronize a local store of a user's contact folders without having to fetch all the contact folders of that mailbox from the server every time.
// returns a ItemContactfoldersItemChildfoldersDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contactfolder-delta?view=graph-rest-1.0
func (m *ItemContactfoldersItemChildfoldersDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetRequestConfiguration)(ItemContactfoldersItemChildfoldersDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContactfoldersItemChildfoldersDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContactfoldersItemChildfoldersDeltaGetResponseable), nil
}
// ToGetRequestInformation get a set of contact folders that have been added, deleted, or removed from the user's mailbox. A delta function call for contact folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the contact folders. This allows you to maintain and synchronize a local store of a user's contact folders without having to fetch all the contact folders of that mailbox from the server every time.
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemContactfoldersItemChildfoldersDeltaRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemContactfoldersItemChildfoldersDeltaRequestBuilder) {
    return NewItemContactfoldersItemChildfoldersDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
