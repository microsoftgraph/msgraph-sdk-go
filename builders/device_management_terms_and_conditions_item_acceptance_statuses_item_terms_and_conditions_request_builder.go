package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder provides operations to manage the termsAndConditions property of the microsoft.graph.termsAndConditionsAcceptanceStatus entity.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderGetQueryParameters navigation link to the terms and conditions that are assigned.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderGetQueryParameters
}
// NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderInternal instantiates a new TermsAndConditionsRequestBuilder and sets the default values.
func NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder) {
    m := &DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus%2Did}/termsAndConditions{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder instantiates a new TermsAndConditionsRequestBuilder and sets the default values.
func NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation navigation link to the terms and conditions that are assigned.
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get navigation link to the terms and conditions that are assigned.
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTermsAndConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsable), nil
}
