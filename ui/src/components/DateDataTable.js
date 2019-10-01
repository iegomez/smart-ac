import React, { Component } from "react";

import Table from '@material-ui/core/Table';
import TablePagination from '@material-ui/core/TablePagination';
import TableBody from '@material-ui/core/TableBody';
import TableHead from '@material-ui/core/TableHead';
import { withStyles } from '@material-ui/core/styles';
import CloudSearch from "mdi-material-ui/CloudSearch";
import TitleBar from "./TitleBar";
import TitleBarButton from "./TitleBarButton";
import Paper from "./Paper";

import moment from "moment";
import { DateTimePicker } from "@material-ui/pickers";

const styles = {

};


class DateDataTable extends Component {
  constructor() {
    super();

    this.state = {
      count: 0,
      rowsPerPage: 10,
      page: 0,
      loaded: {
        rows: false,
      },
      startDate: moment().subtract(90, "minutes").toISOString(),
      endDate: moment().toISOString(),
    };

    this.onChangePage = this.onChangePage.bind(this);
    this.onChangeRowsPerPage = this.onChangeRowsPerPage.bind(this);
    this.onStartChange = this.onStartChange.bind(this);
    this.onEndChange = this.onEndChange.bind(this);
    this.onUpdate = this.onUpdate.bind(this);
  }

  componentDidMount() {
    this.onChangePage(null, 0);
  }

  componentDidUpdate(prevProps) {
    if (this.props === prevProps) {
      return;
    }

    this.onChangePage(null, 0);
  }

  onStartChange(date) {
    this.setState({
      startDate: date.toISOString(),
    });
  }

  onEndChange(date) {
    this.setState({
      endDate: date.toISOString(),
    });
  }

  onUpdate() {
    //Go back to the first page.
    this.props.getPage(this.state.rowsPerPage, 0, this.state.startDate, this.state.endDate, (result) => {
      this.setState({
        count: parseInt(result.totalCount, 10),
        rows: result.result.map((row, i) => this.props.getRow(row)),
      });
    });
  }

  onChangePage(event, page) {
    this.props.getPage(this.state.rowsPerPage, (page) * this.state.rowsPerPage, this.state.startDate, this.state.endDate, (result) => {
      this.setState({
        page: page,
        count: parseInt(result.totalCount, 10),
        rows: result.result.map((row, i) => this.props.getRow(row)),
      });
    });
  }

  onChangeRowsPerPage(event) {
    this.setState({
      rowsPerPage: event.target.value,
    });

    this.props.getPage(event.target.value, 0, this.state.startDate, this.state.endDate, (result) => {
      this.setState({
        page: 0,
        count: result.totalCount,
        rows: result.result.map((row, i) => this.props.getRow(row)),
      });
    });
  }

  render() {
    if (this.state.rows === undefined) {
      return(<div></div>);
    }

    return(
      <Paper>
        <TitleBar key={1}
          buttons={[
            <DateTimePicker 
              key={1} 
              value={this.state.startDate} 
              onChange={this.onStartChange} 
              id="startDate"
              label="from"
            />,
            <DateTimePicker 
              key={2}
              value={this.state.endDate}
              onChange={this.onEndChange} 
              id="endDate"
              label="to"
            />,
            <TitleBarButton
              key={3}
              label="Update"
              icon={<CloudSearch />}
              onClick={this.onUpdate}
            />
          ]}
        >
        </TitleBar>
        <Table className={this.props.classes.table}>
          <TableHead>
            {this.props.header}
          </TableHead>
          <TableBody>
            {this.state.rows}
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
    );
  }
}

export default withStyles(styles)(DateDataTable);
