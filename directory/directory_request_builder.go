package directory

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i47420a6f1d93f934067d08aecf791436625344c23db8a5ebf8ac8c6aa35228f5 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits"
    i608270f0d53df54d261f336ae52901ae79654f5ed54338f92dab3b2188e27aeb "github.com/microsoftgraph/msgraph-sdk-go/directory/deleteditems"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ib6eef29306489af3d4867139be6f8c0d2f07e27cc972bbe619040834fd6bc53c "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item"
    ic11656dccc6289d407b8e6129a339cd67d5b4802e9850fa2b7197192054f419c "github.com/microsoftgraph/msgraph-sdk-go/directory/deleteditems/item"
)

// DirectoryRequestBuilder provides operations to manage the directory singleton.
type DirectoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRequestBuilderGetOptions options for Get
type DirectoryRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DirectoryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DirectoryRequestBuilderGetQueryParameters get directory
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRequestBuilderPatchOptions options for Patch
type DirectoryRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Directoryable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AdministrativeUnits the administrativeUnits property
func (m *DirectoryRequestBuilder) AdministrativeUnits()(*i47420a6f1d93f934067d08aecf791436625344c23db8a5ebf8ac8c6aa35228f5.AdministrativeUnitsRequestBuilder) {
    return i47420a6f1d93f934067d08aecf791436625344c23db8a5ebf8ac8c6aa35228f5.NewAdministrativeUnitsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdministrativeUnitsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.directory.administrativeUnits.item collection
func (m *DirectoryRequestBuilder) AdministrativeUnitsById(id string)(*ib6eef29306489af3d4867139be6f8c0d2f07e27cc972bbe619040834fd6bc53c.AdministrativeUnitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["administrativeUnit_id"] = id
    }
    return ib6eef29306489af3d4867139be6f8c0d2f07e27cc972bbe619040834fd6bc53c.NewAdministrativeUnitItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get directory
func (m *DirectoryRequestBuilder) CreateGetRequestInformation(options *DirectoryRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update directory
func (m *DirectoryRequestBuilder) CreatePatchRequestInformation(options *DirectoryRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeletedItems the deletedItems property
func (m *DirectoryRequestBuilder) DeletedItems()(*i608270f0d53df54d261f336ae52901ae79654f5ed54338f92dab3b2188e27aeb.DeletedItemsRequestBuilder) {
    return i608270f0d53df54d261f336ae52901ae79654f5ed54338f92dab3b2188e27aeb.NewDeletedItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletedItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.directory.deletedItems.item collection
func (m *DirectoryRequestBuilder) DeletedItemsById(id string)(*ic11656dccc6289d407b8e6129a339cd67d5b4802e9850fa2b7197192054f419c.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ic11656dccc6289d407b8e6129a339cd67d5b4802e9850fa2b7197192054f419c.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get directory
func (m *DirectoryRequestBuilder) Get(options *DirectoryRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Directoryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Directoryable), nil
}
// Patch update directory
func (m *DirectoryRequestBuilder) Patch(options *DirectoryRequestBuilderPatchOptions)(error) {
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
