package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder provides operations to manage the media for the teamwork entity.
type DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderGetQueryParameters the content stream, if the item represents a file.
type DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderGetQueryParameters struct {
    // Format of the content
    Format *string `uriparametername:"%24format"`
}
// DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderGetQueryParameters
}
// DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedteamsItemChannelsItemFilesfolderContentRequestBuilderInternal instantiates a new DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemFilesfolderContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) {
    m := &DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/filesFolder/content{?%24format*}", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsItemFilesfolderContentRequestBuilder instantiates a new DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemFilesfolderContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsItemFilesfolderContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the content stream, if the item represents a file.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the content stream, if the item represents a file.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put the content stream, if the item represents a file.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation the content stream, if the item represents a file.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the content stream, if the item represents a file.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the content stream, if the item represents a file.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsItemFilesfolderContentRequestBuilder) {
    return NewDeletedteamsItemChannelsItemFilesfolderContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
