package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChannelsItemFilesfolderFilesFolderRequestBuilder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
type ItemChannelsItemFilesfolderFilesFolderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChannelsItemFilesfolderFilesFolderRequestBuilderGetQueryParameters get the metadata for the location where the files of a channel are stored. 
type ItemChannelsItemFilesfolderFilesFolderRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChannelsItemFilesfolderFilesFolderRequestBuilderGetQueryParameters
}
// NewItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal instantiates a new ItemChannelsItemFilesfolderFilesFolderRequestBuilder and sets the default values.
func NewItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    m := &ItemChannelsItemFilesfolderFilesFolderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}/filesFolder{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemChannelsItemFilesfolderFilesFolderRequestBuilder instantiates a new ItemChannelsItemFilesfolderFilesFolderRequestBuilder and sets the default values.
func NewItemChannelsItemFilesfolderFilesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the team entity.
// returns a *ItemChannelsItemFilesfolderContentRequestBuilder when successful
func (m *ItemChannelsItemFilesfolderFilesFolderRequestBuilder) Content()(*ItemChannelsItemFilesfolderContentRequestBuilder) {
    return NewItemChannelsItemFilesfolderContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the metadata for the location where the files of a channel are stored. 
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-get-filesfolder?view=graph-rest-1.0
func (m *ItemChannelsItemFilesfolderFilesFolderRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ToGetRequestInformation get the metadata for the location where the files of a channel are stored. 
// returns a *RequestInformation when successful
func (m *ItemChannelsItemFilesfolderFilesFolderRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChannelsItemFilesfolderFilesFolderRequestBuilder when successful
func (m *ItemChannelsItemFilesfolderFilesFolderRequestBuilder) WithUrl(rawUrl string)(*ItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    return NewItemChannelsItemFilesfolderFilesFolderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
