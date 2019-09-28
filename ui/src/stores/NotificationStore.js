import { EventEmitter } from "events";
import dispatcher from "../dispatcher";


class NotificationStore extends EventEmitter {
  constructor() {
    super();
    this.notifications = [];
    this.message = {};
  }

  getAll() {
    return this.notifications;
  }


  createNotification(type, message, variant) {
    const id = Date.now();

    this.notifications.push({
      id: id,
      type: type,
      message: message,
      variant: variant,
    });

    //Fow now, let's have a max of 10 notifications so alerts don't fill the page.
    if(this.notifications.length > 10) {
      this.notifications.splice(0, 1);
    }

    this.emit("change");
  }

  deleteNotification(id) {
    let notification = null;

    for(var n of this.notifications) {
      if(n.id === id) {
        notification = n;
      }
    }

    this.notifications.splice(this.notifications.indexOf(notification), 1);
    this.emit("change");
  }

  handleActions(action) {
    switch(action.type) {
      case "CREATE_NOTIFICATION": {
        this.createNotification(action.notification.type, action.notification.message);
        break;
      }
      case "DELETE_NOTIFICATION": {
        this.deleteNotification(action.id);
        break;
      }
      default:
        break;
    }
  }
}


const notificationStore = new NotificationStore();
dispatcher.register(notificationStore.handleActions.bind(notificationStore));

export default notificationStore;
