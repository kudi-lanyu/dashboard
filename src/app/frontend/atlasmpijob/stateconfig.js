import {breadcrumbsConfig} from "../common/components/breadcrumbs/service";
import {stateName as chromeStateName} from "../chrome/state";
// todo: modify parent state to atlastasks
import {stateName, stateUrl} from "./state";

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    url: stateUrl,
    parent: chromeStateName,
    data: {
      [breadcrumbsConfig]: {
        'atlasmpijob': i18n.MSG_BREADCRUMBS_ATLASMPIJOB_LABEL,
      },
    },
    views: {
      '': {
        templateUrl: 'atlasmpijob/deploy.html',
      },
    },
  });
}

const i18n = {
  /** @type {string} @desc Label 'Cluster' that appears as a breadcrumbs on the action bar. */
  MSG_BREADCRUMBS_ATLASMPIJOB_LABEL: goog.getMsg('atlasmpijob'),
};
