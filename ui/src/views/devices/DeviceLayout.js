
import React, { Component } from "react";
import { Route, Switch, Link, withRouter } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Tooltip from '@material-ui/core/Tooltip';

import Delete from "mdi-material-ui/Delete";

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import TitleBarButton from "../../components/TitleBarButton";

import DeviceStore from "../../stores/DeviceStore";

import UpdateDevice from "./UpdateDevice";
import ListData from "../data/ListDeviceData";
import DeviceKey from "./DeviceKey";

import theme from "../../theme";

const styles = {
  tabs: {
    borderBottom: "1px solid " + theme.palette.divider,
    height: "49px",
  },
  tab: {
    minWidth: 60,
  },
};

class DeviceLayout extends Component {
  constructor() {
    super();
    this.state = {
      tab: 0,
    };
    this.onChangeTab = this.onChangeTab.bind(this);
    this.deleteDevice = this.deleteDevice.bind(this);
    this.locationToTab = this.locationToTab.bind(this);
    this.getDevice = this.getDevice.bind(this);
  }

  componentDidMount() {
    this.getDevice();
  }

  componentDidUpdate(oldProps) {
    if (this.props === oldProps) {
      return;
		}
    this.locationToTab();
  }

  getDevice() {
    DeviceStore.get(this.props.match.params.id, (resp) => {
      this.setState({
        device: resp,
      },
        () => {
          this.locationToTab();
        }
      );
    });
  }

  deleteDevice() {
    if (window.confirm("Are you sure you want to delete this device?")) {
      DeviceStore.delete(this.props.match.params.id, resp => {
        this.props.history.push(`/`);
      });
    }
  }

  onChangeTab(e, v) {
    this.setState({
      tab: v,
    });
  }

  locationToTab() {

    let tab = 0;

    if (window.location.href.endsWith("/edit")) {
      tab = 1;
    } else if (window.location.href.endsWith("/data")) {
      tab = 2;
    }

    this.setState({
      tab: tab,
    });
  }

  render() {
    if (this.state.device === undefined) {
      return(<div></div>);
    }

    return(
      <Grid container spacing={4}>
        <TitleBar
          buttons={
            <TitleBarButton
              label="Delete"
              icon={<Delete />}
              color="secondary"
              onClick={this.deleteDevice}
            />
          }
        >
          <TitleBarTitle to={`/`} title="Devices" />
          <TitleBarTitle title="/" />
          <TitleBarTitle to={`/devices/${this.props.match.params.id}`} title={this.state.device.serialNumber} />
        </TitleBar>
        

        <Grid item xs={12}>
          <Tabs
            value={this.state.tab}
            onChange={this.onChangeTab}
            indicatorColor="primary"
            className={this.props.classes.tabs}
            variant="scrollable"
            scrollButtons="auto"
          >
            <Tooltip title="Key">
              <Tab style={styles.tab} label="Key" component={Link} to={`/devices/${this.props.match.params.id}`} />
            </Tooltip>
            <Tooltip title="Edit">
              <Tab style={styles.tab} label="Edit" component={Link} to={`/devices/${this.props.match.params.id}/edit`} />
            </Tooltip>
            <Tooltip title="Data">
              <Tab style={styles.tab} label="Data" component={Link} to={`/devices/${this.props.match.params.id}/data`} />
            </Tooltip>
          </Tabs>
        </Grid>

        <Grid item xs={12}>
          <Switch>
            <Route exact path={`${this.props.match.path}/`} render={props => <DeviceKey device={this.state.device} {...props} />} />
            <Route exact path={`${this.props.match.path}/edit`} render={props => <UpdateDevice device={this.state.device} {...props} />} />
            <Route exact path={`${this.props.match.path}/data`} render={props => <ListData device={this.state.device} {...props} />} />
          </Switch>
        </Grid>
      </Grid>
    );
  }

}

export default withStyles(styles)(withRouter(DeviceLayout));