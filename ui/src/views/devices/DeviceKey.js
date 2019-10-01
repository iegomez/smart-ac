import React, { Component } from "react";
import { withRouter } from 'react-router-dom';

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';

import { CardContent } from "@material-ui/core";

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import DeviceStore from "../../stores/DeviceStore";
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

class DeviceKey extends Component {

  constructor() {
    super();
    this.state = {
      show: false,
    };
    this.onSubmit = this.onSubmit.bind(this);
    this.onToggle = this.onToggle.bind(this);
  }

  componentDidMount() {
    console.log("Mounting device key");
    DeviceStore.getKey(this.props.device.id, (resp) => {
      console.log(resp);
      this.setState({
        apiKey: resp.apiKey,
      });
    });
  }

  onSubmit() {
    DeviceStore.updateKey(this.props.device.id, resp => {
      console.log(resp);
      this.setState({
        apiKey: resp.apiKey,
      });
    })
  }

  onToggle() {
    this.setState({
      show: !this.state.show,
    });
  }

  render() {
    return(
      <Grid container spacing={3}>
        <TitleBar
        >
        <TitleBarTitle title="Device key" />
        </TitleBar>
        <Grid item xs={12}>
          <Card>
            <CardContent>
            <TextField
              ref={(apiKey) => this.apiKey = apiKey}
              id="apiKey"
              label="API key"
              margin="normal"
              value={this.state.apiKey || ""}
              fullWidth
              type={!!this.state.show ? "text" : "password"}
              readOnly
            />
            <Button color="primary" type="submit" onClick={this.onToggle}>{!!this.state.show ? "Hide" : "Show"}</Button>
            <Button color="primary" type="submit" onClick={this.onSubmit}>Regenerate</Button>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    );
  }
}

export default withRouter(DeviceKey);