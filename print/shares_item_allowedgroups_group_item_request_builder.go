package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SharesItemAllowedgroupsGroupItemRequestBuilder builds and executes requests for operations under \print\shares\{printerShare-id}\allowedGroups\{group-id}
type SharesItemAllowedgroupsGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewSharesItemAllowedgroupsGroupItemRequestBuilderInternal instantiates a new SharesItemAllowedgroupsGroupItemRequestBuilder and sets the default values.
func NewSharesItemAllowedgroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedgroupsGroupItemRequestBuilder) {
    m := &SharesItemAllowedgroupsGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/shares/{printerShare%2Did}/allowedGroups/{group%2Did}", pathParameters),
    }
    return m
}
// NewSharesItemAllowedgroupsGroupItemRequestBuilder instantiates a new SharesItemAllowedgroupsGroupItemRequestBuilder and sets the default values.
func NewSharesItemAllowedgroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedgroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharesItemAllowedgroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of print entities.
// returns a *SharesItemAllowedgroupsItemRefRequestBuilder when successful
func (m *SharesItemAllowedgroupsGroupItemRequestBuilder) Ref()(*SharesItemAllowedgroupsItemRefRequestBuilder) {
    return NewSharesItemAllowedgroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *SharesItemAllowedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *SharesItemAllowedgroupsGroupItemRequestBuilder) ServiceProvisioningErrors()(*SharesItemAllowedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewSharesItemAllowedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
