package bundles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibd8f407d60551e9a5ea4d8ad2c73daee536f46f41a45c102bdf79e1803708f0c "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/count"
)

// BundlesRequestBuilder provides operations to manage the bundles property of the microsoft.graph.drive entity.
type BundlesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BundlesRequestBuilderGetQueryParameters collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
type BundlesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// BundlesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BundlesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BundlesRequestBuilderGetQueryParameters
}
// NewBundlesRequestBuilderInternal instantiates a new BundlesRequestBuilder and sets the default values.
func NewBundlesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BundlesRequestBuilder) {
    m := &BundlesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/bundles{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBundlesRequestBuilder instantiates a new BundlesRequestBuilder and sets the default values.
func NewBundlesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BundlesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBundlesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *BundlesRequestBuilder) Count()(*ibd8f407d60551e9a5ea4d8ad2c73daee536f46f41a45c102bdf79e1803708f0c.CountRequestBuilder) {
    return ibd8f407d60551e9a5ea4d8ad2c73daee536f46f41a45c102bdf79e1803708f0c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *BundlesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *BundlesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *BundlesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *BundlesRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *BundlesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *BundlesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemCollectionResponseable), nil
}
