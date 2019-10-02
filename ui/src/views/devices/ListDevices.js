import React, { Component } from "react";
import { Link } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import Button from '@material-ui/core/Button';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Plus from "mdi-material-ui/Plus";

import TableCellLink from "../../components/TableCellLink";
import DataTable from "../../components/DataTable";
import DeviceStore from "../../stores/DeviceStore";
import theme from "../../theme";


const styles = {
  buttons: {
    textAlign: "right",
  },
  button: {
    marginLeft: 2 * theme.spacing(1),
  },
  icon: {
    marginRight: theme.spacing(1),
  },
};


class ListDevices extends Component {
  constructor() {
    super();
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    DeviceStore.list(limit, offset, callbackFunc);
  }

  getRow(obj) {

    return(
      <TableRow key={obj.id}>
        <TableCellLink to={`/devices/${obj.id}`}>{obj.id}</TableCellLink>
        <TableCellLink to={`/devices/${obj.id}`}>{obj.serialNumber}</TableCellLink>
        <TableCell>{obj.firmwareVersion}</TableCell>
        <TableCell>{obj.registeredAt}</TableCell>
      </TableRow>
    );
  }

  render() {
    return(
      <Grid container spacing={4}>
        <Grid item xs={12} className={this.props.classes.buttons}>
          <TitleBar>
          <TitleBarTitle title="Devices" />
          </TitleBar>
          <Button variant="outlined" className={this.props.classes.button} component={Link} to={`/devices/create`}>
            <Plus className={this.props.classes.icon} />
            Create
          </Button>
        </Grid>
        
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell>ID</TableCell>
                <TableCell>Serial Number</TableCell>
                <TableCell>Firmware Version</TableCell>
                <TableCell>Registered At</TableCell>
              </TableRow>
            }
            getPage={this.getPage}
            getRow={this.getRow}
          />
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(ListDevices);
