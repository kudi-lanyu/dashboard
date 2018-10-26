import stateConfig from './stateConfig';
import jobModule from './job/module';
import nodeModule from './node/module';
import podModule from './pod/module';
import namespaceModule from '../common/namespace/module';
import componentsModule from '../common/components/module';

export default angular.module(
  'kubernetesDashboard.atlas',
  [
    'ngMaterial',
    'ui.router',
    componentsModule.name,
    namespaceModule.name,
    jobModule.name,
    nodeModule.name,
    podModule.name,
  ]
).config(stateConfig);
