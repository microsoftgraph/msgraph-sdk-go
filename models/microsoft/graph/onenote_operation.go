package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnenoteOperation struct {
    Operation
    error *OnenoteOperationError;
    percentComplete *string;
    resourceId *string;
    resourceLocation *string;
}
func NewOnenoteOperation()(*OnenoteOperation) {
    m := &OnenoteOperation{
        Operation: *NewOperation(),
    }
    return m
}
func (m *OnenoteOperation) GetError()(*OnenoteOperationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *OnenoteOperation) GetPercentComplete()(*string) {
    if m == nil {
        return nil
    } else {
        return m.percentComplete
    }
}
func (m *OnenoteOperation) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
func (m *OnenoteOperation) GetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLocation
    }
}
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
func (m *OnenoteOperation) SetError(value *OnenoteOperationError)() {
    m.error = value
}
func (m *OnenoteOperation) SetPercentComplete(value *string)() {
    m.percentComplete = value
}
func (m *OnenoteOperation) SetResourceId(value *string)() {
    m.resourceId = value
}
func (m *OnenoteOperation) SetResourceLocation(value *string)() {
    m.resourceLocation = value
}
