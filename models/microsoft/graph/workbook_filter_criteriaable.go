package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookFilterCriteriaable 
type WorkbookFilterCriteriaable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetColor()(*string)
    GetCriterion1()(*string)
    GetCriterion2()(*string)
    GetDynamicCriteria()(*string)
    GetFilterOn()(*string)
    GetIcon()(WorkbookIconable)
    GetOperator()(*string)
    GetValues()(Jsonable)
    SetColor(value *string)()
    SetCriterion1(value *string)()
    SetCriterion2(value *string)()
    SetDynamicCriteria(value *string)()
    SetFilterOn(value *string)()
    SetIcon(value WorkbookIconable)()
    SetOperator(value *string)()
    SetValues(value Jsonable)()
}
