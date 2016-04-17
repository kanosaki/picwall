import * as React from "react";
import LeftNav from "material-ui/lib/left-nav";
import {Wall} from "../containers/Wall";
import {Settings} from "../containers/Settings";
import IconButton from "material-ui/lib/icon-button";
import NavigationMenu from "material-ui/lib/svg-icons/navigation/menu";
import ActionSettings from "material-ui/lib/svg-icons/action/settings";

interface AppState {
  openLeftNav: Boolean;
  openSettingsNav: Boolean;
}


export class App extends React.Component<any, AppState> {
  state: AppState = {
    openLeftNav: false,
    openSettingsNav: false,
  };

  onMenuClicked(open: Boolean) {
    this.setState({
      openLeftNav: open
    } as AppState);
  }

  onSettingsClicked(open: Boolean) {
    this.setState({
      openSettingsNav: open
    } as AppState);
  }

  render() {
    return (
      <div>
        <div style={{
          position: 'absolute',
          zIndex: 10,
        }}>
          <IconButton onClick={() => this.onMenuClicked(true)}>
            <NavigationMenu />
          </IconButton>
          <IconButton onClick={() => this.onSettingsClicked(true)}>
            <ActionSettings />
          </IconButton>
        </div>
        <Wall />
        <LeftNav open={this.state.openLeftNav}>
          <IconButton onClick={() => this.onMenuClicked(false)}>
            <NavigationMenu />
          </IconButton>
        </LeftNav>
        <LeftNav open={this.state.openSettingsNav}>
          <Settings
            onDone={() => this.onSettingsClicked(false)}
          />
        </LeftNav>
      </div>
    );
  }
}

