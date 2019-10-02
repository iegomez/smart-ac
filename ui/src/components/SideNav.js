import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import Divider from '@material-ui/core/Divider';
import Account from "mdi-material-ui/Account";
import Server from "mdi-material-ui/Server";
import RadioTower from "mdi-material-ui/RadioTower";
import Admin from "./Admin";

import theme from "../theme";


const styles = {
  drawerPaper: {
    position: "fixed",
    width: 270,
    paddingTop: theme.spacing(9),
  },
  select: {
    paddingTop: theme.spacing(1),
    paddingLeft: theme.spacing(3),
    paddingRight: theme.spacing(3),
    paddingBottom: theme.spacing(1),
  },
};

class SideNav extends Component {
  constructor() {
    super();

    this.state = {
      open: true,
    };

  }

  componentDidMount() {
    
  }

  componentDidUpdate(prevProps) {
    if (this.props === prevProps) {
      return;
    }
  }

  render() {

    return(
      <Drawer
        variant="persistent"
        anchor="left"
        open={this.props.open}
        classes={{paper: this.props.classes.drawerPaper}}
      >
        <Admin>
          <List>
            <ListItem button component={Link} to="/users">
              <ListItemIcon>
                <Account />
              </ListItemIcon>
              <ListItemText primary="Users" />
            </ListItem>
          </List>
          <Divider />
        </Admin>
        <List>
          <ListItem button component={Link} to="/">
            <ListItemIcon>
              <Server />
            </ListItemIcon>
            <ListItemText primary="Devices" />
          </ListItem>
          <ListItem button component={Link} to="/data">
            <ListItemIcon>
              <RadioTower />
            </ListItemIcon>
            <ListItemText primary="Data" />
          </ListItem>
        </List>
      </Drawer>
    );
  }
}

export default withRouter(withStyles(styles)(SideNav));
