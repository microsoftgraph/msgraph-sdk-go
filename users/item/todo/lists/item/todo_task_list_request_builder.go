package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    if45dc72f3e02b5fd2949bbcfca9e82c09320d6e4ce0c6aa6f9c3507c3d63c1fa "github.com/microsoftgraph/msgraph-sdk-go/users/item/todo/lists/item/extensions"
    ifa2e26d54424700e30d71b507b8ffed3e154c962f224a028c2f02c3d9e3bef39 "github.com/microsoftgraph/msgraph-sdk-go/users/item/todo/lists/item/tasks"
    i2a8ddcb9d28855cae25ca9abfc54cecc8c5aa2d65e13c350c7a83dee4748b97c "github.com/microsoftgraph/msgraph-sdk-go/users/item/todo/lists/item/extensions/item"
    i9f2e822407f0ce614d4bbd55f26b0c76dd2025c6ab5f67b5d2567ed17bbee91c "github.com/microsoftgraph/msgraph-sdk-go/users/item/todo/lists/item/tasks/item"
)

// todoTaskListRequestBuilder builds and executes requests for operations under \users\{user-id}\todo\lists\{todoTaskList-id}
type TodoTaskListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TodoTaskListRequestBuilderDeleteOptions options for Delete
type TodoTaskListRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TodoTaskListRequestBuilderGetOptions options for Get
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
// todoTaskListRequestBuilderGetQueryParameters the task lists in the users mailbox.
type TodoTaskListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// TodoTaskListRequestBuilderPatchOptions options for Patch
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
// NewTodoTaskListRequestBuilderInternal instantiates a new TodoTaskListRequestBuilder and sets the default values.
func NewTodoTaskListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskListRequestBuilder) {
    m := &TodoTaskListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/todo/lists/{todoTaskList_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTodoTaskListRequestBuilder instantiates a new TodoTaskListRequestBuilder and sets the default values.
func NewTodoTaskListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoTaskListRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the task lists in the users mailbox.
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
// CreateGetRequestInformation the task lists in the users mailbox.
func (m *TodoTaskListRequestBuilder) CreateGetRequestInformation(options *TodoTaskListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the task lists in the users mailbox.
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
// Delete the task lists in the users mailbox.
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
func (m *TodoTaskListRequestBuilder) Extensions()(*if45dc72f3e02b5fd2949bbcfca9e82c09320d6e4ce0c6aa6f9c3507c3d63c1fa.ExtensionsRequestBuilder) {
    return if45dc72f3e02b5fd2949bbcfca9e82c09320d6e4ce0c6aa6f9c3507c3d63c1fa.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.todo.lists.item.extensions.item collection
func (m *TodoTaskListRequestBuilder) ExtensionsById(id string)(*i2a8ddcb9d28855cae25ca9abfc54cecc8c5aa2d65e13c350c7a83dee4748b97c.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i2a8ddcb9d28855cae25ca9abfc54cecc8c5aa2d65e13c350c7a83dee4748b97c.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the task lists in the users mailbox.
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
// Patch the task lists in the users mailbox.
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
func (m *TodoTaskListRequestBuilder) Tasks()(*ifa2e26d54424700e30d71b507b8ffed3e154c962f224a028c2f02c3d9e3bef39.TasksRequestBuilder) {
    return ifa2e26d54424700e30d71b507b8ffed3e154c962f224a028c2f02c3d9e3bef39.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.todo.lists.item.tasks.item collection
func (m *TodoTaskListRequestBuilder) TasksById(id string)(*i9f2e822407f0ce614d4bbd55f26b0c76dd2025c6ab5f67b5d2567ed17bbee91c.TodoTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["todoTask_id"] = id
    }
    return i9f2e822407f0ce614d4bbd55f26b0c76dd2025c6ab5f67b5d2567ed17bbee91c.NewTodoTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
