package directoryroles

import (
    i2a93ad836768ad49d953d9dfcefa9cac543730d20058b421e20bc27251dc8156 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/count"
    ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getavailableextensionproperties"
    ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getbyids"
    id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/delta"
    ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// DirectoryRolesRequestBuilder provides operations to manage the collection of directoryRole entities.
type DirectoryRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRolesRequestBuilderGetOptions options for Get
type DirectoryRolesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRolesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRolesRequestBuilderGetQueryParameters get entities from directoryRoles
type DirectoryRolesRequestBuilderGetQueryParameters struct {
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
}
// DirectoryRolesRequestBuilderPostOptions options for Post
type DirectoryRolesRequestBuilderPostOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDirectoryRolesRequestBuilderInternal instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    m := &DirectoryRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles{?skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRolesRequestBuilder instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectoryRolesRequestBuilder) Count()(*i2a93ad836768ad49d953d9dfcefa9cac543730d20058b421e20bc27251dc8156.CountRequestBuilder) {
    return i2a93ad836768ad49d953d9dfcefa9cac543730d20058b421e20bc27251dc8156.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformation(options *DirectoryRolesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformation(options *DirectoryRolesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *DirectoryRolesRequestBuilder) Delta()(*id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.DeltaRequestBuilder) {
    return id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) Get(options *DirectoryRolesRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleCollectionResponseable), nil
}
func (m *DirectoryRolesRequestBuilder) GetAvailableExtensionProperties()(*ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.GetAvailableExtensionPropertiesRequestBuilder) {
    return ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRolesRequestBuilder) GetByIds()(*ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.GetByIdsRequestBuilder) {
    return ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) Post(options *DirectoryRolesRequestBuilderPostOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDirectoryRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleable), nil
}
func (m *DirectoryRolesRequestBuilder) ValidateProperties()(*ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.ValidatePropertiesRequestBuilder) {
    return ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
