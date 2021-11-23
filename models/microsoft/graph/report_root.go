package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// reportRoot 
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
// NewReportRoot instantiates a new reportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
        Entity: *NewEntity(),
    }
    return m
}
// GetDailyPrintUsageByPrinter gets the dailyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByPrinter
    }
}
// GetDailyPrintUsageByUser gets the dailyPrintUsageByUser property value. 
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByUser
    }
}
// GetMonthlyPrintUsageByPrinter gets the monthlyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByPrinter
    }
}
// GetMonthlyPrintUsageByUser gets the monthlyPrintUsageByUser property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByUser
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dailyPrintUsageByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinter, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByPrinter))
            }
            m.SetDailyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByUser))
            }
            m.SetDailyPrintUsageByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinter, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByPrinter))
            }
            m.SetMonthlyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByUser))
            }
            m.SetMonthlyPrintUsageByUser(res)
        }
        return nil
    }
    return res
}
func (m *ReportRoot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDailyPrintUsageByPrinter sets the dailyPrintUsageByPrinter property value. 
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    m.dailyPrintUsageByPrinter = value
}
// SetDailyPrintUsageByUser sets the dailyPrintUsageByUser property value. 
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUser)() {
    m.dailyPrintUsageByUser = value
}
// SetMonthlyPrintUsageByPrinter sets the monthlyPrintUsageByPrinter property value. 
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    m.monthlyPrintUsageByPrinter = value
}
// SetMonthlyPrintUsageByUser sets the monthlyPrintUsageByUser property value. 
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUser)() {
    m.monthlyPrintUsageByUser = value
}
