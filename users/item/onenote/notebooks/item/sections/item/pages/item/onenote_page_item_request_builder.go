package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2f2bb214f1739b3a1d8adbadddf67f6653077483f4e552d281976eb80f8fac45 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/parentnotebook"
    i374a3854fede5c16a128fa7c82eae208a7de81effb7f1d8f274b78911fa48ad1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/content"
    i3c070dbc71b1cbeee6f41784416f45b58bd9b4308ed897a8c671ad0f44f08e01 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/onenotepatchcontent"
    i3c35da6a18d6d95413d165044b88afb586ff7c5c7eba173d4032a95c42c129bd "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/preview"
    iad1ad062d8bf6a85e7598f552dee37c991382cfb2a48b3b5f27d661218e4d173 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/parentsection"
    if6899f5f5dda6c8021139cab7e1caa965984ce2c56373e27d27a4ee984e7bfc1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/copytosection"
)

// OnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type OnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenotePageItemRequestBuilderDeleteOptions options for Delete
type OnenotePageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageItemRequestBuilderGetOptions options for Get
type OnenotePageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenotePageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type OnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenotePageItemRequestBuilderPatchOptions options for Patch
type OnenotePageItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePageable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    m := &OnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/notebooks/{notebook_id}/sections/{onenoteSection_id}/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenotePageItemRequestBuilder) Content()(*i374a3854fede5c16a128fa7c82eae208a7de81effb7f1d8f274b78911fa48ad1.ContentRequestBuilder) {
    return i374a3854fede5c16a128fa7c82eae208a7de81effb7f1d8f274b78911fa48ad1.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageItemRequestBuilder) CopyToSection()(*if6899f5f5dda6c8021139cab7e1caa965984ce2c56373e27d27a4ee984e7bfc1.CopyToSectionRequestBuilder) {
    return if6899f5f5dda6c8021139cab7e1caa965984ce2c56373e27d27a4ee984e7bfc1.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pages for users
func (m *OnenotePageItemRequestBuilder) CreateDeleteRequestInformation(options *OnenotePageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) CreateGetRequestInformation(options *OnenotePageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property pages in users
func (m *OnenotePageItemRequestBuilder) CreatePatchRequestInformation(options *OnenotePageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property pages for users
func (m *OnenotePageItemRequestBuilder) Delete(options *OnenotePageItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) Get(options *OnenotePageItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateOnenotePageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePageable), nil
}
func (m *OnenotePageItemRequestBuilder) OnenotePatchContent()(*i3c070dbc71b1cbeee6f41784416f45b58bd9b4308ed897a8c671ad0f44f08e01.OnenotePatchContentRequestBuilder) {
    return i3c070dbc71b1cbeee6f41784416f45b58bd9b4308ed897a8c671ad0f44f08e01.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageItemRequestBuilder) ParentNotebook()(*i2f2bb214f1739b3a1d8adbadddf67f6653077483f4e552d281976eb80f8fac45.ParentNotebookRequestBuilder) {
    return i2f2bb214f1739b3a1d8adbadddf67f6653077483f4e552d281976eb80f8fac45.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageItemRequestBuilder) ParentSection()(*iad1ad062d8bf6a85e7598f552dee37c991382cfb2a48b3b5f27d661218e4d173.ParentSectionRequestBuilder) {
    return iad1ad062d8bf6a85e7598f552dee37c991382cfb2a48b3b5f27d661218e4d173.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property pages in users
func (m *OnenotePageItemRequestBuilder) Patch(options *OnenotePageItemRequestBuilderPatchOptions)(error) {
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
// Preview provides operations to call the preview method.
func (m *OnenotePageItemRequestBuilder) Preview()(*i3c35da6a18d6d95413d165044b88afb586ff7c5c7eba173d4032a95c42c129bd.PreviewRequestBuilder) {
    return i3c35da6a18d6d95413d165044b88afb586ff7c5c7eba173d4032a95c42c129bd.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
