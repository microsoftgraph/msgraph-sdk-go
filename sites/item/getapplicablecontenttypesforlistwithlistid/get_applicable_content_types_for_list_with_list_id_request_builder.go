package getapplicablecontenttypesforlistwithlistid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetApplicableContentTypesForListWithListIdRequestBuilder provides operations to call the getApplicableContentTypesForList method.
type GetApplicableContentTypesForListWithListIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetApplicableContentTypesForListWithListIdRequestBuilderGetOptions options for Get
type GetApplicableContentTypesForListWithListIdRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetApplicableContentTypesForListWithListIdRequestBuilderInternal instantiates a new GetApplicableContentTypesForListWithListIdRequestBuilder and sets the default values.
func NewGetApplicableContentTypesForListWithListIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, listId *string)(*GetApplicableContentTypesForListWithListIdRequestBuilder) {
    m := &GetApplicableContentTypesForListWithListIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/microsoft.graph.getApplicableContentTypesForList(listId='{listId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if listId != nil {
        urlTplParams[""] = *listId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetApplicableContentTypesForListWithListIdRequestBuilder instantiates a new GetApplicableContentTypesForListWithListIdRequestBuilder and sets the default values.
func NewGetApplicableContentTypesForListWithListIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetApplicableContentTypesForListWithListIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetApplicableContentTypesForListWithListIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getApplicableContentTypesForList
func (m *GetApplicableContentTypesForListWithListIdRequestBuilder) CreateGetRequestInformation(options *GetApplicableContentTypesForListWithListIdRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getApplicableContentTypesForList
func (m *GetApplicableContentTypesForListWithListIdRequestBuilder) Get(options *GetApplicableContentTypesForListWithListIdRequestBuilderGetOptions)(GetApplicableContentTypesForListWithListIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetApplicableContentTypesForListWithListIdResponseable), nil
}
