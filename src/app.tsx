/// <reference path="../typings/bundle.d.ts" />

import * as React from "react";


export class AppStates {

}

export class AppProps {

}


export class AppContainer extends React.Component<AppProps, AppStates> {
    state: AppStates = {};

    render() {
        return (
            <div>
                Hello world!
            </div>
        );
    }

}
