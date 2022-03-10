package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppDiagnosticStatus provides operations to call the getManagedAppDiagnosticStatuses method.
type ManagedAppDiagnosticStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Instruction on how to mitigate a failed validation
    mitigationInstruction *string;
    // The state of the operation
    state *string;
    // The validation friendly name
    validationName *string;
}
// NewManagedAppDiagnosticStatus instantiates a new managedAppDiagnosticStatus and sets the default values.
func NewManagedAppDiagnosticStatus()(*ManagedAppDiagnosticStatus) {
    m := &ManagedAppDiagnosticStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedAppDiagnosticStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppDiagnosticStatusFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagedAppDiagnosticStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedAppDiagnosticStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppDiagnosticStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["mitigationInstruction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMitigationInstruction(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["validationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationName(val)
        }
        return nil
    }
    return res
}
// GetMitigationInstruction gets the mitigationInstruction property value. Instruction on how to mitigate a failed validation
func (m *ManagedAppDiagnosticStatus) GetMitigationInstruction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mitigationInstruction
    }
}
// GetState gets the state property value. The state of the operation
func (m *ManagedAppDiagnosticStatus) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetValidationName gets the validationName property value. The validation friendly name
func (m *ManagedAppDiagnosticStatus) GetValidationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.validationName
    }
}
func (m *ManagedAppDiagnosticStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedAppDiagnosticStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("mitigationInstruction", m.GetMitigationInstruction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("validationName", m.GetValidationName())
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
func (m *ManagedAppDiagnosticStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMitigationInstruction sets the mitigationInstruction property value. Instruction on how to mitigate a failed validation
func (m *ManagedAppDiagnosticStatus) SetMitigationInstruction(value *string)() {
    if m != nil {
        m.mitigationInstruction = value
    }
}
// SetState sets the state property value. The state of the operation
func (m *ManagedAppDiagnosticStatus) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetValidationName sets the validationName property value. The validation friendly name
func (m *ManagedAppDiagnosticStatus) SetValidationName(value *string)() {
    if m != nil {
        m.validationName = value
    }
}
