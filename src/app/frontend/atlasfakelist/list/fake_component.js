

export class AtlasFakeCardController {
  constructor($state){
    this.spec;
    this.state_ = $state;
  }

  /**
   * Returns true if node is in ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInReadyState() {
    return this.spec.ready === 'True';
  }

  /**
   * Returns true if node is in non-ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInNotReadyState() {
    return this.spec.ready === 'False';
  }

  /**
   * Returns true if node is in unknown state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInUnknownState() {
    return this.spec.ready === 'Unknown';
  }
}

export const atlasFakeCardComponent = {
  bindings: {
    'spec': '=',
  },
  controller: AtlasFakeCardController,
  templateUrl: 'atlasfakelist/list/fake.html',
};
