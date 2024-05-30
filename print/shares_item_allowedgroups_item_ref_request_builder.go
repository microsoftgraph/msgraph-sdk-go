package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SharesItemAllowedgroupsItemRefRequestBuilder provides operations to manage the collection of print entities.
type SharesItemAllowedgroupsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SharesItemAllowedgroupsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharesItemAllowedgroupsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSharesItemAllowedgroupsItemRefRequestBuilderInternal instantiates a new SharesItemAllowedgroupsItemRefRequestBuilder and sets the default values.
func NewSharesItemAllowedgroupsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedgroupsItemRefRequestBuilder) {
    m := &SharesItemAllowedgroupsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/shares/{printerShare%2Did}/allowedGroups/{group%2Did}/$ref", pathParameters),
    }
    return m
}
// NewSharesItemAllowedgroupsItemRefRequestBuilder instantiates a new SharesItemAllowedgroupsItemRefRequestBuilder and sets the default values.
func NewSharesItemAllowedgroupsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedgroupsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharesItemAllowedgroupsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete revoke the specified group's access to submit print jobs to the associated printerShare.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/printershare-delete-allowedgroup?view=graph-rest-1.0
func (m *SharesItemAllowedgroupsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *SharesItemAllowedgroupsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation revoke the specified group's access to submit print jobs to the associated printerShare.
// returns a *RequestInformation when successful
func (m *SharesItemAllowedgroupsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SharesItemAllowedgroupsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SharesItemAllowedgroupsItemRefRequestBuilder when successful
func (m *SharesItemAllowedgroupsItemRefRequestBuilder) WithUrl(rawUrl string)(*SharesItemAllowedgroupsItemRefRequestBuilder) {
    return NewSharesItemAllowedgroupsItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
