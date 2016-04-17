import {ActionTypes} from "./actions";
import {AppState} from "./state";

export function mainReducer(state: AppState, action): AppState {
  switch (action.type) {
    case ActionTypes.UpdateWallItems:
      return Object.assign({}, state, {
        items: action.items,
      });
    case ActionTypes.UpdateWindowSize:
      return Object.assign({}, state, {
        wallSize: {
          width: action.width,
          height: action.height,
        },
      });
    case ActionTypes.UpdateCellSize:
      return Object.assign({}, state, {
        cellSize: {
          width: action.width,
          height: action.height,
        }
      });
    case ActionTypes.UpdateWallCapacity:
      return Object.assign({}, state, {
        wallCapacity: action.capacity,
      });
    case ActionTypes.SetFocusEntry:
      return Object.assign({}, state, {
        focusedEntry: action.entry,
      });
    case ActionTypes.SinkUpdated:
      return Object.assign({}, state, {
        items: state.sink.buffer,
      });
    default:
      return state;
  }
}