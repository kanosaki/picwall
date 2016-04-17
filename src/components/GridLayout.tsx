import * as React from "react";
import {GridItem} from "./GridItem";

class GridProps {
  cols: number;
  rows: number;
  colWidth: number;
  rowHeight: number;
}

class GridStates {
}

export class GridLayout extends React.Component<GridProps, GridStates> {
  state: {};

  render() {
    return (
      <div className="grid-layout">
        {_.map(this.props.children, this.wrapItem.bind(this))}
      </div>
    );
  }

  private wrapItem(item: React.ReactNode, key: any): React.Component {
    return (
      <GridItem
        key={key}
        colWidth={this.props.colWidth}
        rowHeight={this.props.rowHeight}
        x={key % this.props.cols}
        y={Math.floor(key / this.props.cols)}>
        {item}
      </GridItem>
    );
  }
}
