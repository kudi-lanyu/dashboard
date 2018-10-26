import {breadcrumbsConfig} from "../common/components/breadcrumbs/service";
import {stateName as chromeStateName} from "../chrome/state";
import {stateName, stateUrl} from "./state";

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    // this state name
    url: stateUrl,
    // add parent state name
    parent: chromeStateName,
    // add other special inrfor to data inhibit from parent can be accept tree struct
    resolve: {
      printState: printState,
    },
    data: {
      [breadcrumbsConfig]: {
        'atlasexample': i18n.MSG_BREADCRUMBS_ATLASEXAMPLE_LABEL,
      },
    },
    views: {
      '': {
        templateUrl: 'atlasexample/example.html',
      },
    },
  });
}

const i18n = {
  /** @type {string} @desc Label 'Cluster' that appears as a breadcrumbs on the action bar. */
  MSG_BREADCRUMBS_ATLASEXAMPLE_LABEL: goog.getMsg('atlasexample'),
};


function printState() {
  console.log("printState");
}
