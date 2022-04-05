package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i3bc7881f8d53ff46b80580bf2bbeca93ee226b9e556c80e1a29474dc201f2a62 "github.com/microsoftgraph/msgraph-sdk-go/applications/validateproperties"
    i7e8a726e09ced7d4436bf674d1a81669775cb1463962eb805fe901152e21a7a1 "github.com/microsoftgraph/msgraph-sdk-go/applications/count"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    id0455efde9c006c43aa69fbd428002d35001851a48cc95647c179d7f5fe714d8 "github.com/microsoftgraph/msgraph-sdk-go/applications/getavailableextensionproperties"
    id1a6f2211b8fade308e722e65256ccffe715e872816ec9b9e12b585d86dd4f21 "github.com/microsoftgraph/msgraph-sdk-go/applications/getbyids"
    ie526c09037227f7fe6fca4d5c202671bc46cb0934b016d8249c0780baf039d18 "github.com/microsoftgraph/msgraph-sdk-go/applications/delta"
)

// ApplicationsRequestBuilder provides operations to manage the collection of application entities.
type ApplicationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ApplicationsRequestBuilderGetOptions options for Get
type ApplicationsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ApplicationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ApplicationsRequestBuilderGetQueryParameters get entities from applications
type ApplicationsRequestBuilderGetQueryParameters struct {
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
// ApplicationsRequestBuilderPostOptions options for Post
type ApplicationsRequestBuilderPostOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewApplicationsRequestBuilderInternal instantiates a new ApplicationsRequestBuilder and sets the default values.
func NewApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsRequestBuilder) {
    m := &ApplicationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationsRequestBuilder instantiates a new ApplicationsRequestBuilder and sets the default values.
func NewApplicationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ApplicationsRequestBuilder) Count()(*i7e8a726e09ced7d4436bf674d1a81669775cb1463962eb805fe901152e21a7a1.CountRequestBuilder) {
    return i7e8a726e09ced7d4436bf674d1a81669775cb1463962eb805fe901152e21a7a1.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from applications
func (m *ApplicationsRequestBuilder) CreateGetRequestInformation(options *ApplicationsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to applications
func (m *ApplicationsRequestBuilder) CreatePostRequestInformation(options *ApplicationsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ApplicationsRequestBuilder) Delta()(*ie526c09037227f7fe6fca4d5c202671bc46cb0934b016d8249c0780baf039d18.DeltaRequestBuilder) {
    return ie526c09037227f7fe6fca4d5c202671bc46cb0934b016d8249c0780baf039d18.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from applications
func (m *ApplicationsRequestBuilder) Get(options *ApplicationsRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApplicationCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApplicationCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ApplicationCollectionResponseable), nil
}
// GetAvailableExtensionProperties the getAvailableExtensionProperties property
func (m *ApplicationsRequestBuilder) GetAvailableExtensionProperties()(*id0455efde9c006c43aa69fbd428002d35001851a48cc95647c179d7f5fe714d8.GetAvailableExtensionPropertiesRequestBuilder) {
    return id0455efde9c006c43aa69fbd428002d35001851a48cc95647c179d7f5fe714d8.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds the getByIds property
func (m *ApplicationsRequestBuilder) GetByIds()(*id1a6f2211b8fade308e722e65256ccffe715e872816ec9b9e12b585d86dd4f21.GetByIdsRequestBuilder) {
    return id1a6f2211b8fade308e722e65256ccffe715e872816ec9b9e12b585d86dd4f21.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to applications
func (m *ApplicationsRequestBuilder) Post(options *ApplicationsRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApplicationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable), nil
}
// ValidateProperties the validateProperties property
func (m *ApplicationsRequestBuilder) ValidateProperties()(*i3bc7881f8d53ff46b80580bf2bbeca93ee226b9e556c80e1a29474dc201f2a62.ValidatePropertiesRequestBuilder) {
    return i3bc7881f8d53ff46b80580bf2bbeca93ee226b9e556c80e1a29474dc201f2a62.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
