import React, { Component } from "react";
import { Link } from "react-router-dom";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import DateDataTable from "../../components/DateDataTable";
import DatumStore from "../../stores/DatumStore";
import DeviceStore from "../../stores/DeviceStore";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import TableCellLink from "../../components/TableCellLink";
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';

import moment from 'moment';

class ListData extends Component {


  constructor() {
    super();
    this.state = {
      devices: [],
      deviceID: -1,
    };
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
    this.onChange = this.onChange.bind(this);
  }

  componentDidMount() {
    DeviceStore.listAll((resp) => {
      this.setState({
        devices: resp.result,
      });
    });
  }


  getPage(limit, offset, startDate, endDate, filters, callbackFunc) {
    if(this.state.deviceID > 0) {
      DatumStore.listForDevice(this.state.deviceID, startDate, endDate, limit, offset, filters, callbackFunc);
    } else {
      DatumStore.list(startDate, endDate, limit, offset, filters, callbackFunc);
    }
  }

  getRow(obj) {

    const val = obj.sensorType == "health_status" ? obj.strVal : obj.val;
    //Add device serial number if we're listing all data.
    if(this.state.deviceID > 0) {
      return (
        <TableRow key={obj.id}>
          <TableCell>{obj.id}</TableCell>
          <TableCell>{obj.createdAt}</TableCell>
          <TableCell>{obj.sensorType}</TableCell>
          <TableCell>{val}</TableCell>
        </TableRow>
      );
    }
    return(
      <TableRow key={obj.id}>
        <TableCell>{obj.id}</TableCell>
        <TableCellLink to={`/devices/${obj.deviceID}`}>{obj.serialNumber}</TableCellLink>
        <TableCell>{obj.createdAt}</TableCell>
        <TableCell>{obj.sensorType}</TableCell>
        <TableCell>{val}</TableCell>
      </TableRow>
    );

  }

  onChange(e) {
    this.setState({
      deviceID: e.target.value,
    });
  }

  render() {

    //Add device column if we're listing all data.
    let vHeader = this.state.deviceID > 0 ?
      <TableRow>
        <TableCell>ID</TableCell>
        <TableCell>Created At</TableCell>
        <TableCell>Sensor type</TableCell>
        <TableCell>Value</TableCell>
      </TableRow>
      :
      <TableRow>
        <TableCell>ID</TableCell>
        <TableCell>Device</TableCell>
        <TableCell>Created At</TableCell>
        <TableCell>Sensor type</TableCell>
        <TableCell>Value</TableCell>
      </TableRow>;

    const devices = this.state.devices.map((device, i) => <MenuItem key={device.id} value={device.id}>{device.serialNumber}</MenuItem>);    

    return (
      <Grid container spacing={3}>
        <TitleBar>
        <TitleBarTitle title="Data" />
        </TitleBar>
        
        <Grid item xs={12}>
          <Select
            id="deviceID"
            value={this.state.deviceID || ""}
            onChange={this.onChange}
            inputProps={{
              name: "deviceID",
              id: "deviceID",
            }}
          >
            <MenuItem value={-1} key={-1}>Select device</MenuItem>
            {devices}
          </Select>
          <DateDataTable
            header={vHeader}
            getPage={this.getPage}
            getRow={this.getRow}
          />
        </Grid>
      </Grid>
    );
  }
}

export default ListData;
