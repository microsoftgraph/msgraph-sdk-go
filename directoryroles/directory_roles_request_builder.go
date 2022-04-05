package directoryroles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i2a93ad836768ad49d953d9dfcefa9cac543730d20058b421e20bc27251dc8156 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/count"
    ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getavailableextensionproperties"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getbyids"
    id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/delta"
    ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/validateproperties"
)

// DirectoryRolesRequestBuilder provides operations to manage the collection of directoryRole entities.
type DirectoryRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRolesRequestBuilderGetOptions options for Get
type DirectoryRolesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DirectoryRolesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDirectoryRolesRequestBuilderInternal instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesRequestBuilder) {
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
func NewDirectoryRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DirectoryRolesRequestBuilder) Count()(*i2a93ad836768ad49d953d9dfcefa9cac543730d20058b421e20bc27251dc8156.CountRequestBuilder) {
    return i2a93ad836768ad49d953d9dfcefa9cac543730d20058b421e20bc27251dc8156.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformation(options *DirectoryRolesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformation(options *DirectoryRolesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *DirectoryRolesRequestBuilder) Delta()(*id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.DeltaRequestBuilder) {
    return id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) Get(options *DirectoryRolesRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleCollectionResponseable), nil
}
// GetAvailableExtensionProperties the getAvailableExtensionProperties property
func (m *DirectoryRolesRequestBuilder) GetAvailableExtensionProperties()(*ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.GetAvailableExtensionPropertiesRequestBuilder) {
    return ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds the getByIds property
func (m *DirectoryRolesRequestBuilder) GetByIds()(*ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.GetByIdsRequestBuilder) {
    return ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) Post(options *DirectoryRolesRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleable), nil
}
// ValidateProperties the validateProperties property
func (m *DirectoryRolesRequestBuilder) ValidateProperties()(*ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.ValidatePropertiesRequestBuilder) {
    return ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
