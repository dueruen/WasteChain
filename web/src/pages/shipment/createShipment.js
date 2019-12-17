import React, { Component } from 'react'
import gql from "graphql-tag";
import { Mutation } from 'react-apollo'

/**
 * Query to create a shipment
 */
const CREATE_SHIPMENT =
gql`
    mutation CreateShipment($wasteType: String!, $currentHolderID: String!, $location: String!, $password: String!)
    {createShipment(request:
        {
            wasteType: $wasteType,
            currentHolderID: $currentHolderID,
            location: $location,
            password: $password
        }
    )}
`;


class CreateShipmentPage extends Component {
  state = {
    wasteType: '',
    currentHolderID: JSON.parse(localStorage.getItem('me'))["id"],
    location: '',
    password: ''
  }

  render() {
    const { currentHolderID, wasteType, location, password } = this.state
    return (
        <section>
            <h2>Create Shipment</h2>
            <form>
                <label>
                    Employee ID
                    <input
                    value={currentHolderID}
                    onChange={e => this.setState({ currentHolderID: e.target.value })}
                    type="text"
                    required
                    />
                </label>
                <br/>
                <label>
                    Password
                    <input
                    value={password}
                    onChange={e => this.setState({ password: e.target.value })}
                    type="password"
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
                    Waste Type
                    <input
                    value={wasteType}
                    onChange={e => this.setState({ wasteType: e.target.value })}
                    type="text"
                    required
                    />
                </label>
            </form>
            <br/>
            <Mutation mutation={CREATE_SHIPMENT} variables={{ currentHolderID, password, location, wasteType }}>
                {createShipment => <button onClick={createShipment}>Create Shipment</button>}
            </Mutation>

        </section>
    )
  }
}

export default CreateShipmentPage
