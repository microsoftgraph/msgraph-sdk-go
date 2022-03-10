package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterStatus provides operations to manage the print singleton.
type PrinterStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A human-readable description of the printer's current processing state. Read-only.
    description *string;
    // The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
    details []PrinterProcessingStateDetail;
    // The current processing state. Valid values are described in the following table. Read-only.
    state *PrinterProcessingState;
}
// NewPrinterStatus instantiates a new printerStatus and sets the default values.
func NewPrinterStatus()(*PrinterStatus) {
    m := &PrinterStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrinterStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterStatusFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrinterStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. A human-readable description of the printer's current processing state. Read-only.
func (m *PrinterStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDetails gets the details property value. The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) GetDetails()([]PrinterProcessingStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfEnumValues(ParsePrinterProcessingStateDetail)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterProcessingStateDetail, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrinterProcessingStateDetail))
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PrinterProcessingState))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The current processing state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) GetState()(*PrinterProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *PrinterStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrinterStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDetails() != nil {
        err := writer.WriteCollectionOfStringValues("details", SerializePrinterProcessingStateDetail(m.GetDetails()))
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *PrinterStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. A human-readable description of the printer's current processing state. Read-only.
func (m *PrinterStatus) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDetails sets the details property value. The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) SetDetails(value []PrinterProcessingStateDetail)() {
    if m != nil {
        m.details = value
    }
}
// SetState sets the state property value. The current processing state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) SetState(value *PrinterProcessingState)() {
    if m != nil {
        m.state = value
    }
}
