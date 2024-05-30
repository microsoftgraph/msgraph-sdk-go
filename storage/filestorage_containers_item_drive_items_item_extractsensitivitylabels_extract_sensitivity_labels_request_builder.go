package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder provides operations to call the extractSensitivityLabels method.
type FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/extractSensitivityLabels", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action extractSensitivityLabels
// returns a ExtractSensitivityLabelsResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ExtractSensitivityLabelsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateExtractSensitivityLabelsResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ExtractSensitivityLabelsResultable), nil
}
// ToPostRequestInformation invoke action extractSensitivityLabels
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
