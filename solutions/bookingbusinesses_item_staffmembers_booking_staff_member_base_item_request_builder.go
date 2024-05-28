package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
type BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderGetQueryParameters get the properties and relationships of a bookingStaffMember in the specified bookingBusiness.
type BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderGetQueryParameters
}
// BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderInternal instantiates a new BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder and sets the default values.
func NewBookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) {
    m := &BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/staffMembers/{bookingStaffMemberBase%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder instantiates a new BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder and sets the default values.
func NewBookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a bookingStaffMember in the specified bookingBusiness.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingstaffmember-delete?view=graph-rest-1.0
func (m *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties and relationships of a bookingStaffMember in the specified bookingBusiness.
// returns a BookingStaffMemberBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingstaffmember-get?view=graph-rest-1.0
func (m *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of a bookingStaffMember in the specified bookingBusiness.
// returns a BookingStaffMemberBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingstaffmember-update?view=graph-rest-1.0
func (m *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, requestConfiguration *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete a bookingStaffMember in the specified bookingBusiness.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of a bookingStaffMember in the specified bookingBusiness.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a bookingStaffMember in the specified bookingBusiness.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, requestConfiguration *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder when successful
func (m *BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) WithUrl(rawUrl string)(*BookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder) {
    return NewBookingbusinessesItemStaffmembersBookingStaffMemberBaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
