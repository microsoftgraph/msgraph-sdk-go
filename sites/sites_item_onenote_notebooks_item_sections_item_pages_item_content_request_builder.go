package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder provides operations to manage the media for the site entity.
type SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderInternal instantiates a new ContentRequestBuilder and sets the default values.
func NewSitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder) {
    m := &SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/content";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder instantiates a new ContentRequestBuilder and sets the default values.
func NewSitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the page's HTML content.
func (m *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePutRequestInformation the page's HTML content.
func (m *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder) CreatePutRequestInformation(ctx context.Context, body []byte, requestConfiguration *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the page's HTML content.
func (m *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder) Get(ctx context.Context, requestConfiguration *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put the page's HTML content.
func (m *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *SitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.CreatePutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
