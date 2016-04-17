
import {applyMiddleware} from "redux";
import {applyMiddleware, Store} from "redux";
import {loggerMiddleware} from "./utils";
import {createStore} from 'redux';
import {AppState} from './state';

export function configureStore(reducer, initState: AppState): Store<AppState> {
  return createStore(reducer, initState, applyMiddleware(loggerMiddleware));  
}