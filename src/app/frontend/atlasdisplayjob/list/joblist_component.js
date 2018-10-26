export class AtlasJobCardListController {
  constructor(){
    this.atlasJobList;
    this.atlasJobListResource;
  }

  // function
}

export const atlasJobCardListComponent = {
  transclude:{
    'header':'?kdHeader',// define it by yourself
    'zerostate':'?kdEmptyListText',
  },
  controller:AtlasJobCardListController,
  bindings:{
    'atlasJobList':'<',
    'atlasJobListResource':'<',
  },
  templateUrl: 'atlasdisplayjob/list/joblist.html',
};
