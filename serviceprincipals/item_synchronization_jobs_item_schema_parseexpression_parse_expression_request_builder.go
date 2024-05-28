package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder provides operations to call the parseExpression method.
type ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilderInternal instantiates a new ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder) {
    m := &ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/jobs/{synchronizationJob%2Did}/schema/parseExpression", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder instantiates a new ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post parse a given string expression into an attributeMappingSource object. For more information about expressions, see Writing Expressions for Attribute Mappings in Microsoft Entra ID.
// returns a ParseExpressionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationschema-parseexpression?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseExpressionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateParseExpressionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseExpressionResponseable), nil
}
// ToPostRequestInformation parse a given string expression into an attributeMappingSource object. For more information about expressions, see Writing Expressions for Attribute Mappings in Microsoft Entra ID.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder when successful
func (m *ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder) {
    return NewItemSynchronizationJobsItemSchemaParseexpressionParseExpressionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
