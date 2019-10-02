import React, { Component } from "react";
import {Router} from "react-router-dom";
import { Route, Switch } from 'react-router-dom';
import classNames from "classnames";

/* eslint-disable */
import { MuiPickersUtilsProvider } from '@material-ui/pickers';
// pick utils
import MomentUtils from '@date-io/moment';
import moment from "@date-io/moment";

import CssBaseline from "@material-ui/core/CssBaseline";
import { MuiThemeProvider, withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';

import history from "./history";
import theme from "./theme";

import TopNav from "./components/TopNav";
import SideNav from "./components/SideNav";
import Notifications from "./components/Notifications";

// users
import Login from "./views/users/Login";
import ListUsers from "./views/users/ListUsers";
import CreateUser from "./views/users/CreateUser";
import UserLayout from "./views/users/UserLayout";
import ChangeUserPassword from "./views/users/ChangeUserPassword";

// devices
import ListDevices from "./views/devices/ListDevices";
import CreateDevice from "./views/devices/CreateDevice";
import DeviceLayout from "./views/devices/DeviceLayout";

// data
import ListData from "./views/data/ListData";

import SessionStore from "./stores/SessionStore";

const drawerWidth = 270;

const styles = {
  root: {
    flexGrow: 1,
    display: "flex",
    minHeight: "100vh",
    flexDirection: "column",
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  main: {
    width: "100%",
    padding: 2 * 24,
    paddingTop: 115,
    flex: 1,
  },

  mainDrawerOpen: {
    paddingLeft: drawerWidth + (2 * 24),
  },
  footerDrawerOpen: {
    paddingLeft: drawerWidth,
  },
};

class App extends Component {
  constructor() {
    super();

    this.state = {
      user: null,
      drawerOpen: false,
    };

    this.setDrawerOpen = this.setDrawerOpen.bind(this);
  }

  componentDidMount() {
    SessionStore.on("change", () => {
      this.setState({
        user: SessionStore.getUser(),
        drawerOpen: SessionStore.getUser() != null,
      });
    });

    this.setState({
      user: SessionStore.getUser(),
      drawerOpen: SessionStore.getUser() != null,
    });
  }

  setDrawerOpen(state) {
    this.setState({
      drawerOpen: state,
    });
  }


  render() {

    let topNav = null;
    let sideNav = null;

    if (this.state.user !== null) {
      topNav = <TopNav setDrawerOpen={this.setDrawerOpen} drawerOpen={this.state.drawerOpen} user={this.state.user} />;
      sideNav = <SideNav open={this.state.drawerOpen} user={this.state.user} />
    }

    return (
      <MuiPickersUtilsProvider utils={MomentUtils}>
        <Router history={history}>
          <React.Fragment>
            <CssBaseline />
            <MuiThemeProvider theme={theme}>
              <div className={this.props.classes.root}>
                {topNav}
                {sideNav}
                <div className={classNames(this.props.classes.main, this.state.drawerOpen && this.props.classes.mainDrawerOpen)}>
                  <Grid container>
                    <Switch>
                      <Route exact path="/" component={ListDevices} />

                      <Route exact path="/login" component={Login} />
                      <Route exact path="/users" component={ListUsers} />
                      <Route exact path="/users/create" component={CreateUser} />
                      <Route exact path="/users/:userID(\d+)" component={UserLayout} />
                      <Route exact path="/users/:userID(\d+)/password" component={ChangeUserPassword} />

                      <Route exact path="/devices/create" component={CreateDevice} />
                      <Route path="/devices/:id(\d+)" component={DeviceLayout} />
                      <Route exact path="/data" component={ListData} />
                    </Switch>
                  </Grid>
                </div>
              </div>
              <Notifications />
            </MuiThemeProvider>
          </React.Fragment>
        </Router>
      </MuiPickersUtilsProvider>
    );
  }
}

export default withStyles(styles)(App);
