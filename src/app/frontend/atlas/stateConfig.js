import {stateName as loginState} from '../login/state';
import {authRequired, namespaceParam, stateName} from './state';

/**
 * Namespace is an abstract state with no path, but with one parameter ?namespace= that
 * is always accepted (since namespace is above all).
 *
 * This state must always be the root in a state tree after login???????. This is enforced during app startup.
 *
 * @param {!ui.router.$stateProvider|kdUiRouter.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName,{
    url: `?${namespaceParam}`,
    abstract: true,
    data:{
      [authRequired]: true,
    },
    views:{
      '':{
        template: '<div ui-view class="kd-content-div-filled"></div>',
      }
    },
  });
  $stateProvider.decorator('parent', requireParentState);
}

/**
 * @param {!Object} stateExtend
 * @param {function(?):!ui.router.$state} parentFn
 * @return {!ui.router.$state}
 */
function requireParentState(stateExtend, parentFn) {
  /** @type {!ui.router.$state} */
  let state = stateExtend['self'];
  if (!state.parent && state.name !== stateName && state.name !== loginState) {
    throw new Error(
      `State "${state.name}" requires parent state to be set to ` +
      `${stateName}. This is likely a programming error.`);
  }
  return parentFn(stateExtend);
}
