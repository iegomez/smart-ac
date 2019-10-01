import { EventEmitter } from "events";

import Swagger from "swagger-client";

import {checkStatus, errorHandler, errorHandlerIgnoreNotFoundWithCallback } from "./helpers";
import dispatcher from "../dispatcher";


class DatumStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/datum.swagger.json", {});
  }

  get(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DatumService.Get({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  delete(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DatumService.Delete({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("deleted");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  list(startDate, endDate, limit, offset, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DatumService.List({
        startDate: startDate,
        endDate: endDate,
        limit: limit,
        offset: offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  listForDevice(deviceID, startDate, endDate, limit, offset, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DatumService.ListForDevice({
        device_id: deviceID,
        startDate: startDate,
        endDate: endDate,
        limit: limit,
        offset: offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  notify(action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: "datum has been " + action,
      },
    });
  }

}

const datumStore = new DatumStore();
export default datumStore;
