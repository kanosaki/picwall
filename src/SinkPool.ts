import * as EventEmitter3 from "eventemitter3";

export class SinkPool extends EventEmitter3 {
  source: SinkSource<any>;
  capacity: number;
  buffer: Array<any>;
  hasError: boolean;
  requesting: Boolean;

  constructor(source: SinkSource<any>, capacity: number) {
    this.hasError = false;
    this.buffer = [];
    this.source = source;
    source.on('message', (msg) => {
      this.push(msg);
    });
    this.capacity = capacity;
    this.requesting = false;
    this.fit();
  }

  changeCapacity(cap: number) {
    this.capacity = cap;
    this.fit();
  }

  protected fit() {
    if (this.requesting) {
      return;
    }
    const shortage = this.capacity - this.buffer.length;
    if (shortage > 0) {
      this.expand(shortage);
    } else if (shortage < 0) {
      this.shrink(-shortage);
    }
  }

  private shrink(count: number) {
    console.log(`Shrinking Pool: ${this.buffer.length} - ${count}`);
    // drop first `count` elements from buffer
    this.buffer = this.buffer.slice(count, this.buffer.length - count);
  }

  private expand(count: number) {
    this.requesting = true;
    console.log(`Expanding Pool: ${this.buffer.length} + ${count}`);
    this.source.request(count);
  }

  protected push(items: Array<any>) {
    const lengthWillBe = this.buffer.length + items.length;
    if (lengthWillBe > this.capacity) {
      this.shrink(lengthWillBe - this.capacity);
    } else if (lengthWillBe < this.capacity) {
      this.fit();
    } else {
      this.requesting = false;
    }
    this.buffer = this.buffer.concat(items);
    this.emit('updated', this.buffer);
  }

  renew(count: number) {
    // drop first `count` elements from buffer
    this.buffer = this.buffer.slice(count, this.buffer.length - count);
    this.fit();
  }
}


// TODO: let this be interface.
export abstract class SinkSource<T> extends EventEmitter3 {
  abstract request(count: number): Promise<T>;
}

export class WebSocketSinkSource extends SinkSource<any> {
  ws: WebSocket;
  isOpened: Boolean;
  openPromise: Promise;
  reqIdCounter: number;

  constructor(src: string) {
    this.isOpened = false;
    this.reqIdCounter = 0;
    let ws = new WebSocket(src);
    this.ws = ws;
    console.log(`Openning WebSocket: ${src}`);
    ws.onmessage = this.onMessage.bind(this);
    ws.onerror = this.onError.bind(this);
    ws.onclose = this.onClose.bind(this);
    this.openPromise = new Promise((resolve, reject) => {
      ws.onopen = () => {
        console.log("Socket Opened.");
        resolve();
        this.isOpened = true;
      };
    });
  }

  request(count: number) {
    let reqId = this.reqIdCounter++;
    console.log("Requesting", count, "items...");
    this.openPromise
      .then(() => {
        this.ws.send(JSON.stringify({
          count: count,
          reqId: reqId,
        }));
      })
      .catch((err) => {
        console.log(err);
      });
  }

  onMessage(msg: MessageEvent) {
    this.emit('message', JSON.parse(msg.data).contents);
  }

  onError(msg) {
    this.emit('error', msg);
  }

  onClose() {
    this.isOpened = false;
  }
}