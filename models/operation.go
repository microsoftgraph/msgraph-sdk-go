package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Operation provides operations to manage the collection of application entities.
type Operation struct {
    Entity
    // The start time of the operation.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time of the last action of the operation.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Possible values are: notStarted, running, completed, failed. Read-only.
    status *OperationStatus
}
// NewOperation instantiates a new operation and sets the default values.
func NewOperation()(*Operation) {
    m := &Operation{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.operation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.onenoteOperation":
                        return NewOnenoteOperation(), nil
                }
            }
        }
    }
    return NewOperation(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The start time of the operation.
func (m *Operation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Operation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*OperationStatus))
        }
        return nil
    }
    return res
}
// GetLastActionDateTime gets the lastActionDateTime property value. The time of the last action of the operation.
func (m *Operation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetStatus gets the status property value. Possible values are: notStarted, running, completed, failed. Read-only.
func (m *Operation) GetStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *Operation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The start time of the operation.
func (m *Operation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The time of the last action of the operation.
func (m *Operation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetStatus sets the status property value. Possible values are: notStarted, running, completed, failed. Read-only.
func (m *Operation) SetStatus(value *OperationStatus)() {
    if m != nil {
        m.status = value
    }
}
