import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";

class DeviceForm extends FormComponent {
  render() {
    if (this.state.object === undefined) {
      return(<div></div>);
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={this.onSubmit}
      >
        <TextField
          id="serialNumber"
          label="Serial Number"
          margin="normal"
          value={this.state.object.serialNumber || ""}
          onChange={this.onChange}
          required
          fullWidth
        />
        <TextField
          id="firmwareVersion"
          label="Firmware Version"
          margin="normal"
          value={this.state.object.firmwareVersion || ""}
          onChange={this.onChange}
          fullWidth
          required
        />
      </Form>
    );
  }
}

export default DeviceForm;
