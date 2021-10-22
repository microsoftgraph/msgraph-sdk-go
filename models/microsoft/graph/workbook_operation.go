package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookOperation struct {
    Entity
    error *WorkbookOperationError;
    resourceLocation *string;
    status *WorkbookOperationStatus;
}
func NewWorkbookOperation()(*WorkbookOperation) {
    m := &WorkbookOperation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookOperation) GetError()(*WorkbookOperationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *WorkbookOperation) GetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLocation
    }
}
func (m *WorkbookOperation) GetStatus()(*WorkbookOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *WorkbookOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookOperationError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*WorkbookOperationError))
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkbookOperationStatus)
        if err != nil {
            return err
        }
        cast := val.(WorkbookOperationStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *WorkbookOperation) IsNil()(bool) {
    return m == nil
}
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
func (m *WorkbookOperation) SetError(value *WorkbookOperationError)() {
    m.error = value
}
func (m *WorkbookOperation) SetResourceLocation(value *string)() {
    m.resourceLocation = value
}
func (m *WorkbookOperation) SetStatus(value *WorkbookOperationStatus)() {
    m.status = value
}
