package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintUsageByPrinter struct {
    PrintUsage
    // 
    printerId *string;
}
// Instantiates a new printUsageByPrinter and sets the default values.
func NewPrintUsageByPrinter()(*PrintUsageByPrinter) {
    m := &PrintUsageByPrinter{
        PrintUsage: *NewPrintUsage(),
    }
    return m
}
// Gets the printerId property value. 
func (m *PrintUsageByPrinter) GetPrinterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.printerId
    }
}
// The deserialization information for the current model
func (m *PrintUsageByPrinter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PrintUsage.GetFieldDeserializers()
    res["printerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrinterId(val)
        return nil
    }
    return res
}
func (m *PrintUsageByPrinter) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintUsageByPrinter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PrintUsage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("printerId", m.GetPrinterId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the printerId property value. 
// Parameters:
//  - value : Value to set for the printerId property.
func (m *PrintUsageByPrinter) SetPrinterId(value *string)() {
    m.printerId = value
}
