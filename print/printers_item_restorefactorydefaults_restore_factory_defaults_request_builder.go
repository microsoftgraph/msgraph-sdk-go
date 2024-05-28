package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder provides operations to call the restoreFactoryDefaults method.
type PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilderInternal instantiates a new PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder and sets the default values.
func NewPrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder) {
    m := &PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/restoreFactoryDefaults", pathParameters),
    }
    return m
}
// NewPrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder instantiates a new PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder and sets the default values.
func NewPrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a printer's default settings to the values specified by the manufacturer.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/printer-restorefactorydefaults?view=graph-rest-1.0
func (m *PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder) Post(ctx context.Context, requestConfiguration *PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation restore a printer's default settings to the values specified by the manufacturer.
// returns a *RequestInformation when successful
func (m *PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder when successful
func (m *PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder) WithUrl(rawUrl string)(*PrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder) {
    return NewPrintersItemRestorefactorydefaultsRestoreFactoryDefaultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
