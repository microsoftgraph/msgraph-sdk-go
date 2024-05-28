package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder provides operations to call the uploadClientCertificate method.
type B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilderInternal instantiates a new B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder) {
    m := &B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/apiConnectorConfiguration/postAttributeCollection/uploadClientCertificate", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder instantiates a new B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a PKCS 12 format key (.pfx) to an API connector's authentication configuration. The input is a base-64 encoded value of the PKCS 12 certificate contents. This method returns an apiConnector.
// returns a IdentityApiConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityapiconnector-uploadclientcertificate?view=graph-rest-1.0
func (m *B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder) Post(ctx context.Context, body B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificatePostRequestBodyable, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityApiConnectorable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityApiConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityApiConnectorable), nil
}
// ToPostRequestInformation upload a PKCS 12 format key (.pfx) to an API connector's authentication configuration. The input is a base-64 encoded value of the PKCS 12 certificate contents. This method returns an apiConnector.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificatePostRequestBodyable, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPostattributecollectionUploadclientcertificateUploadClientCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
