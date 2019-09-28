import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import { withStyles } from "@material-ui/core/styles";

import blue from "@material-ui/core/colors/blue";
import theme from "../theme";


const styles = {
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 10,
  },
  hidden: {
    display: "none",
  },
  flex: {
    flex: 1,
  },
  logo: {
    height: 32,
  },
  search: {
    marginRight: 3 * theme.spacing(1),
    color: theme.palette.common.white,
    background: blue[400],
    width: 450,
    padding: 5,
    borderRadius: 3,
  },
  avatar: {
    background: blue[600],
    color: theme.palette.common.white,
  },
  chip: {
    background: blue[600],
    color: theme.palette.common.white,
    marginRight: theme.spacing(1),
    "&:hover": {
      background: blue[400],
    },
    "&:active": {
      background: blue[400],
    },
  },
  iconButton: {
    color: theme.palette.common.white,
    marginRight: theme.spacing(1),
  },
};


class TopNav extends Component {
  constructor() {
    super();

    this.state = {
    };
  }

  render() {

    return(
      <AppBar className={this.props.classes.appBar}>
        <Toolbar>
          <div className={this.props.classes.flex}>
            <img src="logo192.png" className={this.props.classes.logo} alt="Smart AC" />
          </div>
        </Toolbar>
      </AppBar>
    );
  }
}

export default withStyles(styles)(withRouter(TopNav));
