import React, { Component } from "react";
import { withRouter } from 'react-router-dom';

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';

import { CardContent } from "@material-ui/core";

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import DeviceForm from "./DeviceForm";
import DeviceStore from "../../stores/DeviceStore";


class CreateDevice extends Component {
  constructor() {
    super();
    this.onSubmit = this.onSubmit.bind(this);
  }

  onSubmit(device) {
    console.log(device);
    DeviceStore.create(device.serialNumber, device.firmwareVersion, resp => {
      this.props.history.push("/");
    });
  }

  render() {
    return(
      <Grid container spacing={3}>
        <TitleBar>
          <TitleBarTitle title="Devices" to="/" />
          <TitleBarTitle title="/" />
          <TitleBarTitle title="Create" />
        </TitleBar>
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <DeviceForm
                submitLabel="Create device"
                onSubmit={this.onSubmit}
              />
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    );
  }
}

export default withRouter(CreateDevice);