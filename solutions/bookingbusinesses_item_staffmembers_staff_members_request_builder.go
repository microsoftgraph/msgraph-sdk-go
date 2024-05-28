package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingbusinessesItemStaffmembersStaffMembersRequestBuilder provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
type BookingbusinessesItemStaffmembersStaffMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingbusinessesItemStaffmembersStaffMembersRequestBuilderGetQueryParameters get a list of bookingStaffMember objects in the specified bookingBusiness.
type BookingbusinessesItemStaffmembersStaffMembersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// BookingbusinessesItemStaffmembersStaffMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemStaffmembersStaffMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingbusinessesItemStaffmembersStaffMembersRequestBuilderGetQueryParameters
}
// BookingbusinessesItemStaffmembersStaffMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemStaffmembersStaffMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBookingStaffMemberBaseId provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder when successful
func (m *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) ByBookingStaffMemberBaseId(bookingStaffMemberBaseId string)(*BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if bookingStaffMemberBaseId != "" {
        urlTplParams["bookingStaffMemberBase%2Did"] = bookingStaffMemberBaseId
    }
    return NewBookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBookingbusinessesItemStaffmembersStaffMembersRequestBuilderInternal instantiates a new BookingbusinessesItemStaffmembersStaffMembersRequestBuilder and sets the default values.
func NewBookingbusinessesItemStaffmembersStaffMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) {
    m := &BookingbusinessesItemStaffmembersStaffMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/staffMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBookingbusinessesItemStaffmembersStaffMembersRequestBuilder instantiates a new BookingbusinessesItemStaffmembersStaffMembersRequestBuilder and sets the default values.
func NewBookingbusinessesItemStaffmembersStaffMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingbusinessesItemStaffmembersStaffMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BookingbusinessesItemStaffmembersCountRequestBuilder when successful
func (m *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) Count()(*BookingbusinessesItemStaffmembersCountRequestBuilder) {
    return NewBookingbusinessesItemStaffmembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of bookingStaffMember objects in the specified bookingBusiness.
// returns a BookingStaffMemberBaseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-list-staffmembers?view=graph-rest-1.0
func (m *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingbusinessesItemStaffmembersStaffMembersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingStaffMemberBaseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseCollectionResponseable), nil
}
// Post create a new bookingStaffMember in the specified bookingBusiness.
// returns a BookingStaffMemberBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-post-staffmembers?view=graph-rest-1.0
func (m *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, requestConfiguration *BookingbusinessesItemStaffmembersStaffMembersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingStaffMemberBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable), nil
}
// ToGetRequestInformation get a list of bookingStaffMember objects in the specified bookingBusiness.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BookingbusinessesItemStaffmembersStaffMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new bookingStaffMember in the specified bookingBusiness.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, requestConfiguration *BookingbusinessesItemStaffmembersStaffMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder when successful
func (m *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) WithUrl(rawUrl string)(*BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) {
    return NewBookingbusinessesItemStaffmembersStaffMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
