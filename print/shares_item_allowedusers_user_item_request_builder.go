package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SharesItemAllowedusersUserItemRequestBuilder builds and executes requests for operations under \print\shares\{printerShare-id}\allowedUsers\{user-id}
type SharesItemAllowedusersUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewSharesItemAllowedusersUserItemRequestBuilderInternal instantiates a new SharesItemAllowedusersUserItemRequestBuilder and sets the default values.
func NewSharesItemAllowedusersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedusersUserItemRequestBuilder) {
    m := &SharesItemAllowedusersUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/shares/{printerShare%2Did}/allowedUsers/{user%2Did}", pathParameters),
    }
    return m
}
// NewSharesItemAllowedusersUserItemRequestBuilder instantiates a new SharesItemAllowedusersUserItemRequestBuilder and sets the default values.
func NewSharesItemAllowedusersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedusersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharesItemAllowedusersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MailboxSettings the mailboxSettings property
// returns a *SharesItemAllowedusersItemMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *SharesItemAllowedusersUserItemRequestBuilder) MailboxSettings()(*SharesItemAllowedusersItemMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewSharesItemAllowedusersItemMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of print entities.
// returns a *SharesItemAllowedusersItemRefRequestBuilder when successful
func (m *SharesItemAllowedusersUserItemRequestBuilder) Ref()(*SharesItemAllowedusersItemRefRequestBuilder) {
    return NewSharesItemAllowedusersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *SharesItemAllowedusersItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *SharesItemAllowedusersUserItemRequestBuilder) ServiceProvisioningErrors()(*SharesItemAllowedusersItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewSharesItemAllowedusersItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
