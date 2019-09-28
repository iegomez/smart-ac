import React, { Component } from "react";

import Table from '@material-ui/core/Table';
import TablePagination from '@material-ui/core/TablePagination';
import TableBody from '@material-ui/core/TableBody';
import TableHead from '@material-ui/core/TableHead';
import { withStyles } from '@material-ui/core/styles';

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import TableCellLink from "../../components/TableCellLink";
import Plus from "mdi-material-ui/Plus";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import TitleBarButton from "../../components/TitleBarButton";
import Paper from '@material-ui/core/Paper';
import theme from '../../theme';

import DeviceStore from '../../stores/DeviceStore';


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


class DeviceItem extends Component {
  constructor() {
    super();
  }

  render() {

    const device = this.props.device;

    return(

      <TableRow key={device.id}>
        <TableCellLink to={`/devices/${device.id}`}>{device.id}</TableCellLink>
        <TableCell>{device.serialNumber}</TableCell>
        <TableCell>{device.firmwareVersion}</TableCell>
        <TableCell>{device.registeredAt}</TableCell>
      </TableRow>

    );
  }
}

class ListDevices extends Component {
  constructor(){
    super();
    this.state = {
      count: 0,
      rowsPerPage: 10,
      page: 0,
      loaded: {
        rows: false,
      },
      devices: [],
    }

    this.listDevices = this.listDevices.bind(this);
    this.onChangePage = this.onChangePage.bind(this);
    this.onChangeRowsPerPage = this.onChangeRowsPerPage.bind(this);
  }

  componentDidMount() {
    this.listDevices()
  }

  listDevices() {
    DeviceStore.list(this.state.rowsPerPage, (this.state.page) * this.state.rowsPerPage, (resp) => {
      if (resp.totalCount > 0) {
        this.setState({
          devices: resp.devices,
          count: parseInt(resp.totalCount, 10),
        });
      }
    });
  }

  onChangePage(event, value) {
    this.setState({
      page: value,
    },
      () => {
        this.listDevices();
      }
    );
  }

  onChangeRowsPerPage(event) {
    this.setState({
      rowsPerPage: event.target.value,
    },
      () => {
        this.listDevices();
      }
    );
  }

  render() {

    const devices = this.state.devices.map((device, i) => <DeviceItem key={i} device={device} />);

    return(
      <Grid container spacing={3}>
        <TitleBar
          buttons={[
            <TitleBarButton
              key={1}
              label="Create"
              icon={<Plus />}
              to={`/devices/create`}
            />,
          ]}
        >
        <TitleBarTitle title="Devices" />
        </TitleBar>
        <Grid item xs={12}>
          <Paper className={this.props.classes.root}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>ID</TableCell>
                  <TableCell>Serial Number</TableCell>
                  <TableCell>Firmware Version</TableCell>
                  <TableCell>Registered At</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {devices}
              </TableBody>
            </Table>
            <TablePagination
              component="div"
              count={this.state.count}
              rowsPerPage={this.state.rowsPerPage}
              page={this.state.page}
              onChangePage={this.onChangePage}
              onChangeRowsPerPage={this.onChangeRowsPerPage}
            />
          </Paper>
        </Grid>
      </Grid>
    );

  }

}

export default withStyles(styles)(ListDevices);