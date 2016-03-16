/// <reference path="../typings/bundle.d.ts" />

import * as ReactDOM from "react-dom";
import * as React from "react";
import {AppContainer} from "./app";

export function bootstrap() {
    let container = document.getElementById("container");
    console.log("FOOBAR");
    ReactDOM.render(<AppContainer/>, container);
}

window.addEventListener("load", () => {
    console.log("Initializing application...");
    bootstrap();
});

