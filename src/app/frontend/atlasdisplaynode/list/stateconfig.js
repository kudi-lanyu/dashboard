// atlas
import {stateName as atlasdisplayStateName} from '../../atlasdisplay/state';
// dashboard
import {breadcrumbsConfig} from "../../common/components/breadcrumbs/service";

//
import {AtlasNodeListController} from "./controller";
import {stateName as parentState, stateUrl} from '../state';

const i18n = {
  MSG_BREADCRUMBS_ATLASFAKELIST_LABEL: goog.getMsg('Atlasnodelist')
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
    'nodeList': resolveAtlasNodeList,
  },
  node: {
    [breadcrumbsConfig]: {
      'label': i18n.MSG_BREADCRUMBS_ATLASFAKELIST_LABEL,
      'parent': atlasdisplayStateName,
    },
  },
  views: {
    '': {
      controller: AtlasNodeListController,
      controllerAs: '$ctrl',
      templateUrl: 'atlasdisplaynode/list/list.html',
    },
  },
};


function resolveAtlasNodeList(kdAtlasNodeListResource, kdDataSelectService) {
  console.log("atlasdisplaynodemodule: resolveAtlasNodeList");
  let query = kdDataSelectService.getDefaultResourceQuery();
  // console.log(kdAtlasNodeListResource.get(query).$promise.ListMeta.totalItems);
  return kdAtlasNodeListResource.get(query).$promise;
}


