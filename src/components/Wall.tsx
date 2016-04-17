import * as React from "react";
import {GridLayout} from "./GridLayout";
import {Size} from "../utils";
import {Entry} from "../Entry";
import EntryFocusView from "./EntryFocus";
import EntrySmallView from "./EntrySmall";
import * as _ from "lodash";

export class WallProps {
  items: Entry[];
  cellSize: Size;
  wallSize: Size;
  focusedEntry: Entry;
  changeCapacity: (number) => void;
  focusEntry: (Entry) => void;
}

export class Wall extends React.Component<WallProps, any> {
  componentDidUpdate(prevProps: WallProps, prevState: any) {
    const colWidth = this.props.cellSize.width;
    const rowHeight = this.props.cellSize.height;
    const cols = Math.ceil(this.props.wallSize.width / colWidth);
    const rows = Math.ceil(this.props.wallSize.height / rowHeight);
    if (cols * rows !== this.props.items.length) {
      this.props.changeCapacity(cols * rows);
    }
  }

  renderGrid(): React.ReactElement {
    const self = this;
    const colWidth = this.props.cellSize.width;
    const rowHeight = this.props.cellSize.height;
    const cols = Math.ceil(this.props.wallSize.width / colWidth);
    const rows = Math.ceil(this.props.wallSize.height / rowHeight);
    const picItems = _.map(this.props.items,
      it => (
        <EntrySmallView
          key={it.id} content={it} 
          width={colWidth} height={rowHeight}
          onClick={self.props.focusEntry}
        /> )
    );
    return (
      <GridLayout cols={cols}
                  rows={rows}
                  colWidth={colWidth}
                  rowHeight={rowHeight}
      >
        {picItems}
      </GridLayout>
    );
  }

  renderFocus(): React.ReactElement {
    const entry = this.props.focusedEntry;
    const self = this;

    return (
      <EntryFocusView
        content={entry} 
        close={() => self.props.focusEntry(null)} 
      />
    );
  }

  render() {
    if (this.props.focusedEntry !== null) {
      return this.renderFocus();
    } else {
      return this.renderGrid();
    }
  }

  componentDidMount() {

  }
}

