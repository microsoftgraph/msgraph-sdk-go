package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder provides operations to call the onenotePatchContent method.
type ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilderInternal instantiates a new OnenotePatchContentRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder) {
    m := &ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/microsoft.graph.onenotePatchContent";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder instantiates a new OnenotePatchContentRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action onenotePatchContent
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder) Post(ctx context.Context, body ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBodyable, requestConfiguration *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action onenotePatchContent
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBodyable, requestConfiguration *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
