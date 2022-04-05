package getmanagedappdiagnosticstatuses

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetManagedAppDiagnosticStatusesRequestBuilder provides operations to call the getManagedAppDiagnosticStatuses method.
type GetManagedAppDiagnosticStatusesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetManagedAppDiagnosticStatusesRequestBuilderGetOptions options for Get
type GetManagedAppDiagnosticStatusesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGetManagedAppDiagnosticStatusesRequestBuilderInternal instantiates a new GetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedAppDiagnosticStatusesRequestBuilder) {
    m := &GetManagedAppDiagnosticStatusesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.getManagedAppDiagnosticStatuses()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagedAppDiagnosticStatusesRequestBuilder instantiates a new GetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewGetManagedAppDiagnosticStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedAppDiagnosticStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation gets diagnostics validation status for a given user.
func (m *GetManagedAppDiagnosticStatusesRequestBuilder) CreateGetRequestInformation(options *GetManagedAppDiagnosticStatusesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get gets diagnostics validation status for a given user.
func (m *GetManagedAppDiagnosticStatusesRequestBuilder) Get(options *GetManagedAppDiagnosticStatusesRequestBuilderGetOptions)(GetManagedAppDiagnosticStatusesResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagedAppDiagnosticStatusesResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagedAppDiagnosticStatusesResponseable), nil
}
