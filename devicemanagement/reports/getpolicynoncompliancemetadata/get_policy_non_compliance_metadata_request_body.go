package getpolicynoncompliancemetadata

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GetPolicyNonComplianceMetadataRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    filter *string;
    // 
    groupBy []string;
    // 
    name *string;
    // 
    orderBy []string;
    // 
    search *string;
    // 
    select_escaped []string;
    // 
    sessionId *string;
    // 
    skip *int32;
    // 
    top *int32;
}
// Instantiates a new getPolicyNonComplianceMetadataRequestBody and sets the default values.
func NewGetPolicyNonComplianceMetadataRequestBody()(*GetPolicyNonComplianceMetadataRequestBody) {
    m := &GetPolicyNonComplianceMetadataRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetPolicyNonComplianceMetadataRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the filter property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// Gets the groupBy property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetGroupBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupBy
    }
}
// Gets the name property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the orderBy property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
// Gets the search property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetSearch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.search
    }
}
// Gets the select_escaped property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetSelect_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
// Gets the sessionId property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetSessionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionId
    }
}
// Gets the skip property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
// Gets the top property value. 
func (m *GetPolicyNonComplianceMetadataRequestBody) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// The deserialization information for the current model
func (m *GetPolicyNonComplianceMetadataRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["groupBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupBy(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["orderBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrderBy(res)
        }
        return nil
    }
    res["search"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val)
        }
        return nil
    }
    res["select_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect_escaped(res)
        }
        return nil
    }
    res["sessionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
        }
        return nil
    }
    res["skip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
func (m *GetPolicyNonComplianceMetadataRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetPolicyNonComplianceMetadataRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("groupBy", m.GetGroupBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("select_escaped", m.GetSelect_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
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
func (m *GetPolicyNonComplianceMetadataRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the filter property value. 
// Parameters:
//  - value : Value to set for the filter property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetFilter(value *string)() {
    m.filter = value
}
// Sets the groupBy property value. 
// Parameters:
//  - value : Value to set for the groupBy property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetGroupBy(value []string)() {
    m.groupBy = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetName(value *string)() {
    m.name = value
}
// Sets the orderBy property value. 
// Parameters:
//  - value : Value to set for the orderBy property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetOrderBy(value []string)() {
    m.orderBy = value
}
// Sets the search property value. 
// Parameters:
//  - value : Value to set for the search property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetSearch(value *string)() {
    m.search = value
}
// Sets the select_escaped property value. 
// Parameters:
//  - value : Value to set for the select_escaped property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetSelect_escaped(value []string)() {
    m.select_escaped = value
}
// Sets the sessionId property value. 
// Parameters:
//  - value : Value to set for the sessionId property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetSessionId(value *string)() {
    m.sessionId = value
}
// Sets the skip property value. 
// Parameters:
//  - value : Value to set for the skip property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
// Sets the top property value. 
// Parameters:
//  - value : Value to set for the top property.
func (m *GetPolicyNonComplianceMetadataRequestBody) SetTop(value *int32)() {
    m.top = value
}
