package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailfoldersMailFolderItemRequestBuilder provides operations to manage the mailFolders property of the microsoft.graph.user entity.
type ItemMailfoldersMailFolderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersMailFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersMailFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMailfoldersMailFolderItemRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type ItemMailfoldersMailFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Include Hidden Folders
    IncludeHiddenFolders *string `uriparametername:"includeHiddenFolders"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailfoldersMailFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersMailFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailfoldersMailFolderItemRequestBuilderGetQueryParameters
}
// ItemMailfoldersMailFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersMailFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChildFolders provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
// returns a *ItemMailfoldersItemChildfoldersChildFoldersRequestBuilder when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) ChildFolders()(*ItemMailfoldersItemChildfoldersChildFoldersRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersChildFoldersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailfoldersMailFolderItemRequestBuilderInternal instantiates a new ItemMailfoldersMailFolderItemRequestBuilder and sets the default values.
func NewItemMailfoldersMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersMailFolderItemRequestBuilder) {
    m := &ItemMailfoldersMailFolderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?%24expand,%24select,includeHiddenFolders*}", pathParameters),
    }
    return m
}
// NewItemMailfoldersMailFolderItemRequestBuilder instantiates a new ItemMailfoldersMailFolderItemRequestBuilder and sets the default values.
func NewItemMailfoldersMailFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersMailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *ItemMailfoldersItemCopyRequestBuilder when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) Copy()(*ItemMailfoldersItemCopyRequestBuilder) {
    return NewItemMailfoldersItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property mailFolders for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersMailFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMailfoldersMailFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the user's mail folders. Read-only. Nullable.
// returns a MailFolderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersMailFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailfoldersMailFolderItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMailFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable), nil
}
// MessageRules provides operations to manage the messageRules property of the microsoft.graph.mailFolder entity.
// returns a *ItemMailfoldersItemMessagerulesMessageRulesRequestBuilder when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) MessageRules()(*ItemMailfoldersItemMessagerulesMessageRulesRequestBuilder) {
    return NewItemMailfoldersItemMessagerulesMessageRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
// returns a *ItemMailfoldersItemMessagesRequestBuilder when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) Messages()(*ItemMailfoldersItemMessagesRequestBuilder) {
    return NewItemMailfoldersItemMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Move provides operations to call the move method.
// returns a *ItemMailfoldersItemMoveRequestBuilder when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) Move()(*ItemMailfoldersItemMoveRequestBuilder) {
    return NewItemMailfoldersItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property mailFolders in users
// returns a MailFolderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersMailFolderItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable, requestConfiguration *ItemMailfoldersMailFolderItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMailFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable), nil
}
// ToDeleteRequestInformation delete navigation property mailFolders for users
// returns a *RequestInformation when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersMailFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the user's mail folders. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersMailFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mailFolders in users
// returns a *RequestInformation when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable, requestConfiguration *ItemMailfoldersMailFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemMailfoldersMailFolderItemRequestBuilder when successful
func (m *ItemMailfoldersMailFolderItemRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersMailFolderItemRequestBuilder) {
    return NewItemMailfoldersMailFolderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
