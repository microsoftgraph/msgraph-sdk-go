package builders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder builds and executes requests for operations under \education\classes\{educationClass-id}\assignments\{educationAssignment-id}\categories\{educationCategory-id}
type EducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewEducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    m := &EducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/categories/{educationCategory%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewEducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *EducationClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) Ref()(*EducationClassesItemAssignmentsItemCategoriesItemRefRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemCategoriesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
