package places

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGraphroomlistGraphRoomListRequestBuilder casts the previous resource to roomList.
type ItemGraphroomlistGraphRoomListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGraphroomlistGraphRoomListRequestBuilderGetQueryParameters get the item of type microsoft.graph.place as microsoft.graph.roomList
type ItemGraphroomlistGraphRoomListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemGraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGraphroomlistGraphRoomListRequestBuilderGetQueryParameters
}
// NewItemGraphroomlistGraphRoomListRequestBuilderInternal instantiates a new ItemGraphroomlistGraphRoomListRequestBuilder and sets the default values.
func NewItemGraphroomlistGraphRoomListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGraphroomlistGraphRoomListRequestBuilder) {
    m := &ItemGraphroomlistGraphRoomListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/places/{place%2Did}/graph.roomList{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemGraphroomlistGraphRoomListRequestBuilder instantiates a new ItemGraphroomlistGraphRoomListRequestBuilder and sets the default values.
func NewItemGraphroomlistGraphRoomListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGraphroomlistGraphRoomListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGraphroomlistGraphRoomListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.place as microsoft.graph.roomList
// returns a RoomListable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGraphroomlistGraphRoomListRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoomListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRoomListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoomListable), nil
}
// Rooms provides operations to manage the rooms property of the microsoft.graph.roomList entity.
// returns a *ItemGraphroomlistRoomsRequestBuilder when successful
func (m *ItemGraphroomlistGraphRoomListRequestBuilder) Rooms()(*ItemGraphroomlistRoomsRequestBuilder) {
    return NewItemGraphroomlistRoomsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.place as microsoft.graph.roomList
// returns a *RequestInformation when successful
func (m *ItemGraphroomlistGraphRoomListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGraphroomlistGraphRoomListRequestBuilder when successful
func (m *ItemGraphroomlistGraphRoomListRequestBuilder) WithUrl(rawUrl string)(*ItemGraphroomlistGraphRoomListRequestBuilder) {
    return NewItemGraphroomlistGraphRoomListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
