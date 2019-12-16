import React, { Component } from 'react'
import { CreateShipment } from '../../api/gql/shipments'
import gql from "graphql-tag";
import { useQuery, useMutation } from '@apollo/react-hooks';
import client from '../../api/gql/ApolloClient'
import { Mutation } from 'react-apollo'
//import { withRouter } from 'react-router-dom'


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
    currentHolderID: '',
    location: '',
    password: ''
  }

  render() {
    const { currentHolderID, wasteType, location, password } = this.state
    return (
      <div>
        <div className="flex flex-column mt3">
          <input
            className="mb2"
            value={currentHolderID}
            onChange={e => this.setState({ currentHolderID: e.target.value })}
            type="text"
            placeholder="ID"
          />
          <input
            className="mb2"
            value={password}
            onChange={e => this.setState({ password: e.target.value })}
            type="password"
            placeholder="Password"
          />
          <input
            className="mb2"
            value={location}
            onChange={e => this.setState({ location: e.target.value })}
            type="text"
            placeholder="Location"
          />
          <input
            className="mb2"
            value={wasteType}
            onChange={e => this.setState({ wasteType: e.target.value })}
            type="text"
            placeholder="wastetype"
          />
        </div>
        <Mutation mutation={CREATE_SHIPMENT} variables={{ currentHolderID, password, location, wasteType }}>
            {createShipment => <button onClick={createShipment}>Submit</button>}
        </Mutation>
      </div>
    )
  }
}

export default CreateShipmentPage
