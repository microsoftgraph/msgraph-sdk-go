package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.contactFolder entity.
type ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderGetQueryParameters the collection of child folders in the folder. Navigation property. Read-only. Nullable.
type ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderGetQueryParameters struct {
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
// ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderGetQueryParameters
}
// ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByContactFolderId1 provides operations to manage the childFolders property of the microsoft.graph.contactFolder entity.
// returns a *ItemContactfoldersItemChildfoldersContactFolderItemRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) ByContactFolderId1(contactFolderId1 string)(*ItemContactfoldersItemChildfoldersContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if contactFolderId1 != "" {
        urlTplParams["contactFolder%2Did1"] = contactFolderId1
    }
    return NewItemContactfoldersItemChildfoldersContactFolderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemContactfoldersItemChildfoldersChildFoldersRequestBuilderInternal instantiates a new ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersChildFoldersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) {
    m := &ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/childFolders{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemContactfoldersItemChildfoldersChildFoldersRequestBuilder instantiates a new ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersChildFoldersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContactfoldersItemChildfoldersChildFoldersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemContactfoldersItemChildfoldersCountRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) Count()(*ItemContactfoldersItemChildfoldersCountRequestBuilder) {
    return NewItemContactfoldersItemChildfoldersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemContactfoldersItemChildfoldersDeltaRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) Delta()(*ItemContactfoldersItemChildfoldersDeltaRequestBuilder) {
    return NewItemContactfoldersItemChildfoldersDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of child folders in the folder. Navigation property. Read-only. Nullable.
// returns a ContactFolderCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContactFolderCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderCollectionResponseable), nil
}
// Post create new navigation property to childFolders for users
// returns a ContactFolderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderable, requestConfiguration *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContactFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderable), nil
}
// ToGetRequestInformation the collection of child folders in the folder. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to childFolders for users
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderable, requestConfiguration *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) WithUrl(rawUrl string)(*ItemContactfoldersItemChildfoldersChildFoldersRequestBuilder) {
    return NewItemContactfoldersItemChildfoldersChildFoldersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
