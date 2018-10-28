// atlas
import {stateName as atlasdisplayStateName} from '../../atlasdisplay/state';
// dashboard
import {breadcrumbsConfig} from "../../common/components/breadcrumbs/service";

//
import {AtlasJobListController} from "./controller";
import {stateName as parentState, stateUrl} from '../state';

const i18n = {
  MSG_BREADCRUMBS_ATLASDISPLAYJOB_LABEL: goog.getMsg('Atlasdisplayjob')
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
    'jobList': resolveAtlasJobList,
  },
  job: {
    [breadcrumbsConfig]: {
      'label': i18n.MSG_BREADCRUMBS_ATLASDISPLAYJOB_LABEL,
      'parent': atlasdisplayStateName,
    },
  },
  views: {
    '': {
      controller: AtlasJobListController,
      controllerAs: '$ctrl',
      templateUrl: 'atlasdisplayjob/list/list.html',
    },
  },
};


function resolveAtlasJobList(kdAtlasJobListResource, kdDataSelectService) {
  console.log("atlasdisplayjobmodule: resolveAtlasJobList");
  let query = kdDataSelectService.getDefaultResourceQuery();
  return kdAtlasJobListResource.get(query).$promise;
}


