import React, { Component } from "react";
import { Link } from "react-router-dom";

import classNames from "classnames";

import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';

import theme from "../theme";

const styles = {
  title: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginRight: theme.spacing(1),
    float: "left",
  },

  link: {
    textDecoration: "none",
    color: theme.palette.primary.main,
  },
};


class TitleBarTitle extends Component {
  render() {
    let component = null;
    let combinedStyles = null;

    if (this.props.to !== undefined) {
      component = Link;
      combinedStyles = classNames(this.props.classes.title, this.props.classes.link);
    } else {
      combinedStyles = this.props.classes.title;
    }


    return(
      <Typography variant="h6" className={combinedStyles} to={this.props.to} component={component}>
        {this.props.title}
      </Typography>
    );
  }
}

export default withStyles(styles)(TitleBarTitle);
