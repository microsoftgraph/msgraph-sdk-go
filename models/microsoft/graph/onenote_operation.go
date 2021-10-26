package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnenoteOperation struct {
    Operation
    // The error returned by the operation.
    error *OnenoteOperationError;
    // The operation percent complete if the operation is still in running status.
    percentComplete *string;
    // The resource id.
    resourceId *string;
    // The resource URI for the object. For example, the resource URI for a copied page or section.
    resourceLocation *string;
}
// Instantiates a new onenoteOperation and sets the default values.
func NewOnenoteOperation()(*OnenoteOperation) {
    m := &OnenoteOperation{
        Operation: *NewOperation(),
    }
    return m
}
// Gets the error property value. The error returned by the operation.
func (m *OnenoteOperation) GetError()(*OnenoteOperationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the percentComplete property value. The operation percent complete if the operation is still in running status.
func (m *OnenoteOperation) GetPercentComplete()(*string) {
    if m == nil {
        return nil
    } else {
        return m.percentComplete
    }
}
// Gets the resourceId property value. The resource id.
func (m *OnenoteOperation) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// Gets the resourceLocation property value. The resource URI for the object. For example, the resource URI for a copied page or section.
func (m *OnenoteOperation) GetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLocation
    }
}
// The deserialization information for the current model
func (m *OnenoteOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Operation.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenoteOperationError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*OnenoteOperationError))
        return nil
    }
    res["percentComplete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPercentComplete(val)
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    res["resourceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceLocation(val)
        return nil
    }
    return res
}
func (m *OnenoteOperation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnenoteOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Operation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("percentComplete", m.GetPercentComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceLocation", m.GetResourceLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the error property value. The error returned by the operation.
// Parameters:
//  - value : Value to set for the error property.
func (m *OnenoteOperation) SetError(value *OnenoteOperationError)() {
    m.error = value
}
// Sets the percentComplete property value. The operation percent complete if the operation is still in running status.
// Parameters:
//  - value : Value to set for the percentComplete property.
func (m *OnenoteOperation) SetPercentComplete(value *string)() {
    m.percentComplete = value
}
// Sets the resourceId property value. The resource id.
// Parameters:
//  - value : Value to set for the resourceId property.
func (m *OnenoteOperation) SetResourceId(value *string)() {
    m.resourceId = value
}
// Sets the resourceLocation property value. The resource URI for the object. For example, the resource URI for a copied page or section.
// Parameters:
//  - value : Value to set for the resourceLocation property.
func (m *OnenoteOperation) SetResourceLocation(value *string)() {
    m.resourceLocation = value
}
