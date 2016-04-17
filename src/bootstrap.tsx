/// <reference path="../typings/bundle.d.ts" />

import * as ReactDOM from "react-dom";
import * as React from "react";
import {App} from "./components/App";
import {mainReducer} from "./reducer";
import {configureStore} from "./store";
import {Provider} from "react-redux";
import {SinkPool, WebSocketSinkSource} from "./SinkPool";
import {RootActions} from "./actions";
import Store = Redux.Store;

export function bootstrap() {
  const container = document.getElementById("container");
  const wallPool = new SinkPool(new WebSocketSinkSource(`wss://${window.location.host}/dev/ws`), 0);
  const store = configureStore(mainReducer, {
    items: [],
    sink: wallPool,
    wallCapacity: 0,
    wallSize: {width: window.innerWidth, height: window.innerHeight},
    cellSize: {
      width: 200, height: 200,
    },
    focusedEntry: null,
  });
  store.subscribe(() => {
    wallPool.changeCapacity(store.getState().wallCapacity);
  });
  wallPool.on('updated', (items) => {
    store.dispatch(RootActions.updateWallItems(items));
  });
  store.dispatch(RootActions.updateWallCapacity(1));

  var timeoutId = null;
  var intervalMs = Math.floor(1000);

  window.addEventListener('resize', () => {
    if (timeoutId !== null) {
      clearTimeout(timeoutId);
    }
    timeoutId = setTimeout(() => {
      store.dispatch(RootActions.updateWindowSize(window.innerWidth, window.innerHeight));
    }, intervalMs);
  });

  ReactDOM.render(
    <Provider store={store}>
      <App/>
    </Provider>,
    container
  );
}

window.addEventListener("load", () => {
  console.log("Initializing application...");
  bootstrap();
});

