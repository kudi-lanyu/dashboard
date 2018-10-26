// dashboard module
import chromeModule from "../chrome/module";
import componentsModule from "../common/components/module";
import filtersModule from "../common/filters/module";
// atlas
import {atlasDataCardComponent} from "./list/data_component";
import {atlasDataCardListComponent} from "./list/datalist_component";
import stateConfig from './stateconfig';

export default angular
  .module(
    'kubernetesDashboard.atlasdisplaydata',
    [
      'ngMaterial',
      'ngResource',
      'ui.router',
      chromeModule.name,
      componentsModule.name,
      filtersModule.name,
    ])
  .config(stateConfig)
  .component('kdAtlasDataCard', atlasDataCardComponent)
  .component('kdAtlasDataCardList', atlasDataCardListComponent)
  .factory('kdAtlasDataListResource', atlasDataListResource)

export function atlasDataListResource($resource) {
  console.log("atlasDataListResource");
  return $resource('/api/v1/atlas/data')
}
