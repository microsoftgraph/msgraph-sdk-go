package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2f2bb214f1739b3a1d8adbadddf67f6653077483f4e552d281976eb80f8fac45 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/parentnotebook"
    i374a3854fede5c16a128fa7c82eae208a7de81effb7f1d8f274b78911fa48ad1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/content"
    i3c070dbc71b1cbeee6f41784416f45b58bd9b4308ed897a8c671ad0f44f08e01 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/onenotepatchcontent"
    i3c35da6a18d6d95413d165044b88afb586ff7c5c7eba173d4032a95c42c129bd "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/preview"
    iad1ad062d8bf6a85e7598f552dee37c991382cfb2a48b3b5f27d661218e4d173 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/parentsection"
    if6899f5f5dda6c8021139cab7e1caa965984ce2c56373e27d27a4ee984e7bfc1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/copytosection"
)

// Builds and executes requests for operations under \users\{user-id}\onenote\notebooks\{notebook-id}\sections\{onenoteSection-id}\pages\{onenotePage-id}
type OnenotePageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The collection of pages in the section.  Read-only. Nullable.
type OnenotePageRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new OnenotePageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenotePageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    m := &OnenotePageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/onenote/notebooks/{notebook_id}/sections/{onenoteSection_id}/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OnenotePageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenotePageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenotePageRequestBuilder) Content()(*i374a3854fede5c16a128fa7c82eae208a7de81effb7f1d8f274b78911fa48ad1.ContentRequestBuilder) {
    return i374a3854fede5c16a128fa7c82eae208a7de81effb7f1d8f274b78911fa48ad1.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) CopyToSection()(*if6899f5f5dda6c8021139cab7e1caa965984ce2c56373e27d27a4ee984e7bfc1.CopyToSectionRequestBuilder) {
    return if6899f5f5dda6c8021139cab7e1caa965984ce2c56373e27d27a4ee984e7bfc1.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *OnenotePageRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *OnenotePageRequestBuilder) CreateGetRequestInformation(q func (value *OnenotePageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OnenotePageRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *OnenotePageRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePage, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenotePageRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenotePageRequestBuilder) Get(q func (value *OnenotePageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePage, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenotePage() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePage), nil
}
func (m *OnenotePageRequestBuilder) OnenotePatchContent()(*i3c070dbc71b1cbeee6f41784416f45b58bd9b4308ed897a8c671ad0f44f08e01.OnenotePatchContentRequestBuilder) {
    return i3c070dbc71b1cbeee6f41784416f45b58bd9b4308ed897a8c671ad0f44f08e01.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentNotebook()(*i2f2bb214f1739b3a1d8adbadddf67f6653077483f4e552d281976eb80f8fac45.ParentNotebookRequestBuilder) {
    return i2f2bb214f1739b3a1d8adbadddf67f6653077483f4e552d281976eb80f8fac45.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentSection()(*iad1ad062d8bf6a85e7598f552dee37c991382cfb2a48b3b5f27d661218e4d173.ParentSectionRequestBuilder) {
    return iad1ad062d8bf6a85e7598f552dee37c991382cfb2a48b3b5f27d661218e4d173.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenotePageRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePage, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// Builds and executes requests for operations under \users\{user-id}\onenote\notebooks\{notebook-id}\sections\{onenoteSection-id}\pages\{onenotePage-id}\microsoft.graph.preview()
func (m *OnenotePageRequestBuilder) Preview()(*i3c35da6a18d6d95413d165044b88afb586ff7c5c7eba173d4032a95c42c129bd.PreviewRequestBuilder) {
    return i3c35da6a18d6d95413d165044b88afb586ff7c5c7eba173d4032a95c42c129bd.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
