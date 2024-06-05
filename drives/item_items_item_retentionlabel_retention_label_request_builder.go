package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemRetentionlabelRetentionLabelRequestBuilder provides operations to manage the retentionLabel property of the microsoft.graph.driveItem entity.
type ItemItemsItemRetentionlabelRetentionLabelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemRetentionlabelRetentionLabelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemRetentionlabelRetentionLabelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemRetentionlabelRetentionLabelRequestBuilderGetQueryParameters information about retention label and settings enforced on the driveItem. Read-write.
type ItemItemsItemRetentionlabelRetentionLabelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemRetentionlabelRetentionLabelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemRetentionlabelRetentionLabelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemRetentionlabelRetentionLabelRequestBuilderGetQueryParameters
}
// ItemItemsItemRetentionlabelRetentionLabelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemRetentionlabelRetentionLabelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemRetentionlabelRetentionLabelRequestBuilderInternal instantiates a new ItemItemsItemRetentionlabelRetentionLabelRequestBuilder and sets the default values.
func NewItemItemsItemRetentionlabelRetentionLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) {
    m := &ItemItemsItemRetentionlabelRetentionLabelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/retentionLabel{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemRetentionlabelRetentionLabelRequestBuilder instantiates a new ItemItemsItemRetentionlabelRetentionLabelRequestBuilder and sets the default values.
func NewItemItemsItemRetentionlabelRetentionLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemRetentionlabelRetentionLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a retention label from a driveItem. For information about retention labels from an administrator's perspective, see Use retention labels to manage the lifecycle of documents stored in SharePoint.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-removeretentionlabel?view=graph-rest-1.0
func (m *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemRetentionlabelRetentionLabelRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get information about retention label and settings enforced on the driveItem. Read-write.
// returns a ItemRetentionLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemRetentionlabelRetentionLabelRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemRetentionLabelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemRetentionLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemRetentionLabelable), nil
}
// Patch lock or unlock a retention label on a driveItem that classifies content as records. For information about retention labels from an administrator's perspective, see Use retention labels to manage the lifecycle of documents stored in SharePoint. For more information about how you can lock and unlock retention labels, see Use record versioning to update records stored in SharePoint or OneDrive.
// returns a ItemRetentionLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-lockorunlockrecord?view=graph-rest-1.0
func (m *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemRetentionLabelable, requestConfiguration *ItemItemsItemRetentionlabelRetentionLabelRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemRetentionLabelable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemRetentionLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemRetentionLabelable), nil
}
// ToDeleteRequestInformation remove a retention label from a driveItem. For information about retention labels from an administrator's perspective, see Use retention labels to manage the lifecycle of documents stored in SharePoint.
// returns a *RequestInformation when successful
func (m *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemRetentionlabelRetentionLabelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation information about retention label and settings enforced on the driveItem. Read-write.
// returns a *RequestInformation when successful
func (m *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemRetentionlabelRetentionLabelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation lock or unlock a retention label on a driveItem that classifies content as records. For information about retention labels from an administrator's perspective, see Use retention labels to manage the lifecycle of documents stored in SharePoint. For more information about how you can lock and unlock retention labels, see Use record versioning to update records stored in SharePoint or OneDrive.
// returns a *RequestInformation when successful
func (m *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemRetentionLabelable, requestConfiguration *ItemItemsItemRetentionlabelRetentionLabelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder when successful
func (m *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) {
    return NewItemItemsItemRetentionlabelRetentionLabelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
