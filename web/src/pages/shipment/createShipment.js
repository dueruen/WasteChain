import React, { Component } from 'react'
import { CreateShipment } from '../../api/gql/shipments'
import gql from "graphql-tag";
import { useQuery, useMutation } from '@apollo/react-hooks';
//import { withRouter } from 'react-router-dom'

//export default withRouter(class CreateShipment extends Component {



export default class CreateShipmentPage extends Component {
    state = {
        currentHolderID: "",
        wasteType: "",
        location: "",
        password: ""
    }

    /**
     * Sets state data when changes are made in text-inputs
     * @param {Event} event
     */
  handleChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    })
}


    render() {
        return(
            <section>
                <h2>Create Waste Shipment</h2>
                <form onSubmit={this.submitHandler}>
                    <label>
                        User ID
                        <input
                        type="text"
                        name="currentHolderID"
                        onChange={this.handleChange}
                        required>
                        </input>
                    </label>
                    <label>
                        Password
                        <input
                        type="password"
                        name="password"
                        onChange={this.handleChange}
                        required>
                        </input>
                    </label>
                    <label>
                        Waste Type
                        <input
                        type="text"
                        name="wasteType"
                        onChange={this.handleChange}
                        required>
                        </input>
                    </label>
                    <label>
                        Location
                        <input
                        type="text"
                        name="location"
                        onChange={this.handleChange}
                        required>
                        </input>
                    </label>
                </form>
                <input type="submit" value="Create Shipment" onClick={this.submitHandler} />
            </section>
        )
    }

}
