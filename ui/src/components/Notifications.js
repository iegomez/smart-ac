import React, { Component } from "react";
import { withRouter } from "react-router-dom";

import Snackbar from '@material-ui/core/Snackbar';
import SnackbarContent from '@material-ui/core/SnackbarContent';
import IconButton from '@material-ui/core/IconButton';
import Close from "mdi-material-ui/Close";

import NotificationStore from "../stores/NotificationStore";
import dispatcher from "../dispatcher";

import ErrorIcon from 'mdi-material-ui/AlertCircle';
import WarningIcon from 'mdi-material-ui/Alert';
import InfoIcon from 'mdi-material-ui/Information';
import SuccessIcon from 'mdi-material-ui/CheckCircle';


class Item extends Component {
  constructor() {
    super();
    this.onClose = this.onClose.bind(this);
  }

  onClose(event, reason) {
    dispatcher.dispatch({
      type: "DELETE_NOTIFICATION",
      id: this.props.id,
    });
  }

  render() {

    let icon = <InfoIcon />;
    let sStyle = {
      backgroundColor: "blue",
    };
    const variant = this.props.variant;
    let autoHide = 3000;
    if(variant != null) {
      if(variant === "success") {
        sStyle = {
          backgroundColor: "green",
        };
        icon = <SuccessIcon />;
      } else if (variant === "warning") {
        sStyle = {
          backgroundColor: "#DAA520",
        };
        icon = <WarningIcon />;
        autoHide = 10000;
      } else if (variant === "error") {
        sStyle = {
          backgroundColor: "red",
        };
        icon = <ErrorIcon />;
        autoHide = 15000;
      }
    }

    return(
      <Snackbar
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "left",
        }}
        open={true}
        autoHideDuration={autoHide}
        onClose={this.onClose}
      >
        <SnackbarContent
          style={sStyle}
          message={<span style={{
              display: 'flex',
              alignItems: 'center',
            }}>{icon}&nbsp;{this.props.notification.message}</span>
          }
          action={[
            <IconButton
              key="close"
              aria-label="Close"
              color="inherit"
              onClick={this.onClose}
            >
              <Close />
            </IconButton>
          ]}
        />
      </Snackbar>
    );
  }
}


class Notifications extends Component {
  constructor() {
    super();

    this.state = {
      notifications: NotificationStore.getAll(),
    };
  }

  componentDidMount() {
    NotificationStore.on("change", () => {
      this.setState({
        notifications: NotificationStore.getAll(),
      });
    });
  }

  render() {
    const items = this.state.notifications.map((n, i) => <Item key={n.id} id={n.id} notification={n} variant={n.type} />);

    return (items);
  }
}

export default withRouter(Notifications);
