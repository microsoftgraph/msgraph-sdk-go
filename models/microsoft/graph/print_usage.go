package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintUsage 
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
// NewPrintUsage instantiates a new printUsage and sets the default values.
func NewPrintUsage()(*PrintUsage) {
    m := &PrintUsage{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompletedBlackAndWhiteJobCount gets the completedBlackAndWhiteJobCount property value. 
func (m *PrintUsage) GetCompletedBlackAndWhiteJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedBlackAndWhiteJobCount
    }
}
// GetCompletedColorJobCount gets the completedColorJobCount property value. 
func (m *PrintUsage) GetCompletedColorJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedColorJobCount
    }
}
// GetIncompleteJobCount gets the incompleteJobCount property value. 
func (m *PrintUsage) GetIncompleteJobCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.incompleteJobCount
    }
}
// GetUsageDate gets the usageDate property value. 
func (m *PrintUsage) GetUsageDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usageDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetCompletedBlackAndWhiteJobCount sets the completedBlackAndWhiteJobCount property value. 
func (m *PrintUsage) SetCompletedBlackAndWhiteJobCount(value *int64)() {
    if m != nil {
        m.completedBlackAndWhiteJobCount = value
    }
}
// SetCompletedColorJobCount sets the completedColorJobCount property value. 
func (m *PrintUsage) SetCompletedColorJobCount(value *int64)() {
    if m != nil {
        m.completedColorJobCount = value
    }
}
// SetIncompleteJobCount sets the incompleteJobCount property value. 
func (m *PrintUsage) SetIncompleteJobCount(value *int64)() {
    if m != nil {
        m.incompleteJobCount = value
    }
}
// SetUsageDate sets the usageDate property value. 
func (m *PrintUsage) SetUsageDate(value *string)() {
    if m != nil {
        m.usageDate = value
    }
}
