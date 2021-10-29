package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ReportRoot struct {
    Entity
    // 
    dailyPrintUsageByPrinter []PrintUsageByPrinter;
    // 
    dailyPrintUsageByUser []PrintUsageByUser;
    // 
    monthlyPrintUsageByPrinter []PrintUsageByPrinter;
    // 
    monthlyPrintUsageByUser []PrintUsageByUser;
}
// Instantiates a new reportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the dailyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByPrinter
    }
}
// Gets the dailyPrintUsageByUser property value. 
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByUser
    }
}
// Gets the monthlyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByPrinter
    }
}
// Gets the monthlyPrintUsageByUser property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByUser
    }
}
// The deserialization information for the current model
func (m *ReportRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dailyPrintUsageByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        res := make([]PrintUsageByPrinter, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintUsageByPrinter))
        }
        m.SetDailyPrintUsageByPrinter(res)
        return nil
    }
    res["dailyPrintUsageByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        res := make([]PrintUsageByUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintUsageByUser))
        }
        m.SetDailyPrintUsageByUser(res)
        return nil
    }
    res["monthlyPrintUsageByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        res := make([]PrintUsageByPrinter, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintUsageByPrinter))
        }
        m.SetMonthlyPrintUsageByPrinter(res)
        return nil
    }
    res["monthlyPrintUsageByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        res := make([]PrintUsageByUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintUsageByUser))
        }
        m.SetMonthlyPrintUsageByUser(res)
        return nil
    }
    return res
}
func (m *ReportRoot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ReportRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDailyPrintUsageByPrinter()))
        for i, v := range m.GetDailyPrintUsageByPrinter() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDailyPrintUsageByUser()))
        for i, v := range m.GetDailyPrintUsageByUser() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonthlyPrintUsageByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageByPrinter() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonthlyPrintUsageByUser()))
        for i, v := range m.GetMonthlyPrintUsageByUser() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the dailyPrintUsageByPrinter property value. 
// Parameters:
//  - value : Value to set for the dailyPrintUsageByPrinter property.
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    m.dailyPrintUsageByPrinter = value
}
// Sets the dailyPrintUsageByUser property value. 
// Parameters:
//  - value : Value to set for the dailyPrintUsageByUser property.
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUser)() {
    m.dailyPrintUsageByUser = value
}
// Sets the monthlyPrintUsageByPrinter property value. 
// Parameters:
//  - value : Value to set for the monthlyPrintUsageByPrinter property.
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    m.monthlyPrintUsageByPrinter = value
}
// Sets the monthlyPrintUsageByUser property value. 
// Parameters:
//  - value : Value to set for the monthlyPrintUsageByUser property.
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUser)() {
    m.monthlyPrintUsageByUser = value
}
