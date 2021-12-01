package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookOperation 
type WorkbookOperation struct {
    Entity
    // The error returned by the operation.
    error *WorkbookOperationError;
    // The resource URI for the result.
    resourceLocation *string;
    // The current status of the operation. Possible values are: NotStarted, Running, Completed, Failed.
    status *WorkbookOperationStatus;
}
// NewWorkbookOperation instantiates a new workbookOperation and sets the default values.
func NewWorkbookOperation()(*WorkbookOperation) {
    m := &WorkbookOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetError gets the error property value. The error returned by the operation.
func (m *WorkbookOperation) GetError()(*WorkbookOperationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetResourceLocation gets the resourceLocation property value. The resource URI for the result.
func (m *WorkbookOperation) GetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLocation
    }
}
// GetStatus gets the status property value. The current status of the operation. Possible values are: NotStarted, Running, Completed, Failed.
func (m *WorkbookOperation) GetStatus()(*WorkbookOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookOperationError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*WorkbookOperationError))
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkbookOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(WorkbookOperationStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *WorkbookOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteStringValue("resourceLocation", m.GetResourceLocation())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. The error returned by the operation.
func (m *WorkbookOperation) SetError(value *WorkbookOperationError)() {
    if m != nil {
        m.error = value
    }
}
// SetResourceLocation sets the resourceLocation property value. The resource URI for the result.
func (m *WorkbookOperation) SetResourceLocation(value *string)() {
    if m != nil {
        m.resourceLocation = value
    }
}
// SetStatus sets the status property value. The current status of the operation. Possible values are: NotStarted, Running, Completed, Failed.
func (m *WorkbookOperation) SetStatus(value *WorkbookOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
