package accessreviews

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3ed202ff0cbdf2b2edbf6f858dd5dd49845b59ab278423c07797fe43a9ccd40f "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/historydefinitions"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i9a0fc83be8a156b608969476f16987d6f0c56857cc62f3dec40c0a7819428e17 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions"
    i0fca5fbe98a7c54d24af60b510c16cbc28027514fdcbe3896a60732f5c572e2b "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/historydefinitions/item"
    idcb18c7708e010766ed0434cebbaf8978d559fd3ae0bed76f0413cdd312a9c3f "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item"
)

// AccessReviewsRequestBuilder provides operations to manage the accessReviews property of the microsoft.graph.identityGovernance entity.
type AccessReviewsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewsRequestBuilderDeleteOptions options for Delete
type AccessReviewsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewsRequestBuilderGetOptions options for Get
type AccessReviewsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewsRequestBuilderGetQueryParameters get accessReviews from identityGovernance
type AccessReviewsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewsRequestBuilderPatchOptions options for Patch
type AccessReviewsRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewSetable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAccessReviewsRequestBuilderInternal instantiates a new AccessReviewsRequestBuilder and sets the default values.
func NewAccessReviewsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewsRequestBuilder) {
    m := &AccessReviewsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewsRequestBuilder instantiates a new AccessReviewsRequestBuilder and sets the default values.
func NewAccessReviewsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessReviews for identityGovernance
func (m *AccessReviewsRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation get accessReviews from identityGovernance
func (m *AccessReviewsRequestBuilder) CreateGetRequestInformation(options *AccessReviewsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessReviews in identityGovernance
func (m *AccessReviewsRequestBuilder) CreatePatchRequestInformation(options *AccessReviewsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *AccessReviewsRequestBuilder) Definitions()(*i9a0fc83be8a156b608969476f16987d6f0c56857cc62f3dec40c0a7819428e17.DefinitionsRequestBuilder) {
    return i9a0fc83be8a156b608969476f16987d6f0c56857cc62f3dec40c0a7819428e17.NewDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefinitionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.accessReviews.definitions.item collection
func (m *AccessReviewsRequestBuilder) DefinitionsById(id string)(*idcb18c7708e010766ed0434cebbaf8978d559fd3ae0bed76f0413cdd312a9c3f.AccessReviewScheduleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewScheduleDefinition_id"] = id
    }
    return idcb18c7708e010766ed0434cebbaf8978d559fd3ae0bed76f0413cdd312a9c3f.NewAccessReviewScheduleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property accessReviews for identityGovernance
func (m *AccessReviewsRequestBuilder) Delete(options *AccessReviewsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get accessReviews from identityGovernance
func (m *AccessReviewsRequestBuilder) Get(options *AccessReviewsRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewSetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateAccessReviewSetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewSetable), nil
}
func (m *AccessReviewsRequestBuilder) HistoryDefinitions()(*i3ed202ff0cbdf2b2edbf6f858dd5dd49845b59ab278423c07797fe43a9ccd40f.HistoryDefinitionsRequestBuilder) {
    return i3ed202ff0cbdf2b2edbf6f858dd5dd49845b59ab278423c07797fe43a9ccd40f.NewHistoryDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HistoryDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.accessReviews.historyDefinitions.item collection
func (m *AccessReviewsRequestBuilder) HistoryDefinitionsById(id string)(*i0fca5fbe98a7c54d24af60b510c16cbc28027514fdcbe3896a60732f5c572e2b.AccessReviewHistoryDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewHistoryDefinition_id"] = id
    }
    return i0fca5fbe98a7c54d24af60b510c16cbc28027514fdcbe3896a60732f5c572e2b.NewAccessReviewHistoryDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property accessReviews in identityGovernance
func (m *AccessReviewsRequestBuilder) Patch(options *AccessReviewsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
