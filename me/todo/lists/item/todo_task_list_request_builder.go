package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i481292137b632034979d405e09dc8881fad2d21a4ae3b2ce24d327950ff62851 "github.com/microsoftgraph/msgraph-sdk-go/me/todo/lists/item/tasks"
    id864f4b2f5037b226963100fc0abd2e8d82e7239ae534da41f932afb427e1c81 "github.com/microsoftgraph/msgraph-sdk-go/me/todo/lists/item/extensions"
    ib8333791b315c010726faf186940da1da21dd7fe61a75da6a7df6302b06da1ab "github.com/microsoftgraph/msgraph-sdk-go/me/todo/lists/item/tasks/item"
    idb65638d6b0efd71c34aef4e9a6d575420271266c9ec4e05a57ce031aaf4a359 "github.com/microsoftgraph/msgraph-sdk-go/me/todo/lists/item/extensions/item"
)

// Builds and executes requests for operations under \me\todo\lists\{todoTaskList-id}
type TodoTaskListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type TodoTaskListRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type TodoTaskListRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TodoTaskListRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The task lists in the users mailbox.
type TodoTaskListRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type TodoTaskListRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TodoTaskList;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new TodoTaskListRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTodoTaskListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskListRequestBuilder) {
    m := &TodoTaskListRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/todo/lists/{todoTaskList_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TodoTaskListRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTodoTaskListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoTaskListRequestBuilderInternal(urlParams, requestAdapter)
}
// The task lists in the users mailbox.
// Parameters:
//  - options : Options for the request
func (m *TodoTaskListRequestBuilder) CreateDeleteRequestInformation(options *TodoTaskListRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The task lists in the users mailbox.
// Parameters:
//  - options : Options for the request
func (m *TodoTaskListRequestBuilder) CreateGetRequestInformation(options *TodoTaskListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The task lists in the users mailbox.
// Parameters:
//  - options : Options for the request
func (m *TodoTaskListRequestBuilder) CreatePatchRequestInformation(options *TodoTaskListRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The task lists in the users mailbox.
// Parameters:
//  - options : Options for the request
func (m *TodoTaskListRequestBuilder) Delete(options *TodoTaskListRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TodoTaskListRequestBuilder) Extensions()(*id864f4b2f5037b226963100fc0abd2e8d82e7239ae534da41f932afb427e1c81.ExtensionsRequestBuilder) {
    return id864f4b2f5037b226963100fc0abd2e8d82e7239ae534da41f932afb427e1c81.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.todo.lists.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *TodoTaskListRequestBuilder) ExtensionsById(id string)(*idb65638d6b0efd71c34aef4e9a6d575420271266c9ec4e05a57ce031aaf4a359.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return idb65638d6b0efd71c34aef4e9a6d575420271266c9ec4e05a57ce031aaf4a359.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The task lists in the users mailbox.
// Parameters:
//  - options : Options for the request
func (m *TodoTaskListRequestBuilder) Get(options *TodoTaskListRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TodoTaskList, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTodoTaskList() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TodoTaskList), nil
}
// The task lists in the users mailbox.
// Parameters:
//  - options : Options for the request
func (m *TodoTaskListRequestBuilder) Patch(options *TodoTaskListRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TodoTaskListRequestBuilder) Tasks()(*i481292137b632034979d405e09dc8881fad2d21a4ae3b2ce24d327950ff62851.TasksRequestBuilder) {
    return i481292137b632034979d405e09dc8881fad2d21a4ae3b2ce24d327950ff62851.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.todo.lists.item.tasks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *TodoTaskListRequestBuilder) TasksById(id string)(*ib8333791b315c010726faf186940da1da21dd7fe61a75da6a7df6302b06da1ab.TodoTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["todoTask_id"] = id
    }
    return ib8333791b315c010726faf186940da1da21dd7fe61a75da6a7df6302b06da1ab.NewTodoTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
