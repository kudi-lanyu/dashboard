// dashboard module
import chromeModule from "../chrome/module";
import componentsModule from "../common/components/module";
import filtersModule from "../common/filters/module";
// atlas
import {atlasFakeCardComponent} from "./list/fake_component";
import {atlasFakeCardListComponent} from "./list/fakelist_component";
import stateConfig from './stateconfig';

export default angular
  .module(
    'kubernetesDashboard.atlasfakelist',
    [
      'ngMaterial',
      'ngResource',
      'ui.router',
      chromeModule.name,
      componentsModule.name,
      filtersModule.name,
    ])
  .config(stateConfig)
  .component('kdAtlasFakeCard', atlasFakeCardComponent)
  .component('kdAtlasFakeCardList', atlasFakeCardListComponent)
  .factory('kdAtlasFakeListResource', atlasFakeListResource)

export function atlasFakeListResource($resource) {
  console.log("atlasFakeListResource");
  return $resource('/api/v1/atlas/fake')
}
