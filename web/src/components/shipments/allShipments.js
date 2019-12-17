import React, { Component } from 'react'
import gql from "graphql-tag";
import { Query } from 'react-apollo'
import ShipmentLink from './shipmentLink/shipmentLink'


/**
 * Query to list all shipments
 */
const LIST_ALL_SHIPMENTS =
    gql`
    query {listAllShipments {
    id
    }}
`;

class AllShipmentsList extends Component {


    render() {
        return(
            <Query query={LIST_ALL_SHIPMENTS}>
                {({ loading, error, data }) => {
                    if (loading) return <div>Fetching</div>
                    if (error) return <div>Error</div>

                    const shipmentsToRender = data.listAllShipments

                    return (
                        <section>
                            <h2>Latest Shipments</h2>
                            {shipmentsToRender.map(shipment => <ShipmentLink key={shipment.id} id={shipment.id}/>)}
                        </section>
                    )
                }}
            </Query>
        )
    }
}

export default AllShipmentsList
