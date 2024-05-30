package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder provides operations to call the runHuntingQuery method.
type MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilderInternal instantiates a new MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder and sets the default values.
func NewMicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder) {
    m := &MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/microsoft.graph.security.runHuntingQuery", pathParameters),
    }
    return m
}
// NewMicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder instantiates a new MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder and sets the default values.
func NewMicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilderInternal(urlParams, requestAdapter)
}
// Post queries a specified set of event, activity, or entity data supported by Microsoft 365 Defender to proactively look for specific threats in your environment. This method is for advanced hunting in Microsoft 365 Defender. This method includes a query in Kusto Query Language (KQL). It specifies a data table in the advanced hunting schema and a piped sequence of operators to filter or search that data, and format the query output in specific ways.  Find out more about hunting for threats across devices, emails, apps, and identities. Learn about KQL. For information on using advanced hunting in the Microsoft 365 Defender portal, see Proactively hunt for threats with advanced hunting in Microsoft 365 Defender.
// returns a HuntingQueryResultsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder) Post(ctx context.Context, body MicrosoftgraphsecurityrunhuntingqueryRunHuntingQueryPostRequestBodyable, requestConfiguration *MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HuntingQueryResultsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateHuntingQueryResultsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HuntingQueryResultsable), nil
}
// ToPostRequestInformation queries a specified set of event, activity, or entity data supported by Microsoft 365 Defender to proactively look for specific threats in your environment. This method is for advanced hunting in Microsoft 365 Defender. This method includes a query in Kusto Query Language (KQL). It specifies a data table in the advanced hunting schema and a piped sequence of operators to filter or search that data, and format the query output in specific ways.  Find out more about hunting for threats across devices, emails, apps, and identities. Learn about KQL. For information on using advanced hunting in the Microsoft 365 Defender portal, see Proactively hunt for threats with advanced hunting in Microsoft 365 Defender.
// returns a *RequestInformation when successful
func (m *MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftgraphsecurityrunhuntingqueryRunHuntingQueryPostRequestBodyable, requestConfiguration *MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder when successful
func (m *MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder) WithUrl(rawUrl string)(*MicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder) {
    return NewMicrosoftgraphsecurityrunhuntingqueryMicrosoftGraphSecurityRunHuntingQueryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
