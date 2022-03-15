package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookApplication provides operations to manage the collection of drive entities.
type WorkbookApplication struct {
    Entity
    // Returns the calculation mode used in the workbook. Possible values are: Automatic, AutomaticExceptTables, Manual.
    calculationMode *string;
}
// NewWorkbookApplication instantiates a new workbookApplication and sets the default values.
func NewWorkbookApplication()(*WorkbookApplication) {
    m := &WorkbookApplication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookApplicationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookApplication(), nil
}
// GetCalculationMode gets the calculationMode property value. Returns the calculation mode used in the workbook. Possible values are: Automatic, AutomaticExceptTables, Manual.
func (m *WorkbookApplication) GetCalculationMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.calculationMode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calculationMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculationMode(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookApplication) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("calculationMode", m.GetCalculationMode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCalculationMode sets the calculationMode property value. Returns the calculation mode used in the workbook. Possible values are: Automatic, AutomaticExceptTables, Manual.
func (m *WorkbookApplication) SetCalculationMode(value *string)() {
    if m != nil {
        m.calculationMode = value
    }
}
