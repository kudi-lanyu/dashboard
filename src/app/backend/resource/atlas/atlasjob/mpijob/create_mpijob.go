package mpijob

import (
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasutil"
	"github.com/kubernetes/dashboard/src/app/backend/resource/deployment"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	//"encoding/json"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/engine"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	MPI_CHART_DIR         = "/dashboard/src/app/backend/resource/atlas/helmcharts/mpijob"
	MPI_CHART_VALUES_YAML = "/dashboard/src/app/backend/resource/atlas/helmcharts/mpijob/values.yaml"
)

func DeployAtlasMpiJob(cfg *rest.Config, spec *MPIJobArgs, namespace string) (bool, error) {
	log.Println("deploy atlas mpi job.")
	filecontentwithfilename, err := RenderMpiJobSpecToYaml(*spec)

	filecontent := ""
	filename := ""
	for k,v := range filecontentwithfilename {
	  filename = k
	  filecontent = v
  }

	// two way to implement the function by new client or use deploymentfile function from deploy
	deploymentSpec := &deployment.AppDeploymentFromFileSpec{
		Name:      filename,
		Namespace: namespace,
		Content:   filecontent,
		Validate:  false,
	}

	log.Println("deploymentSepc:", deploymentSpec)

	isDeployed, err := deployment.DeployAppFromFile(cfg, deploymentSpec)

	return isDeployed, err
}

func GetCurrentPathAndPrint() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Println("err:", err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		log.Println("err:", err)
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		err := `error: Can't find "/" or "\".`
		log.Println("err:", err)
	}
	if err == nil {
		log.Println("currentPath:", string(path[0:i+1]))
	}
}

func RenderMpiJobSpecToYaml(mpiJobSpec MPIJobArgs) (map[string]string, error) {
	// charts
	mpi_chart, err := chartutil.Load(MPI_CHART_DIR)
	if err != nil {
		log.Println("load chart dir is failed.", err)
	}

	//GetCurrentPathAndPrint()

	// values
	chart_default_values, err := chartutil.ReadValuesFile(MPI_CHART_VALUES_YAML)
	if err != nil {
		log.Println("value_file is not exist.", err)
	}

	// merge charts default values from values.yaml and other additional values receive from frontend
	mpiJob_values, err := atlasutil.TransferMapToChartValues(mpiJobSpec)
	if err != nil {
		log.Println("transfer spec to chart values failed.", err)
	}

	// merge default and additional values
	merge_values, err := atlasutil.MergeChartValues(chart_default_values, mpiJob_values)
	if err != nil {
		log.Println("merge charts values fail.", err)
	}

	helm_values_yaml, err := merge_values.YAML()
	if err != nil {
		log.Println("merge charts values fail.", err)
	}

	// render
	helm_engine := engine.New()

	// because of engine should use .Values and .Release variable, so must merge values into chart again with top key
	tmpConfig := &chart.Config{Raw: helm_values_yaml}
	helm_values, err := atlasutil.ToHelmRenderValues(mpi_chart, tmpConfig)
	if err != nil {
		log.Println("construct helm values failed.", err)
	}

	render_out, err := helm_engine.Render(mpi_chart, helm_values)
	if err != nil {
		log.Println("render charts failed.", err)
	}
	
  //
	////traversal Render output
	//log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
  //for k, v := range render_out {
  //  log.Println("out", k, v) // file name and content
  //  log.Println("value",v) // file content
  //}

	//result, err := yaml.Marshal(render_out)
	//if err != nil {
	//	log.Println("yaml marshal failed.", err)
	//}

  //result, err := json.Marshal(render_out)
  //if err != nil {
  //  log.Println("json marshal failed.", err)
  //}
  //log.Println(string(result))

	return render_out, err
}
