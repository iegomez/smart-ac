import { EventEmitter } from "events";

import Swagger from "swagger-client";
import { checkStatus, errorHandler, errorHandlerLogin } from "./helpers";

class SessionStore extends EventEmitter {
  constructor() {
    super();
    this.client = null;
    this.user = null;

    this.swagger = Swagger("/swagger/user.swagger.json", this.getClientOpts())
    
    this.swagger.then(client => {
      this.client = client;

      if (this.getToken() !== null) {
        this.fetchProfile(() => {});
      }
    });
  }

  getClientOpts() {
    return {
      requestInterceptor: (req) => {
        if (this.getToken() !== null) {
          req.headers["Grpc-Metadata-Authorization"] = "Bearer " + this.getToken();
        }
        return req;
      },
    }
  }

  setToken(token) {
    localStorage.setItem("jwt", token);
  }

  getToken() {
    return localStorage.getItem("jwt");
  }

  getUser() {
    return this.user;
  }

  isAdmin() {
    if (this.user === undefined || this.user === null) {
      return false;
    }
    return this.user.isAdmin;
  }

  login(login, callBackFunc) {
    this.swagger.then(client => {
      client.apis.UserService.Login({body: login})
        .then(checkStatus)
        .then(resp => {
          this.setToken(resp.obj.jwt);
          this.fetchProfile(callBackFunc);
        })
        .catch(errorHandlerLogin);
    });
  }

  logout(callBackFunc) {
    localStorage.clear();
    this.user = null;
    this.emit("change");
    callBackFunc();
  }

  fetchProfile(callBackFunc) {
    this.swagger.then(client => {
      client.apis.UserService.Profile({})
        .then(checkStatus)
        .then(resp => {
          this.user = resp.obj.user;
          this.emit("change");
          callBackFunc();
        })
        .catch(errorHandler);
    });
  }

}

const sessionStore = new SessionStore();
export default sessionStore;
