package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder provides operations to call the createDownloadUrl method.
type ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal instantiates a new ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    m := &ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/logCollectionRequests/{deviceLogCollectionResponse%2Did}/createDownloadUrl", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder instantiates a new ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createDownloadUrl
// Deprecated: This method is obsolete. Use PostAsCreateDownloadUrlPostResponse instead.
// returns a ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseable), nil
}
// PostAsCreateDownloadUrlPostResponse invoke action createDownloadUrl
// returns a ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) PostAsCreateDownloadUrlPostResponse(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action createDownloadUrl
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
