package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ia70f653993cbe8ba50d07297124b6fccc016858daaec3d574aff57f6dd413715 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/users"
    iaa1b72b2e917a242ccb04624a7df0510273ede01d0eb20e634f0ce118cbe44d1 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/classes"
    id764fdef7d1fc1e96b7dd96ba4ca2de348edde4d6b8c59eb7bc5df6cfdbe1fa3 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/administrativeunit"
    i410e950d0e7e1732c6613b5e229d432d16a6b98c7937bf13ea765450eb216fc9 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/users/item"
    i86542833251cc9d20683beeed83578dc1c733b42152adb9e290928af411e99e5 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/classes/item"
)

// EducationSchoolItemRequestBuilder provides operations to manage the schools property of the microsoft.graph.educationRoot entity.
type EducationSchoolItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSchoolItemRequestBuilderDeleteOptions options for Delete
type EducationSchoolItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EducationSchoolItemRequestBuilderGetOptions options for Get
type EducationSchoolItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EducationSchoolItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EducationSchoolItemRequestBuilderGetQueryParameters get schools from education
type EducationSchoolItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSchoolItemRequestBuilderPatchOptions options for Patch
type EducationSchoolItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSchoolable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AdministrativeUnit the administrativeUnit property
func (m *EducationSchoolItemRequestBuilder) AdministrativeUnit()(*id764fdef7d1fc1e96b7dd96ba4ca2de348edde4d6b8c59eb7bc5df6cfdbe1fa3.AdministrativeUnitRequestBuilder) {
    return id764fdef7d1fc1e96b7dd96ba4ca2de348edde4d6b8c59eb7bc5df6cfdbe1fa3.NewAdministrativeUnitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Classes the classes property
func (m *EducationSchoolItemRequestBuilder) Classes()(*iaa1b72b2e917a242ccb04624a7df0510273ede01d0eb20e634f0ce118cbe44d1.ClassesRequestBuilder) {
    return iaa1b72b2e917a242ccb04624a7df0510273ede01d0eb20e634f0ce118cbe44d1.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.schools.item.classes.item collection
func (m *EducationSchoolItemRequestBuilder) ClassesById(id string)(*i86542833251cc9d20683beeed83578dc1c733b42152adb9e290928af411e99e5.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return i86542833251cc9d20683beeed83578dc1c733b42152adb9e290928af411e99e5.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationSchoolItemRequestBuilderInternal instantiates a new EducationSchoolItemRequestBuilder and sets the default values.
func NewEducationSchoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSchoolItemRequestBuilder) {
    m := &EducationSchoolItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSchoolItemRequestBuilder instantiates a new EducationSchoolItemRequestBuilder and sets the default values.
func NewEducationSchoolItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSchoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSchoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schools for education
func (m *EducationSchoolItemRequestBuilder) CreateDeleteRequestInformation(options *EducationSchoolItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get schools from education
func (m *EducationSchoolItemRequestBuilder) CreateGetRequestInformation(options *EducationSchoolItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schools in education
func (m *EducationSchoolItemRequestBuilder) CreatePatchRequestInformation(options *EducationSchoolItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property schools for education
func (m *EducationSchoolItemRequestBuilder) Delete(options *EducationSchoolItemRequestBuilderDeleteOptions)(error) {
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
// Get get schools from education
func (m *EducationSchoolItemRequestBuilder) Get(options *EducationSchoolItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSchoolable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSchoolFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSchoolable), nil
}
// Patch update the navigation property schools in education
func (m *EducationSchoolItemRequestBuilder) Patch(options *EducationSchoolItemRequestBuilderPatchOptions)(error) {
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
// Users the users property
func (m *EducationSchoolItemRequestBuilder) Users()(*ia70f653993cbe8ba50d07297124b6fccc016858daaec3d574aff57f6dd413715.UsersRequestBuilder) {
    return ia70f653993cbe8ba50d07297124b6fccc016858daaec3d574aff57f6dd413715.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.schools.item.users.item collection
func (m *EducationSchoolItemRequestBuilder) UsersById(id string)(*i410e950d0e7e1732c6613b5e229d432d16a6b98c7937bf13ea765450eb216fc9.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser_id"] = id
    }
    return i410e950d0e7e1732c6613b5e229d432d16a6b98c7937bf13ea765450eb216fc9.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
