package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder provides operations to call the createDownloadUrl method.
type ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal instantiates a new ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    m := &ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/logCollectionRequests/{deviceLogCollectionResponse%2Did}/createDownloadUrl", pathParameters),
    }
    return m
}
// NewManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder instantiates a new ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createDownloadUrl
// Deprecated: This method is obsolete. Use PostAsCreateDownloadUrlPostResponse instead.
// returns a ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseable), nil
}
// PostAsCreateDownloadUrlPostResponse invoke action createDownloadUrl
// returns a ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) PostAsCreateDownloadUrlPostResponse(ctx context.Context, requestConfiguration *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action createDownloadUrl
// returns a *RequestInformation when successful
func (m *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
