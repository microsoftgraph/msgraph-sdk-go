package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintUsageByPrinter provides operations to manage the reportRoot singleton.
type PrintUsageByPrinter struct {
    PrintUsage
    // 
    printerId *string;
}
// NewPrintUsageByPrinter instantiates a new printUsageByPrinter and sets the default values.
func NewPrintUsageByPrinter()(*PrintUsageByPrinter) {
    m := &PrintUsageByPrinter{
        PrintUsage: *NewPrintUsage(),
    }
    return m
}
// CreatePrintUsageByPrinterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintUsageByPrinterFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintUsageByPrinter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintUsageByPrinter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PrintUsage.GetFieldDeserializers()
    res["printerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterId(val)
        }
        return nil
    }
    return res
}
// GetPrinterId gets the printerId property value. 
func (m *PrintUsageByPrinter) GetPrinterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.printerId
    }
}
func (m *PrintUsageByPrinter) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetPrinterId sets the printerId property value. 
func (m *PrintUsageByPrinter) SetPrinterId(value *string)() {
    if m != nil {
        m.printerId = value
    }
}
