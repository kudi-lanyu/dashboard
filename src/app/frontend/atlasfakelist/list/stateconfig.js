// atlas
import {stateName as atlasdisplayStateName} from '../../atlasdisplay/state';
// dashboard
import {breadcrumbsConfig} from "../../common/components/breadcrumbs/service";

//
import {AtlasFakeListController} from "./controller";
import {stateName as parentState, stateUrl} from '../state';

const i18n = {
  MSG_BREADCRUMBS_ATLASFAKELIST_LABEL: goog.getMsg('Atlasfakelist')
};

/**
 * Config state object for the example list view.
 * @type {!ui.router.StateConfig}
 */
export const config = {
  url: stateUrl,
  parent: parentState, //
  //resolve function will inject to controller.<string, function>
  resolve: {
    'fakeList': resolveAtlasFakeList,
  },
  data: {
    [breadcrumbsConfig]: {
      'label': i18n.MSG_BREADCRUMBS_ATLASFAKELIST_LABEL,
      'parent': atlasdisplayStateName,
    },
  },
  views: {
    '': {
      controller: AtlasFakeListController,
      controllerAs: '$ctrl',
      templateUrl: 'atlasfakelist/list/list.html',
    },
  },
};


function resolveAtlasFakeList(kdAtlasFakeListResource, kdDataSelectService) {
  console.log("atlasfakelistmodule: resolveAtlasFakeList");
  let query = kdDataSelectService.getDefaultResourceQuery();
  // console.log(kdAtlasFakeListResource.get(query).$promise.ListMeta.totalItems);
  return kdAtlasFakeListResource.get(query).$promise;
}


