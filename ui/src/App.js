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
import Notifications from "./components/Notifications";

import ListDevices from "./views/devices/ListDevices";
import CreateDevice from "./views/devices/CreateDevice";
import DeviceLayout from "./views/devices/DeviceLayout";
import ListData from "./views/data/ListData";

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
};

class App extends Component {
  constructor() {
    super();

    this.state = {
      user: null,
    };
  }

  componentDidMount() {
    
  }


  render() {
    return (
      <MuiPickersUtilsProvider utils={MomentUtils}>
        <Router history={history}>
          <React.Fragment>
            <CssBaseline />
            <MuiThemeProvider theme={theme}>
              <div style={{
                margin: "40px"
              }}>
                <Grid container>
                  <Switch>
                    <Route exact path="/" component={ListDevices} />
                    <Route exact path="/devices/create" component={CreateDevice} />
                    <Route path="/devices/:id(\d+)" component={DeviceLayout} />
                    <Route exact path="/data" component={ListData} />
                  </Switch>
                </Grid>
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
