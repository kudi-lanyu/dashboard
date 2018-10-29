

export class AtlasJobCardController {
  constructor($state, kdAtlasJobResource, kdDataSelectService){
    this.job;
    this.jobInfo;
    this.state_ = $state;
    this.kdAtlasJobResource_ = kdAtlasJobResource;
    this.kdDataSelectService_ = kdDataSelectService;
  }

  /** @export */
  $onInit() {
    console.log("atlas job componet! ");
    // console.log(this.node.objectMeta.name);
    // todo: handle namespace
    let query = this.kdDataSelectService_.getDefaultResourceQuery(this.job.objectMeta.namespace, this.job.objectMeta.name);
    this.kdAtlasJobResource_.get(query,
      (response) => {
        this.jobInfo = response;
        console.log(this.jobInfo);
      },
      (err) => {
        console.log("resource get err: ", err);
      });
  }

  getJobStatus() {
    if (this.jobInfo.status.launcherStatus == "Succeeded") {
      return "SUCCEEDED";
    }else if (this.jobInfo.status.launcherStatus == "Failed"){
      return "FAILED";
    }else
      return "PENDING";
  }

  /**
   * Returns true if node is in ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInReadyState() {
    return this.jobInfo.status.launcherStatus === "Succeeded";
  }

  /**
   * Returns true if node is in non-ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInNotReadyState() {
    return this.jobInfo.status.launcherStatus === "Failed";
  }

  /**
   * Returns true if node is in unknown state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInUnknownState() {
    if (!this.isInReadyState() && !this.isInNotReadyState()) {
      return true;
    }else {
      return false;
    }
  }
}

export const atlasJobCardComponent = {
  bindings: {
    'job': '=',
  },
  controller: AtlasJobCardController,
  templateUrl: 'atlasdisplayjob/list/job.html',
};
