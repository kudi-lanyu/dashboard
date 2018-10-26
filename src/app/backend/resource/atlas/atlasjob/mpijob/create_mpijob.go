package mpijob

import (
  "log"
  "k8s.io/helm/pkg/chartutil"
  "k8s.io/helm/pkg/proto/hapi/chart"
  "k8s.io/helm/pkg/engine"
  "github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasutil"
  "os"
  "path/filepath"
  "os/exec"
  "strings"
  "github.com/kubernetes/dashboard/src/app/backend/resource/deployment"
  "k8s.io/client-go/rest"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

const (
  MPI_CHART_DIR         = "/dashboard/src/app/backend/resource/atlas/helmcharts/mpijob"
  MPI_CHART_VALUES_YAML = "/dashboard/src/app/backend/resource/atlas/helmcharts/mpijob/values.yaml"
)

func DeployAtlasMpiJob(cfg *rest.Config, spec *MPIJobArgs) (bool, error) {
  log.Println("deploy atlas mpi job.")
  data, err := RenderMpiJobSpecToYaml(*spec)

  // tmpfile
  file, err := ioutil.TempFile("", "tmpfile")
  if err != nil {
    panic(err)
  }
  defer os.Remove(file.Name())

  if _, err := file.Write(data); err != nil {
    panic(err)
  }

  tmpfilename := file.Name()

  filecontent,err := ioutil.ReadFile(tmpfilename)

  log.Println("filecontent:_+_+++++++++++++_+______________+_+\n")
  log.Println(string(filecontent))

  // two way to implement the function by new client or use deploymentfile function from deploy
  deploymentSpec := &deployment.AppDeploymentFromFileSpec{
    Name:      "mpijob",
    Namespace: "specific",
    Content:   string(filecontent),
    Validate:  false,
  }



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
    log.Println("currentPath:", string(path[0: i+1]))
  }
}

func RenderMpiJobSpecToYaml(mpiJobSpec MPIJobArgs) ([]byte, error) {
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

  //traversal Render output
  //log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
  //for k, v := range render_out {
  //  log.Println("out", k, v)
  //}

  result, err := yaml.Marshal(render_out)
  if err != nil {
    log.Println("yaml marshal failed.", err)
  }

  return result, err
}
