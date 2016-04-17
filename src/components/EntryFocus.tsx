import * as React from "react";
import {Entry} from "../Entry";
import NavigationClose from "material-ui/lib/svg-icons/navigation/close";
import IconButton from "material-ui/lib/icon-button";

interface EntryFocusProps {
  content: Entry;
  close: () => void;
}

interface EntryFocusState {
  showMetadata: Boolean;
}

export default class EntryFocusView extends React.Component<EntryFocusProps, EntryFocusState> {
  render() {
    const self = this;
    const it = this.props.content;
    return (
      <div className="pic-focus">
        <div className="pic-focus-actions">
          <IconButton onClick={self.props.close}>
            <NavigationClose />
          </IconButton>
        </div>
        <div className="pic-focus-metadata">
          {it.caption}
        </div>
        <img src={it.thumbnail}/>
      </div>
    );
  }
}
