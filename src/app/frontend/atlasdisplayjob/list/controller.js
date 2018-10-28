

export class AtlasJobListController{
  // construct
  /**
   *
   */
  constructor(jobList, kdAtlasJobListResource){
    this.atlasJobList = jobList;
    this.atlasJobListResource = kdAtlasJobListResource;
    console.log("atlas job list constructor");
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

