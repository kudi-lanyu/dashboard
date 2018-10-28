package test

import (
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasutil"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/engine"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"log"
	"testing"
)

const chart_dir_mpi = "../../helmcharts/mpijob"
const chart_value_test_file_mpi = "../../helmcharts/mpijob/values.yaml"

func testHelmCharts() {
	// charts
	chart_from_dir, err := chartutil.Load(chart_dir_mpi)
	if err != nil {
		log.Panicln("load chart dir is failed.", err)
	}

	// values
	chart_default_values, err := chartutil.ReadValuesFile(chart_value_test_file_mpi)
	if err != nil {
		log.Println("value_file is not exist.", err)
	}
	// merge default values and other additional values
	// receive parameters from frontend and construct new info struct to generate chartutil.Values
	add_values, err := HandleFrontValues()
	if err != nil {
		log.Println("add values fail.", err)
	}
	//log.Println(atlashelm.PrintValues(add_values))

	// merge default and additional values
	merge_values, err := atlasutil.MergeChartValues(chart_default_values, add_values)
	if err != nil {
		log.Println("merge charts values fail.", err)
	}
	//log.Println(atlashelm.PrintValues(helm_values))

	helm_values_yaml, err := merge_values.YAML()
	if err != nil {
		log.Println("merge charts values fail.", err)
	}

	// render
	helm_engine := engine.New()
	// because of engine need user .Values , so must merge values into chart again with top key
	tmpConfig := &chart.Config{Raw: helm_values_yaml}
	helm_values, err := atlasutil.ToHelmRenderValues(chart_from_dir, tmpConfig)
	if err != nil {
		log.Println("construct helm values failed.", err)
	}

	render_out, err := helm_engine.Render(chart_from_dir, helm_values)
	if err != nil {
		log.Println("render charts failed.", err)
	}

	//traversal Render output
	log.Println("+++++++++++++++++++++++++++++++++++++++++++")
	for k, v := range render_out {
		log.Println("out", k, v)
	}
}

func TestHelm_parse(t *testing.T) {
	testHelmCharts()
	return
}
