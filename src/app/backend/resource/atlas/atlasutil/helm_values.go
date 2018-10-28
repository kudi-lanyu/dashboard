package atlasutil

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/strvals"
	"log"
	"os"
	"strings"
)

// valueFiles present path set of values.yaml files
type valueFiles []string

func (v *valueFiles) String() string {
	return fmt.Sprint(*v)
}

func (v *valueFiles) Type() string {
	return "valueFiles"
}

func (v *valueFiles) Set(value string) error {
	for _, filePath := range strings.Split(value, ",") {
		*v = append(*v, filePath)
	}
	return nil
}

// Merges source and destination map, preferring values from the source map
func mergeValues(dest map[string]interface{}, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		// If the key doesn't exist already, then just set the key to that value
		if _, exists := dest[k]; !exists {
			dest[k] = v
			continue
		}
		nextMap, ok := v.(map[string]interface{})
		// If it isn't another map, overwrite the value
		if !ok {
			dest[k] = v
			continue
		}
		// Edge case: If the key exists in the destination, but isn't a map
		destMap, isMap := dest[k].(map[string]interface{})
		// If the source map has a map for this key, prefer it
		if !isMap {
			dest[k] = v
			continue
		}
		// If we got to this point, it is a map in both, so merge them
		dest[k] = mergeValues(destMap, nextMap)
	}
	return dest
}

// istable is a special-purpose function to see if the present thing matches the definition of a YAML table.
func istable(v interface{}) bool {
	_, ok := v.(map[string]interface{})
	return ok
}

// coalesceTables merges a source map into a destination map.
//
// dest is considered authoritative.
func coalesceTables(dst, src map[string]interface{}) map[string]interface{} {
	// Because dest has higher precedence than src, dest values override src
	// values.
	for key, val := range src {
		if istable(val) {
			if innerdst, ok := dst[key]; !ok {
				dst[key] = val
			} else if istable(innerdst) {
				coalesceTables(innerdst.(map[string]interface{}), val.(map[string]interface{}))
			} else {
				log.Printf("warning: cannot overwrite table with non table for %s (%v)", key, val)
			}
			continue
		} else if dv, ok := dst[key]; ok && istable(dv) {
			log.Printf("warning: destination for %s is a table. Ignoring non-table value %v", key, val)
			continue
		} else if !ok { // <- ok is still in scope from preceding conditional.
			dst[key] = val
			continue
		}
	}
	return dst
}

func MergeChartValues(values1, values2 chartutil.Values) (chartutil.Values, error) {
	v1 := values1.AsMap()
	v2 := values2.AsMap()

	for key, val := range v1 {
		if value, ok := v2[key]; ok {
			if value == nil {
				// When the YAML value is null, we remove the value's key.
				// This allows Helm's various sources of values (value files or --set) to
				// remove incompatible keys from any previous chart, file, or set values.
				delete(v2, key)
			} else if dest, ok := value.(map[string]interface{}); ok {
				// if v[key] is a table, merge nv's val table into v[key].
				src, ok := val.(map[string]interface{})
				if !ok {
					log.Printf("warning: skipped value for %s: Not a table.", key)
					continue
				}
				// Because v2 has higher precedence than v1, dest values override src
				// values.
				coalesceTables(dest, src)
			}
		} else {
			// If the key is not in v2, copy it from v1.
			v2[key] = val
		}
	}
	return v2, nil
}

// vals merges values from files specified via -f/--values and
// directly via --set or --set-string or --set-file, marshaling them to YAML
func vals(valueFiles valueFiles, values []string, stringValues []string, fileValues []string) ([]byte, error) {
	base := map[string]interface{}{}

	for _, filePath := range valueFiles {
		currentMap := map[string]interface{}{}

		var bytes []byte
		var err error
		if strings.TrimSpace(filePath) == "-" {
			bytes, err = ioutil.ReadAll(os.Stdin)
		} else {
			bytes, err = readFile(filePath)
		}

		if err != nil {
			return []byte{}, err
		}

		if err := yaml.Unmarshal(bytes, &currentMap); err != nil {
			return []byte{}, fmt.Errorf("failed to parse %s: %s", filePath, err)
		}
		// Merge with the previous map
		base = mergeValues(base, currentMap)
	}

	for _, value := range values {
		if err := strvals.ParseInto(value, base); err != nil {
			return []byte{}, fmt.Errorf("failed parsing --set data: %s", err)
		}
	}

	for _, value := range stringValues {
		if err := strvals.ParseIntoString(value, base); err != nil {
			return []byte{}, fmt.Errorf("failed parsing --set-string data: %s", err)
		}
	}

	for _, value := range fileValues {
		reader := func(rs []rune) (interface{}, error) {
			bytes, err := readFile(string(rs))
			return string(bytes), err
		}
		if err := strvals.ParseIntoFile(value, base, reader); err != nil {
			return []byte{}, fmt.Errorf("failed parsing --set-file data: %s", err)
		}
	}

	return yaml.Marshal(base)
}

//readFile load a file from the local directory.
func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func PrintValues(values chartutil.Values) string {
	data, err := json.Marshal(values.AsMap())
	if err != nil {
		log.Println("printVaules: json marshal failed.", err)
	}
	log.Println(string(data))

	return string(data)
}

// ReleaseOptions represents the additional release options needed
// for the composition of the final values struct
type ReleaseOptions struct {
	Name string
	//Time      *timestamp.Timestamp
	Namespace string
	Revision  int
}

// ToRenderValues composes the struct from the data coming from the Releases, Charts and Values files
//
// WARNING: This function is deprecated for Helm > 2.1.99 Use ToRenderValuesCaps() instead. It will
// remain in the codebase to stay SemVer compliant.
//
// In Helm 3.0, this will be changed to accept Capabilities as a fourth parameter.
func ToRenderValues(chrt *chart.Chart, chrtVals *chart.Config, options ReleaseOptions) (chartutil.Values, error) {
	caps := &chartutil.Capabilities{APIVersions: chartutil.DefaultVersionSet}
	return ToRenderValuesCaps(chrt, chrtVals, options, caps)
}

// ToRenderValuesCaps composes the struct from the data coming from the Releases, Charts and Values files
//
// This takes both ReleaseOptions and Capabilities to merge into the render values.
func ToRenderValuesCaps(chrt *chart.Chart, chrtVals *chart.Config, options ReleaseOptions, caps *chartutil.Capabilities) (chartutil.Values, error) {

	top := map[string]interface{}{
		"Release": map[string]interface{}{
			"Name":      options.Name,
			"Namespace": options.Namespace,
			"Revision":  options.Revision,
		},
		"Chart":        chrt.Metadata,
		"Files":        chartutil.NewFiles(chrt.Files),
		"Capabilities": caps,
	}

	vals, err := chartutil.CoalesceValues(chrt, chrtVals)
	if err != nil {
		return top, err
	}

	top["Values"] = vals
	return top, nil
}

func ToHelmRenderValues(chart *chart.Chart, config *chart.Config) (chartutil.Values, error) {

	c := chart
	v := config
	// construct release options
	o := ReleaseOptions{
		Name: "temporary",
		//Time:timeconv.Now(),
		Namespace: "al Basrah",
		//IsInstall:true,
		Revision: 5,
	}

	caps := &chartutil.Capabilities{
		APIVersions: chartutil.DefaultVersionSet,
	}

	res, err := ToRenderValuesCaps(c, v, o, caps)
	if err != nil {
		log.Panicln("construct helm values failed.", err)
	}

	return res, err
}

func TransferMapToChartValues(mapSpec interface{}) (chartutil.Values, error) {
	// transfer the map to chartutil.Values.
	data, err := yaml.Marshal(mapSpec)
	log.Println("datatoyaml: [[[[[[[[[[[")
	log.Println(string(data))

	//data1,err := json.Marshal(mapSpec)
	//log.Println("datatojson:]]]]]]]")
	//log.Println(string(data1))

	if err != nil {
		log.Println(err)
		return nil, err
	}
	helmValues, err := chartutil.ReadValues(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return helmValues, nil
}
