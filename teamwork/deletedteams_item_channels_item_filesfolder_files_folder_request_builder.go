package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
type DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderGetQueryParameters metadata for the location where the channel's files are stored.
type DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderGetQueryParameters
}
// NewDeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal instantiates a new DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    m := &DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/filesFolder{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder instantiates a new DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the teamwork entity.
// returns a *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) Content()(*DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) {
    return NewDeletedteamsItemChannelsItemFilesfolderContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get metadata for the location where the channel's files are stored.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
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
// ToGetRequestInformation metadata for the location where the channel's files are stored.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    return NewDeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
