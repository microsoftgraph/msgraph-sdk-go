package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ResourceAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Allowed Actions
    allowedResourceActions []string;
    // Not Allowed Actions.
    notAllowedResourceActions []string;
}
// Instantiates a new resourceAction and sets the default values.
func NewResourceAction()(*ResourceAction) {
    m := &ResourceAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowedResourceActions property value. Allowed Actions
func (m *ResourceAction) GetAllowedResourceActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.allowedResourceActions
    }
}
// Gets the notAllowedResourceActions property value. Not Allowed Actions.
func (m *ResourceAction) GetNotAllowedResourceActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notAllowedResourceActions
    }
}
// The deserialization information for the current model
func (m *ResourceAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedResourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedResourceActions(res)
        }
        return nil
    }
    res["notAllowedResourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotAllowedResourceActions(res)
        }
        return nil
    }
    return res
}
func (m *ResourceAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ResourceAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("allowedResourceActions", m.GetAllowedResourceActions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("notAllowedResourceActions", m.GetNotAllowedResourceActions())
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
func (m *ResourceAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowedResourceActions property value. Allowed Actions
// Parameters:
//  - value : Value to set for the allowedResourceActions property.
func (m *ResourceAction) SetAllowedResourceActions(value []string)() {
    m.allowedResourceActions = value
}
// Sets the notAllowedResourceActions property value. Not Allowed Actions.
// Parameters:
//  - value : Value to set for the notAllowedResourceActions property.
func (m *ResourceAction) SetNotAllowedResourceActions(value []string)() {
    m.notAllowedResourceActions = value
}
