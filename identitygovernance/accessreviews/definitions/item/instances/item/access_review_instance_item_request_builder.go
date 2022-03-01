package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i294ddceafef43eaa0b25cedb2839f1cbb0c453d7e7c544c63c38bf3fc140fa18 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/resetdecisions"
    i3011b416e1fbf818fb42fe3d7de3bb3ea3136fb606994268b64ee1859c2b3229 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stop"
    i354cbd130706e9e98deafc9f1d5a9b5cef510bfe73b8928151750e8dc965ff4b "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/sendreminder"
    i5451506c6a51cc37c7bd81ab078505ae9728ac07623e29e52b6897a016d027b2 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/acceptrecommendations"
    i5e1ccb1ad78906f2443f054f53d5d0547e9eceac11b87f43d7d1e7073c4dbfcd "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/batchrecorddecisions"
    iaba5c5f8866d817198badccdab50ac168445b09d85f789a8f760044ae410a919 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions"
    ibef71d2452748b7a1d404bef412390f9d6f26c7062269a40f953a67d2d0bd44e "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/contactedreviewers"
    icc113f5eb4a8c739b83f1517b406812dfa71ea84b8eb7fb9a3ca0629d5b5320a "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/applydecisions"
    ia5228e75938f6cab481e6dee19c7eeb8666192ef9a8d846054e182dcf31deb9e "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item"
    idbe515597f9c38cd733993c3ed77612e84f59bbc489b0a78f8c8a2160b9ac50d "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/contactedreviewers/item"
)

// AccessReviewInstanceItemRequestBuilder builds and executes requests for operations under \identityGovernance\accessReviews\definitions\{accessReviewScheduleDefinition-id}\instances\{accessReviewInstance-id}
type AccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewInstanceItemRequestBuilderDeleteOptions options for Delete
type AccessReviewInstanceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceItemRequestBuilderGetOptions options for Get
type AccessReviewInstanceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewInstanceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceItemRequestBuilderGetQueryParameters if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
type AccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewInstanceItemRequestBuilderPatchOptions options for Patch
type AccessReviewInstanceItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewInstance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*i5451506c6a51cc37c7bd81ab078505ae9728ac07623e29e52b6897a016d027b2.AcceptRecommendationsRequestBuilder) {
    return i5451506c6a51cc37c7bd81ab078505ae9728ac07623e29e52b6897a016d027b2.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*icc113f5eb4a8c739b83f1517b406812dfa71ea84b8eb7fb9a3ca0629d5b5320a.ApplyDecisionsRequestBuilder) {
    return icc113f5eb4a8c739b83f1517b406812dfa71ea84b8eb7fb9a3ca0629d5b5320a.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*i5e1ccb1ad78906f2443f054f53d5d0547e9eceac11b87f43d7d1e7073c4dbfcd.BatchRecordDecisionsRequestBuilder) {
    return i5e1ccb1ad78906f2443f054f53d5d0547e9eceac11b87f43d7d1e7073c4dbfcd.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceItemRequestBuilder) {
    m := &AccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition_id}/instances/{accessReviewInstance_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewInstanceItemRequestBuilder instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*ibef71d2452748b7a1d404bef412390f9d6f26c7062269a40f953a67d2d0bd44e.ContactedReviewersRequestBuilder) {
    return ibef71d2452748b7a1d404bef412390f9d6f26c7062269a40f953a67d2d0bd44e.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.contactedReviewers.item collection
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*idbe515597f9c38cd733993c3ed77612e84f59bbc489b0a78f8c8a2160b9ac50d.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer_id"] = id
    }
    return idbe515597f9c38cd733993c3ed77612e84f59bbc489b0a78f8c8a2160b9ac50d.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewInstanceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) CreateGetRequestInformation(options *AccessReviewInstanceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformation(options *AccessReviewInstanceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AccessReviewInstanceItemRequestBuilder) Decisions()(*iaba5c5f8866d817198badccdab50ac168445b09d85f789a8f760044ae410a919.DecisionsRequestBuilder) {
    return iaba5c5f8866d817198badccdab50ac168445b09d85f789a8f760044ae410a919.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.decisions.item collection
func (m *AccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*ia5228e75938f6cab481e6dee19c7eeb8666192ef9a8d846054e182dcf31deb9e.AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem_id"] = id
    }
    return ia5228e75938f6cab481e6dee19c7eeb8666192ef9a8d846054e182dcf31deb9e.NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) Delete(options *AccessReviewInstanceItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) Get(options *AccessReviewInstanceItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewInstance, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAccessReviewInstance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewInstance), nil
}
// Patch if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) Patch(options *AccessReviewInstanceItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AccessReviewInstanceItemRequestBuilder) ResetDecisions()(*i294ddceafef43eaa0b25cedb2839f1cbb0c453d7e7c544c63c38bf3fc140fa18.ResetDecisionsRequestBuilder) {
    return i294ddceafef43eaa0b25cedb2839f1cbb0c453d7e7c544c63c38bf3fc140fa18.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceItemRequestBuilder) SendReminder()(*i354cbd130706e9e98deafc9f1d5a9b5cef510bfe73b8928151750e8dc965ff4b.SendReminderRequestBuilder) {
    return i354cbd130706e9e98deafc9f1d5a9b5cef510bfe73b8928151750e8dc965ff4b.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceItemRequestBuilder) Stop()(*i3011b416e1fbf818fb42fe3d7de3bb3ea3136fb606994268b64ee1859c2b3229.StopRequestBuilder) {
    return i3011b416e1fbf818fb42fe3d7de3bb3ea3136fb606994268b64ee1859c2b3229.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
