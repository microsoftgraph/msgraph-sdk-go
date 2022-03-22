package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningStep 
type ProvisioningStep struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Summary of what occurred during the step.
    description *string;
    // Details of what occurred during the step.
    details DetailsInfoable;
    // Name of the step.
    name *string;
    // Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export, unknownFutureValue.
    provisioningStepType *ProvisioningStepType;
    // Status of the step. Possible values are: success, warning,  failure, skipped, unknownFutureValue.
    status *ProvisioningResult;
}
// NewProvisioningStep instantiates a new provisioningStep and sets the default values.
func NewProvisioningStep()(*ProvisioningStep) {
    m := &ProvisioningStep{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProvisioningStepFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningStepFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewProvisioningStep(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningStep) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. Summary of what occurred during the step.
func (m *ProvisioningStep) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDetails gets the details property value. Details of what occurred during the step.
func (m *ProvisioningStep) GetDetails()(DetailsInfoable) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetailsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(DetailsInfoable))
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["provisioningStepType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningStepType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStepType(val.(*ProvisioningStepType))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProvisioningResult))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the step.
func (m *ProvisioningStep) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetProvisioningStepType gets the provisioningStepType property value. Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export, unknownFutureValue.
func (m *ProvisioningStep) GetProvisioningStepType()(*ProvisioningStepType) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStepType
    }
}
// GetStatus gets the status property value. Status of the step. Possible values are: success, warning,  failure, skipped, unknownFutureValue.
func (m *ProvisioningStep) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *ProvisioningStep) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningStepType() != nil {
        cast := (*m.GetProvisioningStepType()).String()
        err := writer.WriteStringValue("provisioningStepType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ProvisioningStep) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. Summary of what occurred during the step.
func (m *ProvisioningStep) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDetails sets the details property value. Details of what occurred during the step.
func (m *ProvisioningStep) SetDetails(value DetailsInfoable)() {
    if m != nil {
        m.details = value
    }
}
// SetName sets the name property value. Name of the step.
func (m *ProvisioningStep) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetProvisioningStepType sets the provisioningStepType property value. Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export, unknownFutureValue.
func (m *ProvisioningStep) SetProvisioningStepType(value *ProvisioningStepType)() {
    if m != nil {
        m.provisioningStepType = value
    }
}
// SetStatus sets the status property value. Status of the step. Possible values are: success, warning,  failure, skipped, unknownFutureValue.
func (m *ProvisioningStep) SetStatus(value *ProvisioningResult)() {
    if m != nil {
        m.status = value
    }
}
