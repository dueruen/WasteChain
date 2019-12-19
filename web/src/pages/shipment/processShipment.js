import React, { Component } from 'react'
import gql from "graphql-tag";
import { Mutation } from 'react-apollo'

/**
 * Mutation to process a shipment
 */
const PROCESS_SHIPMENT =
gql`
    mutation ProcessShipment($shipmentID: String!, $ownerID: String!, $location: String!, $password: String!)
    {processShipment(request:
        {
            shipmentID: $shipmentID,
            ownerID: $ownerID,
            location: $location,
            password: $password,
        }
    )}
`;

class ProcessShipmentPage extends Component {
    state = {
        shipmentID: this.props.shipmentID,
        ownerID: JSON.parse(localStorage.getItem('me'))["id"],
        location: '',
        password: ''
      }


    render() {
        const { shipmentID, ownerID, location, password } = this.state
        return(
            <section>
                <h2>Process Shipment {shipmentID}</h2>
                <form>
                <label>
                    Shipment ID
                    <input
                    value={shipmentID}
                    onChange={e => this.setState({ shipmentID: e.target.value })}
                    type="text"
                    required
                    />
                </label>
                <br/>
                <label>
                    Location
                    <input
                    value={location}
                    onChange={e => this.setState({ location: e.target.value })}
                    type="text"
                    required
                    />
                </label>
                <br/>
                <label>
                    Password of Owner
                    <input
                    value={password}
                    onChange={e => this.setState({ password: e.target.value })}
                    type="password"
                    required
                    />
                </label>
                <br/>
            </form>
            <br/>
            <Mutation mutation={PROCESS_SHIPMENT} variables={{ shipmentID, password, location, ownerID }}>
                {processShipment => <button onClick={processShipment}>Process Shipment</button>}
            </Mutation>
        </section>
        )
    }

}

export default ProcessShipmentPage
