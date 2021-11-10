package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintUsage struct {
    Entity
    // 
    completedBlackAndWhiteJobCount *int64;
    // 
    completedColorJobCount *int64;
    // 
    incompleteJobCount *int64;
    // 
    usageDate *string;
}
// Instantiates a new printUsage and sets the default values.
func NewPrintUsage()(*PrintUsage) {
    m := &PrintUsage{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the completedBlackAndWhiteJobCount property value. 
func (m *PrintUsage) GetCompletedBlackAndWhiteJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedBlackAndWhiteJobCount
    }
}
// Gets the completedColorJobCount property value. 
func (m *PrintUsage) GetCompletedColorJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedColorJobCount
    }
}
// Gets the incompleteJobCount property value. 
func (m *PrintUsage) GetIncompleteJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.incompleteJobCount
    }
}
// Gets the usageDate property value. 
func (m *PrintUsage) GetUsageDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usageDate
    }
}
// The deserialization information for the current model
func (m *PrintUsage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedBlackAndWhiteJobCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedBlackAndWhiteJobCount(val)
        }
        return nil
    }
    res["completedColorJobCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedColorJobCount(val)
        }
        return nil
    }
    res["incompleteJobCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteJobCount(val)
        }
        return nil
    }
    res["usageDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageDate(val)
        }
        return nil
    }
    return res
}
func (m *PrintUsage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintUsage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("completedBlackAndWhiteJobCount", m.GetCompletedBlackAndWhiteJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("completedColorJobCount", m.GetCompletedColorJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("incompleteJobCount", m.GetIncompleteJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usageDate", m.GetUsageDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the completedBlackAndWhiteJobCount property value. 
// Parameters:
//  - value : Value to set for the completedBlackAndWhiteJobCount property.
func (m *PrintUsage) SetCompletedBlackAndWhiteJobCount(value *int64)() {
    m.completedBlackAndWhiteJobCount = value
}
// Sets the completedColorJobCount property value. 
// Parameters:
//  - value : Value to set for the completedColorJobCount property.
func (m *PrintUsage) SetCompletedColorJobCount(value *int64)() {
    m.completedColorJobCount = value
}
// Sets the incompleteJobCount property value. 
// Parameters:
//  - value : Value to set for the incompleteJobCount property.
func (m *PrintUsage) SetIncompleteJobCount(value *int64)() {
    m.incompleteJobCount = value
}
// Sets the usageDate property value. 
// Parameters:
//  - value : Value to set for the usageDate property.
func (m *PrintUsage) SetUsageDate(value *string)() {
    m.usageDate = value
}
