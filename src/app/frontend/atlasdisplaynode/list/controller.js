

export class AtlasNodeListController{
  // construct
  /**
   *
   */
  constructor(nodeList, kdAtlasNodeListResource){
    this.atlasNodeList = nodeList;
    this.atlasNodeListResource = kdAtlasNodeListResource;
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

