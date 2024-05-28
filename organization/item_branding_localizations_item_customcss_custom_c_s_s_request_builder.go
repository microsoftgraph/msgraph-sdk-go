package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder provides operations to manage the media for the organization entity.
type ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderInternal instantiates a new ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder and sets the default values.
func NewItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) {
    m := &ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/branding/localizations/{organizationalBrandingLocalization%2Did}/customCSS", pathParameters),
    }
    return m
}
// NewItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder instantiates a new ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder and sets the default values.
func NewItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete cSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get cSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put cSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) Put(ctx context.Context, body []byte, contentType *string, requestConfiguration *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderPutRequestConfiguration)([]byte, error) {
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
// ToDeleteRequestInformation cSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
// returns a *RequestInformation when successful
func (m *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation cSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
// returns a *RequestInformation when successful
func (m *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "image/bmp, image/jpg, image/jpeg, image/gif, image/vnd.microsoft.icon, image/png, image/tiff, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation cSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
// returns a *RequestInformation when successful
func (m *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, contentType *string, requestConfiguration *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder when successful
func (m *ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) WithUrl(rawUrl string)(*ItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder) {
    return NewItemBrandingLocalizationsItemCustomcssCustomCSSRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
