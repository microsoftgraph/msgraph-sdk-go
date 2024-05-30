package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder provides operations to call the allowedCalendarSharingRoles method.
type ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters invoke function allowedCalendarSharingRoles
type ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters struct {
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
// ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters
}
// NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, user *string)(*ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/allowedCalendarSharingRoles(User='{User}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if user != nil {
        m.BaseRequestBuilder.PathParameters["User"] = *user
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function allowedCalendarSharingRoles
// Deprecated: This method is obsolete. Use GetAsAllowedCalendarSharingRolesWithUserGetResponse instead.
// returns a ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable), nil
}
// GetAsAllowedCalendarSharingRolesWithUserGetResponse invoke function allowedCalendarSharingRoles
// returns a ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) GetAsAllowedCalendarSharingRolesWithUserGetResponse(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable), nil
}
// ToGetRequestInformation invoke function allowedCalendarSharingRoles
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
