

export class AtlasJobCardController {
  constructor($state){
    this.job;
    this.state_ = $state;
  }

  /**
   * Returns true if node is in ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInReadyState() {
    return this.job.ready === 'True';
  }

  /**
   * Returns true if node is in non-ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInNotReadyState() {
    return this.job.ready === 'False';
  }

  /**
   * Returns true if node is in unknown state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInUnknownState() {
    return this.job.ready === 'Unknown';
  }
}

export const atlasJobCardComponent = {
  bindings: {
    'job': '=',
  },
  controller: AtlasJobCardController,
  templateUrl: 'atlasdisplayjob/list/job.html',
};
