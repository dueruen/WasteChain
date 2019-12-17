import React, { Component } from 'react'
import styled from '@emotion/styled'
import './shipmentLink.css'

class ShipmentLink extends Component {

    onClick = event => {
        alert("TODO: Transfer to detailed shipment view")
    }

    render() {
        return(
            <div className="shipmentlink"
            onClick={this.onClick}>
                <h3>{this.props.id}</h3>
            </div>
        )
    }

}

export default ShipmentLink

