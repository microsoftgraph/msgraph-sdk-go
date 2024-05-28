package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrintersItemTasktriggersItemDefinitionRequestBuilder provides operations to manage the definition property of the microsoft.graph.printTaskTrigger entity.
type PrintersItemTasktriggersItemDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemTasktriggersItemDefinitionRequestBuilderGetQueryParameters an abstract definition that is used to create a printTask when triggered by a print event. Read-only.
type PrintersItemTasktriggersItemDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintersItemTasktriggersItemDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemTasktriggersItemDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintersItemTasktriggersItemDefinitionRequestBuilderGetQueryParameters
}
// NewPrintersItemTasktriggersItemDefinitionRequestBuilderInternal instantiates a new PrintersItemTasktriggersItemDefinitionRequestBuilder and sets the default values.
func NewPrintersItemTasktriggersItemDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemTasktriggersItemDefinitionRequestBuilder) {
    m := &PrintersItemTasktriggersItemDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/taskTriggers/{printTaskTrigger%2Did}/definition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrintersItemTasktriggersItemDefinitionRequestBuilder instantiates a new PrintersItemTasktriggersItemDefinitionRequestBuilder and sets the default values.
func NewPrintersItemTasktriggersItemDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemTasktriggersItemDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemTasktriggersItemDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get an abstract definition that is used to create a printTask when triggered by a print event. Read-only.
// returns a PrintTaskDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersItemTasktriggersItemDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersItemTasktriggersItemDefinitionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintTaskDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintTaskDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintTaskDefinitionable), nil
}
// ToGetRequestInformation an abstract definition that is used to create a printTask when triggered by a print event. Read-only.
// returns a *RequestInformation when successful
func (m *PrintersItemTasktriggersItemDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersItemTasktriggersItemDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrintersItemTasktriggersItemDefinitionRequestBuilder when successful
func (m *PrintersItemTasktriggersItemDefinitionRequestBuilder) WithUrl(rawUrl string)(*PrintersItemTasktriggersItemDefinitionRequestBuilder) {
    return NewPrintersItemTasktriggersItemDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
