import * as React from "react";
import {Rect} from "../utils";

class GridItemProps {
  rowHeight: number;
  colWidth: number;
  style: Object;
  x: number;
  y: number;
  w: number;
  h: number;
}

class GridItemStates {
}

export class GridItem extends React.Component<GridItemProps, GridItemStates> {
  constructor(props: GridItemProps) {
    super(props);
    this.state = {};
  }

  state: GridItemStates = new GridItemStates();

  render() {
    let style = this.createTransform(
      this.props.y * this.props.rowHeight,
      this.props.x * this.props.colWidth,
      this.props.w * this.props.colWidth,
      this.props.h * this.props.rowHeight
    );

    return (
      <div className="grid-item" style={style}>
        {this.props.children}
      </div>
    );
  }

  protected createTransform(top, left, width, height): Object {
    const translate = `translate(${left}px, ${top}px)`;
    return {
      position: "absolute",
      transform: translate,
      width: `${width}px`,
      height: `${height}px`,
    };
  }
}
