package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenoteOperation provides operations to manage the educationRoot singleton.
type OnenoteOperation struct {
    Operation
    // The error returned by the operation.
    error OnenoteOperationErrorable;
    // The operation percent complete if the operation is still in running status.
    percentComplete *string;
    // The resource id.
    resourceId *string;
    // The resource URI for the object. For example, the resource URI for a copied page or section.
    resourceLocation *string;
}
// NewOnenoteOperation instantiates a new onenoteOperation and sets the default values.
func NewOnenoteOperation()(*OnenoteOperation) {
    m := &OnenoteOperation{
        Operation: *NewOperation(),
    }
    return m
}
// CreateOnenoteOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnenoteOperation(), nil
}
// GetError gets the error property value. The error returned by the operation.
func (m *OnenoteOperation) GetError()(OnenoteOperationErrorable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Operation.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnenoteOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(OnenoteOperationErrorable))
        }
        return nil
    }
    res["percentComplete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentComplete(val)
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["resourceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceLocation(val)
        }
        return nil
    }
    return res
}
// GetPercentComplete gets the percentComplete property value. The operation percent complete if the operation is still in running status.
func (m *OnenoteOperation) GetPercentComplete()(*string) {
    if m == nil {
        return nil
    } else {
        return m.percentComplete
    }
}
// GetResourceId gets the resourceId property value. The resource id.
func (m *OnenoteOperation) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetResourceLocation gets the resourceLocation property value. The resource URI for the object. For example, the resource URI for a copied page or section.
func (m *OnenoteOperation) GetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLocation
    }
}
func (m *OnenoteOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetError sets the error property value. The error returned by the operation.
func (m *OnenoteOperation) SetError(value OnenoteOperationErrorable)() {
    if m != nil {
        m.error = value
    }
}
// SetPercentComplete sets the percentComplete property value. The operation percent complete if the operation is still in running status.
func (m *OnenoteOperation) SetPercentComplete(value *string)() {
    if m != nil {
        m.percentComplete = value
    }
}
// SetResourceId sets the resourceId property value. The resource id.
func (m *OnenoteOperation) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetResourceLocation sets the resourceLocation property value. The resource URI for the object. For example, the resource URI for a copied page or section.
func (m *OnenoteOperation) SetResourceLocation(value *string)() {
    if m != nil {
        m.resourceLocation = value
    }
}
