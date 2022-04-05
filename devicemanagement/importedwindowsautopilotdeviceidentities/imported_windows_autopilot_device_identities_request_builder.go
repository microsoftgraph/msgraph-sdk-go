package importedwindowsautopilotdeviceidentities

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i5ec97fd59e061df65913355a6db15803dc4a5a301962b4c741f6a7499d24792c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/importedwindowsautopilotdeviceidentities/import_escaped"
    iaa29daa3c226776b057f1acdf0964b2ce382e6c77b1be10f0c9b676220be7b8e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/importedwindowsautopilotdeviceidentities/count"
)

// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions options for Get
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters collection of imported Windows autopilot devices.
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions options for Post
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal instantiates a new ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    m := &ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilder instantiates a new ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Count()(*iaa29daa3c226776b057f1acdf0964b2ce382e6c77b1be10f0c9b676220be7b8e.CountRequestBuilder) {
    return iaa29daa3c226776b057f1acdf0964b2ce382e6c77b1be10f0c9b676220be7b8e.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of imported Windows autopilot devices.
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) CreateGetRequestInformation(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) CreatePostRequestInformation(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get collection of imported Windows autopilot devices.
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Get(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateImportedWindowsAutopilotDeviceIdentityCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable), nil
}
// Import_escaped the import property
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Import_escaped()(*i5ec97fd59e061df65913355a6db15803dc4a5a301962b4c741f6a7499d24792c.ImportRequestBuilder) {
    return i5ec97fd59e061df65913355a6db15803dc4a5a301962b4c741f6a7499d24792c.NewImportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Post(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable), nil
}
