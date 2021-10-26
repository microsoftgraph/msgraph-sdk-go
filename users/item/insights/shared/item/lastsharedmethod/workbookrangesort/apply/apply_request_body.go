package apply

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
type ApplyRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    fields []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookSortField;
    // 
    hasHeaders *bool;
    // 
    matchCase *bool;
    // 
    method *string;
    // 
    orientation *string;
}
// Instantiates a new applyRequestBody and sets the default values.
func NewApplyRequestBody()(*ApplyRequestBody) {
    m := &ApplyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the fields property value. 
func (m *ApplyRequestBody) GetFields()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookSortField) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// Gets the hasHeaders property value. 
func (m *ApplyRequestBody) GetHasHeaders()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasHeaders
    }
}
// Gets the matchCase property value. 
func (m *ApplyRequestBody) GetMatchCase()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.matchCase
    }
}
// Gets the method property value. 
func (m *ApplyRequestBody) GetMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.method
    }
}
// Gets the orientation property value. 
func (m *ApplyRequestBody) GetOrientation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
// The deserialization information for the current model
func (m *ApplyRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookSortField() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookSortField, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookSortField))
        }
        m.SetFields(res)
        return nil
    }
    res["hasHeaders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasHeaders(val)
        return nil
    }
    res["matchCase"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMatchCase(val)
        return nil
    }
    res["method"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMethod(val)
        return nil
    }
    res["orientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrientation(val)
        return nil
    }
    return res
}
func (m *ApplyRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ApplyRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("fields", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasHeaders", m.GetHasHeaders())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("matchCase", m.GetMatchCase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("method", m.GetMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("orientation", m.GetOrientation())
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
func (m *ApplyRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the fields property value. 
// Parameters:
//  - value : Value to set for the fields property.
func (m *ApplyRequestBody) SetFields(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookSortField)() {
    m.fields = value
}
// Sets the hasHeaders property value. 
// Parameters:
//  - value : Value to set for the hasHeaders property.
func (m *ApplyRequestBody) SetHasHeaders(value *bool)() {
    m.hasHeaders = value
}
// Sets the matchCase property value. 
// Parameters:
//  - value : Value to set for the matchCase property.
func (m *ApplyRequestBody) SetMatchCase(value *bool)() {
    m.matchCase = value
}
// Sets the method property value. 
// Parameters:
//  - value : Value to set for the method property.
func (m *ApplyRequestBody) SetMethod(value *string)() {
    m.method = value
}
// Sets the orientation property value. 
// Parameters:
//  - value : Value to set for the orientation property.
func (m *ApplyRequestBody) SetOrientation(value *string)() {
    m.orientation = value
}
