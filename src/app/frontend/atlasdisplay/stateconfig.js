

// dashboard remain state or config
import {stateName as chromeStateName} from '../chrome/state';
import {breadcrumbsConfig} from '../common/components/breadcrumbs/service';

// atlas state or config
import {stateName, stateUrl} from './state';
import {AtlasdisplayController} from "./controller";

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    url: stateUrl,
    parent: chromeStateName,
    resolve: {
      'atlasdisplay': resolveAtlasdisplay,
      'cluster': resolveClusterResource,
      'atlasfakelist': resolveAtlasFakeList,
    },
    data: {
      [breadcrumbsConfig]: {
        'label': i18n.MSG_BREADCRUMBS_ATLASDISPLAY_LABEL,
      },
    },
    views: {
      '': ({
        controller: AtlasdisplayController,
        controllerAs: '$ctrl',
        templateUrl: 'atlasdisplay/display.html'
      }),
    },
  });
}

/**
 * @param {!angular.$resource} kdAtlasDisplayResource
 * @param {!./../common/dataselect/service.DataSelectService} kdDataSelectService
 * @return {!angular.$q.Promise}
 * @ngInject
 */
function resolveAtlasdisplay(kdAtlasDisplayResource, kdDataSelectService) {
  console.log("resolveAtlasdisplay");
  let paginationQuery = kdDataSelectService.getDefaultResourceQuery();
  return kdAtlasDisplayResource.get(paginationQuery).$promise;
}

/**
 * @param {!angular.$resource} kdAtlasFakeListResource
 * @param {!./../common/dataselect/service.DataSelectService} kdDataSelectService
 * @return {!angular.$q.Promise}
 * @ngInject
 */
function resolveAtlasFakeList(kdAtlasFakeListResource, kdDataSelectService) {
  console.log("resolveAtlasFakeList");
  let paginationQuery = kdDataSelectService.getDefaultResourceQuery();
  return kdAtlasFakeListResource.get(paginationQuery).$promise;
}

/**
 * @param {!angular.$resource} kdClusterResource
 * @param {!./../common/dataselect/service.DataSelectService} kdDataSelectService
 * @return {!angular.$q.Promise}
 * @ngInject
 */
function resolveClusterResource(kdClusterResource, kdDataSelectService) {
  console.log("resolveClusterResource");
  let paginationQuery = kdDataSelectService.getDefaultResourceQuery();
  return kdClusterResource.get(paginationQuery).$promise;
}

const i18n = {
  /** @type {string} @desc Label 'Workloads' that appears as a breadcrumbs on the action bar. */
  MSG_BREADCRUMBS_ATLASDISPLAY_LABEL: goog.getMsg('Atlasdisplay'),
};
