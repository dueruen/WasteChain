import React, { Component } from 'react'
import gql from "graphql-tag";
import { Query } from 'react-apollo'
import ReactLoading from 'react-loading';
import ShipmentHistoryElement from '../../components/shipments/shipmentHistoryElement/shipmentHistoryElement'
import { navigate } from "@reach/router"


/**
 * Query to get shipments details
 * Has the shipment ID as input parameter
 */
const GET_SHIPMENT_DETAILS =
gql`
    query GetShipmentDetails($shipmentID: ID!){getShipmentDetails(shipmentID: $shipmentID) {
    wasteType, id, currentHolderID, producingCompanyID, history{id, ownerID, event, receiverID, location, timestamp}}}
`;

class DetailedShipmentPage extends Component {
    state = {
        shipmentID: this.props.shipmentID
      }

    renderProcessButton = (ownerID) => {
        let currentUserID
        if(localStorage.getItem('me') !== null){
            currentUserID = JSON.parse(localStorage.getItem('me'))["id"]
        }
        if(currentUserID === null) {return null}
        if(ownerID === null) {return null}
        if(ownerID !== currentUserID) {return null}
        return(
            <button onClick={this.processButtonOnClick}>Process Shipment</button>
        )
    }

    renderTransferButton = (ownerID) => {
        let currentUserID
        if(localStorage.getItem('me') !== null){
            currentUserID = JSON.parse(localStorage.getItem('me'))["id"]
        }
        if(currentUserID === null) {return null}
        if(ownerID === null) {return null}
        if(ownerID !== currentUserID) {return null}
        return(
            <button onClick={this.transferButtonOnClick}>Transfer Shipment</button>
        )
    }

    processButtonOnClick = () => {
        navigate('/shipment/process/' + this.state.shipmentID)
    }

    transferButtonOnClick = () => {
        navigate('/shipment/transfer/' + this.state.shipmentID)
    }

    render() {
        const { shipmentID } = this.state
        return(
            <Query query={GET_SHIPMENT_DETAILS} variables={{shipmentID}}>
                {({ loading, error, data }) => {
                    if (loading) return <ReactLoading type={'spin'}color={'#8bb849'} height={'20%'} width={'20%'}/>
                    if (error) return <div>Error</div>
                    const shipment = data.getShipmentDetails

                    return (
                        <section>
                            <h2>Shipment {shipmentID}</h2>
                            <h3>Waste Type: {shipment.wasteType}</h3>
                            <h3>Producing Company: {shipment.producingCompanyID}</h3>
                            <h3>Current Holder: {shipment.currentHolderID}</h3>
                            <h3>Shipment History: </h3>
                            {shipment.history.map(historyelement =>
                            <ShipmentHistoryElement
                            key={historyelement.id}
                            event={historyelement.event}
                            ownerID={historyelement.ownerID}
                            receiverID={historyelement.receiverID}
                            location={historyelement.location}
                            timestamp={historyelement.timestamp}
                            />)}
                            {this.renderProcessButton(shipment.currentHolderID)}
                            {this.renderTransferButton(shipment.currentHolderID)}
                        </section>
                    )
                }}
            </Query>
        )
    }

}

export default DetailedShipmentPage
