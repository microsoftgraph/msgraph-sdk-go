package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder provides operations to call the allowedCalendarSharingRoles method.
type ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters invoke function allowedCalendarSharingRoles
type ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters struct {
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
// ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetQueryParameters
}
// NewItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal instantiates a new ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder and sets the default values.
func NewItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, user *string)(*ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    m := &ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/allowedCalendarSharingRoles(User='{User}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if user != nil {
        m.BaseRequestBuilder.PathParameters["User"] = *user
    }
    return m
}
// NewItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder instantiates a new ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder and sets the default values.
func NewItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function allowedCalendarSharingRoles
// Deprecated: This method is obsolete. Use GetAsAllowedCalendarSharingRolesWithUserGetResponse instead.
// returns a ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable), nil
}
// GetAsAllowedCalendarSharingRolesWithUserGetResponse invoke function allowedCalendarSharingRoles
// returns a ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) GetAsAllowedCalendarSharingRolesWithUserGetResponse(ctx context.Context, requestConfiguration *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable), nil
}
// ToGetRequestInformation invoke function allowedCalendarSharingRoles
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder when successful
func (m *ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder) {
    return NewItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
