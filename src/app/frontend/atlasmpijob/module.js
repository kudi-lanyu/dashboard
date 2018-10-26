import stateConfig from './stateconfig';
import {AtlasMpiJobService} from './service';


// deploy common
import {deployFromFileComponent} from "./deployfromfile/component";
import {deployFromInputComponent} from "./deployfrominput/component";
import {deployFromSettingsComponent} from "./deployfromsettings/component";
// deploy inject module
import chromeModule from '../chrome/module';
import componentsModule from '../common/components/module';
import csrfTokenModule from '../common/csrftoken/module';
import historyModule from '../common/history/module';
import errorHandlingModule from '../common/errorhandling/module';
import validatorsModule from '../common/validators/module';
import helpSectionModule from '../deploy/deployfromsettings/helpsection/module';
import {mpiJobEnvsComponent} from './deployfromsettings/mpijobenvs/component';
import {mpiJobDataDirsComponent} from './deployfromsettings/mpijobdatadirs/component';
import {mpiJobDataSetComponet} from './deployfromsettings/mpijobdataset/component';

export default angular
  .module(
    'kubernetesDashboard.atlasmpijob',
    [
      'ngMaterial',
      'ui.router',
      'ngResource',
      chromeModule.name,
      errorHandlingModule.name,
      historyModule.name,
      componentsModule.name,
      csrfTokenModule.name,
      validatorsModule.name,
      helpSectionModule.name,
    ])
  .config(stateConfig)
  .component('kdAtlasEnvs', mpiJobEnvsComponent)
  .component('kdAtlasMpiJobDataDirs', mpiJobDataDirsComponent)
  .component('kdAtlasMpiJobDataSet', mpiJobDataSetComponet)
  .component('kdAtlasMpiJobDeployFromFile', deployFromFileComponent)
  .component('kdAtlasMpiJobDeployFromInput', deployFromInputComponent)
  .component('kdAtlasMpiJobDeployFromSettings', deployFromSettingsComponent)
  .service('kdAtlasMpiJobService', AtlasMpiJobService)
