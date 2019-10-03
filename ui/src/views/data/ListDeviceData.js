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

import moment from 'moment';

class ListData extends Component {


  constructor() {
    super();
    this.state = {
    };
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }


  getPage(limit, offset, startDate, endDate, filters, callbackFunc) {
    if (this.props.device !== undefined) {
      DatumStore.listForDevice(this.props.device.id, startDate, endDate, limit, offset, filters, callbackFunc);
    }
  }

  getRow(obj) {

    const val = obj.sensorType == "health_status" ? obj.strVal : obj.val;
    return (
      <TableRow key={obj.id}>
        <TableCell>{obj.id}</TableCell>
        <TableCell>{obj.createdAt}</TableCell>
        <TableCell>{obj.sensorType}</TableCell>
        <TableCell>{val}</TableCell>
      </TableRow>
    );
  }

  render() {

    //Add device column if we're listing all data.
    let vHeader = 
      <TableRow>
        <TableCell>ID</TableCell>
        <TableCell>Created At</TableCell>
        <TableCell>Sensor type</TableCell>
        <TableCell>Value</TableCell>
      </TableRow>;

    return (
      <Grid container spacing={3}>
        <TitleBar>
        <TitleBarTitle title="Data" />
        </TitleBar>
        
        <Grid item xs={12}>

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
