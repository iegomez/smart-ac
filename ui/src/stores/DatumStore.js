import { EventEmitter } from "events";

import Swagger from "swagger-client";

import {checkStatus, errorHandler, errorHandlerIgnoreNotFoundWithCallback } from "./helpers";
import dispatcher from "../dispatcher";
import sessionStore from "./SessionStore";

class DatumStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/datum.swagger.json", sessionStore.getClientOpts());
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

  list(startDate, endDate, limit, offset, filters, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DatumService.List({
        startDate: startDate,
        endDate: endDate,
        limit: limit,
        offset: offset,
        filters: filters,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  listForDevice(deviceID, startDate, endDate, limit, offset, filters, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DatumService.ListForDevice({
        device_id: deviceID,
        startDate: startDate,
        endDate: endDate,
        limit: limit,
        offset: offset,
        filters: filters,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  listForDeviceBySerialNumber(serialNumber, startDate, endDate, limit, offset, filters, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DatumService.ListForDeviceBySerialNumber({
        serial_number: serialNumber,
        startDate: startDate,
        endDate: endDate,
        limit: limit,
        offset: offset,
        filters: filters,
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
