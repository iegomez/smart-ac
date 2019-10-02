import { EventEmitter } from "events";

import Swagger from "swagger-client";

import {checkStatus, errorHandler, errorHandlerIgnoreNotFoundWithCallback } from "./helpers";
import dispatcher from "../dispatcher";
import sessionStore from "./SessionStore";

class DeviceStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/device.swagger.json", sessionStore.getClientOpts());
  }

  create(serialNumber, firmwareVersion, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.Create({
        body: {
          serialNumber: serialNumber,
          firmwareVersion: firmwareVersion,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("created");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  get(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.Get({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  getKey(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.GetAPIKey({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  getBySerialNumber(serialNumber, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.GetBySerialNumber({
        serialNumber: serialNumber,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  update(id, serialNumber, firmwareVersion, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.Update({
        "id": id,
        body: {
          serialNumber: serialNumber,
          firmwareVersion: firmwareVersion,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.emit("update");
        this.notify("updated");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });

  }

  updateKey(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.UpdateAPIKey({
        "id": id,
        body: {},
      })
      .then(checkStatus)
      .then(resp => {
        this.emit("update");
        this.notify("updated");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });

  }

  delete(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.Delete({
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

  list(limit, offset, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.List({
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

  listAll(callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceService.ListAll()
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
        message: "device has been " + action,
      },
    });
  }

}

const deviceStore = new DeviceStore();
export default deviceStore;
