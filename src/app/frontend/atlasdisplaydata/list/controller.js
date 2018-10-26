

export class AtlasDataListController{
  constructor(dataList, kdAtlasDataListResource){
    this.atlasDataList = dataList;
    this.atlasDataListResource = kdAtlasDataListResource;
  }

  // controller function
  /**
   * @return {boolean}
   * @export
   */
  shouldShowZeroState() {
    /** @type {number} */
    // let resourcesLength = this.listexample.ListSpec.nums;
    // return resourcesLength === 0;
    return true;
  }

}

