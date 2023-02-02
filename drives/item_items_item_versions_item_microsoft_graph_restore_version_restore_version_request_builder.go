package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder provides operations to call the restoreVersion method.
type ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilderInternal instantiates a new RestoreVersionRequestBuilder and sets the default values.
func NewItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder) {
    m := &ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/versions/{driveItemVersion%2Did}/microsoft.graph.restoreVersion";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder instantiates a new RestoreVersionRequestBuilder and sets the default values.
func NewItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a previous version of a DriveItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the file.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/driveitemversion-restore?view=graph-rest-1.0
func (m *ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation restore a previous version of a DriveItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the file.
func (m *ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRestoreVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
