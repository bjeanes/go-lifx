package type_registry

import "reflect"

type idToTypeMap map[uint16]reflect.Type
type typeToIdMap map[reflect.Type]uint16

type TypeRegistry struct {
	idToType idToTypeMap
	typeToId typeToIdMap
}

func New() TypeRegistry {
	return TypeRegistry{
		make(idToTypeMap),
		make(typeToIdMap),
	}
}

func (reg TypeRegistry) New(id uint16) interface{} {
	if t := reg.idToType[id]; t != nil {
		return reflect.New(t).Interface()
	}

	return nil
}

func (reg TypeRegistry) Register(id uint16, v interface{}) {
	t := reflect.TypeOf(v).Elem()
	reg.idToType[id] = t
	reg.typeToId[t] = id
}
