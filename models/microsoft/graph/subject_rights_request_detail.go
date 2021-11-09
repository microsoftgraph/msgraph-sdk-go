package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SubjectRightsRequestDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Count of items that are excluded from the request.
    excludedItemCount *int64;
    // Count of items per insight.
    insightCounts []KeyValuePair;
    // Count of items found.
    itemCount *int64;
    // Count of item that need review.
    itemNeedReview *int64;
    // Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams.
    productItemCounts []KeyValuePair;
    // Count of items signed off by the administrator.
    signedOffItemCount *int64;
    // Total item size in bytes.
    totalItemSize *int64;
}
// Instantiates a new subjectRightsRequestDetail and sets the default values.
func NewSubjectRightsRequestDetail()(*SubjectRightsRequestDetail) {
    m := &SubjectRightsRequestDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the excludedItemCount property value. Count of items that are excluded from the request.
func (m *SubjectRightsRequestDetail) GetExcludedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.excludedItemCount
    }
}
// Gets the insightCounts property value. Count of items per insight.
func (m *SubjectRightsRequestDetail) GetInsightCounts()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.insightCounts
    }
}
// Gets the itemCount property value. Count of items found.
func (m *SubjectRightsRequestDetail) GetItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.itemCount
    }
}
// Gets the itemNeedReview property value. Count of item that need review.
func (m *SubjectRightsRequestDetail) GetItemNeedReview()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.itemNeedReview
    }
}
// Gets the productItemCounts property value. Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams.
func (m *SubjectRightsRequestDetail) GetProductItemCounts()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.productItemCounts
    }
}
// Gets the signedOffItemCount property value. Count of items signed off by the administrator.
func (m *SubjectRightsRequestDetail) GetSignedOffItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.signedOffItemCount
    }
}
// Gets the totalItemSize property value. Total item size in bytes.
func (m *SubjectRightsRequestDetail) GetTotalItemSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalItemSize
    }
}
// The deserialization information for the current model
func (m *SubjectRightsRequestDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludedItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedItemCount(val)
        }
        return nil
    }
    res["insightCounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetInsightCounts(res)
        }
        return nil
    }
    res["itemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCount(val)
        }
        return nil
    }
    res["itemNeedReview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemNeedReview(val)
        }
        return nil
    }
    res["productItemCounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetProductItemCounts(res)
        }
        return nil
    }
    res["signedOffItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedOffItemCount(val)
        }
        return nil
    }
    res["totalItemSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalItemSize(val)
        }
        return nil
    }
    return res
}
func (m *SubjectRightsRequestDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SubjectRightsRequestDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("excludedItemCount", m.GetExcludedItemCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInsightCounts()))
        for i, v := range m.GetInsightCounts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("insightCounts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("itemCount", m.GetItemCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("itemNeedReview", m.GetItemNeedReview())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProductItemCounts()))
        for i, v := range m.GetProductItemCounts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("productItemCounts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("signedOffItemCount", m.GetSignedOffItemCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalItemSize", m.GetTotalItemSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SubjectRightsRequestDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the excludedItemCount property value. Count of items that are excluded from the request.
// Parameters:
//  - value : Value to set for the excludedItemCount property.
func (m *SubjectRightsRequestDetail) SetExcludedItemCount(value *int64)() {
    m.excludedItemCount = value
}
// Sets the insightCounts property value. Count of items per insight.
// Parameters:
//  - value : Value to set for the insightCounts property.
func (m *SubjectRightsRequestDetail) SetInsightCounts(value []KeyValuePair)() {
    m.insightCounts = value
}
// Sets the itemCount property value. Count of items found.
// Parameters:
//  - value : Value to set for the itemCount property.
func (m *SubjectRightsRequestDetail) SetItemCount(value *int64)() {
    m.itemCount = value
}
// Sets the itemNeedReview property value. Count of item that need review.
// Parameters:
//  - value : Value to set for the itemNeedReview property.
func (m *SubjectRightsRequestDetail) SetItemNeedReview(value *int64)() {
    m.itemNeedReview = value
}
// Sets the productItemCounts property value. Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams.
// Parameters:
//  - value : Value to set for the productItemCounts property.
func (m *SubjectRightsRequestDetail) SetProductItemCounts(value []KeyValuePair)() {
    m.productItemCounts = value
}
// Sets the signedOffItemCount property value. Count of items signed off by the administrator.
// Parameters:
//  - value : Value to set for the signedOffItemCount property.
func (m *SubjectRightsRequestDetail) SetSignedOffItemCount(value *int64)() {
    m.signedOffItemCount = value
}
// Sets the totalItemSize property value. Total item size in bytes.
// Parameters:
//  - value : Value to set for the totalItemSize property.
func (m *SubjectRightsRequestDetail) SetTotalItemSize(value *int64)() {
    m.totalItemSize = value
}
