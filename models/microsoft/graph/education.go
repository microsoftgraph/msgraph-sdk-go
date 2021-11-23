package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Education 
type Education struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    classes []EducationClass;
    // 
    me *EducationUser;
    // 
    schools []EducationSchool;
    // 
    users []EducationUser;
}
// NewEducation instantiates a new education and sets the default values.
func NewEducation()(*Education) {
    m := &Education{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Education) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClasses gets the classes property value. 
func (m *Education) GetClasses()([]EducationClass) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
// GetMe gets the me property value. 
func (m *Education) GetMe()(*EducationUser) {
    if m == nil {
        return nil
    } else {
        return m.me
    }
}
// GetSchools gets the schools property value. 
func (m *Education) GetSchools()([]EducationSchool) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// GetUsers gets the users property value. 
func (m *Education) GetUsers()([]EducationUser) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Education) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationClass() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClass, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationClass))
            }
            m.SetClasses(res)
        }
        return nil
    }
    res["me"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationUser() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMe(val.(*EducationUser))
        }
        return nil
    }
    res["schools"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSchool() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSchool, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationSchool))
            }
            m.SetSchools(res)
        }
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationUser))
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
func (m *Education) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Education) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClasses()))
        for i, v := range m.GetClasses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("classes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("me", m.GetMe())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchools()))
        for i, v := range m.GetSchools() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("schools", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("users", cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Education) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClasses sets the classes property value. 
func (m *Education) SetClasses(value []EducationClass)() {
    m.classes = value
}
// SetMe sets the me property value. 
func (m *Education) SetMe(value *EducationUser)() {
    m.me = value
}
// SetSchools sets the schools property value. 
func (m *Education) SetSchools(value []EducationSchool)() {
    m.schools = value
}
// SetUsers sets the users property value. 
func (m *Education) SetUsers(value []EducationUser)() {
    m.users = value
}
