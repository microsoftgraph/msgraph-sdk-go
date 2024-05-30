package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAcceptedsendersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\acceptedSenders\{directoryObject-id}
type ItemAcceptedsendersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAcceptedsendersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemAcceptedsendersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAcceptedsendersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAcceptedsendersDirectoryObjectItemRequestBuilder) {
    m := &ItemAcceptedsendersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/acceptedSenders/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemAcceptedsendersDirectoryObjectItemRequestBuilder instantiates a new ItemAcceptedsendersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAcceptedsendersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAcceptedsendersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAcceptedsendersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of group entities.
// returns a *ItemAcceptedsendersItemRefRequestBuilder when successful
func (m *ItemAcceptedsendersDirectoryObjectItemRequestBuilder) Ref()(*ItemAcceptedsendersItemRefRequestBuilder) {
    return NewItemAcceptedsendersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
