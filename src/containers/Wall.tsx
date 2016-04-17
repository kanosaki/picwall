import {connect} from "react-redux";
import {Wall as WallComponent} from "../components/Wall";
import {RootActions as Actions} from "../actions";
import {AppState} from "../state";
import {Entry} from "../Entry";

const mapStateToProps = (state: AppState, ownProps) => {
  return {
    focusedEntry: state.focusedEntry,
    items: state.items,
    cellSize: state.cellSize,
    wallSize: state.wallSize,
  };
};

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    changeCapacity: (capacity: number) => {
      dispatch(Actions.updateWallCapacity(capacity));
    },
    focusEntry: (entry: Entry) => {
      dispatch(Actions.focusEntry(entry));
    }
  };
};

export const Wall = connect(mapStateToProps, mapDispatchToProps)(WallComponent);
