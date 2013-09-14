package Sleep

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"reflect"
	"strings"
)

type Query struct {
	query         interface{}
	selection     interface{}
	skip          int
	limit         int
	sort          []string
	populate      map[string]*Query
	path          string
	z             *Sleep
	populated     map[string]interface{}
	isPopOp       bool
	parentStruct  interface{}
	populateField interface{}
	isSlice       bool
	popModel      string
}

func (q *Query) populateExec(parentStruct interface{}) error {
	for key, val := range q.populate {
		val.parentStruct = parentStruct
		val.findPopulatePath(key)
		model, ok := q.z.models[val.popModel]
		if !ok {
			panic("Unable to find `" + val.popModel + "` model. Was it registered?")
		}

		var schema interface{}
		if val.isSlice {
			ids := val.populateField.([]bson.ObjectId)
			if len(ids) == 0 {
				return nil
			}

			schemaType := reflect.PtrTo(reflect.TypeOf(model.schema))
			slicedType := reflect.SliceOf(schemaType)
			schema = reflect.New(slicedType).Interface()
			val.query = M{"_id": M{"$in": ids}}

		} else {
			schema = &model.schema
			id := val.populateField.(bson.ObjectId)
			va.query = M{"_id": id}
		}

		err := val.Exec(schema)
		if err != nil {
			panic(err)
		}
		parentModel := reflect.ValueOf(val.parentStruct).Elem().FieldByName("Model").Interface().(Model)
		parentModel.populated[key] = schema
	}
	return nil
}

func (q *Query) Populate(fields ...string) *Query {
	for _, elem := range fields {
		q.populate[elem] = &Query{isPopOp: true,
			populate:  make(map[string]*Query),
			populated: make(map[string]interface{}), z: q.z}
	}
	return q
}

func (q *Query) PopulateQuery(field string, query *Query) *Query {
	query.isPopOp = true
	query.populate = make(map[string]*Query)
	query.populated = make(map[string]interface{})
	query.z = q.z
	q.populate[field] = query
	return q
}

func (q *Query) findPopulatePath(path string) {
	parts := strings.Split(path, ".")
	resultVal := reflect.ValueOf(q.parentStruct).Elem()

	var refVal reflect.Value
	partsLen := len(parts)
	for i := 0; i < partsLen; i++ {
		elem := parts[i]
		if i == 0 {
			refVal = resultVal.FieldByName(elem)
			structTag, _ := resultVal.Type().FieldByName(elem)
			q.popModel = structTag.Tag.Get(q.z.modelTag)
		} else if i == partsLen-1 {
			structTag, _ := refVal.Type().FieldByName(elem)
			q.popModel = structTag.Tag.Get(q.z.modelTag)
			refVal = refVal.FieldByName(elem)
		}

		if !refVal.IsValid() {
			panic("field `" + elem + "` not found in populate path `" + path + "`")
		}
	}

	if refVal.Kind() == reflect.Slice {
		q.isSlice = true
	}
	q.populateField = refVal.Interface()
}

func (query *Query) Exec(result interface{}) error {
	if reflect.TypeOf(result).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Expecting a pointer type but recieved %v. If you are passing in a slice, make sure to pass a pointer to it.", reflect.TypeOf(result)))
	}
	typ := reflect.TypeOf(result).Elem()
	var structName string
	isSlice := false
	if typ.Kind() == reflect.Slice {
		structName = typ.Elem().Elem().Name()
		isSlice = true
	} else {
		structName = typ.Name()
	}

	model := query.z.models[structName]
	q := model.C.Find(query.query)

	if query.limit != 0 {
		q = q.Limit(query.limit)
	}

	if query.skip != 0 {
		q = q.Skip(query.skip)
	}

	sortLen := len(query.sort)
	if sortLen != 0 {
		for i := 0; i < sortLen; i++ {
			q = q.Sort(query.sort[i])
		}
	}

	if query.selection != nil {
		q = q.Select(query.selection)
	}

	var err error

	if isSlice == true {
		err = q.All(result)
		if err != nil {
			return err
		}

		val := reflect.ValueOf(result).Elem()
		elemCount := val.Len()
		for i := 0; i < elemCount; i++ {
			modelCpy := query.z.models[structName]
			sliceElem := val.Index(i)
			modelCpy.doc = sliceElem.Interface()
			modelElem := sliceElem.Elem().FieldByName("Model")
			modelElem.Set(reflect.ValueOf(model))
		}
		return err
	}

	err = q.One(result)
	if err != nil {
		return err
	}

	model.doc = result
	val := reflect.ValueOf(result).Elem()
	modelVal := val.FieldByName("Model")
	modelVal.Set(reflect.ValueOf(model))

	query.populateExec(result)

	return err
}

func (q *Query) Select(selection interface{}) *Query {
	q.selection = selection
	return q
}

func (q *Query) Skip(skip int) *Query {
	q.skip = skip
	return q
}

func (q *Query) Limit(lim int) *Query {
	q.limit = lim
	return q
}

func (q *Query) Sort(fields ...string) *Query {
	q.sort = fields
	return q
}