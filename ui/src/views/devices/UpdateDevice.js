import React, { Component } from "react";
import { withRouter } from 'react-router-dom';

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';

import { CardContent } from "@material-ui/core";

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import DeviceStore from "../../stores/DeviceStore";
import DeviceForm from "./DeviceForm";

class UpdateDevice extends Component {

  constructor() {
    super();
    this.state = {};
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentWillMount() {
    DeviceStore.get(this.props.match.params.id, (resp) => {
      this.setState({
        device: resp,
      });
    });
  }

  onSubmit(device) {
    DeviceStore.update(device.id, device.serialNumber, device.firmwareVersion, resp => {
      this.props.history.push("/");
    })
  }

  render() {
    return(
      <Grid container spacing={3}>
        <TitleBar
        >
        <TitleBarTitle title="Update device" />
        </TitleBar>
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <DeviceForm
                submitLabel="Update"
                onSubmit={this.onSubmit}
                object={this.state.device}
              />
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    );
  }
}

export default withRouter(UpdateDevice);