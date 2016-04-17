import {Entry} from "./Entry";
export module ActionTypes {
  export const REDUX_INIT = "@@redux/INIT";
  export const SinkUpdated = "sink/changed";
  export const UpdateWindowSize = "window/resize";
  export const UpdateCellSize = "cell/resize";
  export const UpdateWallItems = "wall/update";
  export const UpdateWallCapacity = "wall/capacity_change";
  export const SetFocusEntry = "wall/focus_entry";
}

export module RootActions {
  export function updateWallItems(items: Array<any>) {
    return {
      type: ActionTypes.UpdateWallItems,
      items: items,
    };
  }

  export function updateWallCapacity(cap: number) {
    return {
      type: ActionTypes.UpdateWallCapacity,
      capacity: cap,
    };
  }

  export function updateWindowSize(width: number, height: number) {
    return {
      type: ActionTypes.UpdateWindowSize,
      width: width,
      height: height,
    };
  }
  
  export function updateFrameSize(width: number, height: number) {
    return {
      type: ActionTypes.UpdateCellSize,
      width: width,
      height: height,
    };
  }
  
  export function focusEntry(entry: Entry) {
    return {
      type: ActionTypes.SetFocusEntry,
      entry: entry,
    };
  }
  
  export function sinkUpdated() {
    return {
      type: ActionTypes.SinkUpdated
    };
  }
}