// dashboard module
import chromeModule from "../chrome/module";
import componentsModule from "../common/components/module";
import filtersModule from "../common/filters/module";
// atlas
import {atlasNodeCardComponent} from "./list/node_component";
import {atlasNodeCardListComponent} from "./list/nodelist_component";
import stateConfig from './stateconfig';

export default angular
  .module(
    'kubernetesDashboard.atlasdisplaynode',
    [
      'ngMaterial',
      'ngResource',
      'ui.router',
      chromeModule.name,
      componentsModule.name,
      filtersModule.name,
    ])
  .config(stateConfig)
  .component('kdAtlasNodeCard', atlasNodeCardComponent)
  .component('kdAtlasNodeCardList', atlasNodeCardListComponent)
  .factory('kdAtlasNodeListResource', atlasNodeListResource)
  .factory('kdAtlasNodeInfoResource', atlasNodeInfoResource)

export function atlasNodeListResource($resource) {
  console.log("atlasNodeListResource");
  return $resource('/api/v1/atlas/node');
}

export function atlasNodeInfoResource($resource) {
  console.log("inject in atlasdisplaynode module.");
  return $resource('/api/v1/atlas/node/{nodename}');
}
