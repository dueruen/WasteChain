import React, { Component } from 'react'
import './shipmentLink.css'
import { navigate } from "@reach/router"

class ShipmentLink extends Component {

    onClick = event => {
        navigate(`/shipment/${this.props.id}`)
    }

    render() {
        return(
            <div className="shipmentlink"
            onClick={this.onClick}>
                <h3>{this.props.id}</h3>
                <h4>Type of waste: {this.props.wasteType}</h4>
            </div>
        )
    }

}

export default ShipmentLink

