package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DataPolicyOperation 
type DataPolicyOperation struct {
    Entity
    // Represents when the request for this data policy operation was completed, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Null until the operation completes.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Specifies the progress of an operation.
    progress *float64;
    // Possible values are: notStarted, running, complete, failed, unknownFutureValue.
    status *DataPolicyOperationStatus;
    // The URL location to where data is being exported for export requests.
    storageLocation *string;
    // Represents when the request for this data operation was submitted, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    submittedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The id for the user on whom the operation is performed.
    userId *string;
}
// NewDataPolicyOperation instantiates a new dataPolicyOperation and sets the default values.
func NewDataPolicyOperation()(*DataPolicyOperation) {
    m := &DataPolicyOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompletedDateTime gets the completedDateTime property value. Represents when the request for this data policy operation was completed, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Null until the operation completes.
func (m *DataPolicyOperation) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetProgress gets the progress property value. Specifies the progress of an operation.
func (m *DataPolicyOperation) GetProgress()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.progress
    }
}
// GetStatus gets the status property value. Possible values are: notStarted, running, complete, failed, unknownFutureValue.
func (m *DataPolicyOperation) GetStatus()(*DataPolicyOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetStorageLocation gets the storageLocation property value. The URL location to where data is being exported for export requests.
func (m *DataPolicyOperation) GetStorageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.storageLocation
    }
}
// GetSubmittedDateTime gets the submittedDateTime property value. Represents when the request for this data operation was submitted, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *DataPolicyOperation) GetSubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.submittedDateTime
    }
}
// GetUserId gets the userId property value. The id for the user on whom the operation is performed.
func (m *DataPolicyOperation) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataPolicyOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["progress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgress(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataPolicyOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DataPolicyOperationStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["storageLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageLocation(val)
        }
        return nil
    }
    res["submittedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedDateTime(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *DataPolicyOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DataPolicyOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("progress", m.GetProgress())
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
    {
        err = writer.WriteStringValue("storageLocation", m.GetStorageLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("submittedDateTime", m.GetSubmittedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. Represents when the request for this data policy operation was completed, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Null until the operation completes.
func (m *DataPolicyOperation) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetProgress sets the progress property value. Specifies the progress of an operation.
func (m *DataPolicyOperation) SetProgress(value *float64)() {
    if m != nil {
        m.progress = value
    }
}
// SetStatus sets the status property value. Possible values are: notStarted, running, complete, failed, unknownFutureValue.
func (m *DataPolicyOperation) SetStatus(value *DataPolicyOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetStorageLocation sets the storageLocation property value. The URL location to where data is being exported for export requests.
func (m *DataPolicyOperation) SetStorageLocation(value *string)() {
    if m != nil {
        m.storageLocation = value
    }
}
// SetSubmittedDateTime sets the submittedDateTime property value. Represents when the request for this data operation was submitted, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *DataPolicyOperation) SetSubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.submittedDateTime = value
    }
}
// SetUserId sets the userId property value. The id for the user on whom the operation is performed.
func (m *DataPolicyOperation) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
