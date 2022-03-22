package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PendingOperations 
type PendingOperations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A property that indicates that an operation that might update the binary content of a file is pending completion.
    pendingContentUpdate PendingContentUpdateable;
}
// NewPendingOperations instantiates a new pendingOperations and sets the default values.
func NewPendingOperations()(*PendingOperations) {
    m := &PendingOperations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePendingOperationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePendingOperationsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPendingOperations(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PendingOperations) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PendingOperations) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["pendingContentUpdate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePendingContentUpdateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingContentUpdate(val.(PendingContentUpdateable))
        }
        return nil
    }
    return res
}
// GetPendingContentUpdate gets the pendingContentUpdate property value. A property that indicates that an operation that might update the binary content of a file is pending completion.
func (m *PendingOperations) GetPendingContentUpdate()(PendingContentUpdateable) {
    if m == nil {
        return nil
    } else {
        return m.pendingContentUpdate
    }
}
// Serialize serializes information the current object
func (m *PendingOperations) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pendingContentUpdate", m.GetPendingContentUpdate())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PendingOperations) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPendingContentUpdate sets the pendingContentUpdate property value. A property that indicates that an operation that might update the binary content of a file is pending completion.
func (m *PendingOperations) SetPendingContentUpdate(value PendingContentUpdateable)() {
    if m != nil {
        m.pendingContentUpdate = value
    }
}
