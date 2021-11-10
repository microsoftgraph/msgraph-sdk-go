package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new printerStatus and sets the default values.
func NewPrinterStatus()(*PrinterStatus) {
    m := &PrinterStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. A human-readable description of the printer's current processing state. Read-only.
func (m *PrinterStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the details property value. The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) GetDetails()([]PrinterProcessingStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// Gets the state property value. The current processing state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) GetState()(*PrinterProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
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
            cast := val.(PrinterProcessingState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *PrinterStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrinterStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("details", SerializePrinterProcessingStateDetail(m.GetDetails()))
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PrinterStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. A human-readable description of the printer's current processing state. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *PrinterStatus) SetDescription(value *string)() {
    m.description = value
}
// Sets the details property value. The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
// Parameters:
//  - value : Value to set for the details property.
func (m *PrinterStatus) SetDetails(value []PrinterProcessingStateDetail)() {
    m.details = value
}
// Sets the state property value. The current processing state. Valid values are described in the following table. Read-only.
// Parameters:
//  - value : Value to set for the state property.
func (m *PrinterStatus) SetState(value *PrinterProcessingState)() {
    m.state = value
}
