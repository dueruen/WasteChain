import React, { Component } from 'react'


class ShipmentHistoryElement extends Component {

    onClick = event => {
        alert("TODO: Transfer to detailed shipment view")
    }

    renderEventType = (typeNumber) => {
        if(typeNumber === 0) { return(<h3>Shipment Created</h3>)}
        if(typeNumber === 1) { return(<h3>Shipment Ownership Transfered</h3>)}
        if(typeNumber === 2) { return(<h3>Shipment Processed</h3>)}
    }

    renderReceiver = (typeNumber) => {
        if(typeNumber === 1) { return(<h3>Receiver of Shipment: {this.props.receiverID}</h3>)}
    }

    renderTimeStamp = (timestamp) => {
        let str = timestamp
        str = str.split("+")
        return(
        <h3>Time: {str[0]}</h3>
        )
    }

    render() {
        return(
            <section>
                <div className="historyelement"
                >
                    {this.renderEventType(this.props.event)}
                    <h3>Responsible for shipment: {this.props.ownerID}</h3>
                    {this.renderReceiver(this.props.event)}
                    <h3>Location: {this.props.location}</h3>
                    {this.renderTimeStamp(this.props.timestamp)}
                </div>
            </section>
        )
    }

}

export default ShipmentHistoryElement

