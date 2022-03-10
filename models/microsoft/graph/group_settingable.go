package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupSettingable 
type GroupSettingable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetTemplateId()(*string)
    GetValues()([]SettingValueable)
    SetDisplayName(value *string)()
    SetTemplateId(value *string)()
    SetValues(value []SettingValueable)()
}
