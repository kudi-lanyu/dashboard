// parent state
import {stateName as chromeStateName,} from '../chrome/state';
// atlas job
import {stateName} from './state';
// list
import {stateName as listState} from './list/state';
import {config as listConfig} from './list/stateconfig';

// detail
// import {stateName as detailStateName} from './detail/state';
// import {config as detailConfig} from './detail/stateconfig';


export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, config)
    .state(listState, listConfig);
  // .state(detailState, detailConfig);
}

/**
 * Config state object for the Ingress abstract state.
 *
 * @type {!ui.router.StateConfig}
 */
const config = {
  abstract: true,
  parent: chromeStateName,
  template: '<ui-view/>',
  resolve: {
    'test': testCall,
  }
};

function testCall() {
  console.log("atlas display job will not be called");
}
