package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i41e6a4f5d8af8cbec70d9d8c4f48278ff47465e70b23d4ee2bbdcafa93bae975 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/assignments"
    i52247cde4874402d4eaf161fa2c9ec2cc4a9257f52df96b9bbbcdf47e1f167f9 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/categories"
    ia9cfc52616110c9bb721b0f95f47d861e5afa2655ecff28a0fc66beb7ba78fe2 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/assign"
    i1de44916490eb612ae701336a2eae5f89d836e0eaae6829aa711a57c559453e1 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/assignments/item"
    i2d1f6abd3c9d6201219162619b1ea91782688f08584c687565b81c2be79dadc0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/categories/item"
)

// MobileAppItemRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MobileAppItemRequestBuilderDeleteOptions options for Delete
type MobileAppItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MobileAppItemRequestBuilderGetOptions options for Get
type MobileAppItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MobileAppItemRequestBuilderGetQueryParameters the mobile apps.
type MobileAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppItemRequestBuilderPatchOptions options for Patch
type MobileAppItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Assign the assign property
func (m *MobileAppItemRequestBuilder) Assign()(*ia9cfc52616110c9bb721b0f95f47d861e5afa2655ecff28a0fc66beb7ba78fe2.AssignRequestBuilder) {
    return ia9cfc52616110c9bb721b0f95f47d861e5afa2655ecff28a0fc66beb7ba78fe2.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *MobileAppItemRequestBuilder) Assignments()(*i41e6a4f5d8af8cbec70d9d8c4f48278ff47465e70b23d4ee2bbdcafa93bae975.AssignmentsRequestBuilder) {
    return i41e6a4f5d8af8cbec70d9d8c4f48278ff47465e70b23d4ee2bbdcafa93bae975.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileApps.item.assignments.item collection
func (m *MobileAppItemRequestBuilder) AssignmentsById(id string)(*i1de44916490eb612ae701336a2eae5f89d836e0eaae6829aa711a57c559453e1.MobileAppAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppAssignment%2Did"] = id
    }
    return i1de44916490eb612ae701336a2eae5f89d836e0eaae6829aa711a57c559453e1.NewMobileAppAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories the categories property
func (m *MobileAppItemRequestBuilder) Categories()(*i52247cde4874402d4eaf161fa2c9ec2cc4a9257f52df96b9bbbcdf47e1f167f9.CategoriesRequestBuilder) {
    return i52247cde4874402d4eaf161fa2c9ec2cc4a9257f52df96b9bbbcdf47e1f167f9.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileApps.item.categories.item collection
func (m *MobileAppItemRequestBuilder) CategoriesById(id string)(*i2d1f6abd3c9d6201219162619b1ea91782688f08584c687565b81c2be79dadc0.MobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory%2Did"] = id
    }
    return i2d1f6abd3c9d6201219162619b1ea91782688f08584c687565b81c2be79dadc0.NewMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMobileAppItemRequestBuilderInternal instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppItemRequestBuilder) {
    m := &MobileAppItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppItemRequestBuilder instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mobileApps for deviceAppManagement
func (m *MobileAppItemRequestBuilder) CreateDeleteRequestInformation(options *MobileAppItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// CreateGetRequestInformation the mobile apps.
func (m *MobileAppItemRequestBuilder) CreateGetRequestInformation(options *MobileAppItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mobileApps in deviceAppManagement
func (m *MobileAppItemRequestBuilder) CreatePatchRequestInformation(options *MobileAppItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
// Delete delete navigation property mobileApps for deviceAppManagement
func (m *MobileAppItemRequestBuilder) Delete(options *MobileAppItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the mobile apps.
func (m *MobileAppItemRequestBuilder) Get(options *MobileAppItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable), nil
}
// Patch update the navigation property mobileApps in deviceAppManagement
func (m *MobileAppItemRequestBuilder) Patch(options *MobileAppItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
