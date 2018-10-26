import chromeModule from '../../chrome/module';
import stateConfig from './stateConfig';
import nodeModule from '../node/module';

/**
 * module for the overview view.
 */
export default angular
  .module(
    'kuberneteDashboard.atlastop',
    [
      'ngMaterial',
      'ngResource',
      'ui.router',
    ])
  .config(stateConfig)
  .factory('kdAtlasTopResource', resource);

function resource($resource) {
  console.log("kdAtlasTopResource============");
  // console.log($resource('api/v1/atlas/atlasctl/top').get());
  return $resource('api/v1/atlas/atlasctl/top')
}
