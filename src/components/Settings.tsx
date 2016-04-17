import * as React from "react";
import RaisedButton from "material-ui/lib/raised-button";
import NavigationCheck from "material-ui/lib/svg-icons/navigation/check";
import NavigationClose from "material-ui/lib/svg-icons/navigation/close";
import Slider from 'material-ui/lib/slider';

interface SettingsProps {
  onDone: () => void;
}

interface SettingsState {

}

export class Settings extends React.Component<SettingsProps, SettingsState> {
  onOK() {
    this.props.onDone();
  }

  onRevert() {
    this.props.onDone();
  }

  render() {
    const self = this; // ??
    const style = {
      margin: '12px'
    };
    return (
      <div className="settings">
        <div>
          <RaisedButton
            label="OK"
            style={style}
            primary={true}
            onClick={() => self.onOK()}
            icon={<NavigationCheck />}
          />
          <RaisedButton
            secondary={true}
            label="Cancel"
            style={style}
            onClick={() => self.onRevert()}
            icon={<NavigationClose />}
          />
        </div>
        <Slider />
      </div>
    );
  }

}
