/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package parsers

import (
	"fmt"
	"reflect"
	"encoding/json"
	"github.com/apache/incubator-openwhisk-wskdeploy/utils"
)


// TODO(): Support other valid Package Manifest types
// TODO(): i.e., timestamp, version, string256, string64, string16
// TODO(): Support JSON schema validation for type: json
// TODO(): Support OpenAPI schema validation

var validParameterNameMap = map[string]string{
	"string":  "string",
	"int":     "integer",
	"float":   "float",
	"bool":    "boolean",
	"int8":    "integer",
	"int16":   "integer",
	"int32":   "integer",
	"int64":   "integer",
	"float32": "float",
	"float64": "float",
	"json":    "json",
	"map":     "json",
}

var typeDefaultValueMap = map[string]interface{}{
	"string":  "",
	"integer": 0,
	"float":   0.0,
	"boolean": false,
	"json":    make(map[string]interface{}),
	// TODO() Support these types + their validation
	// timestamp
	// null
	// version
	// string256
	// string64
	// string16
	// scalar-unit
	// schema
	// object
}

func isValidParameterType(typeName string) bool {
	_, isValid := typeDefaultValueMap[typeName]
	return isValid
}

// TODO(): throw errors
func getTypeDefaultValue(typeName string) interface{} {

	if val, ok := typeDefaultValueMap[typeName]; ok {
		return val
	} else {
		// TODO() throw an error "type not found"
	}
	return nil
}

func ResolveParamTypeFromValue(name string, value interface{}, filePath string) (string, error) {
	// Note: 'string' is the default type if not specified and not resolvable.
	var paramType string = "string"
	var err error = nil

	if value != nil {
		actualType := reflect.TypeOf(value).Kind().String()

		// See if the actual type of the value is valid
		if normalizedTypeName, found := validParameterNameMap[actualType]; found {
			// use the full spec. name
			paramType = normalizedTypeName

		} else {
			// raise an error if parameter's value is not a known type
			// TODO() - move string to i18n
			msgs := []string{"Parameter [" + name + "] has a value that is not a known type. [" + actualType + "]"}
			err = utils.NewYAMLParserErr(filePath, nil, msgs)
		}
	}
	return paramType, err
}


/*
    resolveSingleLineParameter assures that a Parameter's Type is correctly identified and set from its Value.

    Additionally, this function:

    - detects if the parameter value contains the name of a valid OpenWhisk parameter types. if so, the
      - param.Type is set to detected OpenWhisk parameter type.
      - param.Value is set to the zero (default) value for that OpenWhisk parameter type.
 */
func resolveSingleLineParameter(paramName string, param *Parameter, filePath string) (interface{}, error) {
	var errorParser error

	if !param.multiline {
		// We need to identify parameter Type here for later validation
		param.Type, errorParser = ResolveParamTypeFromValue(paramName, param.Value, filePath)

		// In single-line format, the param's <value> can be a "Type name" and NOT an actual value.
		// if this is the case, we must detect it and set the value to the default for that type name.
		if param.Value != nil && param.Type == "string" {
			// The value is a <string>; now we must test if is the name of a known Type
			if isValidParameterType(param.Value.(string)) {
				// If the value is indeed the name of a Type, we must change BOTH its
				// Type to be that type and its value to that Type's default value
				param.Type = param.Value.(string)
				param.Value = getTypeDefaultValue(param.Type)
				//fmt.Printf("EXIT: Parameter [%s] type=[%v] value=[%v]\n", paramName, param.Type, param.Value)
			}
		}

	} else {
		msgs := []string{"Parameter [" + paramName + "] is not single-line format."}
		return param.Value, utils.NewYAMLParserErr(filePath, nil, msgs)
	}

	return param.Value, errorParser
}

/*
    resolveMultiLineParameter assures that the values for Parameter Type and Value are properly set and are valid.

    Additionally, this function:
    - uses param.Default as param.Value if param.Value is not provided
    - uses the actual param.Value data type for param.type if param.Type is not provided

 */
func resolveMultiLineParameter(paramName string, param *Parameter, filePath string) (interface{}, error) {
	var errorParser error

	if param.multiline {
		var valueType string

		// if we do not have a value, but have a default, use it for the value
		if param.Value == nil && param.Default != nil {
			param.Value = param.Default
		}

		// Note: if either the value or default is in conflict with the type then this is an error
		valueType, errorParser = ResolveParamTypeFromValue(paramName, param.Value, filePath)

		// if we have a declared parameter Type, assure that it is a known value
		if param.Type != "" {
			if !isValidParameterType(param.Type) {
				// TODO() - move string to i18n
				msgs := []string{"Parameter [" + paramName + "] has an invalid Type. [" + param.Type + "]"}
				return param.Value, utils.NewYAMLParserErr(filePath, nil, msgs)
			}
		} else {
			// if we do not have a value for the Parameter Type, use the Parameter Value's Type
			param.Type = valueType
		}

		// TODO{} if the declared and actual parameter type conflict, generate TypeMismatch error
		//if param.Type != valueType{
		//	errorParser = utils.NewParameterTypeMismatchError("", param.Type, valueType )
		//}
	} else {
		msgs := []string{"Parameter [" + paramName + "] is not multiline format."}
		return param.Value, utils.NewYAMLParserErr(filePath, nil, msgs)
	}


	return param.Value, errorParser
}


/*
    resolveJSONParameter assure JSON data is converted to a map[string]{interface*} type.

    This function handles the forms JSON data appears in:
    1) a string containing JSON, which needs to be parsed into map[string]interface{}
    2) is a map of JSON (but not a map[string]interface{}
 */
func resolveJSONParameter(paramName string, param *Parameter, value interface{}, filePath string) (interface{}, error) {
	var errorParser error

	if param.Type == "json" {
		// Case 1: if user set parameter type to 'json' and the value's type is a 'string'
		if str, ok := value.(string); ok {
			var parsed interface{}
			errParser := json.Unmarshal([]byte(str), &parsed)
			if errParser == nil {
				//fmt.Printf("EXIT: Parameter [%s] type=[%v] value=[%v]\n", paramName, param.Type, parsed)
				return parsed, errParser
			}
		}

		// Case 2: value contains a map of JSON
		// We must make sure the map type is map[string]interface{}; otherwise we cannot
		// marshall it later on to serialize in the body of an HTTP request.
		if( param.Value != nil && reflect.TypeOf(param.Value).Kind() == reflect.Map ) {
			if _, ok := param.Value.(map[interface{}]interface{}); ok {
				var temp map[string]interface{} =
					utils.ConvertInterfaceMap(param.Value.(map[interface{}]interface{}))
				//fmt.Printf("EXIT: Parameter [%s] type=[%v] value=[%v]\n", paramName, param.Type, temp)
				return temp, errorParser
			}
		} // else TODO{}
	} else {
		msgs := []string{"Parameter [" + paramName + "] is not JSON format."}
		return param.Value, utils.NewYAMLParserErr(filePath, nil, msgs)
	}

	return param.Value, errorParser
}

/*
    ResolveParameter assures that the Parameter structure's values are correctly filled out for
    further processing.  This includes special processing for

    - single-line format parameters
      - deriving missing param.Type from param.Value
      - resolving case where param.Value contains a valid Parameter type name
    - multi-line format parameters:
      - assures that param.Value is set while taking into account param.Default
      - validating param.Type

    Note: parameter values may set later (overridden) by an (optional) Deployment file

 */
func ResolveParameter(paramName string, param *Parameter, filePath string) (interface{}, error) {

	var errorParser error
	// default parameter value to empty string
	var value interface{} = ""

	// Trace Parameter struct before any resolution
	//dumpParameter(paramName, param, "BEFORE")

	// Parameters can be single OR multi-line declarations which must be processed/validated differently
	if !param.multiline {

		// This function will assure that param.Value and param.Type are correctly set
		value, errorParser = resolveSingleLineParameter(paramName, param, filePath)

	} else {

		value, errorParser = resolveMultiLineParameter(paramName, param, filePath)
	}

	// String value pre-processing (interpolation)
	// See if we have any Environment Variable replacement within the parameter's value

	// Make sure the parameter's value is a valid, non-empty string
	if ( param.Value != nil && param.Type == "string") {
		// perform $ notation replacement on string if any exist
		value = utils.GetEnvVar(param.Value)
	}

	// JSON - Handle both cases, where value 1) is a string containing JSON, 2) is a map of JSON
	if param.Type == "json" {
		value, errorParser = resolveJSONParameter(paramName, param, value, filePath)
	}

	// Default value to zero value for the Type
	// Do NOT error/terminate as Value may be provided later by a Deployment file.
	if value == nil {
		value = getTypeDefaultValue(param.Type)
		// @TODO(): Need warning message here to warn of default usage, support for warnings (non-fatal)
		//msgs := []string{"Parameter [" + paramName + "] is not multiline format."}
		//return param.Value, utils.NewParserErr(filePath, nil, msgs)
	}

	// Trace Parameter struct after resolution
	//dumpParameter(paramName, param, "AFTER")
	//fmt.Printf("EXIT: Parameter [%s] type=[%v] value=[%v]\n", paramName, param.Type, value)
	return value, errorParser
}

// Provides debug/trace support for Parameter type
func dumpParameter(paramName string, param *Parameter, separator string) {

	fmt.Printf("%s:\n", separator)
	fmt.Printf("\t%s: (%T)\n", paramName, param)
	if param != nil {
		fmt.Printf("\t\tParameter.Description: [%s]\n", param.Description)
		fmt.Printf("\t\tParameter.Type: [%s]\n", param.Type)
		fmt.Printf("\t\t--> Actual Type: [%T]\n", param.Value)
		fmt.Printf("\t\tParameter.Value: [%v]\n", param.Value)
		fmt.Printf("\t\tParameter.Default: [%v]\n", param.Default)
	}
}
