// dashboard module
import chromeModule from "../chrome/module";
import componentsModule from "../common/components/module";
import filtersModule from "../common/filters/module";
// atlas
import {atlasJobCardComponent} from "./list/job_component";
import {atlasJobCardListComponent} from "./list/joblist_component";
import stateConfig from './stateconfig';

export default angular
  .module(
    'kubernetesDashboard.atlasdisplayjob',
    [
      'ngMaterial',
      'ngResource',
      'ui.router',
      chromeModule.name,
      componentsModule.name,
      filtersModule.name,
    ])
  .config(stateConfig)
  .component('kdAtlasJobCard', atlasJobCardComponent)
  .component('kdAtlasJobCardList', atlasJobCardListComponent)
  .factory('kdAtlasJobListResource', atlasJobListResource)
  .factory('kdAtlasJobResource', atlasJobResource)

export function atlasJobListResource($resource) {
  console.log("atlasJobListResource");
  return $resource('/api/v1/atlas/job');
}

export function atlasJobResource($resource) {
  console.log("atlasJobResource");
  return $resource('/api/v1/atlas/job/:name');
}
