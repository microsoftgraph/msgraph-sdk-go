package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRejectedsendersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\rejectedSenders\{directoryObject-id}
type ItemRejectedsendersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRejectedsendersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemRejectedsendersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRejectedsendersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRejectedsendersDirectoryObjectItemRequestBuilder) {
    m := &ItemRejectedsendersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/rejectedSenders/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemRejectedsendersDirectoryObjectItemRequestBuilder instantiates a new ItemRejectedsendersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRejectedsendersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRejectedsendersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRejectedsendersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of group entities.
// returns a *ItemRejectedsendersItemRefRequestBuilder when successful
func (m *ItemRejectedsendersDirectoryObjectItemRequestBuilder) Ref()(*ItemRejectedsendersItemRefRequestBuilder) {
    return NewItemRejectedsendersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
