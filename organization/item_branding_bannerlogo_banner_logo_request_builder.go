package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemBrandingBannerlogoBannerLogoRequestBuilder provides operations to manage the media for the organization entity.
type ItemBrandingBannerlogoBannerLogoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBrandingBannerlogoBannerLogoRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingBannerlogoBannerLogoRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBrandingBannerlogoBannerLogoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingBannerlogoBannerLogoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBrandingBannerlogoBannerLogoRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingBannerlogoBannerLogoRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBrandingBannerlogoBannerLogoRequestBuilderInternal instantiates a new ItemBrandingBannerlogoBannerLogoRequestBuilder and sets the default values.
func NewItemBrandingBannerlogoBannerLogoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingBannerlogoBannerLogoRequestBuilder) {
    m := &ItemBrandingBannerlogoBannerLogoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/branding/bannerLogo", pathParameters),
    }
    return m
}
// NewItemBrandingBannerlogoBannerLogoRequestBuilder instantiates a new ItemBrandingBannerlogoBannerLogoRequestBuilder and sets the default values.
func NewItemBrandingBannerlogoBannerLogoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingBannerlogoBannerLogoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBrandingBannerlogoBannerLogoRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBrandingBannerlogoBannerLogoRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemBrandingBannerlogoBannerLogoRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBrandingBannerlogoBannerLogoRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBrandingBannerlogoBannerLogoRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBrandingBannerlogoBannerLogoRequestBuilder) Put(ctx context.Context, body []byte, contentType *string, requestConfiguration *ItemBrandingBannerlogoBannerLogoRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, contentType, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToDeleteRequestInformation a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// returns a *RequestInformation when successful
func (m *ItemBrandingBannerlogoBannerLogoRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingBannerlogoBannerLogoRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// returns a *RequestInformation when successful
func (m *ItemBrandingBannerlogoBannerLogoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingBannerlogoBannerLogoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "image/bmp, image/jpg, image/jpeg, image/gif, image/vnd.microsoft.icon, image/png, image/tiff, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// returns a *RequestInformation when successful
func (m *ItemBrandingBannerlogoBannerLogoRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, contentType *string, requestConfiguration *ItemBrandingBannerlogoBannerLogoRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, *contentType)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemBrandingBannerlogoBannerLogoRequestBuilder when successful
func (m *ItemBrandingBannerlogoBannerLogoRequestBuilder) WithUrl(rawUrl string)(*ItemBrandingBannerlogoBannerLogoRequestBuilder) {
    return NewItemBrandingBannerlogoBannerLogoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
