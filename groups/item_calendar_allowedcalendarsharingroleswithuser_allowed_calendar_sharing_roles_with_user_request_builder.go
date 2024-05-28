package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder provides operations to call the allowedCalendarSharingRoles method.
type ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters invoke function allowedCalendarSharingRoles
type ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters struct {
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
// ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters
}
// NewItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal instantiates a new ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder and sets the default values.
func NewItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, user *string)(*ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    m := &ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/allowedCalendarSharingRoles(User='{User}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if user != nil {
        m.BaseRequestBuilder.PathParameters["User"] = *user
    }
    return m
}
// NewItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder instantiates a new ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder and sets the default values.
func NewItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function allowedCalendarSharingRoles
// Deprecated: This method is obsolete. Use GetAsAllowedCalendarSharingRolesWithUserGetResponse instead.
// returns a ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable), nil
}
// GetAsAllowedCalendarSharingRolesWithUserGetResponse invoke function allowedCalendarSharingRoles
// returns a ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) GetAsAllowedCalendarSharingRolesWithUserGetResponse(ctx context.Context, requestConfiguration *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable), nil
}
// ToGetRequestInformation invoke function allowedCalendarSharingRoles
// returns a *RequestInformation when successful
func (m *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder when successful
func (m *ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    return NewItemCalendarAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
