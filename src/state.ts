// application root state
import {SinkPool} from "./SinkPool";
import {Size} from "./utils";
import {Entry} from "./Entry";

export class AppState {
  items: Array<any>;
  sink: SinkPool;
  wallCapacity: number;
  wallSize: Size;
  cellSize: Size;
  focusedEntry: Entry;
}