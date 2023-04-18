package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SharesItemAllowedUsersUserItemRequestBuilder builds and executes requests for operations under \print\shares\{printerShare-id}\allowedUsers\{user-id}
type SharesItemAllowedUsersUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewSharesItemAllowedUsersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewSharesItemAllowedUsersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedUsersUserItemRequestBuilder) {
    m := &SharesItemAllowedUsersUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/shares/{printerShare%2Did}/allowedUsers/{user%2Did}", pathParameters),
    }
    return m
}
// NewSharesItemAllowedUsersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewSharesItemAllowedUsersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedUsersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharesItemAllowedUsersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MailboxSettings the mailboxSettings property
func (m *SharesItemAllowedUsersUserItemRequestBuilder) MailboxSettings()(*SharesItemAllowedUsersItemMailboxSettingsRequestBuilder) {
    return NewSharesItemAllowedUsersItemMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of print entities.
func (m *SharesItemAllowedUsersUserItemRequestBuilder) Ref()(*SharesItemAllowedUsersItemRefRequestBuilder) {
    return NewSharesItemAllowedUsersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
