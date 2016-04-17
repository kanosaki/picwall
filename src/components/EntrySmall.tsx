import * as React from "react";
import {Entry} from "../Entry";


interface EntrySmallProps {
  content: Entry;
  width: number;
  height: number;
  onClick: (Entry) => void;
}

interface EntrySmallState {
}

export default class EntrySmallView extends React.Component<EntrySmallProps, EntrySmallState> {

  render() {
    const it = this.props.content;
    const style = {
      backgroundImage: `url('${it.thumbnail}')`,
      backgroundPosition: "center",
      backgroundSize: "cover",
      zIndex: -1,
      width: `${this.props.width}px`,
      height: `${this.props.height}px`,
    };
    return (
      <div key={it.caption} className="pic-frame" style={style} onClick={() => this.props.onClick(it)}>
        <div className="frame-desc">
          {it.caption}
        </div>
      </div>);
  }
}
