package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder provides operations to call the supportedLanguages method.
type ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetQueryParameters get the list of locales and languages that are supported for the user, as configured on the user's mailbox server. When setting up an Outlook client, the user selects the preferred language from this supported list. You can subsequently get the preferred language bygetting the user's mailbox settings.
type ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetQueryParameters
}
// NewItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderInternal instantiates a new ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder and sets the default values.
func NewItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder) {
    m := &ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook/supportedLanguages(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder instantiates a new ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder and sets the default values.
func NewItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of locales and languages that are supported for the user, as configured on the user's mailbox server. When setting up an Outlook client, the user selects the preferred language from this supported list. You can subsequently get the preferred language bygetting the user's mailbox settings.
// Deprecated: This method is obsolete. Use GetAsSupportedLanguagesGetResponse instead.
// returns a ItemOutlookSupportedlanguagesSupportedLanguagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/outlookuser-supportedlanguages?view=graph-rest-1.0
func (m *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetRequestConfiguration)(ItemOutlookSupportedlanguagesSupportedLanguagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookSupportedlanguagesSupportedLanguagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookSupportedlanguagesSupportedLanguagesResponseable), nil
}
// GetAsSupportedLanguagesGetResponse get the list of locales and languages that are supported for the user, as configured on the user's mailbox server. When setting up an Outlook client, the user selects the preferred language from this supported list. You can subsequently get the preferred language bygetting the user's mailbox settings.
// returns a ItemOutlookSupportedlanguagesSupportedLanguagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/outlookuser-supportedlanguages?view=graph-rest-1.0
func (m *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder) GetAsSupportedLanguagesGetResponse(ctx context.Context, requestConfiguration *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetRequestConfiguration)(ItemOutlookSupportedlanguagesSupportedLanguagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookSupportedlanguagesSupportedLanguagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookSupportedlanguagesSupportedLanguagesGetResponseable), nil
}
// ToGetRequestInformation get the list of locales and languages that are supported for the user, as configured on the user's mailbox server. When setting up an Outlook client, the user selects the preferred language from this supported list. You can subsequently get the preferred language bygetting the user's mailbox settings.
// returns a *RequestInformation when successful
func (m *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder when successful
func (m *ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder) WithUrl(rawUrl string)(*ItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder) {
    return NewItemOutlookSupportedlanguagesSupportedLanguagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
