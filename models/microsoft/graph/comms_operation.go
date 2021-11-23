package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CommsOperation 
type CommsOperation struct {
    Entity
    // Unique Client Context string. Max limit is 256 chars.
    clientContext *string;
    // The result information. Read-only.
    resultInfo *ResultInfo;
    // Possible values are: notStarted, running, completed, failed. Read-only.
    status *OperationStatus;
}
// NewCommsOperation instantiates a new commsOperation and sets the default values.
func NewCommsOperation()(*CommsOperation) {
    m := &CommsOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetClientContext gets the clientContext property value. Unique Client Context string. Max limit is 256 chars.
func (m *CommsOperation) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
// GetResultInfo gets the resultInfo property value. The result information. Read-only.
func (m *CommsOperation) GetResultInfo()(*ResultInfo) {
    if m == nil {
        return nil
    } else {
        return m.resultInfo
    }
}
// GetStatus gets the status property value. Possible values are: notStarted, running, completed, failed. Read-only.
func (m *CommsOperation) GetStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommsOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["resultInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResultInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultInfo(val.(*ResultInfo))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(OperationStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *CommsOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CommsOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
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
// SetClientContext sets the clientContext property value. Unique Client Context string. Max limit is 256 chars.
func (m *CommsOperation) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetResultInfo sets the resultInfo property value. The result information. Read-only.
func (m *CommsOperation) SetResultInfo(value *ResultInfo)() {
    m.resultInfo = value
}
// SetStatus sets the status property value. Possible values are: notStarted, running, completed, failed. Read-only.
func (m *CommsOperation) SetStatus(value *OperationStatus)() {
    m.status = value
}
