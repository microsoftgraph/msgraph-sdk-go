package education

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i11d420057e9890f34c6f87265f6025f3489c0f78de3b99ece6ba73f225d5a75c "github.com/microsoftgraph/msgraph-sdk-go/education/me"
    i82d352ad9c3de7536dcc103a1a2a276e7beba1b35548191dda9e846d8c21e897 "github.com/microsoftgraph/msgraph-sdk-go/education/users"
    i8803b2d448091478d216c51910e6ed01bbc8c76e5b76968b633118267314ebe9 "github.com/microsoftgraph/msgraph-sdk-go/education/schools"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    if6b46a923414b614133630b614b404c3005a11daa1e05f10bb834755518d40dd "github.com/microsoftgraph/msgraph-sdk-go/education/classes"
    i31a8a42bc05cb93b0a8b6dc7e28fb5a481e3934df6dc18de8217d0fe268d617c "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item"
    i6fc2718a806c86250bd9f9aaba5418413286c98ed345a7dbe988c558836bde8e "github.com/microsoftgraph/msgraph-sdk-go/education/users/item"
    ibcc84b6252b9c075d39d4d8641ef3cf01a04613a5da305d7d48d21c122395065 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item"
)

// EducationRequestBuilder provides operations to manage the educationRoot singleton.
type EducationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationRequestBuilderGetOptions options for Get
type EducationRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EducationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EducationRequestBuilderGetQueryParameters get education
type EducationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationRequestBuilderPatchOptions options for Patch
type EducationRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationRootable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Classes the classes property
func (m *EducationRequestBuilder) Classes()(*if6b46a923414b614133630b614b404c3005a11daa1e05f10bb834755518d40dd.ClassesRequestBuilder) {
    return if6b46a923414b614133630b614b404c3005a11daa1e05f10bb834755518d40dd.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item collection
func (m *EducationRequestBuilder) ClassesById(id string)(*i31a8a42bc05cb93b0a8b6dc7e28fb5a481e3934df6dc18de8217d0fe268d617c.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return i31a8a42bc05cb93b0a8b6dc7e28fb5a481e3934df6dc18de8217d0fe268d617c.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationRequestBuilderInternal instantiates a new EducationRequestBuilder and sets the default values.
func NewEducationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationRequestBuilder) {
    m := &EducationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationRequestBuilder instantiates a new EducationRequestBuilder and sets the default values.
func NewEducationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get education
func (m *EducationRequestBuilder) CreateGetRequestInformation(options *EducationRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update education
func (m *EducationRequestBuilder) CreatePatchRequestInformation(options *EducationRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get education
func (m *EducationRequestBuilder) Get(options *EducationRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationRootFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationRootable), nil
}
// Me the me property
func (m *EducationRequestBuilder) Me()(*i11d420057e9890f34c6f87265f6025f3489c0f78de3b99ece6ba73f225d5a75c.MeRequestBuilder) {
    return i11d420057e9890f34c6f87265f6025f3489c0f78de3b99ece6ba73f225d5a75c.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update education
func (m *EducationRequestBuilder) Patch(options *EducationRequestBuilderPatchOptions)(error) {
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
// Schools the schools property
func (m *EducationRequestBuilder) Schools()(*i8803b2d448091478d216c51910e6ed01bbc8c76e5b76968b633118267314ebe9.SchoolsRequestBuilder) {
    return i8803b2d448091478d216c51910e6ed01bbc8c76e5b76968b633118267314ebe9.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.schools.item collection
func (m *EducationRequestBuilder) SchoolsById(id string)(*ibcc84b6252b9c075d39d4d8641ef3cf01a04613a5da305d7d48d21c122395065.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool_id"] = id
    }
    return ibcc84b6252b9c075d39d4d8641ef3cf01a04613a5da305d7d48d21c122395065.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Users the users property
func (m *EducationRequestBuilder) Users()(*i82d352ad9c3de7536dcc103a1a2a276e7beba1b35548191dda9e846d8c21e897.UsersRequestBuilder) {
    return i82d352ad9c3de7536dcc103a1a2a276e7beba1b35548191dda9e846d8c21e897.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item collection
func (m *EducationRequestBuilder) UsersById(id string)(*i6fc2718a806c86250bd9f9aaba5418413286c98ed345a7dbe988c558836bde8e.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser_id"] = id
    }
    return i6fc2718a806c86250bd9f9aaba5418413286c98ed345a7dbe988c558836bde8e.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
