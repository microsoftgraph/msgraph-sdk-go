package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new education and sets the default values.
func NewEducation()(*Education) {
    m := &Education{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Education) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the classes property value. 
func (m *Education) GetClasses()([]EducationClass) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
// Gets the me property value. 
func (m *Education) GetMe()(*EducationUser) {
    if m == nil {
        return nil
    } else {
        return m.me
    }
}
// Gets the schools property value. 
func (m *Education) GetSchools()([]EducationSchool) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// Gets the users property value. 
func (m *Education) GetUsers()([]EducationUser) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// The deserialization information for the current model
func (m *Education) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationClass() })
        if err != nil {
            return err
        }
        res := make([]EducationClass, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationClass))
        }
        m.SetClasses(res)
        return nil
    }
    res["me"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationUser() })
        if err != nil {
            return err
        }
        m.SetMe(val.(*EducationUser))
        return nil
    }
    res["schools"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSchool() })
        if err != nil {
            return err
        }
        res := make([]EducationSchool, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationSchool))
        }
        m.SetSchools(res)
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationUser() })
        if err != nil {
            return err
        }
        res := make([]EducationUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationUser))
        }
        m.SetUsers(res)
        return nil
    }
    return res
}
func (m *Education) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Education) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the classes property value. 
// Parameters:
//  - value : Value to set for the classes property.
func (m *Education) SetClasses(value []EducationClass)() {
    m.classes = value
}
// Sets the me property value. 
// Parameters:
//  - value : Value to set for the me property.
func (m *Education) SetMe(value *EducationUser)() {
    m.me = value
}
// Sets the schools property value. 
// Parameters:
//  - value : Value to set for the schools property.
func (m *Education) SetSchools(value []EducationSchool)() {
    m.schools = value
}
// Sets the users property value. 
// Parameters:
//  - value : Value to set for the users property.
func (m *Education) SetUsers(value []EducationUser)() {
    m.users = value
}
