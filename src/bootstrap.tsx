/// <reference path="../typings/bundle.d.ts" />

import * as ReactDOM from "react-dom";
import * as React from "react";
import {AppContainer} from "./app";

function bootstrap() {
    let container = document.getElementById("container");
    ReactDOM.render(<AppContainer/>, container);
}


interface BrowserWindow extends Window {
    bootstrap(): void;
}

declare var window: BrowserWindow;

window.bootstrap = bootstrap;


