package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

// ConnectionOperation provides operations to manage the collection of externalConnection entities.
type ConnectionOperation struct {
    Entity
    // If status is failed, provides more information about the error that caused the failure.
    error PublicErrorable;
    // Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
    status *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionOperationStatus;
}
// NewConnectionOperation instantiates a new connectionOperation and sets the default values.
func NewConnectionOperation()(*ConnectionOperation) {
    m := &ConnectionOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConnectionOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectionOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConnectionOperation(), nil
}
// GetError gets the error property value. If status is failed, provides more information about the error that caused the failure.
func (m *ConnectionOperation) GetError()(PublicErrorable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectionOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseConnectionOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionOperationStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
func (m *ConnectionOperation) GetStatus()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ConnectionOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConnectionOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. If status is failed, provides more information about the error that caused the failure.
func (m *ConnectionOperation) SetError(value PublicErrorable)() {
    if m != nil {
        m.error = value
    }
}
// SetStatus sets the status property value. Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
func (m *ConnectionOperation) SetStatus(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
