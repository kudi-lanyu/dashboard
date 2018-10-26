import stateConfig from './stateconfig';

// other module
import chromeModule from '../chrome/module';
import namespaceModule from '../namespace/module';
import nodeModule from '../node/module';
import persistentVolumeModule from '../persistentvolume/module';
import roleModule from '../role/module';
import storageClassModule from '../storageclass/module';
import clusterModule from '../cluster/module';
// atlas module
import atlasFakeListModule from '../atlasfakelist/module';
import atlasDisplayDataModule from '../atlasdisplaydata/module';

export default angular
  .module(
    'kubernetesDashboard.atlasdisplay',
    [
      'ngMaterial',
      'ngResource',
      'ui.router',
      // other module
      chromeModule.name,
      nodeModule.name,
      namespaceModule.name,
      persistentVolumeModule.name,
      roleModule.name,
      storageClassModule.name,
      clusterModule.name,
      // atlas module
      atlasFakeListModule.name,
      atlasDisplayDataModule.name,
    ])
  .config(stateConfig)
  .factory('kdAtlasDisplayResource', atlasdisplayResource)

function atlasdisplayResource($resource) {
  return $resource('/api/v1/atlas/display')
}
