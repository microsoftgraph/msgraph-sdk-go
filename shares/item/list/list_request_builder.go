package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i6cc54c104a1f556165cc44589ee14207379ed8c37fa85729d5c314e13f2214d7 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/columns"
    i7a0c6962c8b1589339031c077d13646778071c443927dc1046bc04df31627bf6 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/items"
    i8b1352d9ad171bd1e78c354f011b512c0faed7a8b44a2d0af7a87499c0e9f548 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes"
    iaad175a704d2fb8c88af7ffd0f8ebd7bfd4f12f115e3413d8b3c39946957d581 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/subscriptions"
    id412f117e08220004924a69577571fc52ee563dcb4aa3bb6992790f4a56a8672 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/drive"
    i6f7c6ed0403d181f0f9df70f035fc911fc0235b004fca66648da29a67534a2e2 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item"
    ia2c40a1ff086507585f951d28494158116a4fcb5951ee49a77ded46131865f37 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/subscriptions/item"
    ia5941405950ea17199c6d656adb76e8bf88f3f605415a64747a31ce95e0c07b7 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/items/item"
    ifd1fbca74ca99c59f8151a23008189eb5fa814dfb2bdb94e0592368b94d84ebe "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/columns/item"
)

type ListRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ListRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ListRequestBuilder) Columns()(*i6cc54c104a1f556165cc44589ee14207379ed8c37fa85729d5c314e13f2214d7.ColumnsRequestBuilder) {
    return i6cc54c104a1f556165cc44589ee14207379ed8c37fa85729d5c314e13f2214d7.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ColumnsById(id string)(*ifd1fbca74ca99c59f8151a23008189eb5fa814dfb2bdb94e0592368b94d84ebe.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ifd1fbca74ca99c59f8151a23008189eb5fa814dfb2bdb94e0592368b94d84ebe.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/shares/{sharedDriveItem_id}/list{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i8b1352d9ad171bd1e78c354f011b512c0faed7a8b44a2d0af7a87499c0e9f548.ContentTypesRequestBuilder) {
    return i8b1352d9ad171bd1e78c354f011b512c0faed7a8b44a2d0af7a87499c0e9f548.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ContentTypesById(id string)(*i6f7c6ed0403d181f0f9df70f035fc911fc0235b004fca66648da29a67534a2e2.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i6f7c6ed0403d181f0f9df70f035fc911fc0235b004fca66648da29a67534a2e2.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListRequestBuilder) CreateGetRequestInformation(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ListRequestBuilderGetQueryParameters)
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
func (m *ListRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListRequestBuilder) Drive()(*id412f117e08220004924a69577571fc52ee563dcb4aa3bb6992790f4a56a8672.DriveRequestBuilder) {
    return id412f117e08220004924a69577571fc52ee563dcb4aa3bb6992790f4a56a8672.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) Get(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewList() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List), nil
}
func (m *ListRequestBuilder) Items()(*i7a0c6962c8b1589339031c077d13646778071c443927dc1046bc04df31627bf6.ItemsRequestBuilder) {
    return i7a0c6962c8b1589339031c077d13646778071c443927dc1046bc04df31627bf6.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ItemsById(id string)(*ia5941405950ea17199c6d656adb76e8bf88f3f605415a64747a31ce95e0c07b7.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return ia5941405950ea17199c6d656adb76e8bf88f3f605415a64747a31ce95e0c07b7.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListRequestBuilder) Subscriptions()(*iaad175a704d2fb8c88af7ffd0f8ebd7bfd4f12f115e3413d8b3c39946957d581.SubscriptionsRequestBuilder) {
    return iaad175a704d2fb8c88af7ffd0f8ebd7bfd4f12f115e3413d8b3c39946957d581.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) SubscriptionsById(id string)(*ia2c40a1ff086507585f951d28494158116a4fcb5951ee49a77ded46131865f37.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ia2c40a1ff086507585f951d28494158116a4fcb5951ee49a77ded46131865f37.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
