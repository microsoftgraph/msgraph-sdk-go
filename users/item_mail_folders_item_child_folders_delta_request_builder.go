package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersDeltaRequestBuilder provides operations to call the delta method.
type ItemMailFoldersItemChildFoldersDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersDeltaRequestBuilderGetQueryParameters get a set of mail folders that have been added, deleted, or removed from the user's mailbox. A delta function call for mail folders in a mailbox is similar to a GET request, except that by appropriatelyapplying state tokens in one or more of these calls,you can query for incremental changes in the mail folders. This allows you to maintain and synchronizea local store of a user's mail folders without having to fetch all the mail folders of that mailbox from the server every time.
type ItemMailFoldersItemChildFoldersDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemMailFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemChildFoldersDeltaRequestBuilderGetQueryParameters
}
// NewItemMailFoldersItemChildFoldersDeltaRequestBuilderInternal instantiates a new DeltaRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersDeltaRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersDeltaRequestBuilder instantiates a new DeltaRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of mail folders that have been added, deleted, or removed from the user's mailbox. A delta function call for mail folders in a mailbox is similar to a GET request, except that by appropriatelyapplying state tokens in one or more of these calls,you can query for incremental changes in the mail folders. This allows you to maintain and synchronizea local store of a user's mail folders without having to fetch all the mail folders of that mailbox from the server every time.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailfolder-delta?view=graph-rest-1.0
func (m *ItemMailFoldersItemChildFoldersDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration)(ItemMailFoldersItemChildFoldersDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMailFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMailFoldersItemChildFoldersDeltaResponseable), nil
}
// ToGetRequestInformation get a set of mail folders that have been added, deleted, or removed from the user's mailbox. A delta function call for mail folders in a mailbox is similar to a GET request, except that by appropriatelyapplying state tokens in one or more of these calls,you can query for incremental changes in the mail folders. This allows you to maintain and synchronizea local store of a user's mail folders without having to fetch all the mail folders of that mailbox from the server every time.
func (m *ItemMailFoldersItemChildFoldersDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMailFoldersItemChildFoldersDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemChildFoldersDeltaRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
