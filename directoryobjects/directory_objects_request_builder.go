package directoryobjects

import (
    i4e32c69eb0781ba00301b7881eaffc86307a431dad4cb4b131d227f662210091 "github.com/microsoftgraph/msgraph-sdk-go/directoryobjects/getbyids"
    i9727e1e2847598e4e14a228da4393d7740547c451b9e58df0e82dd81a5ff82bf "github.com/microsoftgraph/msgraph-sdk-go/directoryobjects/getavailableextensionproperties"
    ia634f5aedaa59620c7034e20ae73e42556e3a50cdb23848c1d283d6c828de7e8 "github.com/microsoftgraph/msgraph-sdk-go/directoryobjects/count"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if6c1d96eb675fbef9136f79b5bca6096fe5466cfa751ae9f3724d3fd3e9e9d11 "github.com/microsoftgraph/msgraph-sdk-go/directoryobjects/validateproperties"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// DirectoryObjectsRequestBuilder provides operations to manage the collection of directoryObject entities.
type DirectoryObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryObjectsRequestBuilderGetOptions options for Get
type DirectoryObjectsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryObjectsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryObjectsRequestBuilderGetQueryParameters get entities from directoryObjects
type DirectoryObjectsRequestBuilderGetQueryParameters struct {
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
// DirectoryObjectsRequestBuilderPostOptions options for Post
type DirectoryObjectsRequestBuilderPostOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryObjectable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDirectoryObjectsRequestBuilderInternal instantiates a new DirectoryObjectsRequestBuilder and sets the default values.
func NewDirectoryObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryObjectsRequestBuilder) {
    m := &DirectoryObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryObjects{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectsRequestBuilder instantiates a new DirectoryObjectsRequestBuilder and sets the default values.
func NewDirectoryObjectsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectoryObjectsRequestBuilder) Count()(*ia634f5aedaa59620c7034e20ae73e42556e3a50cdb23848c1d283d6c828de7e8.CountRequestBuilder) {
    return ia634f5aedaa59620c7034e20ae73e42556e3a50cdb23848c1d283d6c828de7e8.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from directoryObjects
func (m *DirectoryObjectsRequestBuilder) CreateGetRequestInformation(options *DirectoryObjectsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation add new entity to directoryObjects
func (m *DirectoryObjectsRequestBuilder) CreatePostRequestInformation(options *DirectoryObjectsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get get entities from directoryObjects
func (m *DirectoryObjectsRequestBuilder) Get(options *DirectoryObjectsRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryObjectCollectionResponseable), nil
}
func (m *DirectoryObjectsRequestBuilder) GetAvailableExtensionProperties()(*i9727e1e2847598e4e14a228da4393d7740547c451b9e58df0e82dd81a5ff82bf.GetAvailableExtensionPropertiesRequestBuilder) {
    return i9727e1e2847598e4e14a228da4393d7740547c451b9e58df0e82dd81a5ff82bf.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryObjectsRequestBuilder) GetByIds()(*i4e32c69eb0781ba00301b7881eaffc86307a431dad4cb4b131d227f662210091.GetByIdsRequestBuilder) {
    return i4e32c69eb0781ba00301b7881eaffc86307a431dad4cb4b131d227f662210091.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directoryObjects
func (m *DirectoryObjectsRequestBuilder) Post(options *DirectoryObjectsRequestBuilderPostOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryObjectable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDirectoryObjectFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryObjectable), nil
}
func (m *DirectoryObjectsRequestBuilder) ValidateProperties()(*if6c1d96eb675fbef9136f79b5bca6096fe5466cfa751ae9f3724d3fd3e9e9d11.ValidatePropertiesRequestBuilder) {
    return if6c1d96eb675fbef9136f79b5bca6096fe5466cfa751ae9f3724d3fd3e9e9d11.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
