export interface Pos {
  x: number; // top
  y: number; // left
}

export interface Size {
  width: number;
  height: number;
}


export interface Rect extends Pos, Size {
}

export const loggerMiddleware = store => next => action => {
  let result = next(action);
  console.log('Dispatch:', action, '-->', store.getState());
  return result
};

export interface RequestPayload {
  reqId: string
}

/// emulates request/response in full-duplex stream.
export class RequestEmulator {
  upstream: WebSocket;
  handlers: Array<Function>;
  sequenceNumber: number;

  constructor(upstream: WebSocket) {
    this.upstream = upstream;
    this.handlers = [];
    this.sequenceNumber = 1;
  }

  request<T>(arg: RequestPayload, timeoutMs: number): Promise<T> {
    const seqNum = this.sequenceNumber++;
    return new Promise((resolve, reject) => {
      if (timeoutMs > 0) {
        let timeoutId = setTimeout(() => {
          reject(); // timeout
          _.remove(this.handlers, h => h(seqNum, null));
        }, timeoutMs);
      }
      this.handlers.push((reqId, data) => {
        if (reqId == seqNum) {
          resolve(data);
        }
      });
    });
  }

  pushMessage(data: RequestPayload) {
    _.remove(this.handlers, h => h(data.reqId, data));
  }
}

