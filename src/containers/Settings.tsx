import {connect} from "react-redux";
import {Settings as SettingsComponent} from "../components/Settings";
import {AppState} from "../state";

const mapStateToProps = (state: AppState, ownProps) => {
  return {
  };
};

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
  };
};

export const Settings = connect(mapStateToProps, mapDispatchToProps)(SettingsComponent);
