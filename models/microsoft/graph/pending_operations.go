package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PendingOperations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A property that indicates that an operation that might update the binary content of a file is pending completion.
    pendingContentUpdate *PendingContentUpdate;
}
// Instantiates a new pendingOperations and sets the default values.
func NewPendingOperations()(*PendingOperations) {
    m := &PendingOperations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PendingOperations) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the pendingContentUpdate property value. A property that indicates that an operation that might update the binary content of a file is pending completion.
func (m *PendingOperations) GetPendingContentUpdate()(*PendingContentUpdate) {
    if m == nil {
        return nil
    } else {
        return m.pendingContentUpdate
    }
}
// The deserialization information for the current model
func (m *PendingOperations) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["pendingContentUpdate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPendingContentUpdate() })
        if err != nil {
            return err
        }
        m.SetPendingContentUpdate(val.(*PendingContentUpdate))
        return nil
    }
    return res
}
func (m *PendingOperations) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PendingOperations) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the pendingContentUpdate property value. A property that indicates that an operation that might update the binary content of a file is pending completion.
// Parameters:
//  - value : Value to set for the pendingContentUpdate property.
func (m *PendingOperations) SetPendingContentUpdate(value *PendingContentUpdate)() {
    m.pendingContentUpdate = value
}
