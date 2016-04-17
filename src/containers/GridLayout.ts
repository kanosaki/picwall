import {connect} from "react-redux";
import {GridLayout as GridLayoutComponent} from "../components/GridLayout";

const mapStateToProps = (state, ownProps) => {
  return {
    currentCapacity: state.sink.capacity,
  };
};

const mapDispatchToProps = (dispatch, ownProps) => {
  return {};
};

export const GridLayout = connect(mapStateToProps, mapDispatchToProps)(GridLayoutComponent);
