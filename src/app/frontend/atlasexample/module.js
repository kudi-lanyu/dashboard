import stateConfig from './stateconfig';
import {AtlasExampleService} from './service';
import csrfTokenModule from '../common/csrftoken/module';
import chromeModule from '../chrome/module';


export default angular
  .module(
    'kubernetesDashboard.example',
    [
      'ngMaterial',
      'ui.router',
      'ngResource',
      chromeModule.name,
      csrfTokenModule.name,
    ])
  .config(stateConfig)
  .service('kdAtlasExampleService', AtlasExampleService)
